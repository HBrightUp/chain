#include <uosio.system/uosio.system.hpp>
#include <uosio.token/uosio.token.hpp>
#include<uosio.bvpay/uosio.bvpay.hpp>

namespace uosiosystem {

   //Begin: modify by alvin
   const int64_t  min_pervote_daily_pay = 20'0000;
   //End: modify by alvin
   
   //Begin: modify by UOS(hml) Date:2018-09-28
   const int64_t  	min_activated_stake   =  300'0000000;
   //End: modify by UOS(hml) Date:2018-09-28
   
   const double   continuous_rate       = 0.04879;          // 5% annual rate
   const double   perblock_rate         = 0.0025;           // 0.25%
   const double   standby_rate          = 0.0075;           // 0.75%
   const uint32_t blocks_per_year       = 52*7*24*2*3600;   // half seconds per year
   const uint32_t seconds_per_year      = 52*7*24*3600;
   
   //Begin: modify by lx
   const uint32_t blocks_per_day        = 24 * 3600;
   const uint32_t blocks_per_hour       = 3600;
   //Begin: modify by lx
   
   const int64_t  useconds_per_day      = 24 * 3600 * int64_t(1000000);
   const int64_t  useconds_per_year     = seconds_per_year*1000000ll;

   void system_contract::onblock( ignore<block_header> ) {
      using namespace uosio;

      require_auth(_self);

      block_timestamp timestamp;
      name producer;
      _ds >> timestamp >> producer;

      // _gstate2.last_block_num is not used anywhere in the system contract code anymore.
      // Although this field is deprecated, we will continue updating it for now until the last_block_num field
      // is eventually completely removed, at which point this line can be removed.
      _gstate2.last_block_num = timestamp;

      /** until activated stake crosses this threshold no new rewards are paid */
      if( _gstate.total_activated_stake < min_activated_stake )
         return;

      if( _gstate.last_pervote_bucket_fill == time_point() )  /// start the presses
         _gstate.last_pervote_bucket_fill = current_time_point();


      /**
       * At startup the initial producer may not be one that is registered / elected
       * and therefore there may be no producer object for them.
       */
      auto prod = _producers.find( producer.value );
      if ( prod != _producers.end() ) {
         _gstate.total_unpaid_blocks++;
         _producers.modify( prod, same_payer, [&](auto& p ) {
               p.unpaid_blocks++;
         });
      }

      /// only update block producers once every minute, block_timestamp is in half seconds
      if( timestamp.slot - _gstate.last_producer_schedule_update.slot > 120 ) {
         update_elected_producers( timestamp );

#if 0    //for mainchain
         if( (timestamp.slot - _gstate.last_name_close.slot) > blocks_per_day ) {
            name_bid_table bids(get_self(), get_self().value);
            auto idx = bids.get_index<"highbid"_n>();
            auto highest = idx.lower_bound( std::numeric_limits<uint64_t>::max()/2 );
            if( highest != idx.end() &&
                highest->high_bid > 0 &&
                (current_time_point() - highest->last_bid_time) > microseconds(useconds_per_day) &&
                _gstate.thresh_activated_stake_time > time_point() &&
                (current_time_point() - _gstate.thresh_activated_stake_time) > microseconds(5 * 7 * useconds_per_day)
            ) {
               _gstate.last_name_close = timestamp;
               channel_namebid_to_rex( highest->high_bid );
               idx.modify( highest, same_payer, [&]( auto& b ){
                  b.high_bid = -b.high_bid;
               });
            }
         }
#endif

#if 1       //for test chain
            const uint32_t blocks_per_day_testchain =  10 * 60;
            const int64_t  useconds_per_day_testchain      = 1 * 60 * int64_t(1000000);


            if( (timestamp.slot - _gstate.last_name_close.slot) > blocks_per_day_testchain ) {
            name_bid_table bids(get_self(), get_self().value);
            auto idx = bids.get_index<"highbid"_n>();
             // **********<begin>This code was modified by camphor 2019-03-14 </begin>*********//
            auto highest = idx.lower_bound( std::numeric_limits<uint64_t>::max()/2 );
             // **********<end>This code was modified by camphor 2019-03-14 </end>*********//
            if( highest != idx.end() &&
                highest->high_bid > 0 &&
                (current_time_point() - highest->last_bid_time) > microseconds(useconds_per_day_testchain) &&
                _gstate.thresh_activated_stake_time > time_point() &&
                (current_time_point() - _gstate.thresh_activated_stake_time) > /*5 * 7 * useconds_per_day*/microseconds(100) ) {
                   _gstate.last_name_close = timestamp;
                   idx.modify( highest, same_payer, [&]( auto& b ){
                         b.high_bid = -b.high_bid;
               });
            }
         }
#endif
      }
   }

