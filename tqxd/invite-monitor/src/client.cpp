#include "common.h"
#include <iostream>
#include "db_mysql.h"
#include "rpc.h"
#include <fstream>
#include <glog/logging.h>
#include <boost/program_options.hpp>
#include <boost/filesystem.hpp>
#include "syncer.h"
#include "tinyformat.h"

#include <gmp.h>
#include <gmpxx.h>

static json s_json_conf;

static bool ParseCmd(int argc,char*argv[])
{
    using namespace boost::program_options;
    std::string conf_file ;

    boost::program_options::options_description opts_desc("All options");
    opts_desc.add_options()
            ("help,h", "help info")
            ("configure,c", value<std::string>(&conf_file)->default_value("../conf/config.json"), "configure file");

    variables_map cmd_param_map;
    try
    {
        store(parse_command_line(argc, argv, opts_desc), cmd_param_map);
    }
    catch(boost::program_options::error_with_no_option_name &ex)
    {
        std::cerr << ex.what() << std::endl;
    }
    notify(cmd_param_map);

    if (cmd_param_map.count("help"))
    {
        std::cout << opts_desc << std::endl;
        return false;
    }
 	std::ifstream jfile(conf_file);

    if (!jfile)
    {
		std::cerr << "No " <<conf_file <<" such config file!\n";
        return false;
    }

    jfile >> s_json_conf;
    if(!s_json_conf.is_object())
    {
		std::cerr << conf_file  <<   "is not json object!\n";
        jfile.close();
        return false;
    }
   
    jfile.close();

    return true;
}

bool GetAddresses(json& json_address_amount)
{
	bool ret = true;

	std::string address_file = s_json_conf["airstake"].get<std::string>();
	std::ifstream jfile(address_file);

    if (!jfile)
    {
        std::cerr << "No " <<address_file <<" such config file!\n";
        return false;
    }

    jfile >> json_address_amount;
   
    jfile.close();
	return ret;
}
void Check5W()
{
	std::string  stake_file = "./airdrop_address_bak.json"; 
	std::ifstream jfile(stake_file);

    if (!jfile)
    {
        std::cerr << "No " <<stake_file <<" such config file!\n";
        return ;
    }
	json json_address_stake;
    jfile >> json_address_stake;
    jfile.close();
	std::map<std::string, uint64_t> map_address_amount;
	for (size_t i =0; i < json_address_stake.size(); i++)
	{
		map_address_amount[json_address_stake[i][0].get<std::string>()] = json_address_stake[i][1].get<uint64_t>();
	}

	std::string lvxin_file = "./airdrop_lvxin.json";
	jfile.open(lvxin_file);
	if (!jfile)
    {
        std::cerr << "No " <<stake_file <<" such config file!\n";
        return ;
    }
	json json_address_lvxin;
	jfile >> json_address_lvxin;
	std::map<std::string, std::string> map_address_lvxin;
	for (size_t i =0; i < json_address_lvxin["RECORDS"].size(); i++)
	{
		json json_value  = json_address_lvxin["RECORDS"][i];
		map_address_lvxin[json_value["to_address"].get<std::string>()] = json_value["airdrop_amount"].get<std::string>();
	}

	std::map<std::string, uint64_t>::iterator iter = map_address_amount.begin();
	uint64_t count =0;
	while (iter != map_address_amount.end())
	{
		if (map_address_lvxin.find(iter->first) == map_address_lvxin.end())
		{
			std::cout <<  iter->first << std::endl;
			std::cout << iter->second << std::endl;
			count += iter->second;
		}
		iter++;
	}
	std::cout << count << std::endl;
}

