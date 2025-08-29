#include <uosio.system/uosio.system.hpp>

#include <uosiolib/uosio.hpp>
#include <uosiolib/crypto.h>
#include <uosiolib/datastream.hpp>
#include <uosiolib/serialize.hpp>
#include <uosiolib/multi_index.hpp>
#include <uosiolib/privileged.hpp>
#include <uosiolib/singleton.hpp>
#include <uosiolib/transaction.hpp>
#include <uosio.token/uosio.token.hpp>

#include <algorithm>
#include <cmath>

namespace uosiosystem {
   using uosio::indexed_by;
   using uosio::const_mem_fun;
   using uosio::singleton;
   using uosio::transaction;

   /**
    *  This method will create a producer_config and producer_info object for 'producer'
    *
    *  @pre producer is not already registered
    *  @pre producer to register is an account
    *  @pre authority of producer to register
    *
    */
   void system_contract::regproducer( const name producer, const uosio::public_key& producer_key, const std::string& url, const std::string& ulord_addr, uint16_t location ) {
      check( url.size() < 512, "url too long" );
      check( producer_key != uosio::public_key(), "public key should not be the default value" );
      require_auth( producer );
	  uosio_assert( prodbw_verf(producer),"producer delegate is too small"); 
      //uosio_assert(ulord_addr.size() == 34, "ulord address is illegal");
      uosio_assert( ulord_addr.size() < 512, "ulord_addr too long" );
      auto prod = _producers.find( producer.value );
      const auto ct = current_time_point();
	  
	  // **********<begin>This code was modified by camphor 2018-12-10 </begin>*********//
	   // Added verification of delete flag
      if ( prod != _producers.end() ) {
          if(prod->is_remove){
              return;
          }
         _producers.modify( prod, producer, [&]( producer_info& info ){
            info.producer_key = producer_key;
            info.is_active    = true;
			info.is_remove    = false;
            info.release      = 0;
            info.url          = url;
			info.ulord_addr   = ulord_addr;
            info.location     = location;
            if ( info.last_claim_time == time_point() )
               info.last_claim_time = ct;
         });

         //Begin:removed by hml Date:2019-11-13
#if 0
         auto prod2 = _producers2.find( producer.value );
         if ( prod2 == _producers2.end() ) {
            _producers2.emplace( producer, [&]( producer_info2& info ){
               info.owner                     = producer;
               info.last_votepay_share_update = ct;
            });
            update_total_votepay_share( ct, 0.0, prod->total_votes );
            // When introducing the producer2 table row for the first time, the producer's votes must also be accounted for in the global total_producer_votepay_share at the same time.
         }
#endif
        //End:removed by hml Date:2019-11-13

      } else {
         _producers.emplace( producer, [&]( producer_info& info ){
            info.owner           = producer;
            info.total_votes     = 0;
            info.producer_key    = producer_key;
            info.is_active       = true;
			   info.is_remove     = false;
            info.release       = 0;
            info.url             = url;
			   info.ulord_addr    = ulord_addr;
            info.location        = location;
            info.last_claim_time = ct;
         });

     //Begin:removed by hml Date:2019-11-13
#if 0
         _producers2.emplace( producer, [&]( producer_info2& info ){
            info.owner                     = producer;
            info.last_votepay_share_update = ct;
         });
#endif
//End:removed by hml Date:2019-11-13

      }

   }

   void system_contract::unregprod( const name producer, const bool release ) {
	if(release){
          require_auth(_self);
          const auto& prod = _producers.get( producer.value, "producer not found" );
          _producers.modify( prod, same_payer, [&]( producer_info& info ){
              info.release = 1;
              info.deactivate();
          });
      } else{
	      require_auth( producer );

	      const auto& prod = _producers.get( producer.value, "producer not found" );
	      _producers.modify( prod, same_payer, [&]( producer_info& info ){
	         info.deactivate();
	      });
	  }
   }

