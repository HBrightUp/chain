/**
 *  @file
 *  @copyright defined in uos/LICENSE.txt
 */
#pragma once

#include <uosio.system/native.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>
#include <uosiolib/privileged.hpp>
#include <uosiolib/singleton.hpp>
#include <uosio.system/exchange_state.hpp>

#include <string>

namespace uosiosystem {

   using uosio::asset;
   using uosio::indexed_by;
   using uosio::const_mem_fun;
   using uosio::block_timestamp;

   struct name_bid {
     account_name            newname;
     account_name            high_bidder;
     int64_t                 high_bid = 0; ///< negative high_bid == closed auction waiting to be claimed
     uint64_t                last_bid_time = 0;

     auto     primary_key()const { return newname;                          }
     uint64_t by_high_bid()const { return static_cast<uint64_t>(-high_bid); }
   };

   typedef uosio::multi_index< N(namebids), name_bid,
                               indexed_by<N(highbid), const_mem_fun<name_bid, uint64_t, &name_bid::by_high_bid>  >
                               >  name_bid_table;


   struct uosio_global_state : uosio::blockchain_parameters {
      uint64_t free_ram()const { return max_ram_size - total_ram_bytes_reserved; }
      //<begin> ==== modify by camphor ==== </begin>
      uint64_t             max_ram_size = 8ll*1024 * 1024 * 1024;
      //<end> ==== modify by camphor ==== </end>
      uint64_t             total_ram_bytes_reserved = 0;
      int64_t              total_ram_stake = 0;

      block_timestamp      last_producer_schedule_update;
      uint64_t             last_pervote_bucket_fill = 0;
      int64_t              pervote_bucket = 0;
      int64_t              perblock_bucket = 0;
      uint32_t             total_unpaid_blocks = 0; /// all blocks which have been produced but not paid
      int64_t              total_activated_stake = 0;
      uint64_t             thresh_activated_stake_time = 0;
      uint16_t             last_producer_schedule_size = 0;
      double               total_producer_vote_weight = 0; /// the sum of all producer votes
      block_timestamp      last_name_close;

      // explicit serialization macro is not necessary, used here only to improve compilation time
      UOSLIB_SERIALIZE_DERIVED( uosio_global_state, uosio::blockchain_parameters,
                                (max_ram_size)(total_ram_bytes_reserved)(total_ram_stake)
                                (last_producer_schedule_update)(last_pervote_bucket_fill)
                                (pervote_bucket)(perblock_bucket)(total_unpaid_blocks)(total_activated_stake)(thresh_activated_stake_time)
                                (last_producer_schedule_size)(total_producer_vote_weight)(last_name_close) )
   };

   struct producer_info {
      account_name          owner;
      double                total_votes = 0;
      uosio::public_key     producer_key; /// a packed public key object
      bool                  is_active = true;
      bool                  is_remove = false;
      int64_t               release = 0;
      std::string           url;
      uint32_t              unpaid_blocks = 0;
      uint64_t              last_claim_time = 0;
      uint16_t              location = 0;
      std::string   	    ulord_addr;

      uint64_t primary_key()const { return owner;                                   }
      double   by_votes()const    { return is_active ? -total_votes : total_votes;  }
      bool     active()const      { return is_active;                               }
      void     deactivate()       { producer_key = public_key(); is_active = false; }

      // explicit serialization macro is not necessary, used here only to improve compilation time
      UOSLIB_SERIALIZE( producer_info, (owner)(total_votes)(producer_key)(is_active)(is_remove)(release)(url)
                        (unpaid_blocks)(last_claim_time)(location)(ulord_addr) )
   };

	//Begin: modify by UOS(hml) Date:2018-09-25
	struct vote_producer {
	   account_name  producer;
	   asset voted;

	};

	typedef std::vector<vote_producer>::iterator iterVoteProducer;

