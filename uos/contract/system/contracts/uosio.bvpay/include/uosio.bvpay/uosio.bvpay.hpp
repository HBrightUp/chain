/**
 *  @file
 *  @copyright defined in uos/LICENSE.txt
 */
#pragma once

#include <uosiolib/asset.hpp>
#include <uosiolib/uosio.hpp>
#include <uosiolib/symbol.hpp>
#include <uosio.system/uosio.system.hpp>

#include <string>
#include <vector>

namespace uosio {

    static constexpr uint64_t BP_VT_PAY_CYCLE = 30 * 24 * 3600;
    #define SYSTEM_COIN ("UOS")

    class [[uosio::contract("uosio.bvpay")]] uosbvpay : public contract {

    public:
        using contract::contract;
        uosbvpay(name receiver, name code,  uosio::datastream<const char*> ds);

        [[uosio::action]]
        void start(name from);

        [[uosio::action]]
        void fillbvpay(name call);

        using start_action = uosio::action_wrapper<"start"_n, &uosbvpay::start>;
        using fillbvpay_action = uosio::action_wrapper<"fillbvpay"_n, &uosbvpay::fillbvpay>;

    public:
         inline asset get_bpvtpay(name call) const;
        void transfer(name from, name to, asset quantity, std::string memo);

    private:

        ///@abi table bpvtpay i64
        struct [[uosio::table]] bp_and_vt_pay{
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
        struct [[uosio::table]] bv_real_pay{
            name             owner;
            asset                    quantity;
            uint64_t primary_key()const { return owner.value; }
            UOSLIB_SERIALIZE( bv_real_pay, (owner)(quantity))
        };
        typedef uosio::multi_index<"bvrealpay"_n , bv_real_pay> bvrealpay;

        //come from uosio.token
        struct [[uosio::table]] account {
            asset    balance;

            uint64_t primary_key()const { return balance.symbol.code().raw(); }
        };

        typedef uosio::multi_index< "accounts"_n, account > accounts;


    private:
        bpvtpay _bpvtpay;
        bvrealpay _bvrealpay;
        uint64_t trash_time;
    };

    asset uosbvpay::get_bpvtpay(name call) const {
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



} /// namespace uosio
