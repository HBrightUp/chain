#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/time.hpp>

namespace conceptionsys {

    using namespace uosio;

    using uosio::time_point;
    using uosio::asset;


    //sort of vote
    const uint8_t opinion_agree  = 1;       // agree
    const uint8_t opinion_oppose  = 2;      // oppose
    const uint8_t opinion_waive   =  3;     // give up

    CONTRACT conception : public uosio::contract {

    private:

        //global variable(key-value)
        TABLE global_var {
            name key;           //key
            uint64_t val;       //value

            uint64_t        primary_key() const { return key.value; }

            UOSLIB_SERIALIZE(global_var, (key)(val));
        };

        //conception information table
        TABLE conception_info {
            name                            proposal;          //  proposal of conception;
            name                            proposer;       // person  offerring a proposal
            uint64_t                        create_time;           // create time of propose
            std::string                     description;         //  description about the conception;

            uint64_t primary_key()const { return proposal.value; }

            UOSLIB_SERIALIZE( conception_info, (proposal)(proposer)(create_time) (description))
            
        };

        //voters information
        TABLE conception_voters {
            uint64_t    key;                // primary key
            name        voter;             // name of voter
            uint8_t     opinion;            // 1:agree   2:oppose 3:waive
            uint64_t    amount;             // amount of vote

            uint64_t primary_key()const { return key; }
            uint64_t by_voter()const { return voter.value; }

            UOSLIB_SERIALIZE( conception_voters, (key)(voter)(opinion)(amount) )
        };

         //statistics votes table
        TABLE conception_statistics {
             name                            proposal;     //  proposal of conception;
             uint64_t                        agree;     //  total votes of agree 
             uint64_t                        oppose;    //  total votes of oppose
             uint64_t                        waive;     //  total votes of waive

             uint64_t primary_key()const { return proposal.value; }

            UOSLIB_SERIALIZE( conception_statistics, (proposal)(agree)(oppose)(waive) )
        };

        //apply table
        TABLE apply_info {
            name            proposer;       // person  offerring a proposal

            uint64_t primary_key()const { return proposer.value; }    

            UOSLIB_SERIALIZE( apply_info, (proposer))
        };

        /*************************************************************
         * 1.this struct come from uosio contract
         * **********************************************************/
        struct vote_producer {
            name  producer;
            asset voted;

        };

        struct voter_info {
            name                owner;      // the voter
            name                proxy;      // the proxy set by the voter, if any
            int64_t                     staked = 0;
            std::vector<vote_producer> vote_producers;      // system vote list
            double                      last_vote_weight = 0;   // the vote weight cast the last time the vote was updated
            double                      proxied_vote_weight= 0;     // the total vote weight delegated to this voter as a proxy
            bool                        is_proxy = 0;       // whether the voter is a proxy for others
            uint32_t                    reserved1 = 0;
            uint32_t                    reserved2 = 0;
            uosio::asset                reserved3;

            uint64_t primary_key()const { return owner.value; }

      UOSLIB_SERIALIZE( voter_info, (owner)(proxy)(vote_producers)(staked)(last_vote_weight)(proxied_vote_weight)(is_proxy)(reserved1)(reserved2)(reserved3) )
   };

        typedef uosio::multi_index< "global"_n, global_var >                global_var_table;
        typedef uosio::multi_index< "concept"_n, conception_info >          conception_info_table;
        typedef uosio::multi_index< "voters"_n, conception_voters,
                    indexed_by<"voter"_n, const_mem_fun<conception_voters, uint64_t, &conception_voters::by_voter> >
                    > conception_voters_table;

        typedef uosio::multi_index< "statistics"_n, conception_statistics > conception_statistics_table;
        typedef uosio::multi_index< "apply"_n, apply_info >                 apply_info_table;

        //quote from uosio.system contract
        //typedef uosio::multi_index< "userres"_n, user_resources>            user_resources_table;
        typedef uosio::multi_index< "voters"_n, voter_info>  voters_table;

        private:
            global_var_table                _global_var;
            conception_info_table           _tb_concept;
            conception_voters_table         _tb_voters;
            conception_statistics_table     _tb_statistics;
            apply_info_table                _tb_apply;
            voters_table                    _tb_system_voters;
        
        private:
            //op   1: add or modify data;   2: delete data;  others: error
            void set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

        public:
            //listen to transfer action of uosio.token contract
            void transfer(const name& from, const name& to, const asset& quantity, const std::string& memo);

        public:
            using contract::contract;
            conception(name receiver, name code,  uosio::datastream<const char*> ds);

            //init contract parameters
            ACTION init();

            //modify contract parameters
            ACTION setdata(const name& key, const uint64_t& val, const uint8_t& op);

            //create proposal
            ACTION create( const name& proposer, const name& proposal, const std::string& description);

            //vote to proposal
            ACTION vote(const name& voter, const name& proposal, const uint8_t& opinion, const asset& quantity );

            //stop vote 
            ACTION stop();

            //remove invalid proposal
            ACTION rmproposal(const name& proposal); 

            ACTION rmvotedata(const uint64_t& items); 

            ACTION reset();  
    };

}
