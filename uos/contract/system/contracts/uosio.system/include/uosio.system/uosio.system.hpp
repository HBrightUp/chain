#pragma once

#include <uosio.system/native.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>
#include <uosiolib/privileged.hpp>
#include <uosiolib/singleton.hpp>
#include <uosio.system/exchange_state.hpp>

#include <string>
#include <deque>
#include <type_traits>
#include <optional>

#ifdef CHANNEL_RAM_AND_NAMEBID_FEES_TO_REX
#undef CHANNEL_RAM_AND_NAMEBID_FEES_TO_REX
#endif
// CHANNEL_RAM_AND_NAMEBID_FEES_TO_REX macro determines whether ramfee and namebid proceeds are
// channeled to REX pool. In order to stop these proceeds from being channeled, the macro must
// be set to 0.
#define CHANNEL_RAM_AND_NAMEBID_FEES_TO_REX 0

static constexpr int64_t      PRODUCE_MIN_BW = 1000000000;

//Begin: add by hml Date:2019-11-21
#define FUNCTION_SWITCH_REX 0
//End: add by hml Date:2019-11-21

namespace uosiosystem {

   using uosio::name;
   using uosio::asset;
   using uosio::symbol;
   using uosio::symbol_code;
   using uosio::indexed_by;
   using uosio::const_mem_fun;
   using uosio::block_timestamp;
   using uosio::time_point;
   using uosio::time_point_sec;
   using uosio::microseconds;
   using uosio::datastream;
   using uosio::check;

   template<typename E, typename F>
   static inline auto has_field( F flags, E field )
   -> std::enable_if_t< std::is_integral_v<F> && std::is_unsigned_v<F> &&
                        std::is_enum_v<E> && std::is_same_v< F, std::underlying_type_t<E> >, bool>
   {
      return ( (flags & static_cast<F>(field)) != 0 );
   }

   template<typename E, typename F>
   static inline auto set_field( F flags, E field, bool value = true )
   -> std::enable_if_t< std::is_integral_v<F> && std::is_unsigned_v<F> &&
                        std::is_enum_v<E> && std::is_same_v< F, std::underlying_type_t<E> >, F >
   {
      if( value )
         return ( flags | static_cast<F>(field) );
      else
         return ( flags & ~static_cast<F>(field) );
   }

   struct [[uosio::table, uosio::contract("uosio.system")]] name_bid {
     name            newname;
     name            high_bidder;
     int64_t         high_bid = 0; ///< negative high_bid == closed auction waiting to be claimed
     time_point      last_bid_time;

     uint64_t primary_key()const { return newname.value;                    }
     uint64_t by_high_bid()const { return static_cast<uint64_t>(-high_bid); }
   };

   struct [[uosio::table, uosio::contract("uosio.system")]] bid_refund {
      name         bidder;
      asset        amount;

      uint64_t primary_key()const { return bidder.value; }
   };

   typedef uosio::multi_index< "namebids"_n, name_bid,
                               indexed_by<"highbid"_n, const_mem_fun<name_bid, uint64_t, &name_bid::by_high_bid>  >
                             > name_bid_table;

   typedef uosio::multi_index< "bidrefunds"_n, bid_refund > bid_refund_table;

   struct [[uosio::table("global"), uosio::contract("uosio.system")]] uosio_global_state : uosio::blockchain_parameters {
      uint64_t free_ram()const { return max_ram_size - total_ram_bytes_reserved; }
      //<begin> ==== modify by camphor ==== </begin>
      uint64_t             max_ram_size = 8ll*1024 * 1024 * 1024;
      //<end> ==== modify by camphor ==== </end>
      uint64_t             total_ram_bytes_reserved = 0;
      int64_t              total_ram_stake = 0;

      block_timestamp      last_producer_schedule_update;
      time_point           last_pervote_bucket_fill;
      int64_t              pervote_bucket = 0;
      int64_t              perblock_bucket = 0;
      uint32_t             total_unpaid_blocks = 0; /// all blocks which have been produced but not paid
      int64_t              total_activated_stake = 0;
      time_point           thresh_activated_stake_time;
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

   /**
    * Defines new global state parameters added after version 1.0
    */
   struct [[uosio::table("global2"), uosio::contract("uosio.system")]] uosio_global_state2 {
      uosio_global_state2(){}

      uint16_t          new_ram_per_block = 0;
      block_timestamp   last_ram_increase;
      block_timestamp   last_block_num; /* deprecated */
      double            total_producer_votepay_share = 0;
      uint8_t           revision = 0; ///< used to track version updates in the future.

      UOSLIB_SERIALIZE( uosio_global_state2, (new_ram_per_block)(last_ram_increase)(last_block_num)
                        (total_producer_votepay_share)(revision) )
   };

   struct [[uosio::table("global3"), uosio::contract("uosio.system")]] uosio_global_state3 {
      uosio_global_state3() { }
      time_point        last_vpay_state_update;
      double            total_vpay_share_change_rate = 0;