   void system_contract::update_elected_producers( block_timestamp block_time ) {
      _gstate.last_producer_schedule_update = block_time;

      auto idx = _producers.get_index<"prototalvote"_n>();

      std::vector< std::pair<uosio::producer_key,uint16_t> > top_producers;
	  
	  //Begin: add by hml
   	  std::vector< std::pair<uosiosystem::producer_ulord,uint16_t>> ulord_producers;
      top_producers.reserve(max_producers_number);
	  ulord_producers.reserve(max_producers_number);

      for ( auto it = idx.cbegin(); it != idx.cend() && top_producers.size() < max_producers_number && 0 < it->total_votes && it->active(); ++it ) {
         top_producers.emplace_back( std::pair<uosio::producer_key,uint16_t>({{it->owner, it->producer_key}, it->location}) );
		 ulord_producers.emplace_back( std::pair<uosiosystem::producer_ulord,uint16_t>({{it->owner, it->ulord_addr}, it->location}) );
      }

    //Begin: removed by hml, Date:2019-11-12
#if 0
      if ( top_producers.size() == 0 || top_producers.size() < _gstate.last_producer_schedule_size ) {
         return;
      }
#endif
    //End: removed by hml, Date:2019-11-12

      //Begin:  add by alvin  for  output bp list begin
      uint32_t  time_now = now();
     
      //clear uosc bp list
      //while(_uosc_bp_list.begin() != _uosc_bp_list.end()) {
       //  _uosc_bp_list.erase(_uosc_bp_list.begin());  
      //}
      //clear uosc bp list
      while(_uosc_bp_list.begin() != _uosc_bp_list.end()) 
      {
         auto  deliter=  _uosc_bp_list.begin();
         if(deliter->bp_valid_time< time_now-6 )
             _uosc_bp_list.erase(_uosc_bp_list.begin());  
         else
             break;
      }
       
	  print("\nclear uosc bp list, ++++++!\n");
	
	   int iconunt = 0;

      //one_minute_prod_bp
      for(auto&  pordit : ulord_producers )
      {
         auto prod_name=  pordit.first.producer_name;

         //add uosc bp list 
         if(iconunt < max_producers_number )
         {
             _uosc_bp_list.emplace( _self, [&]( one_minute_bp_list& producer ){
              producer.ulord_addr            = pordit.first.ulord_addr;
              producer.bpname                = prod_name;
              producer.bp_valid_time        = time_now+ 3*120 +  (iconunt++) ;
             });
         }
      } 

      if ( top_producers.size() < _gstate.last_producer_schedule_size ) {
      	if(top_producers.size() < 1){
			return;
      	}
      }

      /// sort by producer name
      std::sort( top_producers.begin(), top_producers.end() );

      std::vector<uosio::producer_key> producers;

      producers.reserve(top_producers.size());
      for( const auto& item : top_producers )
         producers.push_back(item.first);

      auto packed_schedule = pack(producers);

      if( set_proposed_producers( packed_schedule.data(),  packed_schedule.size() ) >= 0 ) {
         _gstate.last_producer_schedule_size = static_cast<decltype(_gstate.last_producer_schedule_size)>( top_producers.size() );
      }
   }

   double stake2vote( int64_t staked ) {
      /// TODO subtract 2080 brings the large numbers closer to this decade
      double weight = int64_t( (now() - (block_timestamp::block_timestamp_epoch / 1000)) / (seconds_per_day * 7) )  / double( 52 );
      return double(staked) * std::pow( 2, weight );
   }