static bool InitLog(const std::string& log_path)
{
	boost::filesystem::path path_check(log_path);

	if( !boost::filesystem::exists(path_check) )
	{
		boost::filesystem::create_directory(path_check);
	}

	FLAGS_alsologtostderr = false;
	FLAGS_colorlogtostderr = true;
	FLAGS_max_log_size = 100;
	FLAGS_stop_logging_if_full_disk  = true;
	std::string log_exec = "log_exe";
	FLAGS_logbufsecs = 0;
	google::InitGoogleLogging(log_exec.c_str());
	FLAGS_log_dir = log_path;
	std::string log_dest = log_path+"/info_";
	google::SetLogDestination(google::GLOG_INFO,log_dest.c_str());
	log_dest = log_path+"/warn_";

	google::SetLogDestination(google::GLOG_WARNING,log_dest.c_str());
	log_dest = log_path+"/error_";
	google::SetLogDestination(google::GLOG_ERROR,log_dest.c_str());
	log_dest = log_path+"/fatal_";
	google::SetLogDestination(google::GLOG_FATAL,log_dest.c_str());
	google::SetStderrLogging(google::GLOG_ERROR);

	return true;
}
void SetLimit()
{
	//bool CurlPostParams(const CurlParams &params, std::string &response)
	CurlParams params;
	json json_post;
	std::string upper = s_json_conf["upper"].get<std::string>();
	std::string lower = s_json_conf["lower"].get<std::string>();

	json_post["upper"] = upper;
	json_post["lower"] = lower;
	params.url = "http://192.168.2.34:8080/invite_stake_airdrop/set_limit";
	params.auth = "";
	params.need_auth = false;
	params.data = json_post.dump();
	std::string response;
	json json_response;
	std::cout << json_post.dump() << std::endl;
	bool ret = CurlPostParams(params, response);
	std::cout << response << std::endl;
	json_response = json::parse(response);
}


std::string AitdAmount(const  std::string& aitd_amount)
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
void BroadcastTx(const std::string& address, const std::string& amount)
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

void SendTx(const std::map<std::string,std::string>& map_address_amount)
{
	std::map<std::string, std::string>::const_iterator iter = map_address_amount.begin();
	while (iter != map_address_amount.end())
	{
		std::string address = iter->first;
		std::string amount = iter->second;
		BroadcastTx(address, amount);
		iter ++;
	}
}

void AirdropArc()
{
	std::string contract = s_json_conf["contract"].get<std::string>();
	json json_address = s_json_conf["airdrop_address"];
	// SELECT to_address, amount, height FROM `transaction` WHERE contract = '0x4d0fc02ad51c89e2ee1bb903089624e84c96749d' AND FROM_address IN ('0x0000000000000000000000000000000000000000','0x5d5fcd79a7b1fa2a1e590caf9fb4bb7650241695') ORDER BY height;
	std::string sql = strprintf("SELECT to_address, amount, height FROM `transaction` WHERE contract = '%s' AND FROM_address IN ('%s','%s','%s','%s','%s') ORDER BY height;",
					contract, json_address[0].get<std::string>(), json_address[1].get<std::string>(), json_address[2].get<std::string>(),
					json_address[3].get<std::string>(), json_address[4].get<std::string>());
	std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;
    map_col_type[1] = DBMysql::STRING;
	map_col_type[2] = DBMysql::INT;
	json json_data;
    g_db_mysql->getData(sql, map_col_type, json_data);
	mpz_class ex_rate = s_json_conf["rate"].get<uint64_t>();
	std::map<std::string, std::string> map_address_amount;

	for (size_t i = 0; i < json_data.size(); i++)
	{
		std::string address = json_data[i][0].get<std::string>();
		mpz_class usdt_amount;
		usdt_amount.set_str(json_data[i][1].get<std::string>(), 10);
		mpz_class aitd_amount ;
		aitd_amount = usdt_amount * 1000000000000 / ex_rate;
		map_address_amount[address] = AitdAmount(aitd_amount.get_str());
		LOG(INFO) << json_data[i].dump();
	}
	SendTx(map_address_amount);
	std::cout << json_data.dump(2) << std::endl;
}