	const int32_t max_voter_producers = 21;
   const int32_t max_bplist_number = 21;
   const int32_t max_time_interval = 126;
   static const char* pszBase58 = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz";
	//End: modify by UOS(hml) Date:2018-09-25
  

   struct voter_info {
      account_name                owner = 0; /// the voter
      account_name                proxy = 0; /// the proxy set by the voter, if any
      //std::vector<account_name>   producers; /// the producers approved by this voter if no proxy set
      int64_t                     staked = 0;

	  //Begin: modify by UOS(hml) Date:2018-09-19
	  //voted record,one user can vote multiple producer
	  //int64_t 					  voted = 0;
	  std::vector<vote_producer> vote_producers;
	  //End: modify by UOS(hml) Date:2018-09-19

      /**
       *  Every time a vote is cast we must first "undo" the last vote weight, before casting the
       *  new vote weight.  Vote weight is calculated as:
       *
       *  stated.amount * 2 ^ ( weeks_since_launch/weeks_per_year)
       */
      double                      last_vote_weight = 0; /// the vote weight cast the last time the vote was updated

      /**
       * Total vote weight delegated to this voter.
       */
      double                      proxied_vote_weight= 0; /// the total vote weight delegated to this voter as a proxy
      bool                        is_proxy = 0; /// whether the voter is a proxy for others

	  
      uint32_t                    reserved1 = 0;
      time                        reserved2 = 0;
      uosio::asset                reserved3;

      uint64_t primary_key()const { return owner; }

      // explicit serialization macro is not necessary, used here only to improve compilation time
      UOSLIB_SERIALIZE( voter_info, (owner)(proxy)/*(producers)*/(vote_producers)(staked)(last_vote_weight)(proxied_vote_weight)(is_proxy)(reserved1)(reserved2)(reserved3) )
   };

   // 
   // alvin  add begin . 
   const uint64_t MIN_BP_NUM = 1;
   const uint64_t MIN_BP_OUT_NUM = 2;
   
   enum bp_list_op {
        ENU_BPLIST_OUTPUT = 1,      // show all list info
        ENU_BPLIST_INIT,          //  init database of bpglobal/bpoutlist/minprodbp
   };

   struct one_minute_bp_list{
      
      //uosio::public_key     producer_key;   // bp  publik key
      std::string 			ulord_addr;
      uint64_t              bp_valid_time=0;  // bp start valid time,  120 second one record.   
      account_name          bpname = 0;    // bp name
	 
      uint64_t primary_key()const { return bp_valid_time; }
      UOSLIB_SERIALIZE( one_minute_bp_list, (ulord_addr)(bp_valid_time)(bpname) )
   };

   struct one_minute_prod_bp{
      //uosio::public_key     producer_key;   // bp  publik key
      std::string           ulord_addr;
      account_name          bpname =0;      // bp name .  
      uint64_t              on_line_time=0; // bp  on line time. +120s every bp update;   
      uint64_t              uptime =0;      // last time update; 
      uint64_t primary_key()const { return bpname; }
      uint64_t by_ontime()const { return on_line_time;   }
      uint64_t by_uptime()const { return uptime;         }
   
      UOSLIB_SERIALIZE( one_minute_prod_bp, (ulord_addr)(bpname)(on_line_time) (uptime) )
   };   
  
   struct bp_global_var{
      uint64_t        id;
      uint64_t        val;
      uint64_t        primary_key() const { return id; }
      UOSLIB_SERIALIZE(bp_global_var, (id)(val));
   };

     struct bp_info{
      uint64_t   primary;
      bool        is_follow;
      uint64_t    oldest_time;
      uint32_t    bplist_num;
      uint64_t        primary_key() const { return primary; }

      UOSLIB_SERIALIZE(bp_info, (primary)(is_follow)(oldest_time)(bplist_num));
   };

   #if 1
   //add by hml Date:2019-03-04
   struct share_bplist {
      uint64_t   primary;
      account_name bp_name;
      uint64_t count;
      uint64_t time;
      std::string hash_bplist;

