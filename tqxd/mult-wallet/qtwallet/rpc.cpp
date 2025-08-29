#include "rpc.h"
//#include <glog/logging.h>
#include <iostream>
#include <chrono>
#include <fstream>
#include <curl/curl.h>
#include <qlogging.h>
#include <QString>

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
        std::cerr << error_str.c_str() << std::endl;
        qWarning("url: ");
        qWarning(params.url.c_str());
        qWarning("error: ");
        qWarning(error_str.c_str());
        return false;
    }
    return true;
}

static bool CurlPost(const std::string& url, const json &json_post, const std::string& auth, json& json_response)
{
    CurlParams curl_params;
    curl_params.auth = auth;
    curl_params.url = url;
    // curl_params.content_type = "content-type:text/plain";
    curl_params.data = json_post.dump();
    std::string response;
    bool ret = CurlPostParams(curl_params,response);
    if (!ret)
        return ret;

    qInfo(response.c_str());
    json_response = json::parse(response);
    if (!json_response["error"].is_null())
    {
        qWarning(response.c_str());
        qWarning(curl_params.data.c_str());
        ret = false;
        return ret;
    }

    return ret;
}

bool Rpc::structRpc(const std::string& method, const json& json_params, json& json_post)
{
    json_post["jsonrpc"] = "2.0";
    json_post["id"] = "curltest";
    json_post["method"] = method;
    json_post["params"] = json_params;
    return true;
}


Rpc Rpc::single_;
Rpc &Rpc::instance()
{
    return single_;
}

void Rpc::initInstance(const std::string &btc_url, const std::string &btc_auth, const std::string &eth_url)
{
    btc_auth_ = btc_auth;
    btc_url_ = btc_url;
    eth_url_ = eth_url;
}

bool Rpc::getBalance(EthAccount &eth_account, std::string& balance)
{
    bool ret = false;
    json json_post;
    json json_params;

    if (eth_account.token)
    {
        json params;
        params["to"] = eth_account.contract;
        params["data"] = "0x70a08231000000000000000000000000" + eth_account.address.substr(2);
        json_params.push_back(params);
        json_params.push_back("latest");
        structRpc("eth_call", json_params, json_post);
    }
    else
    {
        json_params.push_back(eth_account.address);
        json_params.push_back("latest");
        structRpc("eth_getBalance", json_params, json_post);
    }
//    std::cout << json_post.dump() << std::endl;
    json json_response;
    if ( !rpcEth(json_post, json_response) )
    {
        return ret;
    }

    balance = json_response["result"].get<std::string>();
    return ret;
}

bool Rpc::getUtxo(const std::string& address, std::vector<Rpc::Utxo> &vect_utxo, double &total)
{
    bool ret = false;
    json json_post;
    json json_params;
    //json json_addresses;
   // json_addresses.push_back(address);
    //json_params.push_back(1);
    //json_params.push_back(999999);
    json_params.push_back(address);
    json_post["params"] = json_params;

    structRpc("listunspent", json_params, json_post);
    json json_response;
    if ( !rpcBtc(json_post, json_response) )
    {
        return ret;
    }

    json json_unspent = json_response["result"];
  /*  for(size_t i = 0; i < json_unspent.size(); i++)
    {
        Utxo unspent_tx;
        json json_vin;
        unspent_tx.txid = json_unspent[i]["txid"].get<std::string>();
        unspent_tx.n = json_unspent[i]["vout"].get<int>();
        unspent_tx.amount = json_unspent[i]["amount"].get<double>();
        unspent_tx.pubkey_script = json_unspent[i]["scriptPubKey"].get<std::string>();
        total += unspent_tx.amount;
        vect_utxo.push_back(unspent_tx);
        ret = true;
    } */
    std::cout << json_unspent.dump(4) << std::endl;
    for(size_t i = 0; i < json_unspent.size(); i++)
    {
        Utxo unspent_tx;
        json json_vin;
        unspent_tx.txid = json_unspent[i][0].get<std::string>();
        unspent_tx.n = json_unspent[i][1].get<int>();
        unspent_tx.amount = json_unspent[i][2].get<double>();
        unspent_tx.pubkey_script = json_unspent[i][3].get<std::string>();
        total += unspent_tx.amount;
        vect_utxo.push_back(unspent_tx);
        ret = true;
    }
    return ret;
}

bool Rpc::sendBtcSignTx(const std::string& sign_rawtransaction_hex, std::string &txid)
{
    bool ret = false;
    json json_post;
    json json_params;
    json_params.push_back(sign_rawtransaction_hex);
    json_post["params"] = json_params;
    structRpc("sendrawtransaction", json_params, json_post);
    json json_relay;
    json_params.clear();
    json_params.push_back(json_post);
    structRpc("relay", json_params, json_relay);
    json json_response;
    if ( !rpcBtc(json_relay, json_response) )
    {
        return ret;
    }

    txid = json_response["result"].get<std::string>();
    ret = true;
    return ret;
}

bool Rpc::sendEthSignTx(const std::string sign_rawtransaction_hex, std::string &txid)
{
    bool ret = false;
    json json_post;
    json json_params;
    json_params.push_back(sign_rawtransaction_hex);
    json_post["params"] = json_params;
    structRpc("eth_sendRawTransaction", json_params, json_post);
    json json_response;
    if ( !rpcEth(json_post, json_response) )
    {
        return ret;
    }

    txid = json_response["result"].get<std::string>();
    ret = true;
    return ret;
}

bool Rpc::getGasPrice(std::string &gas_price)
{
    bool ret = false;
    json json_post;
    json json_params = json::array();
    json_post["params"] = json_params;
    structRpc("eth_gasPrice", json_params, json_post);
    json json_response;
    if ( !rpcEth(json_post, json_response) )
    {
        return ret;
    }

    gas_price = json_response["result"].get<std::string>();
    ret = true;
    return ret;
}

bool Rpc::getNonce(const std::string &address, std::string &nonce)
{
    bool ret = false;
    json json_post;
    json json_params = json::array();
    json_params.push_back(address);
    json_params.push_back("latest");
    json_post["params"] = json_params;
    structRpc("eth_getTransactionCount", json_params, json_post);
    json json_response;
    if ( !rpcEth(json_post, json_response) )
    {
        return ret;
    }

    nonce = json_response["result"].get<std::string>();
    return ret;
}

bool Rpc::rpcBtc(const json &json_post, json& json_response)
{
    return CurlPost(btc_url_,json_post,btc_auth_, json_response);
}

bool Rpc::rpcEth(const json &json_post, json& json_response)
{
    return CurlPost(eth_url_,json_post,btc_auth_, json_response);
}







