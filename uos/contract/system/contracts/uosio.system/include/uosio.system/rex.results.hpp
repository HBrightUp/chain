#pragma once

#include <uosiolib/uosio.hpp>
#include <uosiolib/name.hpp>
#include <uosiolib/asset.hpp>

using uosio::name;
using uosio::asset;
using uosio::action_wrapper;

class [[uosio::contract("rex.results")]] rex_results : uosio::contract {
   public:

      using uosio::contract::contract;

      [[uosio::action]]
      void buyresult( const asset& rex_received );

      [[uosio::action]]
      void sellresult( const asset& proceeds );

      [[uosio::action]]
      void orderresult( const name& owner, const asset& proceeds );

      [[uosio::action]]
      void rentresult( const asset& rented_tokens );

      using buyresult_action   = action_wrapper<"buyresult"_n,   &rex_results::buyresult>;
      using sellresult_action  = action_wrapper<"sellresult"_n,  &rex_results::sellresult>;
      using orderresult_action = action_wrapper<"orderresult"_n, &rex_results::orderresult>;
      using rentresult_action  = action_wrapper<"rentresult"_n,  &rex_results::rentresult>;
};
