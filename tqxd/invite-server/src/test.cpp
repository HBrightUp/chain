// Copyright (c) 2009-2010 Satoshi Nakamoto
// Copyright (c) 2009-2017 The Bitcoin Core developers
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

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
#include "config.h"
#ifdef WIN32
// Win32 LevelDB doesn't use filedescriptors, and the ones used for
// accessing block files don't count towards the fd_set size limit
// anyway.
#define MIN_CORE_FILEDESCRIPTORS 0
#else
#define MIN_CORE_FILEDESCRIPTORS 150
#endif
extern json json_connect;
bool InitError(const std::string& str)
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
configManager* configManager::m_ConfigInstance = nullptr;
void Interrupt()
{
    InterruptHTTPServer();
    InterruptHTTPRPC();
    InterruptRPC();
   // InterruptREST();
    
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
    //StopREST();
    StopRPC();
    StopHTTPServer();
    // CScheduler/checkqueue threadGroup
    threadGroup.interrupt_all();
    threadGroup.join_all();

#ifndef WIN32
    try {
        fs::remove(GetPidFile());
    } catch (const fs::filesystem_error& e) {
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
static void registerSignalHandler(int signal, void(*handler)(int))
{
    struct sigaction sa;
    sa.sa_handler = handler;
    sigemptyset(&sa.sa_mask);
    sa.sa_flags = 0;
    sigaction(signal, &sa, nullptr);
}
#endif

static CWaitableCriticalSection cs_GenesisWait;
static CConditionVariable condvar_GenesisWait;

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



