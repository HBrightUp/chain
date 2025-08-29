#ifndef RPC_H_
#define RPC_H_
#include "json.hpp"
#include <iostream>
using json = nlohmann::json;
class Rpc
{
public:
    Rpc()
    {
    }

    ~Rpc()
    {
    }

    struct Utxo
    {
        std::string txid;
        int n ;
        double amount;
        std::string pubkey_script;
    };

    struct EthAccount
    {
        std::string address;
        std::string contract;
        bool token;
        int decimal;

    };
    static Rpc& instance();

    void initInstance(const std::string& btc_url, const std::string& btc_auth, const std::string& eth_url);

    void getEthUrl(std::string& url)
    {
        url = eth_url_;
    }

    void getBtcUrl(std::string& url)
    {
        url = btc_url_;
    }

public:
   
    bool getBalance(EthAccount &eth_account, std::string &balance);//eth

    bool getUtxo(const std::string &address, std::vector<Utxo>& vect_utxo, double& total);//btc

    bool sendBtcSignTx(const std::string& sign_rawtransaction_hex, std::string& txid);//btc

    bool sendEthSignTx(const std::string sign_rawtransaction_hex, std::string& txid);//eth

    bool getGasPrice(std::string& gas_price);//eth

    bool getNonce(const std::string &address, std::string& nonce);//eth

public:
    bool setBtcRpc(std::string node_url, std::string auth)
    {
        btc_url_ = node_url;
        btc_auth_ = auth;
        return true;
    }

    bool setEthRpc(std::string node_url)
    {
        eth_url_ = node_url;
        return true;
    }

    bool rpcBtc(const json& json_post, json& json_response);

    bool rpcEth(const json& json_post, json& json_response);

    bool structRpc(const std::string& method, const json& json_params, json& json_post);	
protected:

    static Rpc single_;
    std::string btc_url_;
    std::string btc_auth_;
    std::string eth_url_;

};

#endif // RPC_H