      UOSLIB_SERIALIZE( uosio_global_state3, (last_vpay_state_update)(total_vpay_share_change_rate) )
   };

   struct [[uosio::table, uosio::contract("uosio.system")]] producer_info {
      name                  owner;
      double                total_votes = 0;
      uosio::public_key     producer_key; /// a packed public key object
      bool                  is_active = true;
      bool                  is_remove = false;
      int64_t               release = 0;
      std::string           url;
      uint32_t              unpaid_blocks = 0;
      time_point            last_claim_time;
      uint16_t              location = 0;
      std::string   	      ulord_addr;

      uint64_t primary_key()const { return owner.value;                             }
      double   by_votes()const    { return is_active ? -total_votes : total_votes;  }
      bool     active()const      { return is_active;                               }
      void     deactivate()       { producer_key = public_key(); is_active = false; }

      // explicit serialization macro is not necessary, used here only to improve compilation time
      UOSLIB_SERIALIZE( producer_info, (owner)(total_votes)(producer_key)(is_active)(is_remove)(release)(url)
                        (unpaid_blocks)(last_claim_time)(location)(ulord_addr) )
   };

   struct [[uosio::table, uosio::contract("uosio.system")]] producer_info2 {
      name            owner;
      double          votepay_share = 0;
      time_point      last_votepay_share_update;

      uint64_t primary_key()const { return owner.value; }

      // explicit serialization macro is not necessary, used here only to improve compilation time
      UOSLIB_SERIALIZE( producer_info2, (owner)(votepay_share)(last_votepay_share_update) )
   };


  //Begin: modify by UOS(hml) Date:2018-09-25
	struct vote_producer {
	   name  producer;
	   asset voted;

	};

	typedef std::vector<vote_producer>::iterator iterVoteProducer;

	const int32_t max_producers_number = 15;
	//End: modify by UOS(hml) Date:2018-09-25

   /**
    * Voter info.
    *
    * @details Voter info stores information about the voter:
    * - `owner` the voter
    * - `proxy` the proxy set by the voter, if any
    * - `producers` the producers approved by this voter if no proxy set
    * - `staked` the amount staked
    */
   struct [[uosio::table, uosio::contract("uosio.system")]] voter_info {
      name                owner;     /// the voter
      name                proxy;     /// the proxy set by the voter, if any
      //std::vector<name>   producers; /// the producers approved by this voter if no proxy set
      int64_t             staked = 0;

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
      double              last_vote_weight = 0; /// the vote weight cast the last time the vote was updated

      /**
       * Total vote weight delegated to this voter.
       */
      double              proxied_vote_weight= 0; /// the total vote weight delegated to this voter as a proxy
      bool                is_proxy = 0; /// whether the voter is a proxy for others


      uint32_t            flags1 = 0;
      uint32_t            reserved2 = 0;
      uosio::asset        reserved3;

      uint64_t primary_key()const { return owner.value; }

      enum class flags1_fields : uint32_t {
         ram_managed = 1,
         net_managed = 2,
         cpu_managed = 4
      };

      // explicit serialization macro is not necessary, used here only to improve compilation time
      UOSLIB_SERIALIZE( voter_info, (owner)(proxy)/*(producers)*/(vote_producers)(staked)/*(vote_producers)*/(last_vote_weight)(proxied_vote_weight)(is_proxy)(flags1)(reserved2)(reserved3) )
   };

   typedef uosio::multi_index< "voters"_n, voter_info >  voters_table;


   typedef uosio::multi_index< "producers"_n, producer_info,
                               indexed_by<"prototalvote"_n, const_mem_fun<producer_info, double, &producer_info::by_votes>  >
                             > producers_table;
   typedef uosio::multi_index< "producers2"_n, producer_info2 > producers_table2;

   typedef uosio::singleton< "global"_n, uosio_global_state >   global_state_singleton;
   typedef uosio::singleton< "global2"_n, uosio_global_state2 > global_state2_singleton;
   typedef uosio::singleton< "global3"_n, uosio_global_state3 > global_state3_singleton;

   static constexpr uint32_t     seconds_per_day = 24 * 3600;

   struct [[uosio::table,uosio::contract("uosio.system")]] rex_pool {
      uint8_t    version = 0;
      asset      total_lent; /// total amount of CORE_SYMBOL in open rex_loans
      asset      total_unlent; /// total amount of CORE_SYMBOL available to be lent (connector)
      asset      total_rent; /// fees received in exchange for lent  (connector)
      asset      total_lendable; /// total amount of CORE_SYMBOL that have been lent (total_unlent + total_lent)
      asset      total_rex; /// total number of REX shares allocated to contributors to total_lendable
      asset      namebid_proceeds; /// the amount of CORE_SYMBOL to be transferred from namebids to REX pool
      uint64_t   loan_num = 0; /// increments with each new loan