   double system_contract::update_total_votepay_share( time_point ct,
                                                       double additional_shares_delta,
                                                       double shares_rate_delta )
   {
      double delta_total_votepay_share = 0.0;
      if( ct > _gstate3.last_vpay_state_update ) {
         delta_total_votepay_share = _gstate3.total_vpay_share_change_rate
                                       * double( (ct - _gstate3.last_vpay_state_update).count() / 1E6 );
      }

      delta_total_votepay_share += additional_shares_delta;
      if( delta_total_votepay_share < 0 && _gstate2.total_producer_votepay_share < -delta_total_votepay_share ) {
         _gstate2.total_producer_votepay_share = 0.0;
      } else {
         _gstate2.total_producer_votepay_share += delta_total_votepay_share;
      }

      if( shares_rate_delta < 0 && _gstate3.total_vpay_share_change_rate < -shares_rate_delta ) {
         _gstate3.total_vpay_share_change_rate = 0.0;
      } else {
         _gstate3.total_vpay_share_change_rate += shares_rate_delta;
      }

      _gstate3.last_vpay_state_update = ct;

      return _gstate2.total_producer_votepay_share;
   }

   double system_contract::update_producer_votepay_share( const producers_table2::const_iterator& prod_itr,
                                                          time_point ct,
                                                          double shares_rate,
                                                          bool reset_to_zero )
   {
      double delta_votepay_share = 0.0;
      if( shares_rate > 0.0 && ct > prod_itr->last_votepay_share_update ) {
         delta_votepay_share = shares_rate * double( (ct - prod_itr->last_votepay_share_update).count() / 1E6 ); // cannot be negative
      }

      double new_votepay_share = prod_itr->votepay_share + delta_votepay_share;
      _producers2.modify( prod_itr, same_payer, [&](auto& p) {
         if( reset_to_zero )
            p.votepay_share = 0.0;
         else
            p.votepay_share = new_votepay_share;

         p.last_votepay_share_update = ct;
      } );

      return new_votepay_share;
   }

