#pragma once

#include <uosiolib/asset.hpp>
#include <uosiolib/uosio.hpp>

#include <string>

namespace uosiosystem {
   class system_contract;
}

namespace uosio {

   using std::string;

   class [[uosio::contract("uosio.token")]] token : public contract {
      public:
         using contract::contract;

         [[uosio::action]]
         void create( name   issuer,
                      asset  maximum_supply);

         [[uosio::action]]
         void issue( name to, asset quantity, string memo );

         [[uosio::action]]
         void retire( name owner, asset quantity, string memo );

         [[uosio::action]]
         void transfer( name    from,
                        name    to,
                        asset   quantity,
                        string  memo );

         [[uosio::action]]
         void open( name owner, const symbol& symbol, name ram_payer );

         [[uosio::action]]
         void close( name owner, const symbol& symbol );

         static asset get_supply( name token_contract_account, symbol_code sym_code )
         {
            stats statstable( token_contract_account, sym_code.raw() );
            const auto& st = statstable.get( sym_code.raw() );
            return st.supply;
         }

         static asset get_balance( name token_contract_account, name owner, symbol_code sym_code )
         {
            accounts accountstable( token_contract_account, owner.value );
            const auto& ac = accountstable.get( sym_code.raw() );
            return ac.balance;
         }

         using create_action = uosio::action_wrapper<"create"_n, &token::create>;
         using issue_action = uosio::action_wrapper<"issue"_n, &token::issue>;
         using retire_action = uosio::action_wrapper<"retire"_n, &token::retire>;
         using transfer_action = uosio::action_wrapper<"transfer"_n, &token::transfer>;
         using open_action = uosio::action_wrapper<"open"_n, &token::open>;
         using close_action = uosio::action_wrapper<"close"_n, &token::close>;
      private:
         struct [[uosio::table]] account {
            asset    balance;

            uint64_t primary_key()const { return balance.symbol.code().raw(); }
         };

         struct [[uosio::table]] currency_stats {
            asset    supply;
            asset    max_supply;
            name     issuer;

            uint64_t primary_key()const { return supply.symbol.code().raw(); }
         };

         typedef uosio::multi_index< "accounts"_n, account > accounts;
         typedef uosio::multi_index< "stat"_n, currency_stats > stats;

         void sub_balance( name owner, asset value );
         void add_balance( name owner, asset value, name ram_payer );
   };

} /// namespace uosio