      uint64_t primary_key()const { return 0; }
   };

   typedef uosio::multi_index< "rexpool"_n, rex_pool > rex_pool_table;

   struct [[uosio::table,uosio::contract("uosio.system")]] rex_fund {
      uint8_t version = 0;
      name    owner;
      asset   balance;

      uint64_t primary_key()const { return owner.value; }
   };

   typedef uosio::multi_index< "rexfund"_n, rex_fund > rex_fund_table;

   struct [[uosio::table,uosio::contract("uosio.system")]] rex_balance {
      uint8_t version = 0;
      name    owner;
      asset   vote_stake; /// the amount of CORE_SYMBOL currently included in owner's vote
      asset   rex_balance; /// the amount of REX owned by owner
      int64_t matured_rex = 0; /// matured REX available for selling
      std::deque<std::pair<time_point_sec, int64_t>> rex_maturities; /// REX daily maturity buckets

      uint64_t primary_key()const { return owner.value; }
   };

   typedef uosio::multi_index< "rexbal"_n, rex_balance > rex_balance_table;

   struct [[uosio::table,uosio::contract("uosio.system")]] rex_loan {
      uint8_t             version = 0;
      name                from;
      name                receiver;
      asset               payment;
      asset               balance;
      asset               total_staked;
      uint64_t            loan_num;
      uosio::time_point   expiration;

      uint64_t primary_key()const { return loan_num;                   }
      uint64_t by_expr()const     { return expiration.elapsed.count(); }
      uint64_t by_owner()const    { return from.value;                 }
   };

   typedef uosio::multi_index< "cpuloan"_n, rex_loan,
                               indexed_by<"byexpr"_n,  const_mem_fun<rex_loan, uint64_t, &rex_loan::by_expr>>,
                               indexed_by<"byowner"_n, const_mem_fun<rex_loan, uint64_t, &rex_loan::by_owner>>
                             > rex_cpu_loan_table;

   typedef uosio::multi_index< "netloan"_n, rex_loan,
                               indexed_by<"byexpr"_n,  const_mem_fun<rex_loan, uint64_t, &rex_loan::by_expr>>,
                               indexed_by<"byowner"_n, const_mem_fun<rex_loan, uint64_t, &rex_loan::by_owner>>
                             > rex_net_loan_table;

   struct [[uosio::table,uosio::contract("uosio.system")]] rex_order {
      uint8_t             version = 0;
      name                owner;
      asset               rex_requested;
      asset               proceeds;
      asset               stake_change;
      uosio::time_point   order_time;
      bool                is_open = true;

      void close()                { is_open = false;    }
      uint64_t primary_key()const { return owner.value; }
      uint64_t by_time()const     { return is_open ? order_time.elapsed.count() : std::numeric_limits<uint64_t>::max(); }
   };

   typedef uosio::multi_index< "rexqueue"_n, rex_order,
                               indexed_by<"bytime"_n, const_mem_fun<rex_order, uint64_t, &rex_order::by_time>>> rex_order_table;

   struct rex_order_outcome {
      bool success;
      asset proceeds;
      asset stake_change;
   };

   
//Begin: add by UOS(hml) 
struct producer_ulord {

      /**
       * Name of the producer
       *
       * @brief Name of the producer
       */
      name     producer_name;

      /**
       * ulord address by this producer
       *
       * @brief 
       */
      std::string       ulord_addr;

      friend bool operator < ( const producer_ulord& a, const producer_ulord& b ) {
         return a.producer_name.value < b.producer_name.value;
      }

      UOSLIB_SERIALIZE( producer_ulord, (producer_name)(ulord_addr) )
   };
   
//End: add by UOS(hml)


  //Begin: add y alvin
   
   enum bp_list_op {
        ENU_BPLIST_OUTPUT = 1,      // show all list info
        ENU_BPLIST_INIT,          //  init database of bpglobal/bpoutlist/minprodbp
   };

   struct [[uosio::table,uosio::contract("uosio.system")]] one_minute_bp_list{
      
      //uosio::public_key     producer_key;   // bp  publik key
      std::string 			ulord_addr;
      uint64_t             bp_valid_time=0;  // bp start valid time,  120 second one record.   
      name                 bpname;    // bp name
	 
      uint64_t primary_key()const { return bp_valid_time; }
      UOSLIB_SERIALIZE( one_minute_bp_list, (ulord_addr)(bp_valid_time)(bpname) )
   };

   //Begin: removed by hml Date:2019-11-19
#if 0
   struct [[uosio::table,uosio::contract("uosio.system")]] one_minute_prod_bp{
      //uosio::public_key     producer_key;   // bp  publik key
      std::string           ulord_addr;
      name                  bpname;      // bp name .  
      uint64_t              on_line_time=0; // bp  on line time. +120s every bp update;   
      uint64_t              uptime =0;      // last time update; 
      uint64_t primary_key()const { return bpname.value; }
      uint64_t by_ontime()const { return on_line_time;   }
      uint64_t by_uptime()const { return uptime;         }
   