   /**
    *  @pre producers must be sorted from lowest to highest and must be registered and active
    *  @pre if proxy is set then no producers can be voted for
    *  @pre if proxy is set then proxy account must exist and be registered as a proxy
    *  @pre every listed producer or proxy must have been previously registered
    *  @pre voter must authorize this action
    *  @pre voter must have previously staked some amount of CORE_SYMBOL for voting
    *  @pre voter->staked must be up to date
    *
    *  @post every producer previously voted for will have vote reduced by previous vote weight
    *  @post every producer newly voted for will have vote increased by new vote amount
    *  @post prior proxy will proxied_vote_weight decremented by previous vote weight
    *  @post new proxy will proxied_vote_weight incremented by new vote weight
    *
    *  If voting for a proxy, the producer votes will not change until the proxy updates their own vote.
    */
   void system_contract::voteproducer( const name voter_name, asset quantity, const name proxy, const std::vector<name>& producers ) {
      require_auth( voter_name );
	  
	  //Begin: reomved by UOS(hml) Date:2019-11-12
      //vote_stake_updater( voter_name );
	  //End: reomved by UOS(hml) Date:2019-11-12
	  
      update_votes( voter_name, quantity, proxy, producers, true );

    //Begin: reomved by UOS(hml) Date:2019-11-12
#if 0
      auto rex_itr = _rexbalance.find( voter_name.value );
      if( rex_itr != _rexbalance.end() && rex_itr->rex_balance.amount > 0 ) {
         check_voting_requirement( voter_name, "voter holding REX tokens must vote for at least 21 producers or for a proxy" );
      }
#endif
    //End: reomved by UOS(hml) Date:2019-11-12

   }

	
	//Begin: modify by UOS(hml) Date:2018-09-21
	//attention: Here is not deal with the case of  proxy, if add furture, the funtion need deal it.
   void system_contract::cancelvote( const name voter_name, asset quantity, const std::vector<name>& producers ) {
      require_auth( voter_name );

	//verify input
	 auto voter = _voters.find(voter_name.value);
     uosio_assert( voter != _voters.end(), "user have not voted before" );

     if(quantity.amount == 0) {
        cancelallvote(voter_name, false);
        return ;
     }

	 uosio_assert(quantity.amount > 100, "min cancel vote is 0.01");
	 uosio_assert(producers.size() > 0 && producers.size() < max_producers_number, "None of cancel vote producer.");
	 uosio_assert( voter->vote_producers.size() > 0 , "not exist in voted producer list.");
	 
#if UOSIOLIB_DEBUG
	print("before cancelvote, map_vote_producer","\n");
	
	 for(auto vp : voter->vote_producers) {
		 print("prudecer: ", vp.producer, " voted: ", vp.voted.amount, "\n");
	 }
#endif
 
	 for( size_t i = 1; i < producers.size(); ++i ) {
        uosio_assert( producers[i-1] < producers[i], "producer votes must be unique and sorted" );
     }

	 bool isFind;


	/*
		1.check all producers in vote_producers, if not, failed.
		2.check all producers have enough voted to cancel, if not, failed.
		3.modify all producers by inputted para of producers.
		4.if zero of voted, clear producer from vector.
	 */

	 for(auto p : producers) {
		isFind = false;
		for(auto vp = voter->vote_producers.begin(); vp != voter->vote_producers.end();) {
			if(vp->producer == p) {
				uosio_assert(vp->voted.amount - quantity.amount >= 0, "amount of cancel vote is greater than voted");

				isFind = true;

#if UOSIOLIB_DEBUG
				print("voted_producer->second=",vp->voted.amount," quantity.amount=",quantity.amount, "\n" );
#endif

				 _voters.modify( voter, same_payer, [&]( auto& av ) {
				 		bool isfind = false;
				 		for(auto it = av.vote_producers.begin(); it != av.vote_producers.end(); ++it) {
							if(p == it->producer) {
								isfind = true;
								it->voted.amount -= quantity.amount;
							}
						} 	
						uosio_assert( isfind, "not exist in voted producer list.");
				 	});

					
				auto list_producer = _producers.find(p.value);
				uosio_assert( list_producer != _producers.end(), "not exist in voted producer list.");
				
				auto new_vote_weight =  quantity.amount;
#if UOSIOLIB_DEBUG
				print("new_vote_weight=",new_vote_weight," list_producer->total_votes=",list_producer->total_votes,"\n" );
#endif

				uosio_assert(list_producer->total_votes - new_vote_weight >= 0, "amount of cancel voted is exceed.");

				_producers.modify( list_producer, same_payer, [&]( auto& producer ) {
		               producer.total_votes -= new_vote_weight;
		               _gstate.total_producer_vote_weight -= new_vote_weight;
		               
           		 });		

				if(vp->voted.amount  == 0) {
					 _voters.modify( voter, same_payer, [&]( auto& av ) {
	         				vp = av.vote_producers.erase(vp);	
					 	});
				}
			
				break;			
			}
			else {
					++vp;
			}
	 	}

		if(!isFind) {
			uosio_assert(false, "The cancel vote list include not existed on voted producer list.");
		}
	 }
	 
#if UOSIOLIB_DEBUG
		print("End cancelvote, map_vote_producer","\n");
		for(auto vp : voter->vote_producers) {
			print("prudecer: ", vp.producer, " voted: ", vp.voted.amount, "\n");
		}
#endif	

   }
   //End: modify by UOS(hml) Date:2018-09-21

   
  //Begin: add by UOS(hml) Date:2018-09-21
void system_contract::cancelallvote( const name voter_name, bool is_sudo) {

    if(is_sudo) {
        require_auth( "uosio"_n );
    }else {
        require_auth(voter_name);
    }

	auto voter = _voters.find(voter_name.value);
    uosio_assert( voter != _voters.end(), "user have not voted before" );

	int64_t new_vote_weight = 0;
	
	for(auto vp = voter->vote_producers.begin(); vp != voter->vote_producers.end(); vp++) {
		
		auto p = _producers.find(vp->producer.value);
		uosio_assert( p != _producers.end(), "not exist in voted producer list.");
		
		new_vote_weight = vp->voted.amount;
		uosio_assert(p->total_votes - new_vote_weight >= 0, "amount of cancel voted is exceed.");

		_producers.modify( p, same_payer, [&]( auto& producer ) {
		               producer.total_votes -= new_vote_weight;
		               _gstate.total_producer_vote_weight -= new_vote_weight;
           		 });	
	}

	_voters.modify( voter, same_payer, [&]( auto& av ) {
	         				 av.vote_producers.clear();	
					 	});
}
//End: add by UOS(hml) Date:2018-09-21

