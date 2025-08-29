#ifndef UT_CONFIG_H
#define UT_CONFIG_H
#include <string>
#include <vector>
#include <set>
#include <json.hpp>

// kafka 中定义的topic名称
#define KA_TOPIC_ATTRIBUTE  "ckAttribute"
#define KA_TOPIC_ACTION     "ckAction"
#define KA_TOPIC_BUSINESS   "ckBusiness"
static const int DEFAULT_HTTP_THREADS=4;
static const int DEFAULT_HTTP_WORKQUEUE=16;
static const int DEFAULT_HTTP_SERVER_TIMEOUT=30;
extern std::set<std::string> setTopic;
struct DB_cfg{
    std::string host;
    int port;
    std::string user;
    std::string pass;
    std::string dbName;
    std::string charset;
};

struct settings{
    std::string m_LogPath;
    bool m_LogTimeStamp;
    bool m_PrintConsole;
    bool m_LogIP;
    int m_LogMicTime;
    bool m_Daemon;
    std::vector<std::string> m_RpcAllowIP;
    int64_t m_RpcTimeout;
    int m_RpcPort;
    int m_RpcWorkQueue;
    int m_HttpThread;
    std::string m_RpcUrl;
    std::string m_RpcPasswd;
    std::string m_RpcUser;
    std::string m_RpcPassword;
    DB_cfg m_Mysql;
    DB_cfg m_Clickhouse;
    std::string m_Kafka;
    std::string m_mkkafka;
};
/**
 * 从配置文件中读取配置类
 **/ 
class configManager{
public:
    ~configManager();
    int init_config(std::string filePath);
    int InitSetting(nlohmann::json &tmpJson);
    settings m_sets;
    static configManager* m_ConfigInstance;
private:
    configManager();
    
public:
    static configManager *GetInstance(){
        if (m_ConfigInstance == NULL)
            m_ConfigInstance = new configManager();
        return m_ConfigInstance;
    }

    class CGarbo{
    public:
        ~CGarbo(){
            if (m_ConfigInstance)
                delete m_ConfigInstance;
        }
    };
    static CGarbo Garbo;
};

#endif