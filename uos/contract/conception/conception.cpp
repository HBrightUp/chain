#include<uosiolib/uosio.hpp>
#include <uosiolib/privileged.h>
#include "conception.hpp"


namespace conceptionsys {
    using namespace uosio;

   conception::conception(name receiver, name code,  uosio::datastream<const char*> ds)
   :contract(receiver, code, ds),
   _global_var(_self, _self.value),
   _tb_concept(_self, _self.value),
   _tb_voters(_self, _self.value),
   _tb_statistics(_self, _self.value),
   _tb_apply(_self, _self.value),
   _tb_system_voters("uosio"_n, ("uosio"_n).value )
   {
   }

   void conception::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {

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

   ACTION conception::init() {

      require_auth(_self);

      //valid time of proposal
      set_global_var("validtime"_n, 3 * 24 * 3600);

      //pay for propose(50.0000 UOS)
      set_global_var("paypropose"_n, 500000);

      //min stake for vote(500.0000 UOS)
      set_global_var("minstake"_n, 5000000);

      //min  number of vote(500.0000 UOS)
      set_global_var("minvotes"_n, 5000000);

      //current vote state of contract
      // 0:idle 1:voting
      set_global_var("votingstate"_n, 0);

   }

   ACTION conception::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
      require_auth(_self);
      set_global_var(key, val, op);
   }

   ACTION conception::create( const name& proposer, const name& proposal, const std::string& description) {
      
      //check auth
      require_auth(proposer);

      //check transfer asset before create proposal
      auto apply_itr = _tb_apply.find(proposer.value);
      if(apply_itr == _tb_apply.end()) {
         uosio_assert(false, "must transfer asset before create proposal");
      }
      else {
         _tb_apply.erase(apply_itr);
      }

      //check length of description
      auto len = description.length();
      uosio_assert(len > 0 && len <= 512 , "length of description information error");

      //check voting state of contract
      auto state_itr = _global_var.find(("votingstate"_n).value);
      uosio_assert(state_itr != _global_var.end(), "not find votingstate field");

      uosio_assert(state_itr->val == 0, "vote is ongoing");
      set_global_var("votingstate"_n, 1);

      //check exist in concept table
      auto concept_itr = _tb_concept.find(proposal.value);
      uosio_assert(concept_itr == _tb_concept.end(), "proposal existed in concept table");

      _tb_concept.emplace( proposer, [&] (auto& c) {
            c.proposal = proposal;
            c.proposer = proposer;
            c.create_time = now();
            c.description = description;
      }); 

      _tb_statistics.emplace( proposer, [&] (auto& s) {
            s.proposal = proposal;
      });
   }

   ACTION conception::vote(const name& voter, const name& proposal, const uint8_t& opinion, const asset& quantity ) {

      require_auth(voter);

      //check valid time
      auto validtime_itr = _global_var.find(("validtime"_n).value);
      uosio_assert(validtime_itr != _global_var.end(), "not find validtime field");
      auto itr_concept = _tb_concept.find(proposal.value);

      auto votingstate_itr = _global_var.find(("votingstate"_n).value);
      uosio_assert(votingstate_itr != _global_var.end(), "not find votingstate field");
      uosio_assert(votingstate_itr->val == 1, "no proposal to vote");

      if(itr_concept->create_time + validtime_itr->val <= now()) {
         if(votingstate_itr->val == 1) {
            set_global_var("votingstate"_n, 0);
            //INLINE_ACTION_SENDER(conceptionsys::conception, stop)( uosio::name(_self), {_self, "active"_n},{ } );
            return ;
         }
         else {
            uosio_assert(false, "proposal overtime" );
         }
      }

      //check opinion
      uosio_assert(opinion >= opinion_agree && opinion <= opinion_waive, "invalid opinion");

      //check quantity
      symbol sym("UOS", 4);
      uosio_assert(quantity.symbol == sym, "invalid code");

      auto minvotes_itr = _global_var.find(("minvotes"_n).value);
      uosio_assert(minvotes_itr != _global_var.end(), "not find minvotes field");
      uosio_assert(quantity.amount >= minvotes_itr->val, "votes not enough");

      //check staked asset

      auto voters_iter = _tb_system_voters.find(voter.value);
      uosio_assert(voters_iter != _tb_system_voters.end(), "voter no stake information");

      //get all tickets information of current account
      auto idx =  _tb_voters.get_index<"voter"_n>();
      auto lower = idx.lower_bound(voter.value);
      auto upper = idx.upper_bound(voter.value);

      uint64_t total_voted = 0;
      for(auto itr = lower; itr != upper; ++itr) {
         if(itr->voter == voter) {
            total_voted += itr->amount;
         }
      }

      print("voter staked: ", voters_iter->staked, " quantity: ", quantity.amount, " total_voted:", total_voted);
      uosio_assert(voters_iter->staked >= total_voted + quantity.amount, " staked asset not enough" );

      _tb_voters.emplace( voter, [&] (auto& v) {
            v.key = _tb_voters.available_primary_key();
            v.voter = voter;
            v.opinion = opinion;
            v.amount = quantity.amount;
      });

      //modify statistics table
      auto statistics_itr = _tb_statistics.find(proposal.value);
      uosio_assert(statistics_itr != _tb_statistics.end(), "proposal not exist in statistics table");

      auto tickets = quantity.amount;
   
      _tb_statistics.modify(statistics_itr, _self, [&](auto& s){
            switch(opinion) {
               case opinion_agree: {
                  s.agree += tickets;
                  break;
               }

               case opinion_oppose: {
                  s.oppose += tickets;
                  break;
               }

               case opinion_waive: {
                  s.waive += tickets;
                  break;
               }

               default: {

               }
            }
      });
   }