   using namespace uosio;

   // issue token to bp and udfs_system  add by alvin
   void system_contract::issuetoken(  ){
   /*

      auto fielditer = _sys_args_list.find(N(issuetime)); 
      uint32_t  time_now =  now();
      if(fielditer == _sys_args_list.end()) {
        setsysargs(N(issuetime),time_now-17*seconds_per_day,0);
        setsysargs(N(udfssupply),2280000000,0);
        setsysargs(N(bpsupply),120000000,0);
        return;
      }
      auto interval = seconds_per_day;
      //auto interval = 60;
      
      if(time_now > fielditer->val+interval)
      {
         //get account udfsreceiver 	 
         if(!is_account(N(udfsreceiver))){
            print( "udfs : ",  time_now, "   no  account: udfsreceiver  ",  "\n");
            return;
         } 
         accounts accountstable( N(uosio.token),N(udfsreceiver) );
         const auto& ac = accountstable.get(uosio::symbol_type(system_token_symbol).name());
         auto udfsiter = _sys_args_list.find(N(udfssupply));
         auto bpiter = _sys_args_list.find(N(bpsupply));
         
         int64_t  udfsfee=0;
         if( ac.balance.amount< udfsiter->val )       
         {
             udfsfee =  udfsiter->val - ac.balance.amount; 
         }
         print( "udfs : ",  udfsiter->val , "  ", time_now, "  ", bpiter->val," udfsfee ", udfsfee,  "\n");
           
         auto  total= udfsfee+ bpiter->val; 

         INLINE_ACTION_SENDER(uosio::token, issue)( N(uosio.token), {{N(uosio),N(active)}},
                                                   {N(uosio), asset(total), std::string("issue tokens for producer pay and udfs")} );

         INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), {N(uosio),N(active)},
                                                       { N(uosio), N(uosio.bvpay), asset(bpiter->val,system_token_symbol), "transfer bvpay" } );
         if(udfsfee>0)
         {
            INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), {N(uosio),N(active)},
                                                       { N(uosio), N(udfsreceiver), asset(udfsfee,system_token_symbol), "udfs fee" } );
         }
 
         setsysargs(N(issuetime),fielditer->val+interval,0);

      }
      else {
         return; 
      }
 */

   }

  void system_contract::setglobalargs(const name& key, const uint64_t& val) {


       require_auth( _self );
       if(key == "maxtrancpu"_n){  // max_transaction_cpu_usage
          _gstate.max_transaction_cpu_usage= val;
          _global.set( _gstate, _self );
       }

  }

  void system_contract::setsysargs(const name& key, const uint64_t& val, uint8_t global) {

      require_auth( _self );
 
      if(global==1)
         return setglobalargs(key,val); 
	  auto fielditer = _sys_args_list.find(key.value);
  
	  //add or modify data
	  if(fielditer == _sys_args_list.end()) {
		  _sys_args_list.emplace( _self, [&] (auto& g) {
			  g.key = key;
			  g.val = val;
		  }); 
	  }
	  else {
		  _sys_args_list.modify(fielditer, _self, [&](auto& g){
			  g.val = val;
		  }); 
	  }
  }

