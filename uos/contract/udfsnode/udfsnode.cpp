/**
 *  @file
 *  @copyright defined in uos/LICENSE.txt
 */


#include <uosiolib/crypto.hpp>
#include <uosiolib/transaction.hpp>
//#include <uosio.token/uosio.token.hpp>
//#include <uosio.bvpay/uosio.bvpay.hpp>
#include "./udfsnode.hpp"


namespace uosio{
    uosio_udfs::uosio_udfs(name receiver, name code,  uosio::datastream<const char*> ds)
            :contract(receiver, code, ds),
            _utstate(_self,_self.value),
             _ethstate(_self,_self.value),
            _utvoter(_self,_self.value),
            _ethvoter(_self,_self.value),
            _uttr(_self,_self.value),
            _ethtr(_self,_self.value),
            _memberreward(_self,_self.value),
            _globalvar(_self, _self.value),
            _udfsnode(_self, _self.value),
            _fee(_self, _self.value) { 

             }

    void uosio_udfs::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {

      uosio_assert(op == 1 || op == 2, "op 1: add or modify; 2: delete");
   
      auto global_itr = _globalvar.find(key.value);

      //delete data
      if(op == 2) {
         uosio_assert(global_itr != _globalvar.end(), "not find key");
         _globalvar.erase(global_itr);
         return ;
      }

      //add or modify data
      if(global_itr == _globalvar.end()) {
         _globalvar.emplace( _self, [&] (auto& g) {
               g.key = key;
               g.val = val;
         }); 
      }
      else {
         _globalvar.modify(global_itr, same_payer, [&](auto& g){
               g.val = val;
         }); 
      }
    }

