#include<uosiolib/uosio.hpp>
#include <uosiolib/crypto.h>
#include <uosiolib/transaction.hpp>
#include "convert.hpp"
#include <cmath>


namespace convertsys {
    using namespace uosio;

    convert::convert(name receiver, name code,  uosio::datastream<const char*> ds)
    :contract(receiver, code, ds),
    _global_var(_self, _self.value),
    _memberreward(_self, _self.value),
    _utuosstate(_self, _self.value),
    _utuosvoter(_self, _self.value),
    _utuostr(_self, _self.value),
    _uosuttr(_self, _self.value),
    _uosutowner(_self, _self.value),
    _uosutstate(_self, _self.value),
    _rate(_self, _self.value)
    {

    }

    void convert::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {

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

    ACTION convert::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
      require_auth(_self);
      set_global_var(key, val, op);
    }

    void convert::init() {
        require_auth(_self);
        
        set_global_var("switch"_n, 0);

        const asset token_supply = get_supply("uosio.token"_n, symbol_code("UOS") );
        print("token_supply:", token_supply.amount, "\n");

        //set_global_var("supply"_n, 200000000 + token_supply.amount);
        set_global_var("supply"_n, 2233562366977);

        if(_rate.begin() == _rate.end()) {
            _rate.emplace(_self, [&](auto &r){
                r.key = 1;
                r.uttouos = std::string("1.0000");
                r.uostout = std::string("1.0000");
            });
        }
        
    }

    void convert::modifymember(std::vector<name> & members) {
        require_auth(_self);
        uosio_assert(members.size() > 2 , "union-member must be greater or equal than 3");

        auto sta = _utuosstate.find(_self.value);
        if(sta == _utuosstate.end()){
            _utuosstate.emplace(_self,[&](auto &a){
                a.owner = _self;
                a.members = members;
            });
        }else {
            _utuosstate.modify(sta , _self, [&](auto &a){
                a.members = members;
            });
        }

        for(auto itr = _utuostr.begin() ; itr != _utuostr.end() ;){
            itr = _utuostr.erase(itr);
        }
        
        for(auto itr = _utuosvoter.begin() ; itr !=_utuosvoter.end() ; ){
            itr = _utuosvoter.erase(itr);
        }

        auto uotst = _uosutstate.find(_self.value);
        if(uotst == _uosutstate.end()) {
            _uosutstate.emplace(_self, [&](auto &a) {
                a.owner = _self;
                a.max_id = 0;
            });
        }

        for(auto itr = _memberreward.begin() ; itr != _memberreward.end(); ) {

            if(itr->reward.amount == 0){
                itr = _memberreward.erase(itr);
                continue;
            }

            if(itr->active) {
                _memberreward.modify(itr , _self, [&](auto &a){
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
                    a.reward = asset(0,CORE_SYMBOL);
                    a.active = 1;
                });
            }
            else{
                _memberreward.modify(mem, _self, [&](auto &a){
                    a.active = 1;
                });
            }
        }
    }

    double convert::get_area(const double x) {
        print("a1+++ \n");
        print("x: ", x, "\n");

        const auto supply_iter = _global_var.find(("supply"_n).value);
        uosio_assert(supply_iter != _global_var.end(), "field supply not exist.");
        const double supply = static_cast<double>(supply_iter->val) / 10000;

        auto calc_area = [](double x, double supply) {
            double area = 45 * log(140) * std::pow(259, std::log((x - supply) / 100000000) / std::log(140)) / log(259) + (x - supply) / 100000000;
            return 100000000 * area;
        };

        double result = 0;

        print("a2+++ \n");
        print("x: ", x, " supply: ", supply, "\n");

        if(x < supply) {
            print("a3+++ \n");
            result = x * 1 ;
        }
        else if(x == supply){
            print("a4+++ \n");
            result = supply * 1 ;
        }
        else {
            print("a5+++ \n");
            result = calc_area(x, supply) + supply * 1 ; 
        }

        print("result: ", result, "\n");
        print("right_uos_value: ", result  - supply, "\n");

        return result; 
    }

