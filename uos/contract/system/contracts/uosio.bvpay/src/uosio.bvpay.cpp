/**
 *  @file
 *  @copyright defined in uos/LICENSE.txt
 */

#include <uosio.bvpay/uosio.bvpay.hpp>
#include <uosiolib/transaction.hpp>
#include <uosio.token/uosio.token.hpp>
#include<uosiolib/dispatcher.hpp>




namespace uosio{

    uosbvpay::uosbvpay(name receiver, name code,  uosio::datastream<const char*> ds):
        contract(receiver, code, ds),
        _bpvtpay(_self, _self.value),
        _bvrealpay(_self, _self.value){
            trash_time = 0;
            auto _global = uosiosystem::global_state_singleton("uosio"_n, ("uosio"_n).value);
            if(_global.exists()){
                auto _gstate =  _global.get();
                trash_time = _gstate.thresh_activated_stake_time.sec_since_epoch();
            }
    }

    /// @abi action
    void uosbvpay::fillbvpay(name call) {
        require_auth(call);
        auto bpvtp = _bpvtpay.find(0);
        auto bpvtp1 = _bpvtpay.find(1);
        if(bpvtp == _bpvtpay.end()){
            return;
        }
        int64_t real_pay = 0;
        uosio_assert(now() >  bpvtp->laster_line , "get_bpvtpay time error");
        if(now() < bpvtp->dead_line){
            double pay = (now() - bpvtp->laster_line ) *  bpvtp->quantity.amount * 1.0 / (bpvtp->dead_line - bpvtp->laster_line);
            real_pay = int64_t(pay);
            real_pay += bpvtp->inherit.amount;
            _bpvtpay.modify(bpvtp, same_payer,[&](auto &a){
                a.quantity.amount -= real_pay;
                a.inherit = asset(0,symbol{SYSTEM_COIN, 4});
                a.laster_line = now();
            });
        } else if( (now() >= bpvtp->dead_line) && (now() < bpvtp1->dead_line) ){
            real_pay = bpvtp->quantity.amount;
            real_pay += bpvtp->inherit.amount;
            uosio_assert(now() >  bpvtp1->laster_line , "get_bpvtpay time error");
            double pay = (now() - bpvtp1->laster_line ) *  bpvtp1->quantity.amount * 1.0 / (bpvtp1->dead_line - bpvtp1->laster_line);
            int64_t next_pay = int64_t(pay);
            real_pay += next_pay;
            _bpvtpay.modify(bpvtp, same_payer,[&](auto &a){
                a.start_line += BP_VT_PAY_CYCLE;
                a.laster_line = now();
                a.dead_line += BP_VT_PAY_CYCLE;
                a.quantity = bpvtp1->quantity;
                a.quantity.amount -= next_pay;
                a.inherit = asset(0, symbol{SYSTEM_COIN, 4});
            });

            _bpvtpay.modify(bpvtp1, same_payer,[&](auto &a){
                a.start_line += BP_VT_PAY_CYCLE;
                a.laster_line +=   BP_VT_PAY_CYCLE;
                a.dead_line += BP_VT_PAY_CYCLE;
                a.quantity =  asset(0, symbol{SYSTEM_COIN, 4});
            });


        }  else if(now() >= bpvtp1->dead_line){
            real_pay = bpvtp->quantity.amount;
            real_pay += bpvtp->inherit.amount;
            real_pay += bpvtp1->quantity.amount;
            _bpvtpay.erase(bpvtp);
            _bpvtpay.erase(bpvtp1);
        }
        bpvtp = _bpvtpay.find(0);
        bpvtp1 = _bpvtpay.find(1);
        if(bpvtp!=_bpvtpay.end()){
            if( (bpvtp->quantity.amount == 0 ) && (bpvtp1->quantity.amount == 0) ){
                _bpvtpay.erase(bpvtp);
                _bpvtpay.erase(bpvtp1);
            }
        }
        auto realpaydb = _bvrealpay.find(_self.value);
        if(realpaydb == _bvrealpay.end()){
            _bvrealpay.emplace(_self,[&](auto &a){
                a.owner = _self;
                a.quantity = asset(real_pay , symbol{SYSTEM_COIN, 4});
            });
        } else {
            _bvrealpay.modify(realpaydb, same_payer,[&](auto &a){
                a.quantity += asset(real_pay , symbol{SYSTEM_COIN, 4});
            });
        }
    }


    /// @abi action
    void uosbvpay::start(name from) {
        uosio_assert(from == "uosio"_n , "only uosio can call uosio.bvpay's start action");
        uosio_assert(trash_time > 0 , "trash_time must bigger than zero");
        require_auth("uosio"_n);
        auto bpvtp =  _bpvtpay.find(0);
        //asset quantity = uosio::token(receiver, code, ds).get_balance(_self, symbol(SYSTEM_COIN, 4).raw());

        accounts bv_acnts( "uosio.token"_n, _self.value );
        auto bv_itr = bv_acnts.find( symbol_code{"UOS"}.raw() );
        uosio_assert(bv_itr != bv_acnts.end(), "can't find uosio.bvpay in accounts");

        asset quantity = bv_itr->balance;

        if(bpvtp == _bpvtpay.end() && quantity.amount>0){
            _bpvtpay.emplace(_self,[&](auto &a){
                a.id = 0;
                a.start_line = now();
                a.laster_line = now();
                a.dead_line = now() + BP_VT_PAY_CYCLE;
                a.quantity = quantity;
                a.inherit = asset(0, symbol{SYSTEM_COIN, 4});
            });
            _bpvtpay.emplace(_self,[&](auto &a){
                a.id = 1;
                a.start_line = now() + BP_VT_PAY_CYCLE;
                a.laster_line = now() + BP_VT_PAY_CYCLE;
                a.dead_line = now() + BP_VT_PAY_CYCLE + BP_VT_PAY_CYCLE;
                a.quantity = asset(0, symbol{SYSTEM_COIN, 4});
                a.inherit = asset(0, symbol{SYSTEM_COIN, 4});
            });
        } else {
            return;
        }


    }