      UOSLIB_SERIALIZE( one_minute_prod_bp, (ulord_addr)(bpname)(on_line_time) (uptime) )
   };   
#endif
   //End: removed by hml Date:2019-11-19

   //global variable
   struct [[uosio::table,uosio::contract("uosio.system")]] sys_args {
	  name            key;
	  uint64_t        val;
	  uint64_t		  primary_key() const { return key.value; }
  
	  UOSLIB_SERIALIZE(sys_args, (key)(val));
   };       

   struct  [[uosio::table,uosio::contract("uosio.system")]] uidaccount    // ���� uosuidwallet ��Լ��������ݽṹ
   {
      name         uid;        //  uid ��   Ψһ����  ������
      name        account;    //  �˻��� 
      uint64_t      creattime;
      uint64_t    spentmoney; //  �����û�����  ����п��ܱ仯  ����Ҫ������ Ϊ����Ǯ��
      uint64_t    storage;    //  �����洢�ֶ�
	   bool        ispayed;    //  �Ƿ���֧��	  	
      uint64_t	  primary_key() const { return uid.value; }
      uint64_t    by_creattime()const { return creattime; }
      
	   UOSLIB_SERIALIZE(uidaccount, (uid)(account)(creattime)(spentmoney)(storage)(ispayed))
   };

   struct[[uosio::table,uosio::contract("uosio.token")]]  account {
      asset    balance;

      uint64_t primary_key()const { return balance.symbol.raw(); }
   };
   typedef uosio::multi_index< "uidaccount"_n, uidaccount,
                 indexed_by<"bycreattime"_n, const_mem_fun<uidaccount, uint64_t, &uidaccount::by_creattime> >
                > uid_account_table;
  
   typedef uosio::multi_index< "uosclist"_n, one_minute_bp_list> uosc_bp_table;  // uosc bp list

   typedef uosio::multi_index< "sysargslist"_n, sys_args > sys_args_table;
   typedef uosio::multi_index<"accounts"_n, account> accounts;

   //End: add by alvin.   bp out list

   //Begin: add by hml Date:2019-11-19
   ///@abi table bpvtpay i64
   struct bp_and_vt_pay{
      uint32_t                 id;   //　id -- 0  : now   ;   id -- 1 : next
      uint64_t                     start_line;
      uint64_t                     laster_line;
      uint64_t                     dead_line;
      asset                    inherit;
      asset                    quantity;
      uint32_t primary_key()const { return id; }
      UOSLIB_SERIALIZE( bp_and_vt_pay, (id)(start_line)(laster_line)(dead_line)(inherit)(quantity))
   };
   typedef uosio::multi_index<"bpvtpay"_n , bp_and_vt_pay> bpvtpay;


   ///@abi table bvrealpay i64
   struct bv_real_pay{
      name             owner;
      asset                    quantity;
      uint64_t primary_key()const { return owner.value; }
      UOSLIB_SERIALIZE( bv_real_pay, (owner)(quantity))
   };
   typedef uosio::multi_index<"bvrealpay"_n , bv_real_pay> bvrealpay;
   //End: add by hml Date:2019-11-19

   /**
    * The UOSIO system contract.
    *
    * @details The UOSIO system contract governs ram market, voters, producers, global state.
    */
   class [[uosio::contract("uosio.system")]] system_contract : public native {

      private:
         voters_table            _voters;
         producers_table         _producers;
         producers_table2        _producers2;
         global_state_singleton  _global;
         global_state2_singleton _global2;
         global_state3_singleton _global3;
         uosio_global_state      _gstate;
         uosio_global_state2     _gstate2;
         uosio_global_state3     _gstate3;
         rammarket               _rammarket;
         rex_pool_table          _rexpool;
         rex_fund_table          _rexfunds;
         rex_balance_table       _rexbalance;
         rex_order_table         _rexorders;
         // add by alvin for other chain dpos output  begin  
          uosc_bp_table         _uosc_bp_list;
          sys_args_table        _sys_args_list;
         // add by alvin for other chain dpos output  end 

         //Begin:add by hml Date:2019-11-19
         bpvtpay _bpvtpay;
         bvrealpay _bvrealpay;
         //Begin:add by hml Date:2019-11-19

      public:
         static constexpr uosio::name active_permission{"active"_n};
         static constexpr uosio::name token_account{"uosio.token"_n};
         static constexpr uosio::name ram_account{"uosio.ram"_n};
         static constexpr uosio::name ramfee_account{"uosio.ramfee"_n};
         static constexpr uosio::name stake_account{"uosio.stake"_n};
         static constexpr uosio::name bpay_account{"uosio.bpay"_n};
         static constexpr uosio::name vpay_account{"uosio.vpay"_n};

         //Begin: add by hml Date:2019-11-13
         static constexpr uosio::name bvpay_account{"uosio.bvpay"_n}; 
         //End: add by hml Date:2019-11-13