      //std::vector< std::pair<uosio::producer_key,uint16_t> > bplist;
      std::vector< uosio::producer_key > bplist;
      uint64_t        primary_key() const { return primary; }

       UOSLIB_SERIALIZE(share_bplist, (primary)(bp_name)(count)(time)(hash_bplist)(bplist));
   };

   typedef uosio::multi_index< N(sharebplist), share_bplist>  share_bplist_table;
#endif

   typedef uosio::multi_index< N(bpglobal), bp_global_var > bp_global_var_table;  // 分钟表

   typedef uosio::multi_index< N(bpoutlist), one_minute_bp_list> min_out_bp_table;  // 分钟表
  
   typedef uosio::multi_index< N(uosclist), one_minute_bp_list> uosc_bp_table;  // uosc bp list

   typedef uosio::multi_index< N(uoslist), one_minute_bp_list> uos_bp_table;  // uos bp list

   typedef uosio::multi_index< N(minprodbp), one_minute_prod_bp,
          indexed_by<N(byontime), const_mem_fun<one_minute_prod_bp, uint64_t, &one_minute_prod_bp::by_ontime>  >,
          indexed_by<N(byuptime), const_mem_fun<one_minute_prod_bp, uint64_t, &one_minute_prod_bp::by_uptime>  >
                               >  min_prod_table;

   // alvin  add end .   bp out list

   typedef uosio::multi_index< N(voters), voter_info>  voters_table;


   typedef uosio::multi_index< N(producers), producer_info,
                               indexed_by<N(prototalvote), const_mem_fun<producer_info, double, &producer_info::by_votes>  >
                               >  producers_table;

   typedef uosio::singleton<N(global), uosio_global_state> global_state_singleton;

   typedef uosio::multi_index< N(bpinfo), bp_info> bpinfo_talbe;

   //   static constexpr uint32_t     max_inflation_rate = 5;  // 5% annual inflation
   static constexpr uint32_t     seconds_per_day = 24 * 3600;
   static constexpr uint64_t     system_token_symbol = CORE_SYMBOL;
   static constexpr int64_t      PRODUCE_MIN_BW = 1000000000;

   class system_contract : public native {
      private:
         voters_table           _voters;
         producers_table        _producers;
         global_state_singleton _global;
        // add by alvin for other chain dpos output  begin  
         min_out_bp_table       _min_out_bp_table; //
         min_prod_table         _minprodtable;     //
         bp_global_var_table    _bp_global_var;
        
          uosc_bp_table         _uosc_bp_list;
         // add by alvin for other chain dpos output  end  
 
         uosio_global_state     _gstate;
         rammarket              _rammarket;

         bpinfo_talbe            _bpinfo;

         //uos_bp_table         _uos_bp_list;
         share_bplist_table     _share_bplist_table;

      public:
         system_contract( account_name s );
         ~system_contract();

         // Actions:
         void onblock( block_timestamp timestamp, account_name producer );
                      // const block_header& header ); /// only parse first 3 fields of block header

         // functions defined in delegate_bandwidth.cpp

         /**
          *  Stakes SYS from the balance of 'from' for the benfit of 'receiver'.
          *  If transfer == true, then 'receiver' can unstake to their account
          *  Else 'from' can unstake at any time.
          */
         void delegatebw( account_name from, account_name receiver,
                          asset stake_net_quantity, asset stake_cpu_quantity, bool transfer );


         /**
          *  Decreases the total tokens delegated by from to receiver and/or
          *  frees the memory associated with the delegation if there is nothing
          *  left to delegate.
          *
          *  This will cause an immediate reduction in net/cpu bandwidth of the
          *  receiver.
          *
          *  A transaction is scheduled to send the tokens back to 'from' after
          *  the staking period has passed. If existing transaction is scheduled, it
          *  will be canceled and a new transaction issued that has the combined
          *  undelegated amount.
          *
          *  The 'from' account loses voting power as a result of this call and
          *  all producer tallies are updated.
          */
         void undelegatebw( account_name from, account_name receiver,
                            asset unstake_net_quantity, asset unstake_cpu_quantity );