   //Begin: add by hml Date:2019-11-18
  asset system_contract::get_bpvtpay(name call) const {
        require_auth(call);
        auto bpvtp = _bpvtpay.find(0);
        auto bpvtp1 = _bpvtpay.find(1);
        auto realpaydb = _bvrealpay.find(_self.value);
        int64_t real_pay = 0;
        if(bpvtp == _bpvtpay.end()){
            if(realpaydb != _bvrealpay.end()){
                return asset(realpaydb->quantity.amount + real_pay, symbol{SYSTEM_COIN, 4});
            }
            return asset(0, symbol{SYSTEM_COIN, 4});
        }

        uosio_assert(now() >  bpvtp->laster_line , "get_bpvtpay time error");
        if(now() < bpvtp->dead_line){
            double pay = (now() - bpvtp->laster_line ) *  bpvtp->quantity.amount * 1.0 / (bpvtp->dead_line - bpvtp->laster_line);
            real_pay = int64_t(pay);
            real_pay += bpvtp->inherit.amount;
        } else if( (now() >= bpvtp->dead_line) && (now() < bpvtp1->dead_line) ){
            real_pay = bpvtp->quantity.amount;
            real_pay += bpvtp->inherit.amount;
            uosio_assert(now() >  bpvtp1->laster_line , "get_bpvtpay time error");
            double pay = (now() - bpvtp1->laster_line ) *  bpvtp1->quantity.amount * 1.0 / (bpvtp1->dead_line - bpvtp1->laster_line);
            int64_t next_pay = int64_t(pay);
            real_pay += next_pay;
        }  else if(now() >= bpvtp1->dead_line){
            real_pay = bpvtp->quantity.amount;
            real_pay += bpvtp->inherit.amount;
            real_pay += bpvtp1->quantity.amount;
        }
        return asset(realpaydb->quantity.amount + real_pay, symbol{SYSTEM_COIN, 4});
    }