    double convert::get_fx(const double x) {

        const auto supply_iter = _global_var.find(("supply"_n).value);
        uosio_assert(supply_iter != _global_var.end(), "field supply not exist.");
        const double supply = static_cast<double>(supply_iter->val) / 10000;

        auto calc_y = [](double x, double supply) {
            uosio_assert(x != supply, "invalid parameter 'x' ");

            double y = 45 * std::pow(1.85, std::log((x - supply) / 100000000) / std::log(140)) + 1;
            return 100000000 * y;
        };
        
        double result = 0;

        if(x <= supply) {
            result =  1;
        }
        else {
            result = calc_y(x, supply);
        }

        print("c1+++ \n");
        print("x: ", x, " result: ", result, "\n");
        return result;
    }

    double convert::get_middile_value(double x1,  const double area,  const double x1_area) {

        const double precision = 0.0001;
        double area_mid = 0;
        int64_t times = 0;

        double x2_end = x1 + area / get_fx(x1);
        double x_mid = (x1 + x2_end) / 2;

        print("b1++++ \n");
        print("x1: ", x1, " area: ", area, " x1_area: ", x1_area, " x2_end: ", x2_end, " x_mid: ", x_mid, "\n");

        while(true) {
            area_mid = get_area(x_mid);

            print("b2++++ \n");
            print("area_mid: ", area_mid, " x1_area: ", x1_area, " area: ", area, " x_mid: ", x_mid, "\n");
            if (abs(area_mid - x1_area  - area  ) <= precision) {
                break;
            }

            if (area_mid - x1_area  - area  > 0) {
                x2_end = x_mid;
                x_mid = (x1 + x_mid) / 2;
            }
            else {
                x1 = x_mid;
                x_mid = (x_mid + x2_end) / 2;
            }

            if(++times >= 100) {
                uosio_assert(false, "too many times");
            }
        }

        return x_mid;
    }

    double convert::calc_uos(const double area){

        print("a++++ \n");
        const asset x1_asset = get_supply("uosio.token"_n, symbol_code("UOS") );
        double x1 = static_cast<double>(x1_asset.amount) / 10000;
        const double x1_area = get_area(x1);

        auto supply_iter = _global_var.find(("supply"_n).value);
        uosio_assert(supply_iter != _global_var.end(), "field supply not exist.");
        double supply = static_cast<double>(supply_iter->val) / 10000;
        const double supply_area = get_area(supply);

        const double ahead_area = x1_area + area ;
        double result = 0;
        //double x2 = 0;

        print("ahead_area: ", ahead_area, " supply_area: ", supply_area, "\n");
        
        if(ahead_area <= supply_area){
            print("b++++ \n");
            result = area * 10000;
        }
        else {
            const double right_area = x1_area + area - supply_area;

            print("right_area: ", right_area, " x1_area: ", x1_area, " area: ", area, " supply_area: ", supply_area, "\n");

            const double x2 = get_middile_value(supply , right_area , supply_area);

            print("c++++ \n");
            print("x1_area: ", x1_area, " supply_area: ", supply_area, "\n");

            if(x1_area >= supply_area) {
                result = (x2  - x1) * 10000;
                print("d++++ \n");
            }
            else {
                result = supply_iter->val - x1_asset.amount;
                result += x2 * 10000  - supply_iter->val;
                print("e++++ \n");
            }
        }

        print("f++++ \n");
        print("result: ", result, "\n");  

        return result;
    }

    double convert::calc_ut(const double distance) {
        const asset x1_asset = get_supply("uosio.token"_n, symbol_code("UOS") );

        const auto supply_iter = _global_var.find(("supply"_n).value);
        uosio_assert(supply_iter != _global_var.end(), "field supply not exist.");
        const double supply = static_cast<double>(supply_iter->val) / 10000;

        const double x1 = static_cast<double>(x1_asset.amount) / 10000;
        const double x2 = x1 - distance;
        double result = 0;

        print("d1+++ \n");
        print("x1: ", x1, " supply: ", supply, " distance: ", distance,  "\n");
        
        if(x1 <= supply) {
            print("d2+++ \n");
            result = distance * 1.0;
        }
        else {
            print("d3+++ \n");
            print("x2: ", x2, " supply: ", supply, " distance: ", distance,  "\n");

            result = get_area(x1) - get_area(x2);    
        }

         print("d4+++ \n");
         print("result: ", result, "\n");
        
        return result;
    }