         /**
          * Increases receiver's ram quota based upon current price and quantity of
          * tokens provided. An inline transfer from receiver to system contract of
          * tokens will be executed.
          */
         void buyram( account_name buyer, account_name receiver, asset tokens );
         void buyrambytes( account_name buyer, account_name receiver, uint32_t bytes );

         /**
          *  Reduces quota my bytes and then performs an inline transfer of tokens
          *  to receiver based upon the average purchase price of the original quota.
          */
         void sellram( account_name receiver, int64_t bytes );

         /**
          *  This action is called after the delegation-period to claim all pending
          *  unstaked tokens belonging to owner
          */
         void refund( account_name owner );

         // functions defined in voting.cpp

         void regproducer( const account_name producer, const public_key& producer_key, const std::string& url, const std::string& ulord_addr, uint16_t location );

         void unregprod( const account_name producer , bool release);

         void setram( uint64_t max_ram_size );

         void setbpout(uint64_t option);

         void voteproducer( const account_name voter,     asset quantity, const account_name proxy, const std::vector<account_name>& producers );

		 //Begin: modify by UOS(hml) Date:2018-09-21
		 void cancelvote( const account_name voter, asset quantity, const std::vector<account_name>& producers );
		//End: modify by UOS(hml) Date:2018-09-21
     
     	//Begin: add by UOS(hml) Date:2018-09-21
		 void cancelallvote( const account_name voter, bool is_sudo);
		//End: add by UOS(hml) Date:2018-09-21
		
         void regproxy( const account_name proxy, bool isproxy );

         void setparams( const uosio::blockchain_parameters& params );

		
         // functions defined in producer_pay.cpp
         void claimrewards( const account_name& owner );

		 //Begin: modify by UOS(hml) Date:2018-10-29
		 //pay fee by net source
		 void payfee( const account_name payer , const uint64_t net_usage , const uint64_t virtule_net_limit , const account_name act_account ,const action_name act_name);
		//End: modify by UOS(hml) Date:2018-10-29
		
         void setpriv( account_name account, uint8_t ispriv );

         void rmvproducer( account_name producer );

         void bidname( account_name bidder, account_name newname, asset bid );

         void setbpinfo(bool is_follow);

         //void setbplist(uosiosystem::share_bplist& share);
         void setbplist(const account_name bp_name, const uint64_t bp_time, const std::string hast_bplist, const std::vector< uosio::producer_key >& bplist);
      private:
         std::string EncodeBase58(const std::vector<unsigned char>& vch);
         std::string EncodeBase58(const unsigned char* pbegin, const unsigned char* pend);

         void update_elected_producers( block_timestamp timestamp );

         // Implementation details:

         //defind in delegate_bandwidth.cpp
         void changebw( account_name from, account_name receiver,
                        asset stake_net_quantity, asset stake_cpu_quantity, bool transfer , bool penaltyprod = false);

         //defind in delegate_bandwidth.cpp
         void penaltyprodbw( account_name producer);

         //verfy produce bandwith
         bool prodbw_verf(account_name prod);

         //defined in voting.hpp
         static uosio_global_state get_default_parameters();

         void update_votes( const account_name voter,     asset quantity,  account_name proxy, const std::vector<account_name>& producers, bool voting );

		//Begin: modify by UOS(hml) Date:2018-09-26
		 void validate_votes(const account_name voter_name,     asset quantity);
		//End: modify by UOS(hml) Date:2018-09-26

         // defined in voting.cpp
         void propagate_weight_change( const voter_info& voter );

         //add by hml date: 2019-03-11
         size_t from_hex( const std::string& hex_str, char* out_data, size_t out_data_len );
         uint8_t from_hex( char c );
   };

} /// uosiosystem
