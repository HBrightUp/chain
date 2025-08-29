#include "uosio.system.hpp"
#include <uosiolib/dispatcher.hpp>
#include <uosiolib/crypto.h>
#include "producer_pay.cpp"
#include "delegate_bandwidth.cpp"
#include "voting.cpp"
#include "exchange_state.cpp"


namespace uosiosystem {

   system_contract::system_contract( account_name s )
   :native(s),
    _voters(_self,_self),
    _producers(_self,_self),
    _global(_self,_self),
    _rammarket(_self,_self),
    _min_out_bp_table(_self, _self),
    _minprodtable(_self,_self),
    _bp_global_var(_self,_self),
    _uosc_bp_list(_self, N(uosclist)),
    _bpinfo(_self, _self),
    //_uos_bp_list(_self, _self)
    _share_bplist_table(_self, _self)
   {
      //print( "construct system\n" );
      _gstate = _global.exists() ? _global.get() : get_default_parameters();

      auto itr = _rammarket.find(S(4,RAMCORE));

      if( itr == _rammarket.end() ) {
         auto system_token_supply   = uosio::token(N(uosio.token)).get_supply(uosio::symbol_type(system_token_symbol).name()).amount;
         if( system_token_supply > 0 ) {
            itr = _rammarket.emplace( _self, [&]( auto& m ) {
               m.supply.amount = 100000000000000ll;
               m.supply.symbol = S(4,RAMCORE);
               m.base.balance.amount = int64_t(_gstate.free_ram());
               m.base.balance.symbol = S(0,RAM);
//               m.quote.balance.amount = system_token_supply / 1000;
               m.quote.balance.amount = 50000000000;
               m.quote.balance.symbol = CORE_SYMBOL;
            });
         }
      } else {
         //print( "ram market already created" );
      }
#if 1
      if( _bpinfo.begin() == _bpinfo.end() ) {
         _bpinfo.emplace( _self, [&] ( auto& info ) {
               info.primary = 1;
               info.is_follow = false;
               info.oldest_time = 0;
               info.bplist_num = 0;    
         });
      }
#endif
   }

   uosio_global_state system_contract::get_default_parameters() {
      uosio_global_state dp;
      get_blockchain_parameters(dp);
      return dp;
   }


   system_contract::~system_contract() {
      //print( "destruct system\n" );
      _global.set( _gstate, _self );
      //uosio_exit(0);
   }

   void system_contract::setram( uint64_t max_ram_size ) {
      require_auth( _self );

      uosio_assert( _gstate.max_ram_size < max_ram_size, "ram may only be increased" ); /// decreasing ram might result market maker issues
      uosio_assert( max_ram_size < 1024ll*1024*1024*1024*1024, "ram size is unrealistic" );
      uosio_assert( max_ram_size > _gstate.total_ram_bytes_reserved, "attempt to set max below reserved" );

      auto delta = int64_t(max_ram_size) - int64_t(_gstate.max_ram_size);
      auto itr = _rammarket.find(S(4,RAMCORE));

      /**
       *  Increase or decrease the amount of ram for sale based upon the change in max
       *  ram size.
       */
      _rammarket.modify( itr, 0, [&]( auto& m ) {
         m.base.balance.amount += delta;
      });

      _gstate.max_ram_size = max_ram_size;
      _global.set( _gstate, _self );
   }

   void system_contract::setparams( const uosio::blockchain_parameters& params ) {
      require_auth( N(uosio) );
      (uosio::blockchain_parameters&)(_gstate) = params;
      uosio_assert( 3 <= _gstate.max_authority_depth, "max_authority_depth should be at least 3" );
      set_blockchain_parameters( params );
   }

   void system_contract::setpriv( account_name account, uint8_t ispriv ) {
      require_auth( _self );
      set_privileged( account, ispriv );
   }

    // **********<begin>This code was modified by camphor 2018-12-10 </begin>*********//
   void system_contract::rmvproducer( account_name producer ) {
      require_auth( _self );
      auto prod = _producers.find( producer );
      uosio_assert( prod != _producers.end(), "producer not found" );
      if(!prod->release){
         _producers.modify( prod, 0, [&](auto& p) {
             p.deactivate();
             p.release = 1;
             p.is_remove = true;
         });
         penaltyprodbw(producer);
      }else{
         _producers.modify( prod, 0, [&](auto& p) {
             p.deactivate();
             p.is_remove = true;
         });
      }

   }
    // **********<end>This code was modified by camphor 2018-12-10 </end>*********//