    void convert::uttouosvote(const name voter, const std::string tr_id, const name owner, const int64_t amount, const std::string memo) {
        require_auth(voter);

        uosio_assert(amount > 0 && amount <= asset::max_amount, "amount  error");
        
        auto& sta = _utuosstate.get(_self.value);
        auto it = std::find(sta.members.begin(), sta.members.end(), voter);
        uosio_assert(it != sta.members.end(),"voter not in union");

        uosio_assert(is_account(owner), "uos haven't this account");
        uosio_assert(tr_id.size() < 70 , "id size must be less than 70");

        std::string buffer(tr_id);
        buffer += std::to_string(owner.value) + std::to_string(amount);
        uint64_t tr_hash = caculate_hash_64(buffer.c_str() , buffer.size());
        auto ver = _utuosvoter.find(voter.value);

        if(ver == _utuosvoter.end()){
            _utuosvoter.emplace(_self,[&](auto &a){
                a.voter = voter;
                a.tr_hash = tr_hash;
                a.active_time = now();
                a.tr_id = tr_id;
            });
        }else{
            auto & tr = _utuostr.get(ver->tr_hash , "can find tr");
            if(tr.votes == 1){
                _utuostr.erase(tr);
            }else{
                _utuostr.modify(tr, _self, [&](auto &a){
                    a.votes --;
                });
            }
            _utuosvoter.modify(ver, _self, [&](auto &a){
                a.tr_hash = tr_hash;
                a.active_time = now();
                a.tr_id = tr_id;
            });
        }
        auto tr = _utuostr.find(tr_hash);
        if(tr == _utuostr.end()){
            _utuostr.emplace(_self,[&](auto &a){
                a.tr_hash = tr_hash;
                a.tr_id = tr_id;
                a.owner = owner;
                a.amount = amount;
                a.votes = 1;
            });
        }else{
            _utuostr.modify(tr, _self, [&](auto &a){
                a.votes ++;
            });
        }

        tr = _utuostr.find(tr_hash);

        if(tr->votes > int32_t(sta.members.size() * 2 / 3)){

            double real_uos = calc_uos(static_cast<double>(tr->amount) / 10000);
            double real_rate_nofee = static_cast<double>(tr->amount)  / real_uos;
            double real_rate_fee = tr->amount  / (static_cast<double>(real_uos) * (1 + 1.0 / 1000) );
            
            if(memo.size() == 0){
                
                action(permission_level{_self, "active"_n}, "uosio.token"_n,
                       "issue"_n, std::make_tuple(tr->owner, asset(real_uos, CORE_SYMBOL),
                                                 std::string("ut_to_uos_utid:")+ tr->tr_id))
                        .send();

                auto rate_iter = _rate.begin();
                uosio_assert(rate_iter != _rate.end(), "rate table no data");

                _rate.modify(rate_iter, _self,[&](auto &r){
                    r.uttouos = std::to_string(real_rate_nofee);
                    //r.uostout = std::to_string(1 / real_rate_fee );
                });
            }

            action(permission_level{_self, "active"_n}, _self,
                       "uttouospass"_n, std::make_tuple(voter, tr_id,
                                                       owner , amount , memo))
                        .send();

            _utuosstate.modify(sta, _self,[&](auto &a){
                    a.laster_id = tr->tr_id;
                    a.laster_owner = tr->owner;
                    a.amount = amount;
                    a.tr_hash = tr->tr_hash;
                });

            for(auto itr = _utuosvoter.begin() ; itr != _utuosvoter.end() ;){
                itr = _utuosvoter.erase(itr);
            }
            for(auto itr = _utuostr.begin() ; itr != _utuostr.end() ; ){
                itr = _utuostr.erase(itr);
            }
        }
    }

    uint64_t convert::caculate_hash_64(const char *buf, uint32_t size) {
        checksum160 calc_hash;
        calc_hash = ripemd160(buf, size );
        auto bytearr =  calc_hash.extract_as_byte_array();
        uint64_t a = * reinterpret_cast<uint64_t *>(bytearr.data());
        uint64_t b = * reinterpret_cast<uint64_t *>(bytearr.data()+8);
        uint32_t c = * reinterpret_cast<uint32_t *>(bytearr.data()+16);
        return a^b^c;
    }