    /// @abi action
    void uosbvpay::transfer(name from, name to, asset quantity, std::string memo) {
        uosio_assert(  quantity.symbol.raw() == symbol{"UOS", 4}.raw(), "must use system coin");
        if(from == _self){
            require_auth(_self);
            auto& realpaydb = _bvrealpay.get(_self.value);
            uosio_assert(realpaydb.quantity >= quantity , "realpaydb.quantity >= quantity");
            _bvrealpay.modify(realpaydb, same_payer,[&](auto &a){
                a.quantity -= quantity;
            });
            return;
        }
        uosio_assert(to == _self , "some one must transfer to this contract");

        if(trash_time == 0){
            return;
        }
        auto bpvtp =  _bpvtpay.find(0);
        if(bpvtp == _bpvtpay.end()){
            _bpvtpay.emplace(_self,[&](auto &a){
                a.id = 0;
                a.start_line = now();
                a.laster_line = now();
                a.dead_line = now() + BP_VT_PAY_CYCLE;
                a.quantity = quantity;
                a.inherit = asset(0, symbol{SYSTEM_COIN, 4});
            });
            _bpvtpay.emplace(_self,[&](auto &a){
                a.id = 1;
                a.start_line = now() + BP_VT_PAY_CYCLE;
                a.laster_line = now() + BP_VT_PAY_CYCLE;
                a.dead_line = now() + BP_VT_PAY_CYCLE + BP_VT_PAY_CYCLE;
                a.quantity = asset(0,symbol{SYSTEM_COIN, 4});
                a.inherit = asset(0, symbol{SYSTEM_COIN, 4});
            });
        } else {
            uosio_assert(now() >= bpvtp->laster_line , " error : bpvtp->laster_line");
            auto bpvtp1 = _bpvtpay.find(1);
            if(now() < bpvtp->dead_line){
                _bpvtpay.modify(bpvtp1, same_payer,[&](auto &a){
                    a.quantity += quantity;
                });
            } else if( (now() >= bpvtp1->start_line) && (now() < bpvtp1->dead_line) ){
                _bpvtpay.modify(bpvtp, same_payer,[&](auto &a){
                    a.start_line += BP_VT_PAY_CYCLE;
                    a.laster_line +=  BP_VT_PAY_CYCLE;
                    a.dead_line += BP_VT_PAY_CYCLE;
                    a.inherit += a.quantity;
                    a.quantity = bpvtp1->quantity;
                });

                _bpvtpay.modify(bpvtp1, same_payer,[&](auto &a){
                    a.start_line += BP_VT_PAY_CYCLE;
                    a.laster_line +=  BP_VT_PAY_CYCLE;
                    a.dead_line += BP_VT_PAY_CYCLE;
                    a.quantity = quantity;

                });

            } else if(now() >= bpvtp1->dead_line){
                _bpvtpay.modify(bpvtp, same_payer,[&](auto &a){
                    a.start_line = now();
                    a.laster_line = now();
                    a.dead_line = now() + BP_VT_PAY_CYCLE;
                    a.inherit += bpvtp1->quantity;
                    a.inherit += a.quantity;
                    a.quantity = quantity;
                });
                _bpvtpay.modify(bpvtp1, same_payer,[&](auto &a){
                    a.start_line = now() + BP_VT_PAY_CYCLE;
                    a.laster_line = now() + BP_VT_PAY_CYCLE;
                    a.dead_line = now() + BP_VT_PAY_CYCLE + BP_VT_PAY_CYCLE;
                    a.quantity = asset(0, symbol{SYSTEM_COIN, 4});
                });
            }
        }
    }



}/// namespace uosio


#define UOSIO_BVPAY_ABI( TYPE, MEMBERS ) \
extern "C" { \
    void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
        auto self = receiver; \
        if( action == ("onerror"_n).value ) { \
            /* onerror is only valid if it is for the "uosio" code account and authorized by "uosio"'s "active permission */ \
            uosio_assert(code == ("uosio"_n).value, "onerror action's are only valid from the \"uosio\" system account"); \
        } \
        if(action == ("transfer"_n).value) { \
            if(code != ("uosio.token"_n).value){\
            return; \
            }\
            uosio::execute_action( uosio::name(receiver), uosio::name(code), &uosio::uosbvpay::transfer); \
            return; \
        } \
        if(code == self || action == ("onerror"_n).value ) { \
            switch( action ) { \
                UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
            } \
        } \
   } \
} \

UOSIO_BVPAY_ABI( uosio::uosbvpay , (fillbvpay)(transfer)(start))


