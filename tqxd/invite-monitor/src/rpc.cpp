#include "rpc.h"
#include <glog/logging.h>
#include <iostream>
#include <chrono>
#include <fstream>
#include <gmp.h>
#include <gmpxx.h>

static Rpc::RpcError CurlPost(const std::string& url, const json &json_post, const std::string& auth, json& json_response)
{
    CurlParams curl_params;
    curl_params.auth = auth;
    curl_params.url = url;
 // curl_params.content_type = "content-type:text/plain";
    curl_params.data = json_post.dump();
    std::string response;
    bool ret = CurlPostParams(curl_params,response);
	if (!ret)
		return Rpc::RPC_NODEDOWN;
    LOG(INFO) << json_post.dump(2);
	
    json_response = json::parse(response);
    LOG(INFO) << json_response.dump(2);
	if (!json_response["error"].is_null())
	{
		LOG(ERROR) << response;
		LOG(ERROR) << curl_params.data;
		return Rpc::RPC_INTERNAL;
	}
	
    return Rpc::RPC_OK;
}

Rpc::RpcError Rpc::callRpc(const std::string& url, const std::string& auth,const json& json_post, json& json_response)
{
	return CurlPost(url, json_post, auth, json_response);
}

bool Rpc::structRpc(const std::string& method, const json& json_params, json& json_post)
{
    json_post["jsonrpc"] = "1.0";
    json_post["id"] = "curltest";
    json_post["method"] = method;
    json_post["params"] = json_params;
    return true;
}

bool  Rpc::eth_blockNumber(uint64_t& height)
{
    json json_response;
    json json_post;
    json json_params = json::array();
    //json_params.push_back("pending");
    structRpc("eth_blockNumber", json_params, json_post);

    Rpc::RpcError ret_type = callRpc(url_, "", json_post, json_response);
    if(ret_type != Rpc::RPC_OK)
    {
        LOG(WARNING) << "coin: "  << coin_  << " url: " << url_;
        LOG(WARNING) << "RPC eth_blockNumber fail!";
        return false; 
    }
    
    mpz_class mp_height;
	std::string hex_height = json_response["result"].get<std::string>();
    mp_height = hex_height;
	//mp_height.set_str(hex_height, 16);
    height = mp_height.get_ui();
    return true;
}

bool  Rpc::eth_newFilter(const uint64_t from_block,
					   const uint64_t to_block,
					   const std::string address,
					   std::string& filter_id)
{
    json json_response;
    json json_post;
    json json_params = json::array();
    json json_filter;
    mpz_class mp_from_block, mp_to_block;
    mp_from_block = from_block;
    mp_to_block = to_block;
    json_filter["address"] = address;
    json_filter["fromBlock"] ="0x" + mp_from_block.get_str(16);
    json_filter["toBlock"] =  "0x" + mp_to_block.get_str(16);
    std::string extern_topic = mp_to_block.get_str(16);
    size_t count = address.size() + extern_topic.size();
    std::string topic = address;
    for (size_t i = count; i < 67; i++)
    {
        topic += "0";
    }
    topic += extern_topic;
    
    //json_filter["topics"].push_back(topic);
    json_filter["topics"] = json::array();
    json_params.push_back(json_filter);
    structRpc("eth_newFilter", json_params, json_post);
    Rpc::RpcError ret_type = callRpc(url_, "", json_post, json_response);
    if(ret_type != Rpc::RPC_OK)
    {
        LOG(WARNING) << "coin: "  << coin_  << " url: " << url_;
        LOG(WARNING) << "RPC eth_newFilter fail!";
        return false; 
    }   
    filter_id = json_response["result"].get<std::string>();
    return true;
}

bool  Rpc::eth_getFilterLogs(const std::string& filter_id , json& json_result)
{
    json json_response;
    json json_post;
    json json_params = json::array();
    json_params.push_back(filter_id);
 
    structRpc("eth_getFilterLogs", json_params, json_post);

    Rpc::RpcError ret_type = callRpc(url_, "", json_post, json_response);
    if(ret_type != Rpc::RPC_OK)
    {
        LOG(WARNING) << "coin: "  << coin_  << " url: " << url_;
        LOG(WARNING) << "RPC eth_getFilterLogs fail!";
        return false; 
    }
    
    json_result = json_response["result"];
    return true;
}

bool  Rpc::eth_getBalance(const std::string& account_address, json& json_result,bool is_token /*= false*/, const std::string  contract/* = ""*/)
{
    if (is_token)
    {
        return eth_getTokenBalance(account_address, contract, json_result);
    }
    json json_response;
    json json_post;
    json json_params = json::array();
    json_params.push_back(account_address);
    json_params.push_back("latest");
    structRpc("eth_getBalance", json_params, json_post);

    Rpc::RpcError ret_type = callRpc(url_, "", json_post, json_response);
    if(ret_type != Rpc::RPC_OK)
    {
        LOG(WARNING) << "coin: "  << coin_  << " url: " << url_;
        LOG(WARNING) << "RPC eth_getBalance fail!";
        return false; 
    }
    
    json_result = json_response["result"];
    return true;
}

bool Rpc::eth_getTokenBalance(const std::string& account_address , const std::string contract , json& json_result)
{
    json json_response;
    json json_post;
    json json_params = json::array();
    json json_token ;
    json_token["to"]= contract;
    json_token["data"] ="0x70a08231000000000000000000000000"+ account_address.substr(2,40);
    json_params.push_back(json_token);
    json_params.push_back("latest");
    structRpc("eth_call", json_params, json_post);

    Rpc::RpcError ret_type = callRpc(url_, "", json_post, json_response);
    if(ret_type != Rpc::RPC_OK)
    {
        LOG(WARNING) << "coin: "  << coin_  << " url: " << url_;
        LOG(WARNING) <<"fail:"<< json_post.dump(2);
        LOG(WARNING) << "RPC eth_call fail!";
        return false; 
    }
    
    json_result = json_response["result"];
    return true;
}





