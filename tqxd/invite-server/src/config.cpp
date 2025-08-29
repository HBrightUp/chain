#include <fstream>
#include <iostream>
#include <config.h>

std::set<std::string> setTopic{KA_TOPIC_ATTRIBUTE , KA_TOPIC_ACTION , KA_TOPIC_BUSINESS};
configManager::configManager()
{
}
configManager::~configManager()
{
}
int configManager::init_config(std::string filePath)
{
    if(filePath.empty())
        filePath = "../conf/conf.json";
    nlohmann::json s_json_conf;
    std::ifstream jfile(filePath);
    if (!jfile)
    {
        std::cerr << "No " <<filePath <<" such config file!\n";
        return -__LINE__;
    }

    jfile >> s_json_conf;
    if(!s_json_conf.is_object())
    {
        std::cerr << filePath  <<   "is not json object!\n";
        jfile.close();
        return -__LINE__;
    }
    jfile.close();
    InitSetting(s_json_conf);
    return 0;
}
int configManager::InitSetting(nlohmann::json& tmpJson ){

    if(tmpJson.find("logpath") == tmpJson.end())
        m_sets.m_LogPath = "./log";
    else
        m_sets.m_LogPath = tmpJson["logpath"];
    
    if(tmpJson.find("daemon") == tmpJson.end())
        m_sets.m_Daemon = true;
    else    
        m_sets.m_Daemon = tmpJson["daemon"] == true ? true : false;

    if(tmpJson.find("rpcservertimeout") == tmpJson.end())
        m_sets.m_RpcTimeout = DEFAULT_HTTP_SERVER_TIMEOUT;
    else
        m_sets.m_RpcTimeout = tmpJson["rpcservertimeout"];

    if(tmpJson.find("rpcworkqueue") == tmpJson.end())
        m_sets.m_RpcWorkQueue = 16;
    else
        m_sets.m_RpcWorkQueue = tmpJson["rpcworkqueue"];

    if(tmpJson.find("rpcpassword") == tmpJson.end())
        m_sets.m_RpcPassword = "a";
    else
        m_sets.m_RpcPassword = tmpJson["rpcpassword"];

    if(tmpJson.find("printtoconsole") == tmpJson.end())
        m_sets.m_PrintConsole = false;
    else
        m_sets.m_PrintConsole = tmpJson["printtoconsole"]== "false"?false:true;
        
    if(tmpJson.find("logtimemicros") == tmpJson.end())
        m_sets.m_LogMicTime = 100;
    else
        m_sets.m_LogMicTime = tmpJson["logtimemicros"];

    if(tmpJson.find("httpthread") == tmpJson.end())
        m_sets.m_HttpThread = 1;
    else
        m_sets.m_HttpThread = tmpJson["httpthread"];

    if(tmpJson.find("rpcport") == tmpJson.end())
        m_sets.m_RpcPort = 8332;
    else
        m_sets.m_RpcPort = tmpJson["rpcport"];

    if(tmpJson.find("logtimestamps") == tmpJson.end())
        m_sets.m_LogTimeStamp = true;
    else
        m_sets.m_LogTimeStamp = tmpJson["logtimestamps"]=="true"?true:false;

    if(tmpJson.find("logips") == tmpJson.end())
        m_sets.m_LogIP = false;
    else
        m_sets.m_LogIP = tmpJson["logips"]=="true"?true:false;

    if(tmpJson.find("mysql") == tmpJson.end()){
        std::cerr <<"do not find mysql config!\n";
        return -__LINE__;
    }

    if(tmpJson.find("rpcurl") == tmpJson.end())
        m_sets.m_RpcUrl = "http://127.0.0.1:8332";
    else
        m_sets.m_RpcUrl = tmpJson["rpcurl"];

    if(tmpJson.find("rpcauth") == tmpJson.end())
        m_sets.m_RpcPasswd = "";
    else
        m_sets.m_RpcPasswd = tmpJson["rpcauth"];

    m_sets.m_Mysql.host = tmpJson["mysql"]["url"];
    m_sets.m_Mysql.user = tmpJson["mysql"]["user"];
    m_sets.m_Mysql.pass = tmpJson["mysql"]["pass"];
    m_sets.m_Mysql.dbName = tmpJson["mysql"]["db"];
    m_sets.m_Mysql.port = tmpJson["mysql"]["port"];

    if(tmpJson.find("clickhouse") == tmpJson.end()){
        std::cerr <<"do not find clickhouse config!\n";
        return -__LINE__;
    }
    m_sets.m_Clickhouse.host = tmpJson["clickhouse"]["url"];
    m_sets.m_Clickhouse.user = tmpJson["clickhouse"]["user"];
    m_sets.m_Clickhouse.pass = tmpJson["clickhouse"]["pass"];
    m_sets.m_Clickhouse.dbName = tmpJson["clickhouse"]["db"];
    m_sets.m_Clickhouse.port = tmpJson["clickhouse"]["port"];
    if(tmpJson.find("kafka") == tmpJson.end()){
        std::cerr << "do not find kafka config !"<< std::endl;
        return -__LINE__;
    }
    m_sets.m_Kafka = tmpJson["kafka"];
    if (tmpJson.find("mkkafka") == tmpJson.end()){
        std::cerr << "do not find kafka config !" << std::endl;
        return -__LINE__;
    }
    m_sets.m_mkkafka = tmpJson["mkkafka"];
}