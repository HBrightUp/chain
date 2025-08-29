#include<iostream>
#include<cstdio>
#include <fcntl.h>
#include<unistd.h>
#include<vector>
#include<memory>
#include<boost/process.hpp>
#include<json/json.h>
#include<./include/types.hpp>
#include<openssl/sha.h>
#include <boost/program_options/parsers.hpp>
#include <boost/program_options/variables_map.hpp>
#include <boost/program_options/options_description.hpp>


namespace bp = ::boost::process;
namespace bpo = boost::program_options;
global_para global;

bool single_instance(const std::string& proc_name);
static int lockfile(int fd);

bool get_parameter(int argc, char* argv[]);
void read_table( std::string filestr, std::vector<std::string>& bplist );

std::string& trim( std::string& str ); 
std::string sha256(const std::string str);

std::string set_bplist_hash(share_bplist& share);

bool excute(const std::string& cmd ,std::string& result);
bool is_success(const std::string& result);

bool getbplist(std::string& result);
bool setbplist(const share_bplist& share );

std::string get_cmd_getbplist(const std::string& rpc);
std::string get_cmd_setbplist( const share_bplist& share, const std::string& rpc );


int main(int argc, char* argv[]) {

	if(!get_parameter(argc,argv)){
		return -1;
	}

#if 0
	if(!single_instance(g_proc_name)) {
		return -1;
	}
#endif

	std::vector<std::string> vec_bplist;
	std::string cmd, result;
	share_bplist share;


	while(true) {

		sleep( VALID_TIME - 1);
		vec_bplist.clear();
		cmd.clear();
		result.clear();
		share.bplist.clear();
		
		//read_table(SCRIPT_PATH, vec_bplist);

		std::cout << "bplist size: " << vec_bplist.size() << std::endl;

		std::string str_bplist;

		if(!getbplist(str_bplist)) {
			continue;
		}

		str_bplist = trim(str_bplist);

		Json::Reader rd;
		Json::Value json_bplist;

		if(!rd.parse(str_bplist, json_bplist)) {
			std::cout << "parse bplist json text failied!" << std::endl;
			continue;
		}

		try {
			int  isize = json_bplist["rows"].size();

			if(isize < 1) {
				continue;
			}

			uint64_t lastest_bp_valid_time = json_bplist["rows"][isize - 1]["bp_valid_time"].asUInt();

			uint64_t bp_valid_time = 0;
			producer_key bpinfo;

			int start_pos = 0;

			for(int i = isize - 1; i >= 0; --i) {
				bp_valid_time = json_bplist["rows"][i]["bp_valid_time"].asUInt();

				//std::cout << "lastest_bp_valid_time: " << lastest_bp_valid_time << " bp_valid_time: " << bp_valid_time << std::endl;
				if( lastest_bp_valid_time - bp_valid_time >= VALID_TIME) {
					break;
				}

				bpinfo.producer_name = json_bplist["rows"][i]["bpname"].asString();
				bpinfo.block_signing_key = json_bplist["rows"][i]["ulord_addr"].asString();
				share.bplist.emplace_back(bpinfo);
				start_pos = i;
			}
			
			//3*120: in ulordchain, this bp_valid_time is ahead of current time. 
			share.time = json_bplist["rows"][isize - 1]["bp_valid_time"].asUInt() - 3 * 120;

		}catch(...) {
			std::cout << "parse json file failed!" << std::endl;
			continue;
		}
		
		set_bplist_hash(share);
		
		if(!setbplist(share)) {
			continue;
		}
	}
	
    return 0;
}

 bool single_instance(const std::string& proc_name) {
	 int  fd;

	std::cout << "aaaa" << std::endl;
	std::cout << "proc_name: " << proc_name << std::endl;

	if(proc_name.length() >= 100) {
		std::cout << "process name too long" << std::endl;
		return false;
	}
	std::cout << "bbbb" << std::endl;
	std::string name = "~/" + proc_name + ".pid";

	std::cout << "name: " << name << std::endl;

    fd = open(name.c_str(), O_RDWR | O_CREAT, (S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH));
    if (fd < 0) {
		std::cout << "open file  failed! name: " << name << std::endl;
        return false;
    }
	std::cout << "cccc" << std::endl;
    if (lockfile(fd) == -1) {                                                 
		std::cout << "file locked! name: " << name << std::endl;
        close(fd);
        return false;
    } else {
		std::cout << "dddd" << std::endl;
        ftruncate(fd, 0);                                                      
		std::string pid = std::to_string(getpid());
        write(fd, pid.c_str(), pid.length());
        return true;
    }

	return true;
 }

 static int lockfile(int fd)
{
    struct flock fl;

    fl.l_type   = F_WRLCK;
    fl.l_start  = 0;
    fl.l_whence = SEEK_SET;
    fl.l_len    = 0;

    return(fcntl(fd, F_SETLK, &fl));
}