   void system_contract::bidname( account_name bidder, account_name newname, asset bid ) {
      require_auth( bidder );
      uosio_assert( uosio::name_suffix(newname) == newname, "you can only bid on top-level suffix" );
      uosio_assert( newname != 0, "the empty name is not a valid account name to bid on" );
      uosio_assert( (newname & 0xFull) == 0, "13 character names are not valid account names to bid on" );
      uosio_assert( (newname & 0x1F0ull) == 0, "accounts with 12 character names and no dots can be created without bidding required" );
      uosio_assert( !is_account( newname ), "account already exists" );
      uosio_assert( bid.symbol == asset().symbol, "asset must be system token" );
      uosio_assert( bid.amount > 0, "insufficient bid" );

      INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), {bidder,N(active)},
                                                    { bidder, N(uosio.names), bid, std::string("bid name ")+(name{newname}).to_string()  } );

      name_bid_table bids(_self,_self);
      print( name{bidder}, " bid ", bid, " on ", name{newname}, "\n" );
      auto current = bids.find( newname );
      if( current == bids.end() ) {
         bids.emplace( bidder, [&]( auto& b ) {
            b.newname = newname;
            b.high_bidder = bidder;
            b.high_bid = bid.amount;
            b.last_bid_time = current_time();
         });
      } else {
         uosio_assert( current->high_bid > 0, "this auction has already closed" );
         uosio_assert( bid.amount - current->high_bid > (current->high_bid / 10), "must increase bid by 10%" );
         uosio_assert( current->high_bidder != bidder, "account is already highest bidder" );

         INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), {N(uosio.names),N(active)},
                                                       { N(uosio.names), current->high_bidder, asset(current->high_bid),
                                                       std::string("refund bid on name ")+(name{newname}).to_string()  } );

         bids.modify( current, bidder, [&]( auto& b ) {
            b.high_bidder = bidder;
            b.high_bid = bid.amount;
            b.last_bid_time = current_time();
         });
      }
   }

   /**
    *  Called after a new account is created. This code enforces resource-limits rules
    *  for new accounts as well as new account naming conventions.
    *
    *  Account names containing '.' symbols must have a suffix equal to the name of the creator.
    *  This allows users who buy a premium name (shorter than 12 characters with no dots) to be the only ones
    *  who can create accounts with the creator's name as a suffix.
    *
    */
   void native::newaccount( account_name     creator,
                            account_name     newact
                            /*  no need to parse authorities
                            const authority& owner,
                            const authority& active*/ ) {

      if( creator != _self ) {
         auto tmp = newact >> 4;
         bool has_dot = false;

         for( uint32_t i = 0; i < 12; ++i ) {
           has_dot |= !(tmp & 0x1f);
           tmp >>= 5;
         }
         if( has_dot ) { // or is less than 12 characters
            auto suffix = uosio::name_suffix(newact);
            if( suffix == newact ) {
               name_bid_table bids(_self,_self);
               auto current = bids.find( newact );
               uosio_assert( current != bids.end(), "no active bid for name" );
               uosio_assert( current->high_bidder == creator, "only highest bidder can claim" );
               uosio_assert( current->high_bid < 0, "auction for name is not closed yet" );
               bids.erase( current );
            } else {
               uosio_assert( creator == suffix, "only suffix may create this account" );
            }
         }
      }

      user_resources_table  userres( _self, newact);

      userres.emplace( newact, [&]( auto& res ) {
        res.owner = newact;
      });

      set_resource_limits( newact, 0, 0, 0 );
   }

    void system_contract::setbpinfo(bool is_follow) {
       require_auth( _self );

       if (_bpinfo.begin() != _bpinfo.end()) {
          const auto info = _bpinfo.begin();

         _bpinfo.modify( info, _self, [&]( auto& m) {
         m.is_follow = is_follow;
         });
       }

    }