         static constexpr uosio::name names_account{"uosio.names"_n};
         static constexpr uosio::name saving_account{"uosio.saving"_n};
         static constexpr uosio::name rex_account{"uosio.rex"_n};
         static constexpr uosio::name null_account{"uosio.null"_n};
         static constexpr symbol ramcore_symbol = symbol(symbol_code("RAMCORE"), 4);
         static constexpr symbol ram_symbol     = symbol(symbol_code("RAM"), 0);
         static constexpr symbol rex_symbol     = symbol(symbol_code("REX"), 4);

         system_contract( name s, name code, datastream<const char*> ds );
         ~system_contract();

         static symbol get_core_symbol( name system_account = "uosio"_n ) {
            rammarket rm(system_account, system_account.value);
            const static auto sym = get_core_symbol( rm );
            return sym;
         }

         // Actions:
         [[uosio::action]]
         void init( unsigned_int version, symbol core );
         [[uosio::action]]
         void onblock( ignore<block_header> header );

         [[uosio::action]]
         void setalimits( name account, int64_t ram_bytes, int64_t net_weight, int64_t cpu_weight );

         [[uosio::action]]
         void setacctram( name account, std::optional<int64_t> ram_bytes );

         [[uosio::action]]
         void setacctnet( name account, std::optional<int64_t> net_weight );

         [[uosio::action]]
         void setacctcpu( name account, std::optional<int64_t> cpu_weight );

         // functions defined in delegate_bandwidth.cpp

         /**
          *  Stakes SYS from the balance of 'from' for the benfit of 'receiver'.
          *  If transfer == true, then 'receiver' can unstake to their account
          *  Else 'from' can unstake at any time.
          */
         [[uosio::action]]
         void delegatebw( name from, name receiver,
                          asset stake_net_quantity, asset stake_cpu_quantity, bool transfer );

         /**
          * Sets total_rent balance of REX pool to the passed value
          */
         [[uosio::action]]
         void setrex( const asset& balance );

         /**
          * Deposits core tokens to user REX fund. All proceeds and expenses related to REX are added to
          * or taken out of this fund. Inline token transfer from user balance is executed.
          */
         [[uosio::action]]
         void deposit( const name& owner, const asset& amount );

         /**
          * Withdraws core tokens from user REX fund. Inline token transfer to user balance is
          * executed.
          */
         [[uosio::action]]
         void withdraw( const name& owner, const asset& amount );

         /**
          * Transfers core tokens from user REX fund and converts them to REX stake.
          * A voting requirement must be satisfied before action can be executed.
          * User votes are updated following this action.
          */
         [[uosio::action]]
         void buyrex( const name& from, const asset& amount );

         /**
          * Use staked core tokens to buy REX.
          * A voting requirement must be satisfied before action can be executed.
          * User votes are updated following this action.
          */
         [[uosio::action]]
         void unstaketorex( const name& owner, const name& receiver, const asset& from_net, const asset& from_cpu );

         /**
          * Converts REX stake back into core tokens at current exchange rate. If order cannot be
          * processed, it gets queued until there is enough in REX pool to fill order.
          * If successful, user votes are updated.
          */
         [[uosio::action]]
         void sellrex( const name& from, const asset& rex );

         /**
          * Cancels queued sellrex order. Order cannot be cancelled once it's been filled.
          */
         [[uosio::action]]
         void cnclrexorder( const name& owner );

         /**
          * Use payment to rent as many SYS tokens as possible and stake them for either CPU or NET for the
          * benefit of receiver, after 30 days the rented SYS delegation of CPU or NET will expire unless loan
          * balance is larger than or equal to payment.
          *
          * If loan has enough balance, it gets renewed at current market price, otherwise, it is closed and
          * remaining balance is refunded to loan owner.
          *
          * Owner can fund or defund a loan at any time before its expiration.
          *
          * All loan expenses and refunds come out of or are added to owner's REX fund.
          */
         [[uosio::action]]
         void rentcpu( const name& from, const name& receiver, const asset& loan_payment, const asset& loan_fund );
         [[uosio::action]]
         void rentnet( const name& from, const name& receiver, const asset& loan_payment, const asset& loan_fund );

         /**
          * Loan owner funds a given CPU or NET loan.
          */
         [[uosio::action]]
         void fundcpuloan( const name& from, uint64_t loan_num, const asset& payment );
         [[uosio::action]]
         void fundnetloan( const name& from, uint64_t loan_num, const asset& payment );
         /**
          * Loan owner defunds a given CPU or NET loan.
          */
         [[uosio::action]]
         void defcpuloan( const name& from, uint64_t loan_num, const asset& amount );
         [[uosio::action]]
         void defnetloan( const name& from, uint64_t loan_num, const asset& amount );

         /**
          * Updates REX vote stake of owner to its current value.
          */
         [[uosio::action]]
         void updaterex( const name& owner );