bool get_parameter(int argc, char* argv[]) {
	bpo::options_description opts("options");
	opts.add_options()
		("help", "help info")
		("p", bpo::value<std::string>(), "producer name of current node")
		("primary-url", bpo::value<std::vector<std::string>>(), "rpc address of primary chain")
		("second-url", bpo::value<std::vector<std::string>>(), "rpc address of second chain");
		
	bpo::variables_map vm;
	try {
		bpo::store(parse_command_line(argc, argv, opts), vm);
	}
	catch (bpo::error_with_no_option_name &ex) {
		std::cout << ex.what() << std::endl;
	}

	bpo::notify(vm);

	if (vm.count("help"))
	{
		std::cout << opts << std::endl;
		return false;
	}

	if (vm.count("p"))
	{
		global.producer_name = vm["p"].as<std::string>();
	}

	if (vm.count("primary-url"))
	{
		global.rpc_primary_chain = vm.at("primary-url").as<std::vector<std::string>>();
	}

	if (vm.count("second-url"))
	{
		global.rpc_second_chain = vm.at("second-url").as<std::vector<std::string>>();
	}

	bool init = false;

	if(global.producer_name.length() > 0 &&
	global.rpc_primary_chain.size() > 0  &&
	global.rpc_second_chain.size() > 0) {
		init = true;
	}

	if(global.producer_name.empty()) {
		std::cout << "please input producer name of current node." << std::endl;
	}

	if(global.rpc_primary_chain.empty()) {
		
		std::cout << "please input rpc address of primary chain." << std::endl;
	}

	if(global.rpc_second_chain.empty()) {
		std::cout << "please input rpc address of second chain." << std::endl;
	}
	
	return init;
}

bool getbplist(std::string& result) {
	std::string cmd;
	bool is_ok = false;
	
	for( auto rpc = global.rpc_primary_chain.begin(); rpc != global.rpc_primary_chain.end(); ++rpc ) {
		
		cmd = get_cmd_getbplist(*rpc);
		if(!excute(cmd, result) ){
			continue;
		}

		if(!is_success(result)) {
		
			continue;
		} else {
			is_ok = true;
			break;
		}
	}
	
	return is_ok;
}

std::string get_cmd_getbplist(const std::string& rpc) {
	return "cluos -u " + rpc + " get table uosio uosclist uosclist -l 441";
}
#if 0
void read_table(std::string filestr, std::vector<std::string>& vec_bplist)
{ 
    std::string line;

    bp::ipstream is; 
    bp::child c(bp::search_path("sh"), filestr,   bp::std_out > is);
	
    
    while (c.running() && std::getline(is, line) && !line.empty())
    {
        vec_bplist.push_back(line);
    }
    c.wait();

    return  ;

} 

#endif

std::string set_bplist_hash( share_bplist& share) {
	std::string hash;
	hash = std::to_string(share.time);

	for(const auto& bp : share.bplist) {
		//hash += bp.producer_name + bp.block_signing_key;
		hash += bp.producer_name;
	}
	share.hash_bplist = sha256(hash);

	return share.hash_bplist;
}

std::string& trim( std::string& str ) 
{
    if (str.empty()) 
    {
        return str;
    }
    
	size_t pos = 0;
	pos = str.find(" ", pos);

	while( pos != std::string::npos) {
		str.replace( pos, 1, "" );
		pos = str.find( " ", pos);
	}

    return str;
}

std::string sha256(const std::string str)
{
	unsigned char hash[SHA256_DIGEST_LENGTH];
	SHA256_CTX context;
	char md[65];
	memset(md, 0, sizeof(md));

	std::string result;

    if(!SHA256_Init(&context))
        return result;

    if(!SHA256_Update(&context, (unsigned char*)str.c_str(), str.length()))
        return result;

    if(!SHA256_Final(hash, &context))
        return result;

	for(int i = 0; i < SHA256_DIGEST_LENGTH; ++i)
    {
        sprintf(md + (i * 2), "%02x", hash[i]);
    }
	
	result.reserve(65);
	for(int i = 0; i < 64; ++i) {
		std::cout << md[i];
		result.push_back(md[i]);
	}

	std::cout << std::endl;

	//if use as follows, memory error occour, why? 
	// so i have to assign data one by one^^^
	//result = md;

	//std::cout << "sha256() result:" << result << std::endl;

	return result;
}

