// Copyright (c) 2009-2010 Satoshi Nakamoto
// Copyright (c) 2009-2017 The Bitcoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

#if defined(HAVE_CONFIG_H)
#include <config/bitcoin-config.h>
#endif

#include <init.h>
#include <compat/sanity.h>
#include <fs.h>
#include <httpserver.h>
#include <httprpc.h>
#include <netbase.h>
#include <rpc/server.h>
#include <rpc/register.h>
#include <rpc/safemode.h>
#include <scheduler.h>
#include <timedata.h>
#include <util.h>
#include <warnings.h>
#include <stdint.h>
#include <stdio.h>
#include <memory>
#include <boost/program_options.hpp>
#include <boost/filesystem.hpp>

#ifndef WIN32
#include <signal.h>
#endif

#include <boost/algorithm/string/classification.hpp>
#include <boost/algorithm/string/replace.hpp>
#include <boost/algorithm/string/split.hpp>
#include <boost/bind.hpp>
#include <boost/interprocess/sync/file_lock.hpp>
#include <boost/thread.hpp>
#include <openssl/crypto.h>
#include <config.h>
#include <title.h>
#ifdef WIN32
// Win32 LevelDB doesn't use filedescriptors, and the ones used for
// accessing block files don't count towards the fd_set size limit
// anyway.
#define MIN_CORE_FILEDESCRIPTORS 0
#else
#define MIN_CORE_FILEDESCRIPTORS 150
#endif
configManager *cf;

const char *__process__ = "DataCenter";
const char *__version__ = "0.1.0";
extern json json_connect;
bool InitError(const std::string &str)
{
    return false;
}

std::atomic<bool> fRequestShutdown(false);

void StartShutdown()
{
    fRequestShutdown = true;
}

bool ShutdownRequested()
{
    return fRequestShutdown;
}

static boost::thread_group threadGroup;
static CScheduler scheduler;
configManager *configManager::m_ConfigInstance = nullptr;
void Interrupt()
{
    InterruptHTTPServer();
    InterruptHTTPRPC();
    InterruptRPC();
}

void Shutdown()
{
    LogPrintf("%s: In progress...\n", __func__);
    static CCriticalSection cs_Shutdown;
    TRY_LOCK(cs_Shutdown, lockShutdown);
    if (!lockShutdown)
        return;
    /// Note: Shutdown() must be able to handle cases in which initialization failed part of the way,
    /// for example if the data directory was found to be locked.
    /// Be sure that anything that writes files or flushes caches only does this if the respective
    /// module was initialized.
    RenameThread("bitcoin-shutoff");
    StopHTTPRPC();
    StopRPC();
    StopHTTPServer();
    threadGroup.interrupt_all();
    threadGroup.join_all();

#ifndef WIN32
    try
    {
        fs::remove(GetPidFile());
    }
    catch (const fs::filesystem_error &e)
    {
        LogPrintf("%s: Unable to remove pidfile: %s\n", __func__, e.what());
    }
#endif
}

/**
 * Signal handlers are very limited in what they are allowed to do.
 * The execution context the handler is invoked in is not guaranteed,
 * so we restrict handler operations to just touching variables:
 */
#ifndef WIN32
static void HandleSIGTERM(int)
{
    fRequestShutdown = true;
}

static void HandleSIGHUP(int)
{
    fReopenDebugLog = true;
}
#else
static BOOL WINAPI consoleCtrlHandler(DWORD dwCtrlType)
{
    fRequestShutdown = true;
    Sleep(INFINITE);
    return true;
}
#endif

#ifndef WIN32
static void registerSignalHandler(int signal, void (*handler)(int))
{
    struct sigaction sa;
    sa.sa_handler = handler;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = 0;
    sigaction(signal, &sa, nullptr);
}
#endif

void OnRPCStarted()
{
}

void OnRPCStopped()
{
}

static CWaitableCriticalSection cs_GenesisWait;
static CConditionVariable condvar_GenesisWait;

bool AppInitServers()
{
     
    RPCServer::OnStarted(&OnRPCStarted);
    RPCServer::OnStopped(&OnRPCStopped);
     
    if (!InitHTTPServer())
        return false;
    if (!StartRPC())
        return false;
    if (!StartHTTPRPC())
        return false;
    if (!StartHTTPServer())
        return false;
    return true;
}
void InitLogging()
{
    // get from config
    fPrintToConsole = cf->m_sets.m_PrintConsole;
    fLogTimestamps = cf->m_sets.m_LogTimeStamp;
    fLogTimeMicros = cf->m_sets.m_LogMicTime;
    fLogIPs = cf->m_sets.m_LogIP;

    std::string version_string = "httpd"; //FormatFullVersion();
#ifdef DEBUG
    version_string += " (debug build)";
#else
    version_string += " (release build)";
#endif
    LogPrintf(PACKAGE_NAME " version %s\n", version_string);
}

