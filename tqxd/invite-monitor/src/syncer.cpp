#include "syncer.h"
#include <glog/logging.h>
#include "db_mysql.h"
#include "tinyformat.h"
#include <gmp.h>
#include <gmpxx.h>

static void SetTimeout(const std::string &name, int second)
{
    struct timeval timeout;
    timeout.tv_sec = second;
    timeout.tv_usec = 0;
    evtimer_add(Job::map_name_event_[name], &timeout);
}

static void ScanAITD(int fd, short kind, void *ctx)
{
    LOG(INFO) << "scan block chain AITD";
    Syncer::instance().scanCoin("AITD");
    SetTimeout("ScanAITD", 1 * 60);
}

static void AirdropInvite(int fd, short kind, void *ctx)
{
    LOG(INFO) << "scan AirdropInvite chain AITD";
    Syncer::instance().airdropInvite();
    SetTimeout("AirdropInvite", 5 * 60);
}

void Syncer::refreshDB()
{
    LOG(INFO) << "refresh DB begin";
    LOG(INFO) << "SQL size: " << vect_sql_.size();
    if (vect_sql_.size() > 0)
    {
        g_db_mysql->batchRefreshDB(vect_sql_);
        vect_sql_.clear();
    }
    LOG(INFO) << "refresh DB end";
}

bool  GetAddressStatus(std::map<std::string, int>& map_address_status)
{
    std::string sql = "SELECT address, status FROM airdrop_status;";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::INT;
    json json_data;
    bool ret = g_db_mysql->getData(sql, map_col_type, json_data);
    if (ret)
    {
        for (size_t i = 0; i < json_data.size();  i++)
        {
            map_address_status[json_data[i][0].get<std::string>()] = json_data[i][1].get<int>();
        }
    }
    return ret;
}

bool  GetInviteShip(std::map<std::string, std::string>& map_low_up)
{
    std::string sql = "SELECT invite_up, invite_low FROM InviteShip;";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
    json json_data;
    bool ret = g_db_mysql->getData(sql, map_col_type, json_data);
    if (ret)
    {
        for (size_t i = 0; i < json_data.size();  i++)
        {
            map_low_up[json_data[i][1].get<std::string>()] = json_data[i][0].get<std::string>();
        }
    }
    return ret;
}

bool  GetDeposit(std::map<std::string, mpz_class>& map_address_amount)
{
    std::string sql = "SELECT address, amount FROM Deposit;";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
    json json_data;
    bool ret = g_db_mysql->getData(sql, map_col_type, json_data);
    if (ret)
    {
        mpz_class amount;
        for (size_t i = 0; i < json_data.size();  i++)
        {
            amount.set_str(json_data[i][1].get<std::string>(), 10);
            if (map_address_amount.find(json_data[i][0].get<std::string>()) !=  map_address_amount.end())
            {
                continue;
            }
            map_address_amount[json_data[i][0].get<std::string>()] = amount;
        }
    }
    return ret;
}


bool GetAirdrop(std::map<std::string, std::string>& map_address_airdrop)
{
    std::string sql = "SELECT recieve, amount FROM AirdropDeposit;";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
    json json_data;
    bool ret = g_db_mysql->getData(sql, map_col_type, json_data);
    if (ret)
    {
        for (size_t i = 0; i < json_data.size();  i++)
        {
            map_address_airdrop[json_data[i][0].get<std::string>()] = json_data[i][1].get<std::string>();
        }
    }
    return ret;
}

bool Syncer::getAddressHeight(const std::string &coin, json &json_data)
{
    std::string sql = "SELECT MAX(to_block), LOWER(address) FROM filterblock WHERE coin= '" + coin + "' GROUP BY LOWER(address);";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::INT;
    map_col_type[1] = DBMysql::STRING;

    g_db_mysql->getData(sql, map_col_type, json_data);
    
    return true;
}

bool GetMaxStatus(int& status)
{
    std::string sql = "SELECT MAX(status) FROM airdrop_status";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::INT;
    json json_data;
    bool ret = g_db_mysql->getData(sql, map_col_type, json_data);
    if (ret)
    {
        status =  json_data[0][0].get<int>();
    }
    
    return ret;
}


