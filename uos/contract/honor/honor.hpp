#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>

namespace honorsys {

    using namespace uosio;

    using uosio::time_point;
    using uosio::asset;
    using uosio::time_point;

    //seconds one day
    const uint64_t seconds_one_day = 24 * 3600; 

    //state of project
    const std::string state_ongoing = "ongoing";
    const std::string state_finished = "finished";
    const std::string state_overtime = "overtime";


    //global variable
    TABLE global_var {
        name key;
        uint64_t val;

        uint64_t        primary_key() const { return key.value; }

        UOSLIB_SERIALIZE(global_var, (key)(val));
    };

    // the define of honor account infomation
    TABLE honor_info {
        name account;               // name of honor account 
        uint8_t vip;                //grade of honor account
        uint64_t credit;            // current credit
        uint64_t last_convert;      // time of last convert credit to uos

        uint64_t primary_key()const { return account.value; }
        uint64_t by_vip()const { return static_cast<uint64_t>(vip); }
        uint64_t by_credit()const { return credit; }

        UOSLIB_SERIALIZE( honor_info, (account)(vip)(credit) (last_convert) )
    };

    //contributor talbe
    TABLE ctribut {
        name contributor;               //  account of transfer asset to contract acccont 
        asset quantity;                //  total asset

        uint64_t primary_key()const { return contributor.value; }

        UOSLIB_SERIALIZE( ctribut, (contributor)(quantity) )
    };

    struct bpcredit {
        name bp;                    // bp name
        uint64_t credit;            // credit given by bp
    };

    //project table 
    TABLE project_info {
        name            project;            // name of project 
        name            owner;              // name of proposer  
        uint64_t        credit;             // average credit of all bp assessed
        uint64_t         update_time;       // time of apply 
        std::string      reason;            // describe about applying award
        std::string       state;             //  state of project(exciting/finished/overtime)
        std::vector<bpcredit>  credit_list;  // credit list of bp assessed

        uint64_t primary_key()const { return project.value; }
        uint64_t by_owner()const { return owner.value; }
        uint64_t by_credit()const { return credit; }

        UOSLIB_SERIALIZE( project_info, (project)(owner)(credit)(update_time)(reason)(state)(credit_list) )
    };

    /***** bp list talbe ******
	1.this talbe come from uosio.system
    ***************************/
    TABLE one_minute_bp_list{
      std::string 			ulord_addr;     
      uint64_t              bp_valid_time = 0;  // bp start valid time,  120 second one record.   
      name                  bpname;    // bp name
	 
      uint64_t primary_key()const { return bp_valid_time; }
      UOSLIB_SERIALIZE( one_minute_bp_list, (ulord_addr)(bp_valid_time)(bpname) )
   };

    typedef uosio::multi_index< "global"_n, global_var > global_var_table;

    typedef uosio::multi_index< "honorinfo"_n, honor_info,
                                indexed_by<"vip"_n, const_mem_fun<honor_info, uint64_t, &honor_info::by_vip> >,
                                indexed_by<"curcredit"_n, const_mem_fun<honor_info, uint64_t, &honor_info::by_credit>  >
                                > honor_info_table;

    //typedef uosio::multi_index< "award"_n, award_credit > award_credit_table;
    typedef uosio::multi_index< "ctribut"_n, ctribut > ctribut_table;

    typedef uosio::multi_index< "projectinfo"_n, project_info,
                                indexed_by<"owner"_n, const_mem_fun<project_info, uint64_t, &project_info::by_owner>  >,
                                indexed_by<"credit"_n, const_mem_fun<project_info, uint64_t, &project_info::by_credit>  >
                                > project_info_table;

    typedef uosio::multi_index< "uosclist"_n, one_minute_bp_list> uosc_bp_table;  // uosc bp list

    CONTRACT honor : public uosio::contract {
        public:
            //
            using contract::contract;
            honor(name receiver, name code,  uosio::datastream<const char*> ds);

            //contract must be inited at first
            ACTION init();

            //bp give a credit to project
            ACTION assess(const name bp, const name project, uint64_t credit);

            //remove a project from table by manual
            ACTION rmproject(const name project);

            //honor account create a project
            ACTION apply(name project, name proposer, std::string reason);

            //honor get the award of project
            ACTION reward(const name project);

            // honor account upgrade
            ACTION upgrade(const name& user, uint64_t credit);

            // honor account convert convert credit to UOS
            ACTION convertuos(const name& user, uint64_t credit);

            //update state of all project
            ACTION update();

            ACTION setdata(const name& key, const uint64_t& val, const uint8_t& op);

            //notification of transfer
            ACTION transfer(const name& from, const name& to, const asset& quantity, const std::string& memo);

        private:
            global_var_table         _global_var;          
            honor_info_table         _honor_list;
            ctribut_table            _ctribut_list;   
            uosc_bp_table            _uosc_bp_list;
            project_info_table       _project_list;
	

	    /**************************
		1.this struct come from uosio.token contract (uosio.cdt 1.5.2)
	    **************************/
            struct [[uosio::table]] account {
                asset    balance;
                uint64_t primary_key()const { return balance.symbol.code().raw(); }
            };

            typedef uosio::multi_index< "accounts"_n, account > accounts;

        private:

            //op   1: add or modify data;   2: delete data;  others: error
            void set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

            // get current bp list 
            bool getbplist(std::vector<name>& bplist);

            // check the authorityt of assess project
            bool check_authority(const name bp);

            //get asset of one account
            asset get_balance( name token_contract_account, name owner, symbol_code sym_code );
    };

}