    void system_contract::validate_votes(const name voter_name, asset quantity) {

		//print(" voter_name= ",name{voter_name}, " quantity=", quantity.amount, "\n");
		
		auto voter = _voters.find(voter_name.value);
		//uosio_assert( voter != _voters.end(), "user must stake before they can vote" );
        if(voter == _voters.end()) {
            return ;
        }

		asset total_voted(0, symbol{"UOS", 4});
		
	  	for(auto vp = voter->vote_producers.begin(); vp != voter->vote_producers.end(); ++vp) 
		{
			total_voted += vp->voted;
		}

		//print(" voter->staked ", voter->staked, " total_voted.amount =", total_voted.amount, "quantity.amount =", quantity.amount, "\n");

		//attention:  voter->staked was changed before excute here, so ignore para of quantity.amount 
		uosio_assert( voter->staked - total_voted.amount /*+ quantity.amount*/ >= 0, "user must cancel vote producer before undelegate." );
   }


   void system_contract::update_votes( const name voter_name, asset quantity, name proxy, const std::vector<name>& producers, bool voting ) {
	//Begin: modify by hml Date:2019-11-12
#if 0
      //validate input
      if ( proxy ) {
         check( producers.size() == 0, "cannot vote for producers and proxy at same time" );
         check( voter_name != proxy, "cannot proxy to self" );
      } else {
         check( producers.size() <= 30, "attempt to vote for too many producers" );
         for( size_t i = 1; i < producers.size(); ++i ) {
            check( producers[i-1] < producers[i], "producer votes must be unique and sorted" );
         }
      }
#endif
      proxy.value = 0;

      auto voter = _voters.find( voter_name.value );
      check( voter != _voters.end(), "user must stake before they can vote" ); /// staking creates voter object
      uosio_assert( producers.size() + voter->vote_producers.size() <= max_producers_number, "attempt to vote for too many producers" );
      //check( !proxy || !voter->is_proxy, "account registered as a proxy is not allowed to use a proxy" );

      /**
       * The first time someone votes we calculate and set last_vote_weight, since they cannot unstake until
       * after total_activated_stake hits threshold, we can use last_vote_weight to determine that this is
       * their first vote and should consider their stake activated.
       */
	   
	      auto new_vote_weight = quantity.amount  * producers.size();
	  _gstate.total_activated_stake +=	new_vote_weight;


	  //get start time of mainnet
	  if( _gstate.total_activated_stake >= min_activated_stake && _gstate.thresh_activated_stake_time == time_point() ) {
            _gstate.thresh_activated_stake_time = current_time_point();
            //<begin> modified by uos camphor
            action(permission_level{_self, active_permission}, bvpay_account,
                     "start"_n, _self)
                     .send();
            //<end> modified by uos camphor
      }

         //check voted is less than staked
        if(voting ) {
            int64_t total_voted = 0;
            for(auto vp = voter->vote_producers.begin(); vp != voter->vote_producers.end(); ++vp) 
            {
                total_voted += vp->voted.amount;
            }
            
            print(" staked =", voter->staked, "total_voted=", total_voted, " amount = ", quantity.amount, "size=",voter->vote_producers.size(),"\n");
            uosio_assert(voter->staked - (total_voted + new_vote_weight) >= 0, "not enough stake asset.");
			
        }
	

      boost::container::flat_map<name, pair<double, bool /*new*/> > producer_deltas;

      if( new_vote_weight >= 0 ) {
         for( const auto& p : producers ) {
            auto& d = producer_deltas[p];
            d.first += quantity.amount;
            d.second = true;
         }
      }

      const auto ct = current_time_point();
      double delta_change_rate         = 0.0;
      double total_inactive_vpay_share = 0.0;

      for( const auto& pd : producer_deltas ) {
         auto pitr = _producers.find( pd.first.value );
         if( pitr != _producers.end() ) {
            check( !voting || pitr->active() || !pd.second.second /* not from new set */, "producer is not currently registered" );
            double init_total_votes = pitr->total_votes;
            _producers.modify( pitr, same_payer, [&]( auto& p ) {
               p.total_votes += pd.second.first;
               if ( p.total_votes < 0 ) { // floating point arithmetics can give small negative numbers
                  p.total_votes = 0;
               }
               _gstate.total_producer_vote_weight += pd.second.first;
               //check( p.total_votes >= 0, "something bad happened" );
            });
        
        //Begin: removed by hml Date:2019-11-13
        #if 0
            auto prod2 = _producers2.find( pd.first.value );
            if( prod2 != _producers2.end() ) {
               const auto last_claim_plus_3days = pitr->last_claim_time + microseconds(3 * useconds_per_day);
               bool crossed_threshold       = (last_claim_plus_3days <= ct);
               bool updated_after_threshold = (last_claim_plus_3days <= prod2->last_votepay_share_update);
               // Note: updated_after_threshold implies cross_threshold

               double new_votepay_share = update_producer_votepay_share( prod2,
                                             ct,
                                             updated_after_threshold ? 0.0 : init_total_votes,
                                             crossed_threshold && !updated_after_threshold // only reset votepay_share once after threshold
                                          );

               if( !crossed_threshold ) {
                  delta_change_rate += pd.second.first;
               } else if( !updated_after_threshold ) {
                  total_inactive_vpay_share += new_votepay_share;
                  delta_change_rate -= init_total_votes;
               }
            }

        #endif
        //End: removed by hml Date:2019-11-13

         } else {
            check( !pd.second.second /* not from new set */, "producer is not registered" ); //data corruption
         }
      }

        //update all voted producer, if not exist,it will ceate new and insert vector.
        vote_producer newvp;
        bool isFind = false;
        
        for(auto p : producers) {
            if(voter->vote_producers.size() > 0) {
                isFind = false;
            
                _voters.modify( voter, same_payer, [&]( auto& av ) {
                    for(auto vp = av.vote_producers.begin(); vp != av.vote_producers.end(); ++vp) {
    #if UOSIOLIB_DEBUG
                        print("p: ", p, " vp.producer: ", vp->producer, "\n");
    #endif
                        if(vp->producer == p) {
                            vp->voted.amount += quantity.amount;
                            isFind = true;
                            break;		
                        }
                    }

                    if(!isFind) {
                        newvp.producer = p;
                        newvp.voted = quantity;
                        av.vote_producers.push_back(newvp);
                    }
                });	
            }
            else {
                newvp.producer = p;
                newvp.voted = quantity;
                _voters.modify( voter, same_payer, [&]( auto& av ) {
                    av.vote_producers.push_back(newvp);
                    });
            }
        }

        #if UOSIOLIB_DEBUG	  
            print("End update_votes, vote_producers:","\n");
            for(auto vp : voter->vote_producers) {
                print("prudecer: ", vp.producer, " voted: ", vp.voted.amount, "\n");
            }
        #endif

    //Begin: removed by hml Date:2019-11-13
	#if 0   
      update_total_votepay_share( ct, -total_inactive_vpay_share, delta_change_rate );

      _voters.modify( voter, same_payer, [&]( auto& av ) {
         av.last_vote_weight = new_vote_weight;
         av.producers = producers;
         av.proxy     = proxy;
      });
    #endif
    //End: removed by hml Date:2019-11-13

        //End: Add by UOS(hml) Date:2018-09-20
    }

