#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>

namespace stoksys {

    using namespace uosio;
    using uosio::time_point;
    using uosio::asset;

    #define STOK_SYMBOL symbol("STOK", 4)

    CONTRACT stok : public uosio::contract {

    private:

        //global variable(key-value)
        TABLE global_var {
            name key;           //key
            uint64_t val;       //value

            uint64_t        primary_key() const { return key.value; }

            UOSLIB_SERIALIZE(global_var, (key)(val));
        };
        typedef uosio::multi_index< "global"_n, global_var >                global_var_table;




        private:
            global_var_table                _global_var;
           
           
        
        private:
            //op   1: add or modify data;   2: delete data;  others: error
            void set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

            //calc y
            double fx(uint64_t x);

            //calc area
            double get_area(uint64_t x);

        public:
            //listen to transfer action of uosio.token contract
            void transfer(const name& from, const name& to, const asset& quantity, const std::string& memo);

        public:
            using contract::contract;
            stok(name receiver, name code,  uosio::datastream<const char*> ds);

            //init contract parameters
            ACTION init();

            //modify contract parameters
            ACTION setdata(const name& key, const uint64_t& val, const uint8_t& op);

            ACTION issue(const uint64_t block_number);

            ACTION clear();  
    };

}   // end stoksys