    //End: add by hml Date:2019-11-18

// add by alvin end
   void system_contract::claimrewards( const name owner ) {
      require_auth(owner);

      const auto& prod = _producers.get( owner.value );
      check( prod.active(), "producer does not have an active key" );

      check( _gstate.total_activated_stake >= min_activated_stake,
                    "cannot claim rewards until the chain is activated (at least 15% of all tokens participate in voting)" );

      const auto ct = current_time_point();

      check( ct - prod.last_claim_time > microseconds(useconds_per_day), "already claimed rewards within past day" );

      //const asset token_bvpay  = uosio::uosbvpay(N(uosio.bvpay)).get_bpvtpay(owner);
      //datastream<const char*> ds = datastream<const char*>(nullptr, 0);
      //const asset token_bvpay  = uosio::uosbvpay("uosio.bvpay"_n, "uosio.bvpay"_n, ds).get_bpvtpay(owner);
      const asset token_bvpay = get_bpvtpay(owner);

      uosio_assert(token_bvpay.amount >= 0 ,"No useful money in the pool of funds");
      const auto usecs_since_last_fill = (ct - _gstate.last_pervote_bucket_fill).count();

      if( usecs_since_last_fill > 0 && _gstate.last_pervote_bucket_fill > time_point() && token_bvpay.amount > 5 ) {
          
         int64_t to_producers = token_bvpay.amount;
         auto  to_per_vote_pay  = 2 * (to_producers / 5);
         auto  to_per_block_pay = to_producers - to_per_vote_pay;

         auto t_data1 = std::make_tuple(owner);
         action(permission_level{owner, active_permission}, bvpay_account,
               "fillbvpay"_n, t_data1)
                  .send();

         auto t_data2 = std::make_tuple( bvpay_account, bpay_account, asset(to_per_block_pay, core_symbol()), std::string("fund per-block bucket"));
         action(permission_level{bvpay_account, active_permission}, token_account,
               "transfer"_n, t_data2)
                  .send();
		 
         auto t_data3 = std::make_tuple( bvpay_account, vpay_account, asset(to_per_vote_pay, core_symbol()), std::string("fund per-vote bucket"));
         action(permission_level{bvpay_account, active_permission}, token_account,
               "transfer"_n, t_data3)
                  .send();
		 
         _gstate.pervote_bucket          += to_per_vote_pay;
         _gstate.perblock_bucket         += to_per_block_pay;
         _gstate.last_pervote_bucket_fill = ct;
      }

     //Begin: removed by hml Date:2019-11-13   
#if 0
      auto prod2 = _producers2.find( owner.value );

      /// New metric to be used in pervote pay calculation. Instead of vote weight ratio, we combine vote weight and
      /// time duration the vote weight has been held into one metric.
      const auto last_claim_plus_3days = prod.last_claim_time + microseconds(3 * useconds_per_day);

      bool crossed_threshold       = (last_claim_plus_3days <= ct);
      bool updated_after_threshold = true;
      if ( prod2 != _producers2.end() ) {
         updated_after_threshold = (last_claim_plus_3days <= prod2->last_votepay_share_update);
      } else {
         prod2 = _producers2.emplace( owner, [&]( producer_info2& info  ) {
            info.owner                     = owner;
            info.last_votepay_share_update = ct;
         });
      }
#endif
     //End: removed by hml Date:2019-11-13 

      // Note: updated_after_threshold implies cross_threshold (except if claiming rewards when the producers2 table row did not exist).
      // The exception leads to updated_after_threshold to be treated as true regardless of whether the threshold was crossed.
      // This is okay because in this case the producer will not get paid anything either way.
      // In fact it is desired behavior because the producers votes need to be counted in the global total_producer_votepay_share for the first time.

      int64_t producer_per_block_pay = 0;
      if( _gstate.total_unpaid_blocks > 0 ) {
         producer_per_block_pay = (_gstate.perblock_bucket * prod.unpaid_blocks) / _gstate.total_unpaid_blocks;
      }

     //Begin: removed by hml Date:2019-11-13   
#if 0
      double new_votepay_share = update_producer_votepay_share( prod2,
                                    ct,
                                    updated_after_threshold ? 0.0 : prod.total_votes,
                                    true // reset votepay_share to zero after updating
                                 );
#endif
   //End: removed by hml Date:2019-11-13 

      int64_t producer_per_vote_pay = 0;

     //Begin: removed by hml Date:2019-11-13 
    #if 0
      if( _gstate2.revision > 0 ) {
         double total_votepay_share = update_total_votepay_share( ct );
         if( total_votepay_share > 0 && !crossed_threshold ) {
            producer_per_vote_pay = int64_t((new_votepay_share * _gstate.pervote_bucket) / total_votepay_share);
            if( producer_per_vote_pay > _gstate.pervote_bucket )
               producer_per_vote_pay = _gstate.pervote_bucket;
         }
      } else {
         if( _gstate.total_producer_vote_weight > 0 ) {
            producer_per_vote_pay = int64_t((_gstate.pervote_bucket * prod.total_votes) / _gstate.total_producer_vote_weight);
         }
      }
    #endif
     //End: removed by hml Date:2019-11-13 

      if( producer_per_vote_pay < min_pervote_daily_pay ) {
         producer_per_vote_pay = 0;
      }

      _gstate.pervote_bucket      -= producer_per_vote_pay;
      _gstate.perblock_bucket     -= producer_per_block_pay;
      _gstate.total_unpaid_blocks -= prod.unpaid_blocks;
		
		 //Begin: removed by hml Date:2019-11-13 
      //update_total_votepay_share( ct, -new_votepay_share, (updated_after_threshold ? prod.total_votes : 0.0) );
		//End: removed by hml Date:2019-11-13 
		
      _producers.modify( prod, same_payer, [&](auto& p) {
         p.last_claim_time = ct;
         p.unpaid_blocks   = 0;
      });

      if( producer_per_block_pay > 0 ) {
         INLINE_ACTION_SENDER(uosio::token, transfer)(
            token_account, { {bpay_account, active_permission}, {owner, active_permission} },
            { bpay_account, owner, asset(producer_per_block_pay, core_symbol()), std::string("producer block pay") }
         );
      }
      if( producer_per_vote_pay > 0 ) {
         INLINE_ACTION_SENDER(uosio::token, transfer)(
            token_account, { {vpay_account, active_permission}, {owner, active_permission} },
            { vpay_account, owner, asset(producer_per_vote_pay, core_symbol()), std::string("producer vote pay") }
         );
      }
   }

	 //Begin: modify by UOS(hml) Date:2018-10-29
	void system_contract::payfee( const name payer , const uint64_t net_usage , const uint64_t virtule_net_limit , const name act_account , const name act_name) {
		#if 0
		   using namespace uosio;
		   uosio_assert(int64_t(net_usage)>0,"net_usage overflow");
		   print( "payer :	", name{payer},"\n" );
		   print( "net_usage : ", net_usage,"\n");
		   print( "virtule_net_limit : ", virtule_net_limit,"\n");
		   print( "act_account : ", name{act_account},"\n");
		   print( "act_name : ", name{act_name},"\n");



		   INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), {payer,N(active)},
					{ payer, N(uosio.league), asset(int64_t(net_usage) *100), "transaction fee" } );

		#endif
		//		 if( _gstate.total_activated_stake < min_activated_stake){
		//			 return;
		//		 }


	}
	   //End: modify by UOS(hml) Date:2018-10-29

} //namespace uosiosystem
