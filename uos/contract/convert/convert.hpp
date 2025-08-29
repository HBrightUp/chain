#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>
#include <uosiolib/symbol.hpp>



namespace convertsys {

    using namespace uosio;
    using uosio::time_point;
    using uosio::asset;

    #define CORE_SYMBOL symbol("UOS", 4)
    const uint8_t direction_uos_to_ut = 1;
    const uint8_t direction_ut_to_uos = 2;

    const int64_t switch_range = 1000000;

    const int64_t FX_UOS_UNIT = 100000000;
    const int64_t UOS_ASSET_UNIT = 10000;

    CONTRACT convert : public uosio::contract {

    private:

        //global variable(key-value)
        TABLE global_var {
            name key;           //key
            uint64_t val;       //value

            uint64_t        primary_key() const { return key.value; }

            UOSLIB_SERIALIZE(global_var, (key)(val));
        };

        typedef uosio::multi_index< "global"_n, global_var >                global_var_table;

         ///@abi table memberreward i64
        struct member_reward{
            name                     owner;
            asset                    reward = asset(0,CORE_SYMBOL);
            uint64_t                 active = 0;

            uint64_t primary_key()const { return owner.value; }
            uint64_t by_active()const {return active;}

            UOSLIB_SERIALIZE( member_reward, (owner)(reward)(active))
        };

        typedef uosio::multi_index< "memberreward"_n , member_reward ,
                indexed_by< "byactive"_n, const_mem_fun<member_reward, uint64_t, &member_reward::by_active> >
                                > memberreward;

        ///@abi table utuosstate i64
        struct last_state {
            name    owner;
            std::string     laster_id;
            name    laster_owner;
            int64_t         amount;
            uint64_t        tr_hash;
            std::vector<name> members;
            uint64_t primary_key()const { return owner.value; }
            UOSLIB_SERIALIZE( last_state, (owner)(laster_id)(laster_owner)(amount)(tr_hash)(members))
        };
        typedef uosio::multi_index<"utstate"_n , last_state> utstate;

        ///@abi table utvoter i64
        struct current_voter{
            name  voter;
            uint64_t      tr_hash;
            uint64_t          active_time;
            std::string   tr_id;
            uint64_t primary_key()const { return voter.value; }
            UOSLIB_SERIALIZE( current_voter, (voter)(tr_hash)(active_time)(tr_id))
        };
        typedef uosio::multi_index< "utvoter"_n , current_voter> utvoter;

        ///@abi table uttr i64
        struct current_transaction {
            uint64_t        tr_hash;
            std::string     tr_id;
            name            owner;
            int64_t         amount;
            int32_t         votes;

            uint64_t primary_key()const { return tr_hash; }
            UOSLIB_SERIALIZE( current_transaction, (tr_hash)(tr_id)(owner)(amount)(votes))
        };
        typedef uosio::multi_index<"uttr"_n , current_transaction> uttr;

        struct uos_to_ut_tr{
            name             owner;
            uint64_t                 id = 0;
            int64_t                  tpn;
            int64_t                  amount;
            std::string              ut_address;
            std::vector<std::string> uosutid;
            std::vector<std::string> trac;

            uint64_t primary_key()const { return owner.value; }
            uint64_t indexbyid()const{ return id;}

            UOSLIB_SERIALIZE( uos_to_ut_tr, (owner)(id)(tpn)(amount)(ut_address)(uosutid)(trac))
        };
        typedef uosio::multi_index< "uosuttr"_n , uos_to_ut_tr ,
                            indexed_by< "indexbyid"_n, const_mem_fun<uos_to_ut_tr, uint64_t, &uos_to_ut_tr::indexbyid> >
                        > uosuttr;

        ///@abi table uosutowner i64
        struct uos_to_ut_owner{
            name             owner;
            int64_t                  old_amount = 0;
            int64_t                  new_amount = 0;

            uint64_t primary_key()const { return owner.value; }
            UOSLIB_SERIALIZE( uos_to_ut_owner, (owner)(old_amount)(new_amount))
        };
        typedef uosio::multi_index< "uosutowner"_n , uos_to_ut_owner > uosutowner;

        ///@abi table uosutstate i64
        struct uos_to_ut_state{
            name                    owner;
            uint64_t                 max_id;
            uint64_t primary_key()const { return owner.value; }
            UOSLIB_SERIALIZE( uos_to_ut_state, (owner)(max_id))
        };
        typedef uosio::multi_index<"uosutstate"_n , uos_to_ut_state> uosutstate;

        ///@abi table uosutstate i64
        struct current_rate{
            uint64_t                    key;
            std::string                 uttouos;
            std::string                 uostout;
            uint64_t                    converted;
            std::string                 reserve1;

            uint64_t primary_key()const { return key; }
            UOSLIB_SERIALIZE( current_rate, (key)(uttouos)(uostout)(converted)(reserve1))
        };
        typedef uosio::multi_index<"rate"_n , current_rate> rateinfo;

        ///@abi table stat i64
         struct currency_stats {
            asset          supply;
            asset          max_supply;
            name           issuer;

            uint64_t primary_key()const { return supply.symbol.code().raw(); }
        };
        typedef uosio::multi_index<"stat"_n, currency_stats> stats;

        static asset get_supply( name token_contract_account, symbol_code sym_code )
        {
            stats statstable( token_contract_account, sym_code.raw() );
            const auto& st = statstable.get( sym_code.raw() );
            return st.supply;
        }


        private:
            global_var_table                _global_var;
            memberreward                    _memberreward;

            //ut to uos  table
            utstate     _utuosstate;
            utvoter     _utuosvoter;
            uttr        _utuostr;
            
            //uos to ut table
            uosuttr      _uosuttr;
            uosutowner   _uosutowner;
            uosutstate   _uosutstate;

            rateinfo     _rate;
            
        private:
            //op   1: add or modify data;   2: delete data;  others: error
            void set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

            uint64_t caculate_hash_64(const char *buf, uint32_t size);

            inline asset get_supply( symbol_code sym )const;

            asset caculate_reward(uosio::asset quantity, std::string memo);

            double calc_uos(const double area);
            double get_middile_value(double x1,  const double area,  const double x1_area);

            double calc_ut(const double distance);

            double get_area(const double x);

            double get_fx(const double x);

        public:
            //listen to transfer action of uosio.token contract
            void transfer(const name& from, const name& to, const asset& quantity, const std::string& memo);

        public:
            using contract::contract;
            convert(name receiver, name code,  uosio::datastream<const char*> ds);

             void init();

            //modify contract parameters
            void setdata(const name& key, const uint64_t& val, const uint8_t& op);

            void modifymember(std::vector<name > &members);

            void uttouosvote(const name voter, const std::string tr_id, const name owner,  const int64_t amount, const std::string memo);

            void setousutid(name voter, name tr, std::string id);

            void uttouospass(name voter,std::string tr_id , name owner , int64_t amount , std::string memo);
            void uostoutpass(name user ,std::string ut_id ,std::string ut_address, int64_t amount ,uint64_t buffer_id);

            void delaytf(name from, asset quantity, int64_t tpn , std::string memo );

            void delaytf1(name from, uosio::asset quantity, std::string memo);

            void setuttr(name voter, name tr, std::string uttr);

            //direction(1: ut->uos, 2: uos->ut) is_self(true: used by _self, false: used by others)
            void updaterate(const name voter, const name tr);

            void clear();  
            void modifyuos(const uint64_t converted); //just for test
    };


}