void system_contract::setbplist(const account_name bp_name, const uint64_t bp_time, const std::string hash_bplist, const std::vector< uosio::producer_key >& bplist) {

   uosio_assert(hash_bplist.length() == 64, "invalid hash");
   uosio_assert( _bpinfo.begin() != _bpinfo.end(), "bpinfo table no exist" );

   const auto& bpinfo = _bpinfo.begin();

    //first: check whether exist  of bp_name
   if(!bpinfo->is_follow) {
      return ;
   }

   require_auth(bp_name);

   //check authorization
   bool is_auth = false;
   uosio::name prod;
    prod.value = bp_name;
   for(auto bp : bplist) {
      if(bp.producer_name == bp_name) {
         is_auth = true;
         break;
      }
   }

   if(!is_auth) {
      uosio_assert(false, "setter not in bp list");
   }

   uint64_t curTime = now();

   //second: check time
   if(bp_time > curTime) {
       print("bp_time", bp_time, " curTime", curTime);
       print("bp_time  ahead by block");
      return ;
   }

   
   if(bp_time < curTime - max_time_interval) {
      print("bp_time: ", bp_time, "curTime - max_time_interval:", curTime - max_time_interval );
      print("bp_time  behind by block");
      return ;
   }

   //three: compare with oldest bp_time
   if( bpinfo->oldest_time >= bp_time) {
      print("time  behind by the oldest bp_time");
      print("bpinfo->oldest_time", bpinfo->oldest_time, "bp_time: ", bp_time );
      return ;
   }

   //four:check hash
   
   prod.value = bp_name;

   //std::string hash_bplist_para = prod.to_string() + std::to_string(bp_time);
   std::string hash_bplist_para;
   std::vector<unsigned char> bp_key;

   hash_bplist_para = std::to_string(bp_time);
   for(const auto& bp : bplist) {
      prod.value = bp.producer_name;
      //hash_bplist_para += prod.to_string() + std::string(bp.block_signing_key );
      hash_bplist_para += prod.to_string();
#if 0
      print("Begin: block_signing_key: ");
      for(int i = 0; i < sizeof(bp.block_signing_key.data); i++) {
         printi(bp.block_signing_key.data[i]);
      }

      print("End: block_signing_key: ");
#endif

      //printhex((void * )&bp.block_signing_key.data , sizeof(bp.block_signing_key.data) );
      //print("\n");

      
      bp_key.clear();
      for(int i = 0; i < sizeof(bp.block_signing_key.data); ++i) {
         bp_key.emplace_back(bp.block_signing_key.data[i]);
      }

   }
   
   checksum256 ck_input,ck_para;
   from_hex(hash_bplist, (char*)ck_input.hash, sizeof(ck_input.hash));
   sha256(hash_bplist_para.c_str(), hash_bplist_para.length(), &ck_para);

   for(int i = 0; i < sizeof(ck_input); i++) {
      uosio_assert(ck_input.hash[i] == ck_para.hash[i], "check bplist hash failed!");
   }

   //if exist,only add times
   bool bfind = false;
   uint64_t max_count = 0;
   uint64_t total_count = 0;
   auto  iter_max_count = _share_bplist_table.end();
   uint64_t oldest_time = curTime;

   // whether share table add item
   int32_t share_item_changed = 0;


   if(_share_bplist_table.begin() == _share_bplist_table.end()) {

      _share_bplist_table.emplace( _self, [&](auto& share) {
         share.primary = _share_bplist_table.available_primary_key();
         share.count = 1; 
         share.time = bp_time;
         share.bp_name = bp_name;
         share.bplist = bplist;
         share.hash_bplist = hash_bplist;
      });
      max_count = 1;
      total_count = 1;
      iter_max_count = _share_bplist_table.begin();
      oldest_time = bp_time;
      ++share_item_changed;

   }else {
      for(auto it = _share_bplist_table.begin(); it != _share_bplist_table.end();) {

         //delete old bplist
         print("curTime: ", curTime, "   it->time:", it->time);
         if( curTime - it->time >  max_time_interval * 2) {
               it = _share_bplist_table.erase(it); 
               --share_item_changed;
               continue;
         }

         //find exist or not
         if(it->hash_bplist == hash_bplist) {
            bfind = true;
            _share_bplist_table.modify(it, _self, [&](auto& share) {
               share.count++; 
               share.time = bp_time;
               share.bp_name = bp_name;
            });
         }

         //get the largest count
         if(it->count > max_count) {
            max_count = it->count;
            iter_max_count = it;
         }

         if(oldest_time > it->time) {
            oldest_time = it->time;
         }

         //only calc  correct bplist
         if(now() - it->time <= max_time_interval ) {
            total_count += it->count; 
         }
         
         ++it;
      }

      if(!bfind) {
         bool is_empty = false; 
         if(_share_bplist_table.begin() == _share_bplist_table.end()) {
            is_empty = true;
         }

         _share_bplist_table.emplace( _self, [&](auto& share) {
            share.primary = _share_bplist_table.available_primary_key();
            share.count = 1; 
            share.time = bp_time;
            share.bp_name = bp_name;
            share.bplist = bplist;
            share.hash_bplist = hash_bplist;
         });

         if(oldest_time > bp_time) {
            oldest_time = bp_time;
         }

         ++total_count;
         ++share_item_changed;

         //only one item
         if(is_empty) {
            iter_max_count = _share_bplist_table.begin();
            max_count = iter_max_count->count;
         }
      }  
   }

   _bpinfo.modify(bpinfo, _self, [&](auto& info) {
      info.oldest_time = oldest_time;
      info.bplist_num += share_item_changed;
   });

   // consensus of 2/3
   if((double(max_count) / double(total_count)) < 2.00/3.00) {
      //print("max_count:", max_count, " total_count: ", total_count, "\n");
      uosio_assert(false, "warning::  cannot get 2/3 consensus bplist" );
      return ;
   }

   std::vector<uosio::producer_key> producers;
   producers.reserve(iter_max_count->bplist.size());
   for(const auto& bp : iter_max_count->bplist ) {
      if(is_account(bp.producer_name)) {
         producers.push_back(bp);
      }
   }
   std::sort(producers.begin(), producers.end());
   bytes packed_schedule = pack(producers);
   
   if( set_proposed_producers( packed_schedule.data(),  packed_schedule.size() ) >= 0 ) {
      _gstate.last_producer_schedule_size = static_cast<decltype(_gstate.last_producer_schedule_size)>( producers.size() );
   }
   
}