         /**
          * Processes max CPU loans, max NET loans, and max queued sellrex orders.
          * Action does not execute anything related to a specific user.
          */
         [[uosio::action]]
         void rexexec( const name& user, uint16_t max );

         /**
          * Consolidate REX maturity buckets into one that can be sold only 4 days
          * from the end of today.
          */
         [[uosio::action]]
         void consolidate( const name& owner );

         /**
          * Moves a specified amount of REX into savings bucket. REX savings bucket
          * never matures. In order for it to be sold, it has to be moved explicitly
          * out of that bucket. Then the moved amount will have the regular maturity
          * period of 4 days starting from the end of the day.
          */
         [[uosio::action]]
         void mvtosavings( const name& owner, const asset& rex );

         /**
          * Moves a specified amount of REX out of savings bucket. The moved amount
          * will have the regular REX maturity period of 4 days.
          */
         [[uosio::action]]
         void mvfrsavings( const name& owner, const asset& rex );

         /**
          * Deletes owner records from REX tables and frees used RAM.
          * Owner must not have an outstanding REX balance.
          */
         [[uosio::action]]
         void closerex( const name& owner );

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
         [[uosio::action]]
         void undelegatebw( name from, name receiver,
                            asset unstake_net_quantity, asset unstake_cpu_quantity );


         /**
          * Increases receiver's ram quota based upon current price and quantity of
          * tokens provided. An inline transfer from receiver to system contract of
          * tokens will be executed.
          */
         [[uosio::action]]
         void buyram( name payer, name receiver, asset quant );
         [[uosio::action]]
         void buyrambytes( name payer, name receiver, uint32_t bytes );

         /**
          *  Reduces quota my bytes and then performs an inline transfer of tokens
          *  to receiver based upon the average purchase price of the original quota.
          */
         [[uosio::action]]
         void sellram( name account, int64_t bytes );

         /**
          *  This action is called after the delegation-period to claim all pending
          *  unstaked tokens belonging to owner
          */
         [[uosio::action]]
         void refund( name owner );

         // functions defined in voting.cpp

         [[uosio::action]]
         void regproducer( const name producer, const public_key& producer_key, const std::string& url, const std::string& ulord_addr, uint16_t location );

         [[uosio::action]]
         void unregprod( const name producer, const bool release);

         [[uosio::action]]
         void setram( uint64_t max_ram_size );
		 
		 //Begin: modify by UOS(hml) Date:2018-09-21
         /**
          * Cancel vote action
          * @details  Cancel vote from producer
          * @param voter - the account of cancel vote
          * @param quantity - amount of cancel vote
          * @param producers - the producer list of voted by voter
          **/
         [[uosio::action]]
        void cancelvote( const name voter, asset quantity, const std::vector<name>& producers );
        //End: modify by UOS(hml) Date:2018-09-21
		
         [[uosio::action]]
         void setramrate( uint16_t bytes_per_block );

         [[uosio::action]]
         void voteproducer( const name voter, asset quantity, const name proxy, const std::vector<name>& producers );

         [[uosio::action]]
         void regproxy( const name proxy, bool isproxy );

         [[uosio::action]]
         void setparams( const uosio::blockchain_parameters& params );

         // functions defined in producer_pay.cpp
         [[uosio::action]]
         void claimrewards( const name owner );

         [[uosio::action]]
         void setpriv( name account, uint8_t is_priv );

         [[uosio::action]]
         void rmvproducer( name producer );

         [[uosio::action]]
         void updtrevision( uint8_t revision );

         [[uosio::action]]
         void bidname( name bidder, name newname, asset bid );

         [[uosio::action]]
         void bidrefund( name bidder, name newname );

         //Begin: modify by UOS(hml) Date:2018-10-29
        //pay fee by net source
        [[uosio::action]]
        void payfee( const name payer , const uint64_t net_usage , const uint64_t virtule_net_limit , const name act_account ,const name act_name);
        //End: modify by UOS(hml) Date:2018-10-29

          // Begin:add by(alvin) Date: 20190810
          [[uosio::action]]
          void setsysargs(const name& key, const uint64_t& val, uint8_t global);
          //End:add by(alvin) Date: 20190810