   /**
    *  An account marked as a proxy can vote with the weight of other accounts which
    *  have selected it as a proxy. Other accounts must refresh their voteproducer to
    *  update the proxy's weight.
    *
    *  @param isproxy - true if proxy wishes to vote on behalf of others, false otherwise
    *  @pre proxy must have something staked (existing row in voters table)
    *  @pre new state must be different than current state
    */
   void system_contract::regproxy( const name proxy, bool isproxy ) {
      require_auth( proxy );

   //Begin: removed by hml Date:2019-11-20
   #if 0

      auto pitr = _voters.find( proxy.value );
      if ( pitr != _voters.end() ) {
         check( isproxy != pitr->is_proxy, "action has no effect" );
         check( !isproxy || !pitr->proxy, "account that uses a proxy is not allowed to become a proxy" );
         _voters.modify( pitr, same_payer, [&]( auto& p ) {
               p.is_proxy = isproxy;
            });
         propagate_weight_change( *pitr );
      } else {
         _voters.emplace( proxy, [&]( auto& p ) {
               p.owner  = proxy;
               p.is_proxy = isproxy;
            });
      }
   #endif
   //Begin: removed by hml Date:2019-11-20

   }

   void system_contract::propagate_weight_change( const voter_info& voter ) {
      check( !voter.proxy || !voter.is_proxy, "account registered as a proxy is not allowed to use a proxy" );
      	
      //Begin: removed by UOS(hml) Date:2018-09-20
#if 0
      double new_weight = stake2vote( voter.staked );
      if ( voter.is_proxy ) {
         new_weight += voter.proxied_vote_weight;
      }

      /// don't propagate small changes (1 ~= epsilon)
      if ( fabs( new_weight - voter.last_vote_weight ) > 1 )  {
         if ( voter.proxy ) {
            auto& proxy = _voters.get( voter.proxy.value, "proxy not found" ); //data corruption
            _voters.modify( proxy, same_payer, [&]( auto& p ) {
                  p.proxied_vote_weight += new_weight - voter.last_vote_weight;
               }
            );
            propagate_weight_change( proxy );
         } else {
            auto delta = new_weight - voter.last_vote_weight;
            const auto ct = current_time_point();
            double delta_change_rate         = 0;
            double total_inactive_vpay_share = 0;
            for ( auto acnt : voter.producers ) {
               auto& prod = _producers.get( acnt.value, "producer not found" ); //data corruption
               const double init_total_votes = prod.total_votes;
               _producers.modify( prod, same_payer, [&]( auto& p ) {
                  p.total_votes += delta;
                  _gstate.total_producer_vote_weight += delta;
               });
               auto prod2 = _producers2.find( acnt.value );
               if ( prod2 != _producers2.end() ) {
                  const auto last_claim_plus_3days = prod.last_claim_time + microseconds(3 * useconds_per_day);
                  bool crossed_threshold       = (last_claim_plus_3days <= ct);
                  bool updated_after_threshold = (last_claim_plus_3days <= prod2->last_votepay_share_update);
                  // Note: updated_after_threshold implies cross_threshold

                  double new_votepay_share = update_producer_votepay_share( prod2,
                                                ct,
                                                updated_after_threshold ? 0.0 : init_total_votes,
                                                crossed_threshold && !updated_after_threshold // only reset votepay_share once after threshold
                                             );

                  if( !crossed_threshold ) {
                     delta_change_rate += delta;
                  } else if( !updated_after_threshold ) {
                     total_inactive_vpay_share += new_votepay_share;
                     delta_change_rate -= init_total_votes;
                  }
               }
            }

            update_total_votepay_share( ct, -total_inactive_vpay_share, delta_change_rate );
         }
      }
      _voters.modify( voter, same_payer, [&]( auto& v ) {
            v.last_vote_weight = new_weight;
         }
      );

      #endif
	//End: removed by UOS(hml) Date:2018-09-20
   }

} /// namespace uosiosystem
