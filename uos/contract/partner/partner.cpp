#include<uosiolib/uosio.hpp>
#include <uosiolib/privileged.h>
#include "partner.hpp"


namespace partnersys {
    using namespace uosio;

   partner::partner(name receiver, name code,  uosio::datastream<const char*> ds)
   :contract(receiver, code, ds),
   _global_var(_self, _self.value),
   _partners(_self, _self.value)
   {
   }

   void partner::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {

      uosio_assert(op == 1 || op == 2, "op 1: add or modify; 2: delete");
   
      auto global_itr = _global_var.find(key.value);

      //delete data
      if(op == 2) {
         uosio_assert(global_itr != _global_var.end(), "not find key");
         _global_var.erase(global_itr);
         return ;
      }

      //add or modify data
      if(global_itr == _global_var.end()) {
         _global_var.emplace( _self, [&] (auto& g) {
               g.key = key;
               g.val = val;
         }); 
      }
      else {
         _global_var.modify(global_itr, _self, [&](auto& g){
               g.val = val;
         }); 
      }
   }

   ACTION partner::init() {

      require_auth(_self);

      //min asset required joining to partners(default: 10000.0000 UOS)
      set_global_var("join"_n, 100000000);

      //seconds  of  one year (3600 * 24 * 365)
      set_global_var("oneyear"_n, 3600 * 24 * 365);

      //set the max release number once time
      set_global_var("maxrelease"_n, 1000);

   }

   ACTION partner::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
      require_auth(_self);
      set_global_var(key, val, op);
   }

   ACTION partner::invite(const name owner, const std::vector<name> invited_list) {
      require_auth(owner);

      auto owner_iter = _partners.find(owner.value);
      uosio_assert(owner_iter != _partners.end(), "owner not in partners table");

      if(owner_iter->asset_list.begin() == owner_iter->asset_list.end()) {
         uosio_assert(false, "owner not transfer asset");
      }

      auto ct = now();

      for(auto act : invited_list) {

         //whether exist on blockchain 
         uosio_assert(is_account(act), "account not exist on blockchain");
         uosio_assert(owner != act, "can't invited yourself");

         //not invited before
         auto act_iter = _partners.find(act.value);
         uosio_assert(act_iter == _partners.end(), "account invited already");

         _partners.emplace( _self, [&] (auto& p) {
               p.owner = act;
               p.prev = owner;
               p.join_time = ct;
               p.update_time = ct;
         }); 
      }

      //add to partners table
      auto partner_iter = _partners.find(owner.value);
      _partners.modify(partner_iter, _self, [&](auto& p){
         for(auto act : invited_list) {
            p.invited++;
            p.update_time = ct;
         }
      }); 
   }

   ACTION partner::clear() {
      require_auth(_self);

      while(_partners.begin() != _partners.end()) {
         _partners.erase(_partners.begin());
      }

   }

   void partner::transfer(const name& from, const name& to, const asset& quantity, const std::string& memo) {

#if 0
      symbol sym("UOS", 4);
      uosio_assert(quantity.symbol == sym, "transfer must system coin");

      const auto join_iter = _global_var.find(("join"_n).value);
      uosio_assert(join_iter != _global_var.end(), "field join not exist");

      const auto year_iter = _global_var.find(("oneyear"_n).value);
      uosio_assert(year_iter != _global_var.end(), "field oneyear not exist");

      const auto maxrelease_iter = _global_var.find(("maxrelease"_n).value);
      uosio_assert(maxrelease_iter != _global_var.end(), "field maxrelease not exist");

      

      const auto ct =  now();

      if(from == _self) {
         auto partner_iter = _partners.find(to.value);
         uosio_assert(partner_iter != _partners.end(), "not find account of 'to' ");

         bool is_find = false;
         for(auto history_iter = partner_iter->asset_list.begin(); history_iter != partner_iter->asset_list.end();) {
            if(history_iter->amount == quantity.amount &&
               ct - history_iter->time >= year_iter->val ) {
                  is_find = true;
                  _partners.modify(partner_iter, _self, [&](auto& p){
                     history_iter = p.asset_list.erase(history_iter);
                     p.update_time = ct;

                  });
            }
            else {
               ++history_iter;
            }
         }
         
         uosio_assert(is_find, "not find vaild history transaction");

         partner_iter = _partners.find(to.value);
         uosio_assert(partner_iter != _partners.end(), "not find account of 'to' ");
         if(partner_iter->asset_list.begin() == partner_iter->asset_list.end()) {
            auto prev_iter =  _partners.get_index<"prev"_n>();
            auto lower = prev_iter.lower_bound(to.value);
            auto upper = prev_iter.upper_bound(to.value);

            std::vector<name>  need_to_modify;
            for(auto iter = lower; iter != upper; ++iter) {
               if(iter->prev == to) {
                  //if exceed 1000 item, we need other method to deal it,
                  //not deal at here
                  need_to_modify.push_back(iter->owner);
                  if(need_to_modify.size() >= maxrelease_iter->val) {
                     uosio_assert(false, "maybe cpu exceed");
                  }
               }
            }

            for(auto user_iter = need_to_modify.begin(); user_iter != need_to_modify.end(); ++user_iter) {
               auto modify_iter = _partners.find((*user_iter).value);
               uosio_assert(modify_iter != _partners.end(), "user not exist");

               _partners.modify(modify_iter, _self, [&](auto& p){
                  p.prev = _self;
                  p.invited = 0;
                  p.update_time = now();
               });
            }

            need_to_modify.clear();
         }

      }

      if(to == _self) {
         uosio_assert(quantity.amount >= join_iter->val, "joining to partner at least 10000 UOS"); 
;
         auto partner_iter = _partners.find(from.value);

         uosio_assert(memo.size() > 0, "no section data");
         const uint8_t section = std::stoi(memo, nullptr,10);
         uosio_assert(section > 0 && section <= 255, "invalid section");

         if(partner_iter == _partners.end()) {

            

            _partners.emplace( _self, [&] (auto& p) {
               p.owner = from;
               p.prev = _self;
               p.join_time = ct;
               p.update_time = ct;
               p.asset_list.push_back(transfer_list{static_cast<uint64_t>(quantity.amount), section, ct});
            }); 
         }else {
            _partners.modify(partner_iter, _self, [&](auto& p){
               p.asset_list.push_back(transfer_list{static_cast<uint64_t>(quantity.amount), section, ct});
               p.update_time = ct;
               
            });
         }
      }

  #endif    
   }

}// end partnersys

#define UOSIO_DISPATCH_PARTNER( TYPE, MEMBERS ) \
extern "C" { \
   void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
      auto self = receiver; \
      if( code == self || code == ("uosio.token"_n).value ) { \
      	 if( code == ("uosio.token"_n).value && action == ("transfer"_n).value ){ \
            uosio::execute_action(uosio::name(receiver), uosio::name(code), &partnersys::partner::transfer); \
            return ; \
      	 } \
         switch( action ) { \
            UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
         } \
      } \
   } \
}

UOSIO_DISPATCH_PARTNER(partnersys::partner, (init)(setdata)(clear)(invite))






















