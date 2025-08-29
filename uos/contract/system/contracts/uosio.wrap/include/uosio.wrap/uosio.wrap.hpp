#pragma once

#include <uosiolib/uosio.hpp>
#include <uosiolib/ignore.hpp>
#include <uosiolib/transaction.hpp>

namespace uosio {

   class [[uosio::contract("uosio.wrap")]] wrap : public contract {
      public:
         using contract::contract;

         [[uosio::action]]
         void exec( ignore<name> executer, ignore<transaction> trx );

         using exec_action = uosio::action_wrapper<"exec"_n, &wrap::exec>;
   };

} /// namespace uosio