void checkShip(const json& json_result)
{
    std::string deposit = "0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c";
    std::string ship  = "0x8a2147350b859703c1a5e651e6b18e1873870808c8b7e1ed450d352b3f1fbd86";
	std::string airdrop = "0x39a7d1d43def1f7fcd2dbfbbf5d2ffcfdcffdece3776dad1ac48a1c9652ae936";
    std::map<std::string, std::string> map_low_up;
    std::map<std::string, std::vector<std::string> > map_up_low;
    std::map<std::string, mpz_class> map_address_deposit;
    std::map<std::string, std::string> map_address_airdrop;
	mpz_class all_total_deposit;
    for(size_t i =0; i < json_result.size(); i++)
    {
        json json_data = json_result[i];
        std::string topic0 = json_data["topics"][0].get<std::string>();
        if (topic0 == deposit)
        {
            std::string data = json_data["data"].get<std::string>();
            std::string address = "0x" + json_data["topics"][1].get<std::string>().substr(26);
            mpz_class amount;
            amount.set_str(data.substr(2), 16);
			if(map_address_deposit.find(address) != map_address_deposit.end())
			{
				continue;
			}
            map_address_deposit[address] = amount;//amount.get_str(10);
			all_total_deposit += amount;

        }
        else if (topic0 == ship)
        {

            std::string upper = "0x" + json_data["topics"][1].get<std::string>().substr(26);
            std::string lower = "0x" + json_data["topics"][2].get<std::string>().substr(26);
			if (lower == "0xb88e097f2003123abf7106a0051437f11e4840a4")
			{
				std::cout << "ship" << std::endl;
				std::cout << json_data.dump(2) << std::endl;
			}
			map_low_up[lower] = upper;
			map_up_low[upper].push_back(lower);

     	}
		else if (topic0 == airdrop)
		{
			std::string data = json_data["data"].get<std::string>();
			std::string address = "0x" + json_data["topics"][2].get<std::string>().substr(26);
			mpz_class amount;
			amount.set_str(data.substr(2), 16);
			map_address_airdrop[address] = amount.get_str(10);
			if (address == "0xb88e097f2003123abf7106a0051437f11e4840a4")
			{
				std::cout << "airdrop" << std::endl;
				std::cout << json_data.dump(2) << std::endl;
			}
		}
        else
        {
            //std::cout << topic0 << std::endl;
        }
    }



	std::cout << all_total_deposit.get_str(10) << std::endl;
	std::vector<std::string> vect_lower;
	std::map<std::string, mpz_class>::iterator iter_lower = map_address_deposit.begin();
	std::map<std::string, mpz_class> map_upper_back;

	std::map<std::string, mpz_class> map_upper_back_bak;

	std::vector<std::string> vect_sql;
	mpz_class ceshi;
	while( iter_lower != map_address_deposit.end())
	{	
		if (iter_lower->first == "0x1fbebea698272bd7debf4b6d1864b65a2f4509e6")
		{
			if (map_address_airdrop.find(iter_lower->first) != map_address_airdrop.end() &&
		    map_low_up.find(iter_lower->first)  != map_low_up.end())
			{
				std::cout << "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww" << std::endl;
				std::cout << iter_lower->second.get_str() << std::endl;
			}
			else
			{
				std::cout << "-----------------------------------------" << std::endl;
			}
		}
		
		if (map_address_airdrop.find(iter_lower->first) != map_address_airdrop.end() &&
		    map_low_up.find(iter_lower->first)  != map_low_up.end())
		{
			vect_lower.push_back(iter_lower->first);
			
			{
				map_upper_back[map_low_up[iter_lower->first]]  +=  iter_lower->second ;
				ceshi += iter_lower->second;

				if (map_upper_back_bak.find(iter_lower->first) != map_upper_back_bak.end())
				{			
					map_upper_back_bak[map_low_up[iter_lower->first]]  =  map_upper_back_bak[map_low_up[iter_lower->first]]  + iter_lower->second  ;
				}
				else
				{
					map_upper_back_bak[map_low_up[iter_lower->first]] =  iter_lower->second ;
				}
			}
			
			std::string sql =  strprintf("INSERT INTO `airdrop_status` (`address`, `status`) VALUES ('%s', '0');",
                                iter_lower->first);
			vect_sql.push_back(sql);
		}
		iter_lower++;
	}
	std::map<std::string, mpz_class>::iterator iter_up = map_upper_back.begin();
	mpz_class total;
	total = 0;
	std::map<std::string, std::string> map_address_sendvalue;
	size_t temp_count =0;
	mpz_class sub_sum ;
	while( iter_up != map_upper_back.end())
	{	
		total += iter_up->second ;

		if (iter_up->second > map_upper_back_bak[iter_up->first])
		{
			if (iter_up->first == "0x95b30b49bc94d206356e30c053b4d12b313650ad")
			{
				std::cout << "address "  << iter_up->first << std::endl;
				std::cout << iter_up->second.get_str() << std::endl;
				std::cout << map_upper_back_bak[iter_up->first].get_str() << std::endl;
			}
			
			mpz_class tmp = iter_up->second - map_upper_back_bak[iter_up->first];
			mpz_class sub = tmp*15/100;
			sub_sum += sub;

			//std::cout << sub.get_str()  << std::endl;
			map_address_sendvalue[iter_up->first] = AitdAmount(sub.get_str());
			//std::cout << map_address_sendvalue[iter_up->first] << std::endl;
			temp_count ++;
		}
		iter_up++;
	
	}

	std::cout << "total: " << sub_sum.get_str(10) << std::endl;
	std::cout << "send size: " <<temp_count << std::endl;
 	//SendTx(map_address_sendvalue);

}