std::string get_cmd_setbplist( const share_bplist& share, const std::string& rpc )
{
	std::string cmd,strTemp;

#if 1
	cmd = "cluos -u " + rpc + " push action uosio setbplist \'{\"bp_name\":\"";
	cmd += global.producer_name + "\",\"bp_time\":\"";
	cmd += std::to_string(share.time) + "\", \"hash_bplist\":\"" + share.hash_bplist + "\", \"bplist\":[";

	for(const auto& bp : share.bplist) {
		cmd += "{\"producer_name\":\"" + bp.producer_name + "\", \"block_signing_key\":\"" + bp.block_signing_key + "\"},";
	}

	cmd.erase(cmd.end() - 1);
	cmd += "]}\' -x 2000  -p " + global.producer_name + "@active";
	
#endif
	//cmd = "env";
#if 0
	std::string exp = "cluos -u http://10.186.11.112:8000 push action uosio setbplist \'{\"bp_name\":\"marsaccount3\",\"bp_time\":\"1553582165\", \"hash_bplist\":\"b7f0d42e2e0555bacbc9a7672aa378727fbd564cbe71f37a4ae1447bb629bab3\", \"bplist\":[{\"producer_name\":\"marsaccount3\", \"block_signing_key\":\"UOS6aWfdf6tHWCepZpP6MzdynuNMAkNKr6nbNMguTuCatq88LyG4G\"},{\"producer_name\":\"dragonexsafe\", \"block_signing_key\":\"UOS7DVNg9bsq1zUZtUbWwA1UyhaqFBmNrCRVnPFeGYgsx7kzfAtiC\"},{\"producer_name\":\"uosvegasjack\", \"block_signing_key\":\"UOS5aTdkbRaPH5WKYPJaxX8HfqGtX8hKx8p6FDPAzpkyJiYnuBu5c\"}]}\' -x 2000  -p marsaccount3@active";
	
	if(exp == cmd) {
		std::cout << "Equal+++++++++++++++++++++++++" << std::endl;
	} else {
		std::cout << "exp size: " << exp.size() << std::endl;
		std::cout << "cmd size: " << cmd.size() << std::endl;

		int pos = exp.size() > cmd.size() ? cmd.size() : exp.size();

		for(int i = 0; i < pos; ++i) {

			if(exp.at(i) != cmd.at(i)) {
				std::cout << "i: " << i << std::endl;
				std::cout << "exp: " << exp.at(i) << "cmd: " << cmd.at(i) << std::endl;
				
			}
		}



	}
	
	std::cout << "exp:" << std::endl;
	std::cout << exp << std::endl;
	std::cout << "cmd:" << std::endl;
	std::cout << cmd << std::endl;
#endif
	return cmd;
}

bool setbplist(const share_bplist& share) {

	std::string cmd,result;
	bool is_ok = false;

	for( auto rpc = global.rpc_second_chain.begin(); rpc != global.rpc_second_chain.end(); ++rpc ) {
		cmd = get_cmd_setbplist(share, *rpc);
		if(!excute(cmd, result)) {
			continue;
		}

		if(!is_success(result)) {
			continue;
		} else {
			is_ok = true;
			break;
		}
	}

	return is_ok;
}

bool excute(const std::string& cmd, std::string& result ) {

	std::cout << cmd << std::endl;

	std::shared_ptr<char> buf(new char[MAX_BUFFER]);
	memset(buf.get(), 0, MAX_BUFFER);

	FILE * fp;
	if((fp = popen(cmd.data(), "r")) == NULL) {
		std::cout << "error: " << errno << std::endl;
		return false;
	}
	std::string strTemp;

	while( fgets( buf.get(), MAX_BUFFER, fp) != NULL) {
		strTemp = buf.get();
		result += strTemp;
	
	}
	
	pclose(fp);

	return true;
}

bool is_success(const std::string& result) {

	if(result.length() < 5 ) {
		return false;
	}

	std::string sub = result.substr(0,5);

	if(sub == std::string("Error") || sub == std::string("Faile")) {
		return false;
	}

	return true;
}
