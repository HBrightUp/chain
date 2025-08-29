#include <rpc/server.h>
#include <utilstrencodings.h>
#include <rpc/register.h>
#include <curl/curl.h>
#include <sstream>
#include <iostream>
#include <database/db_mysql.h>
#include <config.h>
#include <mutex>
#include <util.h>

struct CurlParams
{
    std::string url;
    std::string auth;
    std::string data;
    std::string content_type;
    bool need_auth;
    CurlParams()
    {
        need_auth = true;
        content_type = "content-type:application/json";
    }
};

static size_t ReplyCallback(void *ptr, size_t size, size_t nmemb, void *stream)
{
    std::string *str = (std::string*)stream;
    (*str).append((char*)ptr, size*nmemb);
    return size * nmemb;
}
static bool CurlPostParams(const CurlParams &params, std::string &response)
{
    CURL *curl = curl_easy_init();
    struct curl_slist *headers = NULL;
    CURLcode res;
    response.clear();
    std::string error_str ;
    if (curl)
    {
        headers = curl_slist_append(headers, params.content_type.c_str());
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
        curl_easy_setopt(curl, CURLOPT_URL, params.url.c_str());
        curl_easy_setopt(curl, CURLOPT_POSTFIELDSIZE, (long)params.data.size());
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, params.data.c_str());

        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, ReplyCallback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *)&response);

        if(params.need_auth)
        {
            curl_easy_setopt(curl, CURLOPT_USERPWD, params.auth.c_str());
            curl_easy_setopt(curl, CURLOPT_HTTPAUTH, CURLAUTH_ANY);
        }

        curl_easy_setopt(curl, CURLOPT_USE_SSL, CURLUSESSL_TRY);
        curl_easy_setopt(curl, CURLOPT_CONNECTTIMEOUT, 120);
        curl_easy_setopt(curl, CURLOPT_TIMEOUT, 120);
        res = curl_easy_perform(curl);
    }
    curl_slist_free_all(headers);
    curl_easy_cleanup(curl);

    if (res != CURLE_OK)
    {
        error_str = curl_easy_strerror(res);
        return false;
    }
    return true;

}

static bool CurlPost(const json &json_post, json& json_response)
{
    CurlParams curl_params;
    configManager* cf = configManager::GetInstance();
    curl_params.auth = cf->m_sets.m_RpcUrl;
    curl_params.url = cf->m_sets.m_RpcPasswd;
    //curl_params.content_type = "content-type:text/plain";
    curl_params.data = json_post.dump();
    std::string response;
    bool ret = CurlPostParams(curl_params,response);
    if (!ret)
        return false;
    json_response = json::parse(response);
    if (!json_response["error"].is_null())
    {
        return false;
    }
    return true;
}
struct despoitInfo {
    std::string txid;
    std::string address;
    std::string amount;
};
std::map<std::string , despoitInfo> mapDespoitData;
json getDeposit(const JSONRPCRequest& request )
{
    json json_result , json_data;
    DBMysql db_mysql;
    if (!db_mysql.openDB())
	{
		throw std::runtime_error("open db failed");
	}
    // txid   address  amount
    std::string address = request.params[0].get<std::string>();
    std::map<std::string , despoitInfo>::iterator iter = mapDespoitData.find(address);
    if(iter != mapDespoitData.end()){
        LogPrintf("return from cache !\n");
        json_result["status"] = "success";
        json_data["address"] = iter->second.address;
        json_data["txid"] = iter->second.txid;
        json_data["amount"] = iter->second.amount;
        json_result["data"] .push_back(json_data);
        return json_result;
    }

    std::stringstream sql ;
    sql << "SELECT txid, address , amount FROM Deposit WHERE address = LOWER('" <<  address << "');" ;
    std::map<int , DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
    map_col_type[2] = DBMysql::STRING;

    json json_d;
    db_mysql.getData(sql.str(), map_col_type, json_d);
    db_mysql.closeDB();
    json json_ret;
    std::stringstream tmp;
    if(json_d.size() > 0)
    {
        json_result["status"] = "success";
        json_data["txid"] = json_d[0][0].get<std::string>();
        json_data["address"] = json_d[0][1].get<std::string>();
        json_data["amount"] = json_d[0][2].get<std::string>();
        despoitInfo inf ;
        inf.txid = json_data["txid"].get<std::string>();
        inf.address = json_data["address"].get<std::string>();
        inf.amount = json_data["amount"].get<std::string>();
        json_result["data"] .push_back(json_data);
        mapDespoitData[address] = inf;
    }
    else
    {
        json_result["status"] = "failed";
    }
    LogPrintf("return from DB !\n");
    return json_result;

}
// TODO Add RPC Method connection here
static const CRPCCommand commands[] =
{ 
    //  category              name                    actor (function)         argNames
      { "ka2ch",            "getDeposit",          &getDeposit,         {} }
};

void RegisterP2PRPCCommands(CRPCTable &t)
{
    for (unsigned int vcidx = 0; vcidx < ARRAYLEN(commands); vcidx++)
        t.appendCommand(commands[vcidx].name, &commands[vcidx]);
}
