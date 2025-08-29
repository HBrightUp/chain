#ifndef RPC_H_
#define RPC_H_
#include "common.h"
#include <iostream>

class Rpc
{
public:
    Rpc()
    {
    }

    ~Rpc()
    {
    }

	enum  RpcError
	{
		RPC_OK = 0,
		RPC_NODEDOWN = 1,
		RPC_INTERNAL = 2
	};

public:
	void setNode(const std::string& coin, const std::string& url)
	{
		coin_ = coin;
		url_ = url;
	}

public:
   
    bool eth_blockNumber(uint64_t& height);

	bool eth_newFilter(const uint64_t from_block,
					   const uint64_t to_block,
					   const std::string address,
					   std::string& filter_id);

	bool eth_getFilterLogs(const std::string& filter_id , json& json_result);

	bool eth_getBalance(const std::string& account_address, json& json_result,bool is_token = false, const std::string  contract = "");

protected:

	bool eth_getTokenBalance(const std::string& account_address , const std::string contract , json& json_result);

protected:
	RpcError callRpc(const std::string& url, const std::string& auth,const json& json_post, json& json_response);

    bool structRpc(const std::string& method, const json& json_params, json& json_post);	
private:
	std::string url_;
	std::string coin_;
};
#endif // RPC_H

