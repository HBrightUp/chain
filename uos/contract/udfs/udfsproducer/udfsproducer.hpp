#include "config.hpp"
#include <cstring>
#include <string>
#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/print.hpp>

namespace udfsuossys {
	using namespace uosio;

	CONTRACT udfsproducer : public contract {
	public:
		using contract::contract;
		udfsproducer(name receiver, name code, datastream<const char*> ds);

		ACTION init();

		ACTION loginnode(name owner, std::string strip, uint8_t level);	

		ACTION updatenode(name owner, int8_t status, bool del);	

		//ACTION setip(std::string strip, std::string strnewip);

		ACTION logoutnode(name owner);

		ACTION unbindnode(name owner);

		ACTION setbonus(name owner, uint64_t phase, uint8_t level, uint64_t score, asset award);

		ACTION clearbonus(uint64_t phase, uint32_t num);

		ACTION paybonus(name owner, uint64_t phase);

		ACTION punishnode(name owner);

		ACTION setperiod(uint64_t phase);

		ACTION updatekey(name key, uint64_t value, bool del);					

		//ACTION clearreward(uint64_t phase);

		//ACTION cleannodes();						

		ACTION clear(uint32_t num);				

		void transfer(name from, name to, asset quantity, std::string memo);

		void nodevoter(uint64_t owner);

		void apply(uint64_t receiver, uint64_t code, uint64_t action);


	private:
		//global variable
 		TABLE udfs_global {		
   			name		key;
   			uint64_t	val;
			
   			uint64_t    primary_key() const { return key.value; }
   			UOSLIB_SERIALIZE(udfs_global, (key)(val));
  		};

		typedef multi_index<"udfsglobal"_n, udfs_global>	udfsglobal_t;
		udfsglobal_t _udfsglobal;

		//reward of every period
		TABLE udfs_reward{
			uint64_t 	phase;			//如20190101
			uint64_t 	nodes;
			//std::string	rate;			//UOS : UT，查询联盟合约获取    //10.17
			uint64_t 	score;
			//uint64_t 	store;		//10.17
			//uint64_t 	in;			//10.17
			//uint64_t 	out;		//10.17
			uint8_t	 	status = 0;		//0可插入记录， 1 调用了clearbonus， 2结算中
			asset	 	bill = ZERO_ASSET;
			asset	 	profit = ZERO_ASSET;	
			asset	 	supply = ZERO_ASSET;	//10.17
			
			uint64_t primary_key() const { return phase; }
			UOSLIB_SERIALIZE(udfs_reward, (phase)(nodes)(score)(status)(bill)(profit));
		};

		typedef multi_index<"udfsreward"_n, udfs_reward>	udfsreward_t;
		udfsreward_t _udfsreward;

		//node variable
		TABLE udfs_node {
			name		owner;		   // account name,primary key
			uint32_t	ip;			   // ip
        		uint64_t    	amount;        // asset of spend for udfs
	        	uint64_t	coin_type;	   // source coin	type:  1-->UT; 2-->ETH
        		//uint64_t	apply_at;
			uint64_t	login_at;	   // timestamp of buy node
			uint8_t		level = 0;		   // node type: 1 super，2 supreme
			int8_t		status = 0;		   // node status，0 unbonuded，1 unregistered，2 working，-1 to -4 exception
			asset		award = ZERO_ASSET;		//award already got
			uint8_t		punish = 0;		//punish level
			uint64_t	score = 0;		//punish score(if score equal zero, punish level change to zero)
			//uint8_t		exp_phase = 0;  //punish periods continue
			
			uint64_t	primary_key() const { return owner.value; }
			uint64_t	get_ip() const { return (uint64_t)ip; }
			uint64_t	get_level() const { return (uint64_t)level; }
			uint64_t	get_login_at() const { return login_at; }
			UOSLIB_SERIALIZE(udfs_node, (owner)(ip)(amount)(coin_type)(login_at)(level)(status)(award)(punish)(score)/*(exp_phase)*/);
		};

		typedef multi_index<"udfsnodes"_n, udfs_node,
			indexed_by<"byip"_n, const_mem_fun<udfs_node, uint64_t, &udfs_node::get_ip> >,
			indexed_by<"bylevel"_n, const_mem_fun<udfs_node, uint64_t, &udfs_node::get_level> >,
			indexed_by<"byloginat"_n, const_mem_fun<udfs_node, uint64_t, &udfs_node::get_login_at> >
			>	udfsnode_t;
		udfsnode_t _udfsnode;

		//list of node bonus
		TABLE udfs_bonus {		//scope
			name	 owner;		//account name, primary key
			uint32_t ip;		//ip
			uint8_t  level;		//node type
			uint64_t score;		//score of the phase
			//uint64_t store;		//store of the phase//10.17
			//uint64_t in;		//in flow//10.17
			//uint64_t out;		//out flow//10.17
			asset	 award = ZERO_ASSET;		//reward of the phase
			
			uint64_t primary_key() const { return owner.value; }
			uint64_t get_ip() const { return (uint32_t)ip; }
			uint64_t get_score() const { return score; }
			uint64_t get_award() const { return award.amount; }
			UOSLIB_SERIALIZE(udfs_bonus, (owner)(ip)(level)(score)(award));	
			
		};
		typedef multi_index<"udfsbonus"_n, udfs_bonus,
			indexed_by<"byscore"_n, const_mem_fun<udfs_bonus, uint64_t, &udfs_bonus::get_ip> >,
			indexed_by<"byaward"_n, const_mem_fun<udfs_bonus, uint64_t, &udfs_bonus::get_score> >,
			indexed_by<"byaccount"_n, const_mem_fun<udfs_bonus, uint64_t, &udfs_bonus::get_award> >
			>  udfsbonus_t;

		struct st_transfer {
			name		from;
			name		to;
			asset		quantity;
			std::string	memo;
		};

		struct st_nodevoter {
			uint64_t voter;
			std::string tr_id;
			uint64_t owner;
			int64_t amount;
			std::string coin_name;
		};

		struct st_addnode{
			uint64_t tr_hash;
			uint64_t owner;
		};

		//udfs节点资质表
    	TABLE udfs_token {
        	uint64_t    tr_hash;       // primary key
        	uint64_t	owner;         // account of apply for udfs
        	uint64_t    amount;        // asset of spend for udfs
        	uint64_t    apply_time;    // time of apply for 
        	uint64_t	coin_type;	   //source coin	type:  1-->UT; 2-->ETH
        	uint64_t    reserve1;      // reserve
        	uint64_t    reserve2;      // reserve

        	uint64_t	primary_key() const { return tr_hash; }
			uint64_t	by_owner() const { return owner; }
        	uint64_t    by_apply_time() const { return apply_time; }
        	UOSLIB_SERIALIZE(udfs_token, (tr_hash)(owner)(amount)(apply_time)(coin_type)(reserve1)(reserve2));
    	};

    	typedef uosio::multi_index< "udfsnode"_n, udfs_token,  
                        indexed_by< "owner"_n, const_mem_fun<udfs_token, uint64_t, &udfs_token::by_owner> >,
                        indexed_by< "applytime"_n, const_mem_fun<udfs_token, uint64_t, &udfs_token::by_apply_time> >
                    >   udfstoken_t;

		void setKey(name key, uint64_t value, bool del = false);

		void sendInline(name to, asset quantity, std::string memo);
		
		void addValue(name key, uint64_t plus, bool negative = false);

		std::string intToIp(uint32_t num);

		uint32_t ipToInt(std::string strIp);

		bool isPublic(uint32_t num);

		//void checkFreeze();

		std::string phaseToStr(uint64_t phase);		

	};

}