void FilterInvter()
{
    Rpc rpc;
	//http://113.31.105.131:18545
    rpc.setNode("AITD", "http://127.0.0.1:8545");
	//rpc.setNode("AITD", "http://172.63.1.44:8545");
    uint64_t from_block = 1540000;//1583123;// 1540000;
    uint64_t to_block =  1583123;//1581036

    //rpc.eth_blockNumber(to_block);
	std::cout << to_block << std::endl;
    std::string address = "0x9b48A05Bf2671Ee8Ea4927482b7c60c145dFc55c";
    std::string filter_id = "0xd2b4f2e591b9ff466eefa6657bb7814e" ;
    bool ret =  rpc.eth_newFilter(from_block, to_block, address, filter_id);
    json json_result;

    rpc.eth_getFilterLogs(filter_id, json_result);

    LOG(INFO) << json_result.dump(2);
	checkShip(json_result);
}

void AirdropStake()
{
	json json_address_amount;
	bool ret  = GetAddresses(json_address_amount);
	uint64_t total = 0;
	std::map<std::string, std::string> map_address_amount;
    for (size_t i = 0; i < json_address_amount.size(); i++)
    {
        std::string address = json_address_amount[i][0].get<std::string>();
        uint64_t amount = json_address_amount[i][1].get<uint64_t>();
		total += amount;
		map_address_amount[address] = std::to_string(amount);   
    }
    SendTx(map_address_amount);
    std::cout << total << std::endl;
}

static bool OpenDB()
{
	json json_connect = s_json_conf["mysql"];

	if (!g_db_mysql->openDB(json_connect))
	{
		std::cerr << "open db fail!" << std::endl;
		return false;
	}

	return true;
}


void CheckTxid(const std::string& txid)
{
	CurlParams params;
	json json_post;
	json_post["hash"] = txid;
	params.url = "http://192.168.2.34:8080/invite_stake_airdrop/get_tx_status";
	params.auth = "";
	params.need_auth = false;
	params.data = json_post.dump();
	std::string response;
	json json_response;
	std::cout << json_post.dump() << std::endl;
	bool ret = CurlPostParams(params, response);
	int status = 0;
	std::cout << response << std::endl;
	if (ret)
	{
		json_response = json::parse(response);
		status = json_response["status"].get<int>();
	}
	std::cout << response << std::endl;
}

void CheckResult()
{
	DBMysql db_airdrop;
	json json_connect = s_json_conf["airdrop"];
	if (!db_airdrop.openDB(json_connect))
	{
		std::cerr  << json_connect.dump() <<std::endl;
		return ;
	}
	std::string sql = strprintf("SELECT hash FROM airdrop WHERE tx_status = 0 AND airdrop_status =1;");
	std::map<int, DBMysql::DataType> map_col_type;
    map_col_type[0] = DBMysql::STRING;

	json json_data;
    db_airdrop.getData(sql, map_col_type, json_data);
	db_airdrop.closeDB();
	for (size_t i=0; i < json_data.size(); i++)
	{
		std::string txid = json_data[i][0].get<std::string>();
		CheckTxid(txid);
	}
}

int main (int argc,char*argv[])
{
	
	assert(ParseCmd(argc,argv));
	std::string log_path = s_json_conf["logpath"].get<std::string>();
	assert(InitLog(log_path));
	assert(OpenDB());
	bool check = s_json_conf["check"].get<bool>();
	if (!check)
	{
		//Check5W();
		//SetLimit();
		//AirdropStake();
		FilterInvter();
		//AirdropArc();
	}
	else
	{
		CheckResult();
	}
	
	
	bool back_run = s_json_conf["daemon"].get<bool>();
	if (back_run)
	{
		fprintf(stdout, "Syncer server starting\n");

		// Daemonize
		if (daemon(1, 0)) 
		{ // don't chdir (1), do close FDs (0)
			fprintf(stderr, "Error: daemon() failed: %s\n", strerror(errno));
			return 0;
		}
	}
	return 0;
}