size_t system_contract::from_hex( const std::string& hex_str, char* out_data, size_t out_data_len ) {
   std::string::const_iterator i = hex_str.begin();
   uint8_t* out_pos = (uint8_t*)out_data;
   uint8_t* out_end = out_pos + out_data_len;
   while( i != hex_str.end() && out_end != out_pos ) {
      *out_pos = from_hex( *i ) << 4;   
      ++i;
      if( i != hex_str.end() )  {
         *out_pos |= from_hex( *i );
         ++i;
      }
      ++out_pos;
   }
   return out_pos - (uint8_t*)out_data;
}

uint8_t system_contract::from_hex( char c ) {
   if( c >= '0' && c <= '9' )
      return c - '0';
   if( c >= 'a' && c <= 'f' )
         return c - 'a' + 10;
   if( c >= 'A' && c <= 'F' )
         return c - 'A' + 10;
   //FC_THROW_EXCEPTION( exception, "Invalid hex character '${c}'", ("c", fc::string(&c,1) ) );
   uosio_assert(0, "Invalid hex character ");
   return 0;
}

std::string system_contract::EncodeBase58(const std::vector<unsigned char>& vch)
{
	return EncodeBase58(vch.data(), vch.data() + vch.size());
}

std::string system_contract::EncodeBase58(const unsigned char* pbegin, const unsigned char* pend)
{
	// Skip & count leading zeroes.
	int zeroes = 0;
	int length = 0;
	while (pbegin != pend && *pbegin == 0) {
		pbegin++;
		zeroes++;
	}
	// Allocate enough space in big-endian base58 representation.
	int size = (pend - pbegin) * 138 / 100 + 1; // log(256) / log(58), rounded up.
	std::vector<unsigned char> b58(size);
	// Process the bytes.
	while (pbegin != pend) {
		int carry = *pbegin;
		int i = 0;
		// Apply "b58 = b58 * 256 + ch".
		for (std::vector<unsigned char>::reverse_iterator it = b58.rbegin(); (carry != 0 || i < length) && (it != b58.rend()); it++, i++) {
			carry += 256 * (*it);
			*it = carry % 58;
			carry /= 58;
		}

		assert(carry == 0);
		length = i;
		pbegin++;
	}
	// Skip leading zeroes in base58 result.
	std::vector<unsigned char>::iterator it = b58.begin() + (size - length);
	while (it != b58.end() && *it == 0)
		it++;
	// Translate the result into a string.
	std::string str;
	str.reserve(zeroes + (b58.end() - it));
	str.assign(zeroes, '1');
	while (it != b58.end())
		str += pszBase58[*(it++)];
	return str;
}


} /// uosio.system


UOSIO_ABI( uosiosystem::system_contract,
     // native.hpp (newaccount definition is actually in uosio.system.cpp)
     (newaccount)(updateauth)(deleteauth)(linkauth)(unlinkauth)(canceldelay)(onerror)
     // uosio.system.cpp
     (setram)(setparams)(setpriv)(rmvproducer)(bidname)
     // delegate_bandwidth.cpp
     (buyrambytes)(buyram)(sellram)(delegatebw)(undelegatebw)(refund)
     // voting.cpp
     (regproducer)(unregprod)(voteproducer)(cancelvote)(regproxy)
     // producer_pay.cpp
     (onblock)(claimrewards)(payfee)
     (setbpout)(setbpinfo)(setbplist)
)