    void uosio_udfs::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
        require_auth(_self);
        set_global_var(key, val, op);
    }


    void uosio_udfs::clear(const name& owner) {
        require_auth(_self);
#if 0
        while(_udfsnode.begin() != _udfsnode.end()) {
             _udfsnode.erase(_udfsnode.begin());
        }
#endif
        auto node = _udfsnode.get_index<"owner"_n>();
        auto begin = node.lower_bound( owner.value );
        if(begin != node.end() && begin->owner == owner) {
            node.erase(begin);
        }

    }

    void uosio_udfs::addnode(const uint64_t tr_hash, const name owner) {
        require_auth(_self);

        uosio_assert(is_account(owner), "uos haven't this account");

        auto node = _udfsnode.get_index<"owner"_n>();
        auto begin = node.lower_bound( owner.value );

        if(begin != node.end() && begin->owner == owner) {
            uosio_assert(false , "current account applied");
        }

        _udfsnode.emplace(_self,[&](auto &u){
                u.tr_hash = tr_hash;
                u.owner = owner;
                u.amount = 0;
                u.apply_time = now();
                u.reserve1 = 0;
                u.reserve2 = 0;
                u.coin_type = 1;
        });

        require_recipient( "udfsproducer"_n );

    }

    template<typename muti_uosstate, typename muti_touostr, typename muti_uosvoter>
    void uosio_udfs::storemember(muti_uosstate& _tb_uosstate, muti_touostr& _tb_touostr, muti_uosvoter& _tb_uosvoter, std::vector<name > &members) {
        auto sta = _tb_uosstate.find(_self.value);
        if(sta == _tb_uosstate.end()){
            _tb_uosstate.emplace(_self,[&](auto &a){
                a.owner = _self;
                a.members = members;
            });
        }else {
            _tb_uosstate.modify(sta , _self, [&](auto &a){
                a.members = members;
            });
        }
        for(auto itr = _tb_touostr.begin() ; itr != _tb_touostr.end() ;){
            itr = _tb_touostr.erase(itr);
        }
        
        for(auto itr = _tb_uosvoter.begin() ; itr !=_tb_uosvoter.end() ; ){
            itr = _tb_uosvoter.erase(itr);
        }

    }

    // @abi action
    void uosio_udfs::modifymember(std::vector<name> & members) {
        require_auth(_self);
        uosio_assert(members.size() > 2 , "union-member must be greater or equal than 3");

        storemember(_utstate, _uttr, _utvoter, members);
        storemember(_ethstate, _ethtr, _ethvoter, members);

        for(auto itr = _memberreward.begin() ; itr != _memberreward.end(); ) {

            if(itr->reward.amount == 0){
                itr = _memberreward.erase(itr);
                continue;
            }
            if(itr->active) {
                _memberreward.modify(itr , same_payer, [&](auto &a){
                    a.active = 0;
                });
            }

            ++itr;
        }

        for(auto member :members) {

            auto mem = _memberreward.find(member.value);
            if(mem == _memberreward.end()){
                _memberreward.emplace(_self, [&](auto &a) {
                    a.owner = member;
                    a.reward = asset(0, {"UOS", 4});
                    a.active = 1;
                });
            }else{
                _memberreward.modify(mem, same_payer,[&](auto &a){
                    a.active = 1;
                });
            }
        }
    }

    template<typename muti_uosstate, typename muti_touostr, typename muti_uosvoter>
    void uosio_udfs::deal(muti_uosstate& _tb_uosstate, muti_touostr& _tb_touostr, muti_uosvoter& _tb_uosvoter, const name voter, const std::string tr_id, const name owner,  const int64_t amount, const std::string coin_name) {
        auto& sta = _tb_uosstate.get(_self.value);
        auto it = std::find(sta.members.begin(), sta.members.end(), voter);
        uosio_assert(it != sta.members.end(),"voter not in union");

        uosio_assert(is_account(owner), "uos haven't this account");
        uosio_assert(tr_id.size() < 70 , "id size must be less than 70");

        //check apply repeatly
        auto node = _udfsnode.get_index<"owner"_n>();
        auto begin = node.lower_bound( owner.value );
        if(begin != node.end() && begin->owner == owner) {
            uosio_assert(false , "current account applied");
        }

        std::string buffer(tr_id);
        buffer += std::to_string(owner.value) + std::to_string(amount);
        uint64_t tr_hash = caculate_hash_64(buffer.c_str() , buffer.size());
        auto ver = _tb_uosvoter.find(voter.value);

        if(ver == _tb_uosvoter.end()){
            _tb_uosvoter.emplace(_self,[&](auto &a){
                a.voter = voter;
                a.tr_hash = tr_hash;
                a.active_time = now();
                a.tr_id = tr_id;
            });
        }else{
            auto & tr = _tb_touostr.get(ver->tr_hash , "can find tr");
            if(tr.votes == 1){
                _tb_touostr.erase(tr);
            }else{
                _tb_touostr.modify(tr, same_payer,[&](auto &a){
                    a.votes --;
                });
            }
            _tb_uosvoter.modify(ver, same_payer,[&](auto &a){
                a.tr_hash = tr_hash;
                a.active_time = now();
                a.tr_id = tr_id;
            });
        }
        auto tr = _tb_touostr.find(tr_hash);
        if(tr == _tb_touostr.end()){
            _tb_touostr.emplace(_self,[&](auto &a){
                a.tr_hash = tr_hash;
                a.tr_id = tr_id;
                a.owner = owner;
                a.amount = amount;
                a.votes = 1;
            });
        }else{
            _tb_touostr.modify(tr, same_payer,[&](auto &a){
                a.votes ++;
            });
        }

        tr = _tb_touostr.find(tr_hash);

        if(tr->votes > int32_t(sta.members.size() * 2 / 3)){
            auto tr_udfs = _udfsnode.find(tr_hash);
            if(tr_udfs == _udfsnode.end()){
                _udfsnode.emplace(_self,[&](auto &u){
                    u.tr_hash = tr_hash;
                    u.owner = owner;
                    u.amount = static_cast<uint64_t>(amount);
                    u.apply_time = now();
                    u.reserve1 = 0;
                    u.reserve2 = 0;

                    if(coin_name == "ut") {
                        u.coin_type = 1;
                    }
                    else if(coin_name == "eth") {
                        u.coin_type = 2;
                    }
                    else {
                        uosio_assert(false, "invalid coin type");
                    }
                });

                action(permission_level{_self, "active"_n}, _self,
                       "applypass"_n, std::make_tuple(voter, tr_id,
                                                       owner , amount , std::string("pass")))
                        .send();

                require_recipient( "udfsproducer"_n );
            }
            else{
                uosio_assert(false, "duplicated transaction");
            }

            _tb_uosstate.modify(sta, same_payer,[&](auto &a){
                    a.laster_id = tr->tr_id;
                    a.laster_owner = tr->owner;
                    a.amount = amount;
                    a.tr_hash = tr->tr_hash;
                });

            for(auto itr = _tb_uosvoter.begin() ; itr != _tb_uosvoter.end() ;){
                itr = _tb_uosvoter.erase(itr);
            }
            for(auto itr = _tb_touostr.begin() ; itr != _tb_touostr.end() ; ){
                itr = _tb_touostr.erase(itr);
            }
        }
    }

    void uosio_udfs::voter(const name voter, const std::string tr_id, const name owner, const int64_t amount, const std::string coin_name) {
        require_auth(voter);

        uosio_assert(amount > 0 && amount <= asset::max_amount, "amount  error");
        
        if(coin_name == "ut") {
            deal(_utstate, _uttr, _utvoter, voter, tr_id, owner, amount, coin_name);
        }
        else if(coin_name == "eth") {
            deal(_ethstate, _ethtr, _ethvoter, voter, tr_id, owner, amount, coin_name);
        }
        else {
            uosio_assert(false, "check name of coin failed");
        }
    }

    // @abi action
    void uosio_udfs::applypass(name voter,std::string tr_id , name owner , int64_t amount , std::string memo){

    }

    // @abi action
    void uosio_udfs::transfer(name from, name to, asset quantity, std::string memo) {
          if(from == _self){
            require_auth(_self);
            return;
        }

        require_auth(from);       
    }

    uint64_t uosio_udfs::caculate_hash_64(const char *buf, uint32_t size) {
        checksum160 calc_hash;
        calc_hash = ripemd160(buf, size);
   
        uint64_t a = * reinterpret_cast<uint64_t *>(calc_hash.data());
        uint64_t b = * reinterpret_cast<uint64_t *>(calc_hash.data() + 8);
        uint32_t c = * reinterpret_cast<uint32_t *>(calc_hash.data() + 16);
        return a^b^c;
    }

    void uosio_udfs::feetableop(const name owner, const uint8_t op) {
        require_auth(owner);

        uosio_assert(op >= 1 && op <= 2, "invalid operation" );

        
    }



    // issue token to bp and udfs_system  add by alvin
    void uosio_udfs::issuetoken() {
        
#if 0
        auto fielditer = _globalvar.find(N(issuetime)); 
        uint32_t  time_now =  now();
        if(fielditer == _globalvar.end()) {
            set_global_var(N(issuetime),time_now);
            set_global_var(N(udfssupply),2280000000);
            set_global_var(N(bpsupply),120000000);
            return;
        }

        //auto interval = seconds_per_day;
        auto interval = 60;
        
        if(time_now > fielditer->val + interval)
        {
            //get account udfsreceiver 	 
            if(!is_account(N(udfsreceiver))){
                print( "udfs : ",  time_now, "   no  account: udfsreceiver  ",  "\n");
                return;
            } 

            accounts accountstable( N(uosio.token),N(udfsreceiver) );
            const auto& ac = accountstable.get(uosio::symbol_type(uosiosystem::system_token_symbol).name());
            auto udfsiter = _globalvar.find(N(udfssupply));
            auto bpiter = _globalvar.find(N(bpsupply));
            
            int64_t  udfsfee = 0;
            if( ac.balance.amount < udfsiter->val )       
            {
                udfsfee =  udfsiter->val - ac.balance.amount; 
            }

            print( "udfs : ",  udfsiter->val , "  ", time_now, "  ", bpiter->val," udfsfee ", udfsfee,  "\n");
            
            auto  total = udfsfee + bpiter->val; 

            INLINE_ACTION_SENDER(uosio::token, issue)( N(uosio.token), {{ _self,N(active)}},
                                                   { _self, asset(total), std::string("issue tokens for producer pay and udfs")} );
/*
            INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), { _self,N(active)},
                                                        { _self, N(uosio.bvpay), asset(bpiter->val, uosiosystem::system_token_symbol), "transfer bvpay" } );
            if(udfsfee>0)
            {
                INLINE_ACTION_SENDER(uosio::token, transfer)( N(uosio.token), { _self,N(active)},
                                                        { _self, N(udfsreceiver), asset(udfsfee, uosiosystem::system_token_symbol), "udfs fee" } );
            }
*/
            set_global_var(N(issuetime),fielditer->val+interval);
        }
        else {
            return; 
        }
#endif
    }

}// namespace uosio


#define UOSIO_UDFS_ABI( TYPE, MEMBERS ) \
extern "C" { \
   void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
      auto _self = receiver; \
      if( action == ("onerror"_n).value) { \
         uosio_assert(code == ("uosio"_n).value, "onerror action's are only valid from the \"uosio\" system account"); \
      } \
      if(action == ("transfer"_n).value) { \
         if(code != ("uosio.token"_n).value){\
           return; \
         }\
         uosio::execute_action( uosio::name(receiver), uosio::name(code), &uosio::uosio_udfs::transfer); \
         return; \
      } \
      if(code == _self || action == ("onerror"_n).value ) { \
         switch( action ) { \
            UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
         } \
      } \
   } \
} \


UOSIO_UDFS_ABI(uosio::uosio_udfs,(setdata)(clear)(modifymember)(transfer)(voter)(issuetoken)(addnode)(feetableop))