std::string SyncerAitdAmount(const  std::string& aitd_amount)
{
	std::string left, right;
	if (aitd_amount.size() > 18)
	{
		left = aitd_amount.substr(0, aitd_amount.size() -18) + ".";
		right = aitd_amount.substr(aitd_amount.size() -18, aitd_amount.size());
	}
	else
	{
		left = "0.";
		size_t i = aitd_amount.size();
		right = "";
		for (; i < 18; i++)
		{
			right += "0";
		}
		right += aitd_amount;
	}

	return left + right;
} 

static std::vector<std::string> s_vect_txid;
void  SyncerBroadcastTx(const std::string& address, const std::string& amount)
{
	//bool CurlPostParams(const CurlParams &params, std::string &response)
	CurlParams params;
	json json_post;
	json_post["address"] = address;
	json_post["value"] = amount;
	params.url = "http://192.168.2.34:8080/invite_stake_airdrop/airdrop";
	params.auth = "";
	params.need_auth = false;
	params.data = json_post.dump();
	std::string response;
	json json_response;
	std::cout << json_post.dump() << std::endl;
	bool ret = CurlPostParams(params, response);
	int status = 0;
	if (ret)
	{
		json_response = json::parse(response);
		status = json_response["status"].get<int>();
	}
	std::cout << response << std::endl;
	if (1 == status)
	{
		s_vect_txid.push_back(json_response["hash"].get<std::string>());
	}	
}

void  SyncerSendTx(const std::map<std::string,std::string>& map_address_amount)
{
	std::map<std::string, std::string>::const_iterator iter = map_address_amount.begin();
	while (iter != map_address_amount.end())
	{
		std::string address = iter->first;
		std::string amount = iter->second;
		SyncerBroadcastTx(address, amount);
		iter ++;
	}
}

void  SendTx(const std::string& upper,  const std::string& lower, const std::string& amount)
{
	//bool CurlPostParams(const CurlParams &params, std::string &response)
	CurlParams params;
	json json_post;
	json_post["address"] = upper;
    json_post["inviteeAddress"] = lower;
	json_post["value"] = amount;
	params.url = "http://192.168.1.11:8080/invite_stake_airdrop/airdrop_rebate";
	params.auth = "";
	params.need_auth = false;
	params.data = json_post.dump();
	std::string response;
	json json_response;
	std::cout << json_post.dump() << std::endl;
	bool ret = CurlPostParams(params, response);
	int status = 0;
	if (ret)
	{
		json_response = json::parse(response);
		status = json_response["status"].get<int>();
	}
    else
    {
        LOG(ERROR) << json_post.dump();
    }	
}


void Syncer::airdropInvite()
{
    LOG(INFO) << "airdrop invite  begin";
    try
    {
        std::map<std::string, int> map_address_status;

        std::map<std::string, std::string> map_low_up;
        std::map<std::string, mpz_class> map_address_deposit;
        std::map<std::string, std::string> map_address_airdrop;
        bool  ret = GetAddressStatus(map_address_status);
        if (!ret)
        {
            LOG(ERROR) << "GetAddressStatus";
            return ;
        }
        ret = GetInviteShip( map_low_up);
        if (!ret)
        {
            LOG(ERROR) << "GetInviteShip";
            return ;
        }
        ret = GetDeposit(map_address_deposit);
        if (!ret)
        {
            LOG(ERROR) << "GetDeposit";
            return ;
        }
        ret = GetAirdrop( map_address_airdrop);
        if (!ret)
        {
            LOG(ERROR) << "GetAirdrop";
            return ;
        }

        int status =0;
        ret = GetMaxStatus(status);
        if (!ret)
        {
            LOG(ERROR) << "GetMaxStatus";
            return ;
        }

        std::map<std::string, mpz_class>::iterator iter_lower = map_address_deposit.begin();
        std::map<std::string, mpz_class> map_upper_back;

        std::vector<std::string> vect_sql;
        struct AirdropData
        {
            std::string lower;
            std::string upper;
            mpz_class amount;
        };
    
        std::vector<AirdropData> vect_airdropdata;
        while( iter_lower != map_address_deposit.end())
        {	
            if (map_address_airdrop.find(iter_lower->first) != map_address_airdrop.end() &&
                map_low_up.find(iter_lower->first)  != map_low_up.end()  && 
                map_address_status.find(iter_lower->first) == map_address_status.end())
            {
                map_upper_back[map_low_up[iter_lower->first]] +=  iter_lower->second;
                AirdropData airdrop_data;
                airdrop_data.lower = iter_lower->first;
                airdrop_data.upper = map_low_up[iter_lower->first];
                airdrop_data.amount = iter_lower->second;
                vect_airdropdata.push_back(airdrop_data);
                std::string sql =  strprintf("INSERT INTO `airdrop_status` (`address`, `status`) VALUES ('%s', '%d');",
                                    iter_lower->first, status +1);
                // vect_sql_.push_back(sql);
            }
            iter_lower++;
        }
        //refreshDB();
        mpz_class total;
        for(size_t i = 0; i < vect_airdropdata.size(); i++)
        {
            AirdropData airdrop_data = vect_airdropdata[i];
            mpz_class tmp =  airdrop_data.amount * 15 /100;
            LOG(INFO) << airdrop_data.upper ;
            LOG(INFO) << airdrop_data.lower ;
            LOG(INFO) << airdrop_data.amount.get_str() ;
            LOG(INFO) << tmp.get_str() ;
            total += tmp;
            std::string balance = SyncerAitdAmount(tmp.get_str());
            std::cout << balance << std::endl;
            //SendTx(airdrop_data.upper, airdrop_data.lower, balance);
        }
        std::cout << vect_airdropdata.size() << std::endl;
        std::cout << total.get_str() << std::endl;
    }
    catch(const std::exception& e)
    {
        LOG(ERROR) << "Syncer::airdropInvite()" <<e.what() << '\n';
    }
    catch(...)
    {
        LOG(ERROR) << "Syncer::airdropInvite()" ;
    }
  
	
}

