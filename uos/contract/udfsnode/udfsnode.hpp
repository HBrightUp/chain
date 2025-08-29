/**
 *  @file
 *  @copyright defined in uos/LICENSE.txt
 */
#pragma once

#include <uosiolib/asset.hpp>
#include <uosiolib/uosio.hpp>
#include <uosiolib/symbol.hpp>

#include <string>
#include <vector>

namespace uosio {   

static constexpr uint32_t     seconds_per_day = 24 * 3600;
    
class [[uosio::contract("udfs.uos")]] uosio_udfs : public contract {

private:

    ///@abi table global_var i64
    struct [[uosio::table]] global_var {
        name        key;           //key
        uint64_t val;       //value

        uint64_t        primary_key() const { return key.value; }
        UOSLIB_SERIALIZE(global_var, (key)(val))
    };

    typedef uosio::multi_index< "global"_n, global_var >   globalvar;

    ///@abi table udfs_node i64
    struct [[uosio::table]] udfs_node {
        uint64_t    tr_hash;        // primary key
        name        owner;       // account of apply for udfs
        uint64_t    amount;       //  assert of spend for udfs
        uint64_t    apply_time;    // time of apply for 
        uint64_t    coin_type;      //source coin   type:  1-->UT; 2-->ETH
        uint64_t    reserve1;      //reserve
        uint64_t    reserve2;      //reserve

        uint64_t        primary_key() const { return tr_hash; }
        uint64_t        by_owner() const { return owner.value; }
        uint64_t        by_apply_time() const { return apply_time; }
        UOSLIB_SERIALIZE(udfs_node, (tr_hash)(owner)(amount)(apply_time)(coin_type)(reserve1)(reserve2))
    };

    typedef uosio::multi_index< "udfsnode"_n, udfs_node,
                        indexed_by< "owner"_n, const_mem_fun<udfs_node, uint64_t, &udfs_node::by_owner> >,
                        indexed_by< "applytime"_n, const_mem_fun<udfs_node, uint64_t, &udfs_node::by_apply_time> >
                    >   udfsnode;

    ///@abi table uttr i64
    struct [[uosio::table]] current_transaction {
        uint64_t        tr_hash;
        std::string     tr_id;
        name            owner;
        int64_t         amount;
        int32_t         votes;
        uint64_t primary_key()const { return tr_hash; }
        UOSLIB_SERIALIZE( current_transaction, (tr_hash)(tr_id)(owner)(amount)(votes))
    };
    typedef uosio::multi_index<"uttr"_n , current_transaction> uttr;
    typedef uosio::multi_index<"ethtr"_n , current_transaction> ethtr;

    ///@abi table utvoter i64
    struct [[uosio::table]] current_voter{
        name            voter;
        uint64_t      tr_hash;
        uint64_t          active_time;
        std::string   tr_id;
        uint64_t primary_key()const { return voter.value; }
        UOSLIB_SERIALIZE( current_voter, (voter)(tr_hash)(active_time)(tr_id))
    };
    typedef uosio::multi_index<"utvoter"_n , current_voter> utvoter;
    typedef uosio::multi_index<"ethvoter"_n , current_voter> ethvoter;

     ///@abi table utuosstate i64
    struct [[uosio::table]] last_state {
        name            owner;
        std::string     laster_id;
        name            laster_owner;
        int64_t         amount;
        uint64_t        tr_hash;
        std::vector<name> members;
        uint64_t primary_key()const { return owner.value; }
        UOSLIB_SERIALIZE( last_state, (owner)(laster_id)(laster_owner)(amount)(tr_hash)(members))
    };
    typedef uosio::multi_index<"utstate"_n , last_state> utstate;
    typedef uosio::multi_index<"ethstate"_n , last_state> ethstate;

    ///@abi table memberreward i64
    struct [[uosio::table]] member_reward{
        name                    owner;
        asset                    reward = asset(0,{"UOS", 4});
        uint64_t                 active = 0;
        uint64_t primary_key()const { return owner.value; }
        uint64_t by_active()const {return active;}
        UOSLIB_SERIALIZE( member_reward, (owner)(reward)(active))
    };

    typedef uosio::multi_index< "memberreward"_n , member_reward ,
            indexed_by< "byactive"_n, const_mem_fun<member_reward, uint64_t, &member_reward::by_active> >
                              > memberreward;

    struct [[uosio::table]] account {
      asset    balance;

      uint64_t primary_key()const { return balance.symbol.raw(); }
    };
    typedef uosio::multi_index<"accounts"_n, account> accounts;

    ///@abi table utuosstate i64
    struct [[uosio::table]] whitelist_fee {
        name    owner;
        
        uint64_t primary_key()const { return owner.value; }
        UOSLIB_SERIALIZE(whitelist_fee, (owner))
    };
    typedef uosio::multi_index<"whitelistfee"_n , whitelist_fee> whitelistfee;

    utstate     _utstate;
    ethstate    _ethstate;
    utvoter      _utvoter;
    ethvoter    _ethvoter;
    uttr        _uttr;
    ethtr       _ethtr;
    memberreward _memberreward;
    globalvar    _globalvar;
    udfsnode     _udfsnode;
    whitelistfee _fee;

public:
    using contract::contract;
    uosio_udfs(name receiver, name code,  uosio::datastream<const char*> ds);

    [[uosio::action]]
    void modifymember(std::vector<name > &members);

    [[uosio::action]]
    void setdata(const name& key, const uint64_t& val, const uint8_t& op);

    [[uosio::action]]
    void clear(const name& owner);

    [[uosio::action]]
    void voter(const name voter, const std::string tr_id, const name owner,  const int64_t amount, const std::string coin_name);
    
    [[uosio::action]]
    void applypass(name voter,std::string tr_id , name owner , int64_t amount , std::string memo);

    template<typename muti_uosstate, typename muti_touostr, typename muti_uosvoter>
    void deal(muti_uosstate& _tb_uosstate, muti_touostr& _tb_touostr, muti_uosvoter& _tb_uosvoter, const name voter, const std::string tr_id, const name owner,  const int64_t amount, const std::string coin_name);
    
    [[uosio::action]]
    void transfer( name from, name to,asset quantity,std::string memo);

    [[uosio::action]]
    void issuetoken();
    
    [[uosio::action]]
    void addnode(const uint64_t tr_hash, const name owner);

    //op:   1-->add    2-->remove
    void feetableop(const name owner, const uint8_t op);

private:

    //op   1: add or modify data;   2: delete data;  others: error
    void set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

     uint64_t caculate_hash_64(const char *buf , uint32_t size);

    template<typename muti_uosstate, typename muti_touostr, typename muti_uosvoter>
    void storemember(muti_uosstate& _tb_uosstate, muti_touostr& _tb_touostr, muti_uosvoter& _tb_uosvoter, std::vector<name > &members);
};



} /// namespace uosio