    void convert::setousutid(name voter, name tr, std::string id) {
        
        require_auth(voter);
        auto& sta = _utuosstate.get(_self.value);
        auto it = std::find(sta.members.begin(), sta.members.end(), voter);

        uosio_assert(id.size() <= 64 ,"ut id size maxmum is 64 ");
        uosio_assert(it != sta.members.end(),"voter was not in union");

        uint32_t index = uint32_t(it - sta.members.begin());
        auto ustr =  _uosuttr.find(tr.value);
        uosio_assert(ustr!=_uosuttr.end(),"No information for this user");

        _uosuttr.modify(ustr, _self,[&](auto &a) {
            a.uosutid[index] = name{sta.members[index]}.to_string() + std::string(" : ");
            a.uosutid[index] += id;
        });

        uint32_t size = 0;
        ustr =  _uosuttr.find(tr.value);
        for(auto itr = ustr->uosutid.begin();itr!= ustr->uosutid.end();itr++){
            if(*itr == std::string()){
                continue;
            }
            uint32_t pos = name{sta.members[uint32_t(itr - ustr->uosutid.begin())]}.to_string().size() + 3;
            if(id.compare(itr->substr(pos)) == 0){
                size++;
            }
        }

        if(size > sta.members.size() * 2 / 3){

            auto t_data = make_tuple(tr, id, ustr->ut_address, ustr->amount, ustr->id);
            action(permission_level{_self, "active"_n}, _self,
                   "uostoutpass"_n, t_data)
                    .send();
           
            auto owner_iter = _uosutowner.find(tr.value);

            auto rate_iter = _rate.begin();
            uosio_assert(rate_iter != _rate.end(), "rate table no data");

            auto s_iter = _global_var.find(("supply"_n).value);
            uosio_assert(s_iter != _global_var.end(), "field supply not exist");


            const auto realut_iter = _global_var.find(("realut"_n).value);
            uosio_assert(realut_iter != _global_var.end(), "field realut not exist.");
#if 0
            const double real_rate_fee = owner_iter->old_amount / static_cast<double>(realut_iter->val);
            
            _rate.modify(rate_iter, _self,[&](auto &r){
                r.converted += owner_iter->old_amount;
                r.uostout = std::to_string( real_rate_fee );
            }); 
#endif
            uosio::asset dispose(owner_iter->old_amount, CORE_SYMBOL);
            auto t_data1 = make_tuple(_self, dispose, std::string(""));
            action(permission_level{_self, "active"_n}, "uosio.token"_n,
                "withdraw"_n, t_data1)
                    .send();
            
            if(owner_iter != _uosutowner.end()){
                _uosutowner.erase(owner_iter);
            }
            _uosuttr.erase(ustr);
        }
    }

    /// @abi action
    void convert::setuttr(name voter, name tr, std::string uttr) {
        require_auth(voter);

        uosio_assert(uttr != std::string() , "trac can't be empty");
        uosio_assert(uttr.size() < 20*1024,"Sign up to 20K bytes");

        auto& sta = _utuosstate.get(_self.value);
        auto it = std::find(sta.members.begin(), sta.members.end(), voter);
        uosio_assert(it != sta.members.end(),"voter was not in union");

        uint32_t index = uint32_t(it - sta.members.begin());
        auto &ustr =  _uosuttr.get(tr.value,"No information for this user");

        _uosuttr.modify(ustr, _self,[&](auto &a){
            a.trac[index] = name{sta.members[index]}.to_string() + std::string(" : ");
            a.trac[index] += uttr;
        });
    }