void Syncer::scanCoin(const std::string &coin)
{
    LOG(INFO) << "Syncer " << coin << " begin";
    try
    {
        resetEvent();
        Rpc rpc;
        rpc.setNode("AITD", "http://127.0.0.1:8545");
        uint64_t from_block = 0, to_block =0;
        
        json json_height;
        bool ret = getAddressHeight("AITD", json_height);
        if(!ret)
        {
            LOG(ERROR) << "getAddressHeight  error";
            throw std::runtime_error("getAddressHeight  error");
        }
        from_block = json_height[0][0].get<uint64_t>() + 1;
        std::string address =  json_height[0][1].get<std::string>();
        ret = rpc.eth_blockNumber(to_block);
        if(!ret)
        {
            LOG(ERROR) << "eth_blockNumber  error";
            throw std::runtime_error("eth_blockNumber  error");
        }
        std::string filter_id;
        ret = rpc.eth_newFilter(from_block, to_block, address, filter_id);
        if(!ret)
        {
            LOG(ERROR) << "eth_newFilter  error";
            throw std::runtime_error("eth_newFilter  error");
        }
        json json_result;
        ret  = rpc.eth_getFilterLogs(filter_id, json_result);
        if(!ret)
        {
            LOG(ERROR) << "eth_getFilterLogs  error";
            throw std::runtime_error("eth_getFilterLogs  error");
        }

        std::string sql = strprintf("INSERT INTO `filterblock` (`coin`, `from_block`, `to_block`, `filter_id`, `address`) VALUES ('AITD', %d, %d, '%s', '%s');"
                                    ,from_block, to_block, filter_id,  address);
        vect_sql_.push_back(sql);
        for (size_t i =0; i < json_result.size(); i++)
        {
            json json_value = json_result[i];
            std::string event_hash = json_value["topics"][0].get<std::string>();
            std::string data = json_value["data"].get<std::string>();
            std::string event_name = "unknow";
            std::string txid = json_value["transactionHash"].get<std::string>();
            if (hash_event_.find(event_hash) != hash_event_.end())
                event_name = hash_event_[event_hash];
            else
                std::cout << event_hash << std::endl;

            mpz_class amount;
            if(event_name == "AirdropDeposit(address,address,uint256)")
            {
                amount.set_str(data.substr(2), 16);
                std::string send = "0x" + json_value["topics"][1].get<std::string>().substr(26);
                std::string recieve = "0x" + json_value["topics"][2].get<std::string>().substr(26);
                std::string balance = amount.get_str(10);
                sql = strprintf("INSERT INTO `AirdropDeposit` (`txid`, `send`, `recieve`, `amount`) VALUES ('%s', '%s', '%s', '%s');",
                                txid, send, recieve, balance);
                vect_sql_.push_back(sql);
                
            }
            else if (event_name == "Deposit(address,uint256)")
            {
                amount.set_str(data.substr(2), 16);
                std::string address = "0x" + json_value["topics"][1].get<std::string>().substr(26);
                std::string balance = amount.get_str(10);
                sql = strprintf("INSERT INTO `Deposit` (`txid`, `address`, `amount`) VALUES ('%s', '%s', '%s');",
                                txid, address, balance);
                vect_sql_.push_back(sql);

            }
            else if (event_name == "InviteShip(address,address)")
            {
                //["0x8a2147350b859703c1a5e651e6b18e1873870808c8b7e1ed450d352b3f1fbd86","0x00000000000000000000000041d01755b9e37e399c4675089dab1d498725e63c","0x000000000000000000000000712e23ae5bd292b32dfb5bf9148e7ae7ed9cbdc6"]
                std::string invite_up = "0x" + json_value["topics"][1].get<std::string>().substr(26);
                std::string invite_low = "0x" + json_value["topics"][2].get<std::string>().substr(26);
                sql = strprintf("INSERT INTO `InviteShip` (`txid`, `invite_up`, `invite_low`) VALUES ('%s', '%s', '%s');",
                                txid, invite_up, invite_low);
                vect_sql_.push_back(sql);

            }
			else if (event_name == "WithdrawProfit(address,uint256)")
			{
				//["0x010d214e8adbe593eebc2e78d29e88f08ddcb363fac75a9ef8c9455ba3c72dcc","0x000000000000000000000000e7d554a2df5581a39f7cdae00b27cac9f4bdf25e"]
				amount.set_str(data.substr(2), 16);
                std::string address = "0x" + json_value["topics"][1].get<std::string>().substr(26);
                std::string balance = amount.get_str(10);
                sql = strprintf("INSERT INTO `WithdrawProfit` (`txid`, `address`, `amount`) VALUES ('%s', '%s', '%s');",
                                txid, address, balance);
                vect_sql_.push_back(sql);
			}
			else if (event_name == "InviterReward(address,uint256)")
			{
				amount.set_str(data.substr(2), 16);
                std::string address = "0x" + json_value["topics"][1].get<std::string>().substr(26);
                std::string balance = amount.get_str(10);
                sql = strprintf("INSERT INTO `InviterReward` (`txid`, `address`, `amount`) VALUES ('%s', '%s', '%s');",
                                txid, address, balance);
                vect_sql_.push_back(sql);
			}
	
/*            sql = strprintf("INSERT INTO `event_data` (`event_name`, `event_hash`, `data`, `txid`, `topics`, `filter_id`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');",
                                    event_name, event_hash, json_value["data"].get<std::string>(), json_value["transactionHash"].get<std::string>(),
                                    json_value["topics"].dump(),  filter_id);*/
            //vect_sql_.push_back(sql);
        }
        refreshDB();
        
        //checkShip(json_result);
		
    }
    catch (const std::exception &e)
    {
        LOG(ERROR) << "Syncer Scan "  << " error : " << e.what();
        std::cerr << e.what() << '\n';
    }
    catch (...)
    {
        LOG(ERROR) << "Syncer Scan "  << " error unknow ";
    }
    refreshDB();
}

bool Syncer::resetContract()
{
    std::string sql = "SELECT distinct NAME, LOWER(addr) FROM event_contract;";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
	json json_data;
    g_db_mysql->getData(sql, map_col_type, json_data);
    for (size_t i = 0; i < json_data.size(); i++)
    {
        address_contract_[json_data[i][1].get<std::string>()] = json_data[i][0].get<std::string>();
    }
    return true;
}

bool Syncer::resetEvent()
{
    std::string sql = "SELECT func, CONCAT('0x' ,HASH) FROM func_keccak256;";
    std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
	json json_data;
    g_db_mysql->getData(sql, map_col_type, json_data);
    for (size_t i = 0; i < json_data.size(); i++)
    {
        hash_event_[json_data[i][1].get<std::string>()] = json_data[i][0].get<std::string>();
    }
    return true;
}


Syncer Syncer::single_;
void Syncer::registerTask(map_event_t & name_events, map_job_t & name_tasks)
{
	 REFLEX_TASK(ScanAITD);
    // REFLEX_TASK(AirdropInvite);
     
    //REFLEX_TASK(ScanETH);
	 //REFLEX_TASK(ScanBalance);
}
