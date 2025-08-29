#include<uosiolib/uosio.hpp>
#include <uosiolib/privileged.h>
#include "stokadominer.hpp"


namespace stoksys {
    using namespace uosio;

   stok::stok(name receiver, name code,  uosio::datastream<const char*> ds)
   :contract(receiver, code, ds),
   _global_var(_self, _self.value)
   {
   }

   void stok::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {

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

   ACTION stok::init() {

      require_auth(_self);

      //block number of issued stok coin
      set_global_var("currentblock"_n, 3600 * 24 * 365 );

      //block number of issued stok coin
      set_global_var("startblock"_n, 3600 * 24 * 365 );
   }

   ACTION stok::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
      require_auth(_self);
      set_global_var(key, val, op);
   }

   ACTION stok::clear() {
      require_auth(_self);
   }

   ACTION stok::issue(const uint64_t block_number) {
      require_auth(_self);

      const auto cur_iter = _global_var.find(("currentblock"_n).value);
      uosio_assert(cur_iter != _global_var.end(), "field currentblock not exist");
      
      print("currentblock: ", cur_iter->val, " block_number: ",block_number);
      uosio_assert(cur_iter->val < block_number, "invaild block nubmer");

      double amount = get_area(block_number) - get_area(cur_iter->val);
      int64_t real_stok = static_cast<int64_t>(amount * 10000);  

      print("block interval: ", block_number - cur_iter->val, " real_stok: ", real_stok);

      if(real_stok <= 0) {
         print("invalid stok quality");
         return ;
      }

      auto action_data1 = make_tuple(_self, "adominer1111"_n, asset(real_stok, STOK_SYMBOL), std::string("issue stok"));
      action(permission_level{_self, "active"_n}, "cloudstak123"_n, "transfer"_n, action_data1).send();

      set_global_var("currentblock"_n, block_number);
   }
   
   double stok::fx(uint64_t x) {

      const auto start_iter = _global_var.find(("startblock"_n).value);
      uosio_assert(start_iter != _global_var.end(), "field startblock not exist");
      uosio_assert(x >= start_iter->val, "invalid x value");

      const double a = 0.3519288176787673;
      const double b = 2.219685438863521;

      return -a * (x - start_iter->val) / 100000000 + b;
      
   }

   double stok::get_area(uint64_t x) {
      const auto start_iter = _global_var.find(("startblock"_n).value);
      uosio_assert(start_iter != _global_var.end(), "field startblock not exist");
      uosio_assert(x >= start_iter->val, "invalid x value");

      return double(0.5) * (fx(start_iter->val) + fx(x)) * (x - start_iter->val);
   }

   void stok::transfer(const name& from, const name& to, const asset& quantity, const std::string& memo) {


      //symbol sym("UOS", 4);
      //uosio_assert(quantity.symbol == sym, "transfer must system coin");
 
   }

}// end stoksys

#define UOSIO_DISPATCH_STOK( TYPE, MEMBERS ) \
extern "C" { \
   void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
      auto self = receiver; \
      if( code == self || code == ("uosio.token"_n).value ) { \
      	 if( code == ("uosio.token"_n).value && action == ("transfer"_n).value ){ \
            uosio::execute_action(uosio::name(receiver), uosio::name(code), &stoksys::stok::transfer); \
            return ; \
      	 } \
         switch( action ) { \
            UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
         } \
      } \
   } \
}

UOSIO_DISPATCH_STOK(stoksys::stok, (init)(setdata)(clear)(issue))






