    void convert::updaterate(const name voter, const name tr) {
        require_auth(voter);
        print("updaterate() voter: ", voter, "tr", tr,  "\n");

        const auto lastaccount_iter = _global_var.find(("lastaccount"_n).value);
        //uosio_assert(lastaccount_iter != _global_var.end(), "field lastaccount not exist.");

        if(lastaccount_iter != _global_var.end() && lastaccount_iter->val == tr.value) {
            print("calc rate already for same account \n");
            //return ;
        }

        auto owner_iter = _uosutowner.find(tr.value);
        if(owner_iter == _uosutowner.end()) {
            print("no convert from uos to ut  \n");
            return ;
        }

        auto rate_iter = _rate.begin();
        uosio_assert(rate_iter != _rate.end(), "rate table no data");

        auto s_iter = _global_var.find(("supply"_n).value);
        uosio_assert(s_iter != _global_var.end(), "field supply not exist");


        double real_ut = 0;
        //const asset token_supply = get_supply("uosio.token"_n, symbol_code("UOS"));
        const int64_t uos_converted = rate_iter->converted;

        if(uos_converted < s_iter->val &&  uos_converted + owner_iter->old_amount <= s_iter->val) {
            real_ut = static_cast<double>(owner_iter->new_amount / 10000);
        }
        else if(uos_converted < s_iter->val &&  uos_converted + owner_iter->old_amount > s_iter->val) {
            int64_t extra = uos_converted + owner_iter->new_amount - s_iter->val;
            if(extra < 0) {
                extra = 0;
            }

            double extra_ut = calc_ut(static_cast<double>(extra) / 10000);
            real_ut = extra_ut + static_cast<double>(owner_iter->new_amount - extra) / 10000;

        }
        else{
            real_ut = calc_ut(static_cast<double>(owner_iter->new_amount) / 10000);
        }
        
        const double real_rate_fee = owner_iter->new_amount / (real_ut * 10000);
        //const double real_rate_fee = owner_iter->old_amount / static_cast<double>(realut_iter->val);
        
        print("real_ut: ", real_ut, " real_rate_fee: ", real_rate_fee, "\n");
        
        _rate.modify(rate_iter, _self,[&](auto &r){
            r.converted += owner_iter->old_amount;
            r.uostout = std::to_string( real_rate_fee );
        }); 

        uint64_t realut = static_cast<uint64_t>(real_ut *10000);
        set_global_var("realut"_n, realut);
        set_global_var("lastaccount"_n, tr.value);
    }

    void convert::clear() {
        require_auth(_self);
#if 0
        while(_uosutowner.begin() != _uosutowner.end()) {
            _uosutowner.erase(_uosutowner.begin());
        }

        while(_uosuttr.begin() != _uosuttr.end()) {
            _uosuttr.erase(_uosuttr.begin());
        }

        while(_rate.begin() != _rate.end()) {
            _rate.erase(_rate.begin());
        }

#endif

   }

   void convert::modifyuos(const uint64_t converted) {
        require_auth(_self);

        auto rate_iter = _rate.begin();
        uosio_assert(rate_iter != _rate.end(), "can't find rate table");

        _rate.modify(rate_iter, _self,[&](auto &r){
            r.converted = converted;
            
        });
   }

    void convert::transfer(const name& from, const name& to, const asset& quantity, const std::string& memo) {
        if(from == _self){
            require_auth(_self);
            return;
        }
        require_auth(from);
        uosio_assert(  to == _self , "must transfer to this contract");
        uosio_assert(  quantity.symbol == CORE_SYMBOL, "must use system coin");
        if(memo.size() == 4){
            auto t_data = make_tuple(_self,quantity,memo);
            action(permission_level{_self, "active"_n}, "uosio.token"_n,
                   "withdraw"_n, t_data)
                    .send();
            return;
        }
        uosio_assert(  memo.size() < 40 , "The ut address max size equal 40");
        uosio_assert(  memo.size() > 10 , "The ut address min size equal 10");
        uosio_assert(  quantity.amount >= 10 * 10000 , "minimum transfer 10 uos");
        uosio_assert(  quantity.amount <= 10000 * 10000 , "maxmum transfer 10000 uos");
        asset old_quantity = quantity;
        asset new_quantity = caculate_reward(quantity , memo);

        uosio_assert(_uosutowner.find(from.value) == _uosutowner.end(), "user has some withdraw not complete");
        _uosutowner.emplace(_self,[&](auto & a){
            a.owner = from;
            a.old_amount = old_quantity.amount;
            a.new_amount = new_quantity.amount;
        });

        auto t_data = make_tuple(from, new_quantity, memo);
        uosio::transaction tr_delay;
        tr_delay.actions.emplace_back( permission_level{ _self, "active"_n }, _self, "delaytf1"_n, t_data );
        tr_delay.delay_sec = 1;
        cancel_deferred( from.value );
        tr_delay.send(from.value , _self , 1);
    }