         using init_action = uosio::action_wrapper<"init"_n, &system_contract::init>;
         using setacctram_action = uosio::action_wrapper<"setacctram"_n, &system_contract::setacctram>;
         using setacctnet_action = uosio::action_wrapper<"setacctnet"_n, &system_contract::setacctnet>;
         using setacctcpu_action = uosio::action_wrapper<"setacctcpu"_n, &system_contract::setacctcpu>;
         using delegatebw_action = uosio::action_wrapper<"delegatebw"_n, &system_contract::delegatebw>;
         using deposit_action = uosio::action_wrapper<"deposit"_n, &system_contract::deposit>;
         using withdraw_action = uosio::action_wrapper<"withdraw"_n, &system_contract::withdraw>;
         using buyrex_action = uosio::action_wrapper<"buyrex"_n, &system_contract::buyrex>;
         using unstaketorex_action = uosio::action_wrapper<"unstaketorex"_n, &system_contract::unstaketorex>;
         using sellrex_action = uosio::action_wrapper<"sellrex"_n, &system_contract::sellrex>;
         using cnclrexorder_action = uosio::action_wrapper<"cnclrexorder"_n, &system_contract::cnclrexorder>;
         using rentcpu_action = uosio::action_wrapper<"rentcpu"_n, &system_contract::rentcpu>;
         using rentnet_action = uosio::action_wrapper<"rentnet"_n, &system_contract::rentnet>;
         using fundcpuloan_action = uosio::action_wrapper<"fundcpuloan"_n, &system_contract::fundcpuloan>;
         using fundnetloan_action = uosio::action_wrapper<"fundnetloan"_n, &system_contract::fundnetloan>;
         using defcpuloan_action = uosio::action_wrapper<"defcpuloan"_n, &system_contract::defcpuloan>;
         using defnetloan_action = uosio::action_wrapper<"defnetloan"_n, &system_contract::defnetloan>;
         using updaterex_action = uosio::action_wrapper<"updaterex"_n, &system_contract::updaterex>;
         using rexexec_action = uosio::action_wrapper<"rexexec"_n, &system_contract::rexexec>;
         using setrex_action = uosio::action_wrapper<"setrex"_n, &system_contract::setrex>;
         using mvtosavings_action = uosio::action_wrapper<"mvtosavings"_n, &system_contract::mvtosavings>;
         using mvfrsavings_action = uosio::action_wrapper<"mvfrsavings"_n, &system_contract::mvfrsavings>;
         using consolidate_action = uosio::action_wrapper<"consolidate"_n, &system_contract::consolidate>;
         using closerex_action = uosio::action_wrapper<"closerex"_n, &system_contract::closerex>;
         using undelegatebw_action = uosio::action_wrapper<"undelegatebw"_n, &system_contract::undelegatebw>;
         using buyram_action = uosio::action_wrapper<"buyram"_n, &system_contract::buyram>;
         using buyrambytes_action = uosio::action_wrapper<"buyrambytes"_n, &system_contract::buyrambytes>;
         using sellram_action = uosio::action_wrapper<"sellram"_n, &system_contract::sellram>;
         using refund_action = uosio::action_wrapper<"refund"_n, &system_contract::refund>;
         using regproducer_action = uosio::action_wrapper<"regproducer"_n, &system_contract::regproducer>;
         using unregprod_action = uosio::action_wrapper<"unregprod"_n, &system_contract::unregprod>;
         using setram_action = uosio::action_wrapper<"setram"_n, &system_contract::setram>;
         using setramrate_action = uosio::action_wrapper<"setramrate"_n, &system_contract::setramrate>;
         using voteproducer_action = uosio::action_wrapper<"voteproducer"_n, &system_contract::voteproducer>;
         using regproxy_action = uosio::action_wrapper<"regproxy"_n, &system_contract::regproxy>;
         using claimrewards_action = uosio::action_wrapper<"claimrewards"_n, &system_contract::claimrewards>;
         using rmvproducer_action = uosio::action_wrapper<"rmvproducer"_n, &system_contract::rmvproducer>;
         using updtrevision_action = uosio::action_wrapper<"updtrevision"_n, &system_contract::updtrevision>;
         using bidname_action = uosio::action_wrapper<"bidname"_n, &system_contract::bidname>;
         using bidrefund_action = uosio::action_wrapper<"bidrefund"_n, &system_contract::bidrefund>;
         using setpriv_action = uosio::action_wrapper<"setpriv"_n, &system_contract::setpriv>;
         using setalimits_action = uosio::action_wrapper<"setalimits"_n, &system_contract::setalimits>;
         using setparams_action = uosio::action_wrapper<"setparams"_n, &system_contract::setparams>;

         //Begin:add by hml Date:2019-11-18
         using setsysargs_action = uosio::action_wrapper<"setsysargs"_n, &system_contract::setsysargs>;
         using payfee_action = uosio::action_wrapper<"payfee"_n, &system_contract::setparams>;
         //End:add by hml Date:2019-11-18

      private:

         // Implementation details:


         //Begin: add by UOS(hml) Date:2018-09-21
         /**
          * details cancel all vote 
          * @param voter - account of cancel vote
          * @param is_sudo - is uosio account
          * 
          * **/
          void cancelallvote( const name voter, bool is_sudo);
          //End: add by UOS(hml) Date:2018-09-21

          //Begin: modify by UOS(hml) Date:2018-09-26
          void validate_votes(const name voter_name, asset quantity);

          asset get_bpvtpay(name call) const;
          //End: modify by UOS(hml) Date:2018-09-26

          // Begin:add by(alvin
          void issuetoken();
          void setglobalargs(const name& key, const uint64_t& val);
          //End:add by(alvin)

