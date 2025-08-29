#ifndef _BPLIST_TYPES_HPP_
#define _BPLIST_TYPES_HPP_

#include<iostream>
#include<stdint.h>
#include<vector>


const int32_t MAX_BP_NUM = 21; 
const std::string SCRIPT_PATH = "./test.sh";
const int32_t VALID_TIME = 121;
const uint32_t MAX_BUFFER = 10 * 1024;
const std::string g_proc_name = "bpforward";
//std::string g_prod_name = "";

//const std::string MAIN_RPC_1 = " http://10.186.11.112:8000 ";
//const std::string SECONDARY_RPC_1 = " http://10.186.11.112:8000 ";



struct producer_key {

    std::string     producer_name;
    std::string    block_signing_key;

    friend bool operator < ( const producer_key& a, const producer_key& b ) {
        return a.producer_name < b.producer_name;
    }
};


struct share_bplist {
    uint64_t time;
    std::string hash_bplist;
    std::vector< producer_key > bplist;
};

struct global_para {
    std::string producer_name;
    std::vector<std::string> rpc_primary_chain;
    std::vector<std::string> rpc_second_chain;
};













#endif