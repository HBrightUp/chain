/**
 *  @file
 *  @copyright defined in uos/LICENSE.txt
 */
#include "uosio.system.hpp"

#include <uosiolib/uosio.hpp>
#include <uosiolib/crypto.h>
#include <uosiolib/print.hpp>
#include <uosiolib/datastream.hpp>
#include <uosiolib/serialize.hpp>
#include <uosiolib/multi_index.hpp>
#include <uosiolib/privileged.hpp>
#include <uosiolib/singleton.hpp>
#include <uosiolib/transaction.hpp>
#include <uosio.token/uosio.token.hpp>
#include <uosiolib/action.h>

#include <algorithm>
#include <cmath>

namespace uosiosystem {
   using uosio::indexed_by;
   using uosio::const_mem_fun;
   using uosio::bytes;
   using uosio::print;
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
   void system_contract::regproducer( const account_name producer, const uosio::public_key& producer_key, const std::string& url, const std::string& ulord_addr, uint16_t location ) {
      uosio_assert( url.size() < 512, "url too long" );
      uosio_assert( producer_key != uosio::public_key(), "public key should not be the default value" );
      require_auth( producer );
      uosio_assert( prodbw_verf(producer),"producer delegate is too small"); 
      //uosio_assert(ulord_addr.size() == 34, "ulord address is illegal");
      uosio_assert( ulord_addr.size() < 512, "ulord_addr too long" );
      auto prod = _producers.find( producer );

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
            });
      } else {
         _producers.emplace( producer, [&]( producer_info& info ){
               info.owner         = producer;
               info.total_votes   = 0;
               info.producer_key  = producer_key;
               info.is_active     = true;
               info.is_remove     = false;
               info.release       = 0;
               info.url           = url;
               info.ulord_addr    = ulord_addr;
               info.location      = location;
         });
      }

	   // **********<end>This code was modified by camphor 2018-12-10 </end>*********//
   }

   void system_contract::unregprod( const account_name producer , bool release) {
      if(release){
          require_auth(_self);
          const auto& prod = _producers.get( producer, "producer not found" );
          _producers.modify( prod, 0, [&]( producer_info& info ){
              info.release = 1;
              info.deactivate();
          });
      } else{
          require_auth( producer );
          const auto& prod = _producers.get( producer, "producer not found" );
          _producers.modify( prod, 0, [&]( producer_info& info ){
              info.deactivate();
          });
      }

   }

   void system_contract::update_elected_producers( block_timestamp block_time ) {
      
       
       print("\n update_elected_producers+++++++++++++++++++++++++++++             \n");
       
       _gstate.last_producer_schedule_update = block_time;

      auto idx = _producers.get_index<N(prototalvote)>();

      std::vector< std::pair<uosio::producer_key,uint16_t> > top_producers;
      std::vector< std::pair<uosio::producer_ulord,uint16_t>> ulord_producers;
      top_producers.reserve(21);
      ulord_producers.reserve(21);

      for ( auto it = idx.cbegin(); it != idx.cend() && top_producers.size() < 21 && 0 < it->total_votes && it->active(); ++it ) {
         top_producers.emplace_back( std::pair<uosio::producer_key,uint16_t>({{it->owner, it->producer_key}, it->location}) );
         ulord_producers.emplace_back( std::pair<uosio::producer_ulord,uint16_t>({{it->owner, it->ulord_addr}, it->location}) );
	}

//  add by alvin  for  output bp list begin

     auto bp_global_var_itr = _bp_global_var.begin();
     if( bp_global_var_itr ==  _bp_global_var.end())
     {
         _bp_global_var.emplace(_self, [&](auto& g){
              g.id = MIN_BP_NUM;
              g.val = 0;
          });
                 
         _bp_global_var.emplace(_self, [&](auto& g){
              g.id = MIN_BP_OUT_NUM;
              g.val = 0;
          });
     }

      uint32_t  time_now = now();
      uint32_t  sec= time_now%120;
      time min_now = 0;
      if(sec>60) min_now=time_now+ (120-sec);
      else  min_now=time_now-sec;          // 取整秒 120s
     
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

      //更新 更新one_minute_prod_bp里面的元素
      for(auto&  pordit : ulord_producers )
      {
         auto prod_name=  pordit.first.producer_name;

         auto min_prod =  _minprodtable.find(prod_name);

         if( min_prod != _minprodtable.end())
         {
             _minprodtable.modify( min_prod, _self, [&]( auto& producer ) {
		producer.ulord_addr           = pordit.first.ulord_addr;
                producer.uptime = time_now;
                producer.on_line_time+=120; // 120 second
                
             });                     
             prints("\n modify  "); printn(prod_name); prints("  ");  printn(_self); prints(" \n"); 
         }
         else
         {
            //print("\n emplace ", name(prod_name).tostring(),   " self ", name(_self).tostring(),   "\n");  
            prints("\n emplace  "); printn(prod_name);prints("  ");  printn(_self); prints(" \n"); 
            _minprodtable.emplace( _self, [&]( one_minute_prod_bp& producer ){
               producer.ulord_addr           = pordit.first.ulord_addr;
	       	   producer.uptime               = time_now;
               producer.bpname               = prod_name;
               producer.on_line_time         = 120;

            });
	
            auto globalvars_itr = _bp_global_var.find(MIN_BP_NUM);
            _bp_global_var.modify(globalvars_itr, _self, [&](auto& g){
                g.val++;
            });
            
         }

         //add uosc bp list 
         if(iconunt<21 )
         {
             _uosc_bp_list.emplace( _self, [&]( one_minute_bp_list& producer ){
              producer.ulord_addr            = pordit.first.ulord_addr;
              producer.bpname                = prod_name;
              producer.bp_valid_time        = time_now+ 3*120 +  (iconunt++) ;
             });
         }
      } 
      auto idx_time = _minprodtable.get_index<N(byontime)>();
      
      auto iter = idx_time.rbegin(); // 找一个 在线最久的 加入到 one_minute_bp_list  
      if(iter != idx_time.rend())
      {      
          //auto out_iter  =  _min_out_bp_table.cbegin( );
          _min_out_bp_table.emplace( _self, [&]( one_minute_bp_list& producer ){
               //producer.producer_key         = iter->producer_key;
               producer.ulord_addr		       = iter->ulord_addr;
	       producer.bpname				   = iter->bpname;
	       producer.bp_valid_time        = min_now+20*60;
          });
           
          auto globalvars_itr = _bp_global_var.find(MIN_BP_OUT_NUM);
          if(globalvars_itr->val > 20)
          {
             auto iter_ = _min_out_bp_table.begin();
             _min_out_bp_table.erase(*iter_);
          }
		  else
          {
             _bp_global_var.modify(globalvars_itr, _self, [&](auto& g){
                g.val++;
              });
          }  
          
          _minprodtable.modify( *iter, _self, [&]( auto& producer ) {
                producer.on_line_time=0; // 120 second              
          });              
      }

      auto globalvar_itr = _bp_global_var.find(MIN_BP_NUM);
      auto bpsize = globalvar_itr->val;
      if(bpsize>25)
      {
          _bp_global_var.modify(globalvar_itr, _self, [&](auto& g){
                g.val=25;
          });
      }
      for(;bpsize >25 ; bpsize--)
      {       
         auto idx_uptime = _minprodtable.get_index<N(byuptime)>();
         auto iter_uptime = idx_uptime.begin();
         _minprodtable.erase(*iter_uptime);   
      }

// add by alvin for output bp list end 
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

      bytes packed_schedule = pack(producers);

      //modifid gy hml
      bool is_follow = false;
      if( _bpinfo.begin() != _bpinfo.end()) {
         is_follow = _bpinfo.begin()->is_follow;
      }

      print("update_elected_producers() is_follow == ", is_follow);
      
      if(!is_follow) {
         if( set_proposed_producers( packed_schedule.data(),  packed_schedule.size() ) >= 0 ) {
             _gstate.last_producer_schedule_size = static_cast<decltype(_gstate.last_producer_schedule_size)>( top_producers.size() );
        }
      }
   }

   double stake2vote( int64_t staked ) {
      /// TODO subtract 2080 brings the large numbers closer to this decade
      double weight = int64_t( (now() - (block_timestamp::block_timestamp_epoch / 1000)) / (seconds_per_day * 7) )  / double( 52 );
      return double(staked) * std::pow( 2, weight );
   }
   /**
    *  @pre producers must be sorted from lowest to highest and must be registered and active
    *  @pre if proxy is set then no producers can be voted for
    *  @pre if proxy is set then proxy account must exist and be registered as a proxy
    *  @pre every listed producer or proxy must have been previously registered
    *  @pre voter must authorize this action
    *  @pre voter must have previously staked some UOS for voting
    *  @pre voter->staked must be up to date
    *
    *  @post every producer previously voted for will have vote reduced by previous vote weight
    *  @post every producer newly voted for will have vote increased by new vote amount
    *  @post prior proxy will proxied_vote_weight decremented by previous vote weight
    *  @post new proxy will proxied_vote_weight incremented by new vote weight
    *
    *  If voting for a proxy, the producer votes will not change until the proxy updates their own vote.
    */
   void system_contract::voteproducer( const account_name voter_name, asset quantity, const account_name proxy, const std::vector<account_name>& producers ) {
      require_auth( voter_name );
      update_votes( voter_name, quantity, proxy, producers, true );
   }

	//Begin: modify by UOS(hml) Date:2018-09-21
	//attention: Here is not deal with the case of  proxy, if add furture, the funtion need deal it.
   void system_contract::cancelvote( const account_name voter_name, asset quantity, const std::vector<account_name>& producers ) {
      require_auth( voter_name );

	//verify input
	 auto voter = _voters.find(voter_name);
     uosio_assert( voter != _voters.end(), "user have not voted before" );

     if(quantity.amount == 0) {
        cancelallvote(voter_name, false);
        return ;
     }

	 uosio_assert(quantity.amount > 100, "min cancel vote is 0.01");
	 uosio_assert(producers.size() > 0 && producers.size() < max_voter_producers, "None of cancel vote producer.");
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

				 _voters.modify( voter, 0, [&]( auto& av ) {
				 		bool isfind = false;
				 		for(auto it = av.vote_producers.begin(); it != av.vote_producers.end(); ++it) {
							if(p == it->producer) {
								isfind = true;
								it->voted.amount -= quantity.amount;
							}
						} 	
						uosio_assert( isfind, "not exist in voted producer list.");
				 	});

					
				auto list_producer = _producers.find(p);
				uosio_assert( list_producer != _producers.end(), "not exist in voted producer list.");
				
				auto new_vote_weight =  quantity.amount;
#if UOSIOLIB_DEBUG
				print("new_vote_weight=",new_vote_weight," list_producer->total_votes=",list_producer->total_votes,"\n" );
#endif

				uosio_assert(list_producer->total_votes - new_vote_weight >= 0, "amount of cancel voted is exceed.");

				_producers.modify( list_producer, 0, [&]( auto& producer ) {
		               producer.total_votes -= new_vote_weight;
		               _gstate.total_producer_vote_weight -= new_vote_weight;
		               
           		 });		

				if(vp->voted.amount  == 0) {
					 _voters.modify( voter, 0, [&]( auto& av ) {
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
void system_contract::cancelallvote( const account_name voter_name, bool is_sudo) {

    if(is_sudo) {
        require_auth( N(uosio) );
    }else {
        require_auth(voter_name);
    }

	auto voter = _voters.find(voter_name);
    uosio_assert( voter != _voters.end(), "user have not voted before" );

	int64_t new_vote_weight = 0;
	
	for(auto vp = voter->vote_producers.begin(); vp != voter->vote_producers.end(); vp++) {
		
		auto p = _producers.find(vp->producer);
		uosio_assert( p != _producers.end(), "not exist in voted producer list.");
		
		new_vote_weight = vp->voted.amount;
		uosio_assert(p->total_votes - new_vote_weight >= 0, "amount of cancel voted is exceed.");

		_producers.modify( p, 0, [&]( auto& producer ) {
		               producer.total_votes -= new_vote_weight;
		               _gstate.total_producer_vote_weight -= new_vote_weight;
           		 });	
	}

	_voters.modify( voter, 0, [&]( auto& av ) {
	         				 av.vote_producers.clear();	
					 	});
}
//End: add by UOS(hml) Date:2018-09-21


void system_contract::validate_votes(const account_name voter_name, asset quantity) {

		print(" voter_name= ",name{voter_name}, " quantity=", quantity.amount, "\n");
		
		auto voter = _voters.find(voter_name);
		uosio_assert( voter != _voters.end(), "user must stake before they can vote" );

		asset total_voted(0);
		
	  	for(auto vp = voter->vote_producers.begin(); vp != voter->vote_producers.end(); ++vp) 
		{
			total_voted += vp->voted;
		}

		print(" voter->staked ", voter->staked, " total_voted.amount =", total_voted.amount, "quantity.amount =", quantity.amount, "\n");

		//attention:  voter->staked was changed before excute here, so ignore para of quantity.amount 
		uosio_assert( voter->staked - total_voted.amount /*+ quantity.amount*/ >= 0, "user must cancel vote producer before undelegate." );
   }

   void system_contract::update_votes( const account_name voter_name,  asset quantity,  account_name proxy, const std::vector<account_name>& producers, bool voting ) {

		 print(" voter_name= ",name{voter_name}, " quantity=", quantity.amount, " proxy = ", name{proxy}, " voting =", voting, "\n");
		 for(auto p : producers) {
			print("update_votes() producers:", p, "\n");
		 }
		 
	//Begin: modify by UOS(hml) Date:2018-09-20

		//cancel the interface of proxy by RPC .
		//uosio_assert(proxy == 0, "not support proxy vote.");
		proxy = 0;
		
	  //validate input
      auto voter = _voters.find(voter_name);
      uosio_assert( voter != _voters.end(), "user must stake before they can vote" ); /// staking creates voter object
      uosio_assert( producers.size() + voter->vote_producers.size() <= max_voter_producers, "attempt to vote for too many producers" );

      /**
       * The first time someone votes we calculate and set last_vote_weight, since they cannot unstake until
       * after total_activated_stake hits threshold, we can use last_vote_weight to determine that this is
       * their first vote and should consider their stake activated.
       */

	  auto new_vote_weight = quantity.amount  * producers.size();
	  _gstate.total_activated_stake +=	new_vote_weight;


	  //get start time of mainnet
	  if( _gstate.total_activated_stake >= min_activated_stake && _gstate.thresh_activated_stake_time == 0 ) {
            _gstate.thresh_activated_stake_time = current_time();

            //<begin> modified by uos camphor
            action(permission_level{_self, N(active)}, N(uosio.bvpay),
                     N(start), _self)
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

      boost::container::flat_map<account_name, pair<double, bool /*new*/> > producer_deltas;

	  if ( voter->last_vote_weight > 0 ){
		for( const auto& p : voter->vote_producers ) {
           auto& d = producer_deltas[p.producer];
           d.first -= voter->last_vote_weight;
           d.second = false;
        }

	  }
     
     if( new_vote_weight >= 0 ) {
        for( const auto& p : producers ) {
           auto& d = producer_deltas[p];
           d.first += new_vote_weight;
           d.second = true;
        }
     }

      for( const auto& pd : producer_deltas ) {
         auto pitr = _producers.find( pd.first );
         if( pitr != _producers.end() ) {
            uosio_assert( !voting || pitr->active() || !pd.second.second /* not from new set */, "producer is not currently registered" );
            _producers.modify( pitr, 0, [&]( auto& p ) {
               p.total_votes += pd.second.first;
               if ( p.total_votes < 0 ) { // floating point arithmetics can give small negative numbers
                  p.total_votes = 0;
               }
               _gstate.total_producer_vote_weight += pd.second.first;
               //uosio_assert( p.total_votes >= 0, "something bad happened" );
            });
         } else {
            uosio_assert( !pd.second.second /* not from new set */, "producer is not registered" ); //data corruption
         }
      }

	//update all voted producer, if not exist,it will ceate new and insert vector.
	vote_producer newvp;
	bool isFind = false;
	
	for(auto p : producers) {
		if(voter->vote_producers.size() > 0) {
			isFind = false;
		
			_voters.modify( voter, 0, [&]( auto& av ) {
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
			_voters.modify( voter, 0, [&]( auto& av ) {
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
   void system_contract::regproxy( const account_name proxy, bool isproxy ) {
      require_auth( proxy );

      auto pitr = _voters.find(proxy);
      if ( pitr != _voters.end() ) {
         uosio_assert( isproxy != pitr->is_proxy, "action has no effect" );
         uosio_assert( !isproxy || !pitr->proxy, "account that uses a proxy is not allowed to become a proxy" );
         _voters.modify( pitr, 0, [&]( auto& p ) {
               p.is_proxy = isproxy;
            });
         propagate_weight_change( *pitr );
      } else {
         _voters.emplace( proxy, [&]( auto& p ) {
               p.owner  = proxy;
               p.is_proxy = isproxy;
            });
      }
   }

  void system_contract::setbpout(uint64_t option)
  {
      require_auth( _self );

        switch(option) {

            //case ENU_BPLIST_OUTPUT:
            case 1:
                 {
                    print("\n", "one_minute_bp_list: ", "\n");
                    for(auto& item : _min_out_bp_table) {
                        print( "  bpname: ", item.bpname, "\n", "  ulord_addr: ", item.ulord_addr, "\n", "  bp_valid_time: ", item.bp_valid_time, "\n");
                     }

                    print("\n", "one_minute_bp_list: ", "\n");
                    for(auto& item : _minprodtable) {
                        print( "  bpname: ", item.bpname, "\n", "  ulord_addr: ", item.ulord_addr, "\n","  on_line_time: ", item.on_line_time, "\n", "  uptime: ", item.uptime, "\n" );
                    }


                    print("\n", "bp_global_var: ", "\n");
                    std::string strTemp;
                    for(auto& item : _bp_global_var) {
                        if(item.id == MIN_BP_NUM) {
                            strTemp = "MIN_BP_NUM";
                        }else if (item.id == MIN_BP_OUT_NUM) {
                            strTemp = "MIN_BP_OUT_NUM";
                        }
                        
                        print( "  id: ", strTemp.c_str(), "\n", "  val: ", item.val, "\n");
                    }
                 }
                  break;

            //case ENU_BPLIST_INIT: 
             case 2:
                  {

				 int icount = 0;
                 while(_minprodtable.begin() != _minprodtable.end()) {
                       _minprodtable.erase(_minprodtable.begin());  
						if(++icount > 100) {
							break;
						}
                  }
				  
				   icount = 0;
                 while(_min_out_bp_table.begin() != _min_out_bp_table.end()) {
                       _min_out_bp_table.erase(_min_out_bp_table.begin());  
						if(++icount > 100) {
							break;
						}
                  }
		
                 auto globalvars_itr = _bp_global_var.find(MIN_BP_NUM);
                 _bp_global_var.modify(globalvars_itr, _self, [&](auto& g){
                       g.val = 0;
                 });


                 auto globalvars_itr1 = _bp_global_var.find(MIN_BP_OUT_NUM);
                 _bp_global_var.modify(globalvars_itr1, _self, [&](auto& g){ 
                       g.val = 0;
                  });

               }
              break;

            default:
                  uosio_assert(0, "invalid paruint8_tsetbpout,option. ");

        }
  }

   void system_contract::propagate_weight_change( const voter_info& voter ) {
      uosio_assert( voter.proxy == 0 || !voter.is_proxy, "account registered as a proxy is not allowed to use a proxy" );

	//Begin: removed by UOS(hml) Date:2018-09-20
#if 0
	  double new_weight = stake2vote( voter.staked );
      if ( voter.is_proxy ) {
         new_weight += voter.proxied_vote_weight;
      }

      /// don't propagate small changes (1 ~= epsilon)
      if ( fabs( new_weight - voter.last_vote_weight ) > 1 )  {
         if ( voter.proxy ) {
            auto& proxy = _voters.get( voter.proxy, "proxy not found" ); //data corruption
            _voters.modify( proxy, 0, [&]( auto& p ) {
                  p.proxied_vote_weight += new_weight - voter.last_vote_weight;
               }
            );
            propagate_weight_change( proxy );
         } else {
            auto delta = new_weight - voter.last_vote_weight;
            for ( auto acnt : voter.producers ) {
               auto& pitr = _producers.get( acnt, "producer not found" ); //data corruption
               _producers.modify( pitr, 0, [&]( auto& p ) {
                     p.total_votes += delta;
                     _gstate.total_producer_vote_weight += delta;
               });
            }
         }
      }

	  _voters.modify( voter, 0, [&]( auto& v ) {
            v.last_vote_weight = new_weight;
         }
      );

	  #endif
	//End: removed by UOS(hml) Date:2018-09-20
      
   }

} /// namespace uosiosystem