   ACTION conception::stop() {
      require_auth(_self);

      auto state_itr = _global_var.find(("votingstate"_n).value);
      uosio_assert(state_itr != _global_var.end(), "not find votingstate field");
      set_global_var("votingstate"_n, 0);
   }

   ACTION conception::rmproposal(const name& proposal) {
      require_auth(_self);

      auto statistics_itr = _tb_statistics.find(proposal.value);
      if(statistics_itr != _tb_statistics.end()) {
         _tb_statistics.erase(statistics_itr);
      }

      auto concept_itr = _tb_concept.find(proposal.value);
      if(concept_itr != _tb_concept.end()) {
         _tb_concept.erase(concept_itr);
      }
   }

   ACTION conception::reset() {
      require_auth(_self);

      auto idx1 = _tb_concept.begin();
      while(idx1 != _tb_concept.end()) {
         idx1 = _tb_concept.erase(idx1);
      }

      auto idx2 = _tb_voters.begin();
      while(idx2 != _tb_voters.end()) {
         idx2 = _tb_voters.erase(idx2);
      }

      auto idx3 = _tb_statistics.begin();
      while(idx3 != _tb_statistics.end()) {
         idx3 = _tb_statistics.erase(idx3);
      }

      auto idx4 = _tb_apply.begin();
      while(idx4 != _tb_apply.end()) {
         idx4 = _tb_apply.erase(idx4);
      }
   }

   ACTION conception::rmvotedata(const uint64_t& items) {

      if(_tb_voters.begin() == _tb_voters.end()) {
         uosio_assert(false, "empty vote data");
      }

      uint64_t counts = 0;

      //erase vote data by given  items
      auto vote_itr = _tb_voters.begin();
      while(vote_itr != _tb_voters.end()) {
         vote_itr = _tb_voters.erase(vote_itr);

         if(items <= ++counts) {
            break;
         }
      }
   }

   void conception::transfer(const name& from, const name& to, const asset& quantity, const std::string& memo) {

      symbol sym("UOS", 4);
      uosio_assert(quantity.symbol == sym, "transfer must system coin");

      if(to == _self) {
         auto paypropose_itr = _global_var.find(("paypropose"_n).value);
      
         if(paypropose_itr != _global_var.end() && memo == "propose") {
            
            uosio_assert(quantity.amount == paypropose_itr->val, "asset pay for propose error");
            auto apply_itr = _tb_apply.find( from.value);
            
            if(apply_itr  == _tb_apply.end()) {
               _tb_apply.emplace( _self, [&](auto& a) {
                  a.proposer = from;
               });
            }
            else {
               uosio_assert(false, "account applied ");
            }    
         }
      }
   }
}

#define UOSIO_DISPATCH_CONCEPTION( TYPE, MEMBERS ) \
extern "C" { \
   void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
      auto self = receiver; \
      if( code == self || code == ("uosio.token"_n).value ) { \
      	 if( code == ("uosio.token"_n).value && action == ("transfer"_n).value ){ \
            uosio::execute_action(uosio::name(receiver), uosio::name(code), &conceptionsys::conception::transfer); \
            return ; \
      	 } \
         switch( action ) { \
            UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
         } \
      } \
   } \
}

UOSIO_DISPATCH_CONCEPTION(conceptionsys::conception, (init)(setdata)(create)(vote)(rmproposal)(reset)(stop)(rmvotedata))






