           // Begin:add by(lx) 
          //defind in delegate_bandwidth.cpp
          void penaltyprodbw( name producer);

          //verfy produce bandwith
          bool prodbw_verf(name prod);
          //End:add by(lx)


         static symbol get_core_symbol( const rammarket& rm ) {
            auto itr = rm.find(ramcore_symbol.raw());
            check(itr != rm.end(), "system contract must first be initialized ");
            return itr->quote.balance.symbol;
         }

         //defined in uosio.system.cpp
         static uosio_global_state get_default_parameters();
         static time_point current_time_point();
         static time_point_sec current_time_point_sec();
         static block_timestamp current_block_time();
         symbol core_symbol()const;
         void update_ram_supply();

         // defined in rex.cpp
         void runrex( uint16_t max );
         void update_resource_limits( const name& from, const name& receiver, int64_t delta_net, int64_t delta_cpu );
         void check_voting_requirement( const name& owner,
                                        const char* error_msg = "must vote for at least 21 producers or for a proxy before buying REX" )const;
         rex_order_outcome fill_rex_order( const rex_balance_table::const_iterator& bitr, const asset& rex );
         asset update_rex_account( const name& owner, const asset& proceeds, const asset& unstake_quant, bool force_vote_update = false );
         void channel_to_rex( const name& from, const asset& amount );
         void channel_namebid_to_rex( const int64_t highest_bid );
         template <typename T>
         int64_t rent_rex( T& table, const name& from, const name& receiver, const asset& loan_payment, const asset& loan_fund );
         template <typename T>
         void fund_rex_loan( T& table, const name& from, uint64_t loan_num, const asset& payment );
         template <typename T>
         void defund_rex_loan( T& table, const name& from, uint64_t loan_num, const asset& amount );
         void transfer_from_fund( const name& owner, const asset& amount );
         void transfer_to_fund( const name& owner, const asset& amount );
         bool rex_loans_available()const;
         bool rex_system_initialized()const { return _rexpool.begin() != _rexpool.end(); }
         bool rex_available()const { return rex_system_initialized() && _rexpool.begin()->total_rex.amount > 0; }
         static time_point_sec get_rex_maturity();
         asset add_to_rex_balance( const name& owner, const asset& payment, const asset& rex_received );
         asset add_to_rex_pool( const asset& payment );
         void process_rex_maturities( const rex_balance_table::const_iterator& bitr );
         void consolidate_rex_balance( const rex_balance_table::const_iterator& bitr,
                                       const asset& rex_in_sell_order );
         int64_t read_rex_savings( const rex_balance_table::const_iterator& bitr );
         void put_rex_savings( const rex_balance_table::const_iterator& bitr, int64_t rex );
         void update_rex_stake( const name& voter );

         void add_loan_to_rex_pool( const asset& payment, int64_t rented_tokens, bool new_loan );
         void remove_loan_from_rex_pool( const rex_loan& loan );
         template <typename Index, typename Iterator>
         int64_t update_renewed_loan( Index& idx, const Iterator& itr, int64_t rented_tokens );

         // defined in delegate_bandwidth.cpp
         void changebw( name from, name receiver,
                        asset stake_net_quantity, asset stake_cpu_quantity, bool transfer, bool penaltyprod = false );
         void update_voting_power( const name& voter, const asset& total_update );

         // defined in voting.hpp
         void update_elected_producers( block_timestamp timestamp );
         void update_votes( const name voter, asset quantity, name proxy, const std::vector<name>& producers, bool voting );
         void propagate_weight_change( const voter_info& voter );
         double update_producer_votepay_share( const producers_table2::const_iterator& prod_itr,
                                               time_point ct,
                                               double shares_rate, bool reset_to_zero = false );
         double update_total_votepay_share( time_point ct,
                                            double additional_shares_delta = 0.0, double shares_rate_delta = 0.0 );

         template <auto system_contract::*...Ptrs>
         class registration {
            public:
               template <auto system_contract::*P, auto system_contract::*...Ps>
               struct for_each {
                  template <typename... Args>
                  static constexpr void call( system_contract* this_contract, Args&&... args )
                  {
                     std::invoke( P, this_contract, args... );
                     for_each<Ps...>::call( this_contract, std::forward<Args>(args)... );
                  }
               };
               template <auto system_contract::*P>
               struct for_each<P> {
                  template <typename... Args>
                  static constexpr void call( system_contract* this_contract, Args&&... args )
                  {
                     std::invoke( P, this_contract, std::forward<Args>(args)... );
                  }
               };

               template <typename... Args>
               constexpr void operator() ( Args&&... args )
               {
                  for_each<Ptrs...>::call( this_contract, std::forward<Args>(args)... );
               }

               system_contract* this_contract;
         };

         registration<&system_contract::update_rex_stake> vote_stake_updater{ this };
   };

} /// uosiosystem