[[noreturn]] static void new_handler_terminate()
{
    // Rather than throwing std::bad-alloc if allocation fails, terminate
    // immediately to (try to) avoid chain corruption.
    // Since LogPrintf may itself allocate memory, set the handler directly
    // to terminate first.
    std::set_new_handler(std::terminate);
    LogPrintf("Error: Out of memory. Terminating.\n");
    // The log was successful, terminate now.
    std::terminate();
};

bool AppInitBasicSetup()
{
    // ********************************************************* Step 1: setup
#ifdef _MSC_VER
    // Turn off Microsoft heap dump noise
    _CrtSetReportMode(_CRT_WARN, _CRTDBG_MODE_FILE);
    _CrtSetReportFile(_CRT_WARN, CreateFileA("NUL", GENERIC_WRITE, 0, nullptr, OPEN_EXISTING, 0, 0));
    // Disable confusing "helpful" text message on abort, Ctrl-C
    _set_abort_behavior(0, _WRITE_ABORT_MSG | _CALL_REPORTFAULT);
#endif
#ifdef WIN32
    // Enable Data Execution Prevention (DEP)
    // Minimum supported OS versions: WinXP SP3, WinVista >= SP1, Win Server 2008
    // A failure is non-critical and needs no further attention!
#ifndef PROCESS_DEP_ENABLE
    // We define this here, because GCCs winbase.h limits this to _WIN32_WINNT >= 0x0601 (Windows 7),
    // which is not correct. Can be removed, when GCCs winbase.h is fixed!
#define PROCESS_DEP_ENABLE 0x00000001
#endif
    typedef BOOL(WINAPI * PSETPROCDEPPOL)(DWORD);
    PSETPROCDEPPOL setProcDEPPol = (PSETPROCDEPPOL)GetProcAddress(GetModuleHandleA("Kernel32.dll"), "SetProcessDEPPolicy");
    if (setProcDEPPol != nullptr)
        setProcDEPPol(PROCESS_DEP_ENABLE);
#endif

#ifndef WIN32
    // get from config
    // TODO
    bool sysPerms = true;
    if (!sysPerms)
    {
        umask(077);
    }

    // Clean shutdown on SIGTERM
    registerSignalHandler(SIGTERM, HandleSIGTERM);
    registerSignalHandler(SIGINT, HandleSIGTERM);

    // Reopen debug.log on SIGHUP
    registerSignalHandler(SIGHUP, HandleSIGHUP);

    // Ignore SIGPIPE, otherwise it will bring the daemon down if the client closes unexpectedly
    signal(SIGPIPE, SIG_IGN);
#else
    SetConsoleCtrlHandler(consoleCtrlHandler, true);
#endif

    std::set_new_handler(new_handler_terminate);
    return true;
}

static bool LockDataDirectory(bool probeOnly)
{
    // Make sure only a single Bitcoin process is using the data directory.
    fs::path datadir = GetDataDir();
    if (!LockDirectory(datadir, ".lock", probeOnly))
    {
        return InitError(strprintf(_("Cannot obtain a lock on data directory %s. %s is probably already running."), datadir.string(), _(PACKAGE_NAME)));
    }
    return true;
}

bool AppInitMain()
{
    // Start the lightweight task scheduler thread
    CScheduler::Function serviceLoop = boost::bind(&CScheduler::serviceQueue, &scheduler);
     
    threadGroup.create_thread(boost::bind(&TraceThread<CScheduler::Function>, "scheduler", serviceLoop));
     
    /* Register RPC commands regardless of -server setting so they will be
     * available in the GUI RPC console even if external calls are disabled.
     */
    RegisterAllCoreRPCCommands(tableRPC);
     
    /* Start the RPC server already.  It will be started in "warmup" mode
     * and not really process calls already (but it will signify connections
     * that the server is there and will be ready later).  Warmup mode will
     * be disabled when initialisation is finished.
     */
     
    if (!AppInitServers())
        return InitError(_("Unable to start HTTP server. See debug log for details."));
     
    SetRPCWarmupFinished();
     
    while (true)
    {
        sleep(10);
    }
    return true;
}

