#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>

namespace partnersys {

    using namespace uosio;

    using uosio::time_point;
    using uosio::asset;



    CONTRACT partner : public uosio::contract {

    private:

        //global variable(key-value)
        TABLE global_var {
            name key;           //key
            uint64_t val;       //value

            uint64_t        primary_key() const { return key.value; }

            UOSLIB_SERIALIZE(global_var, (key)(val));
        };
        typedef uosio::multi_index< "global"_n, global_var >                global_var_table;



        TABLE invited_persion {
            name invited;

            uint64_t        primary_key() const { return invited.value; }

            UOSLIB_SERIALIZE(invited_persion, (invited));
        };
        typedef uosio::multi_index< "invitation"_n, invited_persion >                invited_persion_table;

        struct transfer_list {
            uint64_t  amount;
            uint8_t   section;
            uint64_t  time;
        };

        TABLE partner_info {
            name                    owner;
            name                    prev;
            uint64_t               invited;
            uint64_t               join_time;
            uint64_t               update_time;
            std::vector<transfer_list>      asset_list;

            uint64_t        primary_key() const { return owner.value; }
            uint64_t by_join_time()const { return join_time; }
            uint64_t by_prev()const { return prev.value; }
            uint64_t by_update_time()const { return update_time; }

            UOSLIB_SERIALIZE(partner_info, (owner)(prev)(invited)(join_time)(update_time)(asset_list));
        };
        typedef uosio::multi_index< "partners"_n, partner_info,
                indexed_by<"jointime"_n, const_mem_fun<partner_info, uint64_t, &partner_info::by_join_time> >,
                indexed_by<"prev"_n, const_mem_fun<partner_info, uint64_t, &partner_info::by_prev> >,
                indexed_by<"updatetime"_n, const_mem_fun<partner_info, uint64_t, &partner_info::by_update_time> >
                >   partner_info_table;
        

        private:
            global_var_table                _global_var;
            partner_info_table              _partners;

           
        
        private:
            //op   1: add or modify data;   2: delete data;  others: error
            void set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

        public:
            //listen to transfer action of uosio.token contract
            void transfer(const name& from, const name& to, const asset& quantity, const std::string& memo);

        public:
            using contract::contract;
            partner(name receiver, name code,  uosio::datastream<const char*> ds);

            //init contract parameters
            ACTION init();

            //modify contract parameters
            ACTION setdata(const name& key, const uint64_t& val, const uint8_t& op);

            ACTION invite(const name owner, const std::vector<name> invited_list);



            ACTION clear();  
    };

}   // end partnersys