    asset convert::caculate_reward(uosio::asset quantity, std::string memo) {
        uosio::asset reward(0,CORE_SYMBOL);
        
        reward = quantity/1000;
        if(reward.amount > 1000000){
            reward.amount = 1000000;
        }

        uosio_assert(reward.amount > 0 , "union member reward too small");
        quantity -= reward ;
        uosio_assert(quantity.amount > 0 , "real withdraw : too small");


        auto sta = _utuosstate.get(_self.value);
        auto members_by_active =  _memberreward.get_index<"byactive"_n>();

        int member_size = sta.members.size();
        uosio_assert(member_size > 0 ,"no union members");

        int64_t extra = reward.amount % member_size;
        int64_t gain = reward.amount / member_size;


        for(auto itr = members_by_active.rbegin() ; itr != members_by_active.rend(); itr++ ){
            if(!itr->active){
                break;
            }
            if(itr == members_by_active.rbegin()){
                _memberreward.modify(*itr, _self,[&](auto &a){
                    a.reward.amount += extra;
                    a.reward.amount += gain;

                });
                continue;
            }
            if((itr->active) && (gain > 0)){
                _memberreward.modify(*itr, _self,[&](auto &a){
                    a.reward.amount += gain;

                });
            }
        }
        return quantity;
    }


     /// @abi action
    void convert::delaytf(name from, asset quantity, int64_t tpn , std::string memo ){
        require_auth(_self);
        auto& sta = _utuosstate.get(_self.value);
        auto &uotst = _uosutstate.get(_self.value);
        auto uot = _uosuttr.find(from.value);
        if(uot == _uosuttr.end() ){
            _uosuttr.emplace(_self,[&](auto & a){
                a.owner = from;
                a.tpn = tpn;
                a.amount = quantity.amount;
                a.id = uotst.max_id + 1;
                a.ut_address = memo;
                a.uosutid = std::vector<std::string>(sta.members.size());
                a.trac =  std::vector<std::string>(sta.members.size());
            });
            _uosutstate.modify(uotst, _self,[&](auto &a){
                a.max_id ++;
            });
        }
    }

    void convert::delaytf1(name from, uosio::asset quantity, std::string memo) {
        require_auth(_self);
        int64_t tpn = tapos_block_num();
        auto t_data = make_tuple(from, quantity, tpn ,memo );
        uosio::transaction out;
        out.actions.emplace_back( permission_level{ _self, "active"_n }, _self, "delaytf"_n, t_data );
        //11*6 + 11*6 + 11*6
        out.delay_sec = 200;
        cancel_deferred( from.value );
        out.send(from.value , _self , 1);
    }

    // @abi action
    void convert::uttouospass(name voter,std::string tr_id , name owner , int64_t amount , std::string memo){
        
    }

    // @abi action
    void convert::uostoutpass(name user ,std::string ut_id ,std::string ut_address, int64_t amount ,uint64_t buffer_id){
        print(name{user}," uostoutpass id is : ",ut_id.c_str());
    }

} // end convertsys namespace 

#define UOSIO_DISPATCH_CONVERT( TYPE, MEMBERS ) \
extern "C" { \
   void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
      auto self = receiver; \
      if( code == self || code == ("uosio.token"_n).value ) { \
      	 if( code == ("uosio.token"_n).value && action == ("transfer"_n).value ){ \
            uosio::execute_action(uosio::name(receiver), uosio::name(code), &convertsys::convert::transfer); \
            return ; \
      	 } \
         switch( action ) { \
            UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
         } \
      } \
   } \
}

UOSIO_DISPATCH_CONVERT(convertsys::convert, (init)(setdata)(modifymember)(uttouosvote)(setousutid)(clear) (uostoutpass)(uttouospass)(delaytf)(delaytf1)(setuttr)(transfer)(updaterate)(modifyuos))
