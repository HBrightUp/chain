#ifndef SYNCER_H_
#define SYNCER_H_

#include "job/job.h"
#include "job/task.h"
#include "rpc.h"

class Syncer:public Task
{
public:
    Syncer()
    {
    }

    virtual ~Syncer()
    {

    }

    static Syncer& instance()
    {
        return single_;
    }

	void init();
   	
    void refreshDB();

    void scanCoin(const std::string& coin);
    void scanBalance(const std::string& coin);
    bool getFromData(json &json_data);

    void airdropInvite();

    void walletToClickhouse();

	void initClickHouse();
protected:

    bool resetToken(const std::string& coin);

	bool resetContract();

    bool resetEvent();
	
	bool getNodeInfo(const std::string &coin, NodeInfo &node);

	bool getAddressHeight(const std::string& coin, json& json_data);
    
    bool getAddress(const std::string&coin,  std::vector<std::string>&vect_address);
    
public:

    void registerTask(map_event_t& name_events, map_job_t& name_tasks);

    void setRpc(const Rpc& rpc)
	{
		rpc_ = rpc; 
	}

    void setClickHouseConnect(const json& json_connect)
    {
	    json_connect_clickhouse_ = json_connect;
    }

    void setWalletConnect(const json& json_connect)
    {
	    json_connect_wallet_ = json_connect;
    }
protected:
    static Syncer single_;
	Rpc rpc_;
	std::vector<std::string> vect_sql_;
    json json_connect_wallet_;
    json json_connect_clickhouse_;
    std::map<std::string, std::string> contract_token_;
    std::map<std::string, std::string> hash_event_;
    std::map<std::string, std::string> address_contract_;
};

#endif