void InitParameterInteraction()
{
    // when specifying an explicit binding address, you want to listen on it
    // even when -connect or -proxy is specified
    if (gArgs.IsArgSet("-bind"))
    {
        if (gArgs.SoftSetBoolArg("-listen", true))
            LogPrintf("%s: parameter interaction: -bind set -> setting -listen=1\n", __func__);
    }
    if (gArgs.IsArgSet("-whitebind"))
    {
        if (gArgs.SoftSetBoolArg("-listen", true))
            LogPrintf("%s: parameter interaction: -whitebind set -> setting -listen=1\n", __func__);
    }

    if (gArgs.IsArgSet("-connect"))
    {
        // when only connecting to trusted nodes, do not seed via DNS, or listen by default
        if (gArgs.SoftSetBoolArg("-dnsseed", false))
            LogPrintf("%s: parameter interaction: -connect set -> setting -dnsseed=0\n", __func__);
        if (gArgs.SoftSetBoolArg("-listen", false))
            LogPrintf("%s: parameter interaction: -connect set -> setting -listen=0\n", __func__);
    }

    if (gArgs.IsArgSet("-proxy"))
    {
        // to protect privacy, do not listen by default if a default proxy server is specified
        if (gArgs.SoftSetBoolArg("-listen", false))
            LogPrintf("%s: parameter interaction: -proxy set -> setting -listen=0\n", __func__);
        // to protect privacy, do not use UPNP when a proxy is set. The user may still specify -listen=1
        // to listen locally, so don't rely on this happening through -listen below.
        if (gArgs.SoftSetBoolArg("-upnp", false))
            LogPrintf("%s: parameter interaction: -proxy set -> setting -upnp=0\n", __func__);
        // to protect privacy, do not discover addresses by default
        if (gArgs.SoftSetBoolArg("-discover", false))
            LogPrintf("%s: parameter interaction: -proxy set -> setting -discover=0\n", __func__);
    }

    if (gArgs.IsArgSet("-externalip"))
    {
        // if an explicit public IP is specified, do not try to find others
        if (gArgs.SoftSetBoolArg("-discover", false))
            LogPrintf("%s: parameter interaction: -externalip set -> setting -discover=0\n", __func__);
    }
}

static bool ParseCmd(int argc, char *argv[])
{
    using namespace boost::program_options;
    std::string conf_file;
    json js_conf;
    boost::program_options::options_description opts_desc("All options");
    opts_desc.add_options()("help,h", "help info")("configure,c", value<std::string>(&conf_file)->default_value("../conf/conf.json"), "configure file");

    variables_map cmd_param_map;
    try
    {
        store(parse_command_line(argc, argv, opts_desc), cmd_param_map);
    }
    catch (boost::program_options::error_with_no_option_name &ex)
    {
        std::cerr << ex.what() << std::endl;
    }
    notify(cmd_param_map);

    if (cmd_param_map.count("help"))
    {
        std::cout << opts_desc << std::endl;
        return false;
    }
    std::ifstream jfile(conf_file);

    if (!jfile)
    {
        std::cerr << "No " << conf_file << " such config file!\n";
        return false;
    }

    jfile >> js_conf;
    if (!js_conf.is_object())
    {
        std::cerr << conf_file << "is not json object!\n";
        jfile.close();
        return false;
    }
    // json_connect = js_conf["mysql"];
    jfile.close();

    return true;
}


int main(int argc, char *argv[])
{
    std::string filePath;
    if (argc < 2)
        filePath = "../conf/conf.json";
    else
        filePath = argv[1];

    char * processName = "data_center";
    utTitle::my_initproctitle(argv , &processName);
    cf = configManager::GetInstance();
    cf->init_config(filePath);

    umask(0);
    {
        int pid = fork();
        if (pid < 0)
        {
            std::cout<< "fork error "<< std::endl;
        }
        else if (pid == 0)
        {
            daemon(1, 1);
            utTitle::my_setproctitle(argv , &processName,"data_center_server123");
            InitLogging();
            setLogFileSuffix("server");
            OpenDebugLog();
            try
            {
                InitParameterInteraction();
                AppInitMain();
            }
            catch (const std::exception &e)
            {
                PrintExceptionContinue(&e, "AppInit()");
            }
            catch (...)
            {
                PrintExceptionContinue(nullptr, "AppInit()");
            }
            return 0;
        }
    }
    LogPrintf("init httpserver success\n");
    return 0;
}
