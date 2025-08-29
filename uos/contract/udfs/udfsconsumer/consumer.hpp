#include <ctime>
#include <uosiolib/uosio.hpp>
#include <uosiolib/time.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/contract.hpp>
#include <uosiolib/transaction.hpp>
#include <uosiolib/crypto.h>
#include <uosiolib/fixed_bytes.hpp>

#include "config.hpp"

namespace udfs_consumer {

using namespace uosio;

std::string& trim(std::string &s);
size_t from_hex( const std::string& hex_str, char* out_data, size_t out_data_len ); 
checksum256 hex_to_sha256(const std::string& hex_str); 
 

CONTRACT consumer : public uosio::contract {
public:
	using contract::contract;
	consumer(name receiver, name code,  uosio::datastream<const char*> ds);
	
	//contract must be inited at first
	ACTION init();
		
    //notification of transfer
    void  transfer(const name& from, const name& to, const asset& quantity, const std::string& memo);

	//设置参数数据
    ACTION  setdata(const name& key, const uint64_t& val, const uint8_t& op);

	ACTION  setstorage(const name &user, const int64_t &storage);

	ACTION  setused(const name &user, const int64_t &storage);

    //给udfs基金会 
    //ACTION  giveudfsfund( const asset& money ) ;
	
	ACTION  clear(const uint64_t &num);
		
    void apply(uint64_t receiver, uint64_t code, uint64_t action);

private:
  //global variable
  TABLE udfs_global {
      name     key;
      uint64_t val;
      uint64_t primary_key() const { return key.value; }
  
     UOSLIB_SERIALIZE(udfs_global, (key)(val));
  };

  //infund of the period
  /*
  TABLE udfs_fund{
	  uint64_t phase;		//期数
	  asset	 infund;	//收入
	  asset	 outfund;
	  uint64_t primary_key() const { return phase; }

	  UOSLIB_SERIALIZE(udfs_fund, (phase)(infund)(outfund));
  };
  */

  //udfs user
  TABLE  udfs_user { 
      name        user;          //  消费者
      uint64_t    create_time;       //  注册时间
      uint64_t	  exp_time;			 //  到期时间
      uint8_t	  pack_type;		 //  套餐包类型
      uint64_t    storage;           //  总的存储大小（购买+赠送） 
      uint64_t    used;              //  已使用的存储大小，有增减
      uint64_t	  gived;			 //  赠送的存储
	  uint64_t    spentmoney;        //  累加,充值金额

      uint64_t	  primary_key() const { return user.value; }
      uint64_t    by_create_time()const { return create_time; }
	  uint64_t    by_exp_time()const { return exp_time; }
	  uint64_t	  by_spent() const { return spentmoney; }
	  uint64_t	  by_pack_type() const {return (uint64_t)pack_type;}
	  UOSLIB_SERIALIZE(udfs_user, (user)(create_time)(exp_time)(pack_type)(storage)(used)(gived)(spentmoney));
  };
/*
  TABLE udfs_pack{
	  uint64_t id;			//套餐id
	  uint64_t storage;		//存储空间
	  uint64_t duration;	//时长
	  asset    quantity;	//金额
	  
	  //asset    month;		//30天
	  //asset    season;		//90天
	  //asset    half_year;	//182天
	  //asset    year;		//365天
	  
	  uint64_t primary_key() const { return id; }
	  uint64_t get_storage() const { return storage; }
	  UOSLIB_SERIALIZE(udfs_pack, (id)(storage)(duration)(quantity) );
  };
*/

   typedef uosio::multi_index< "udfsglobal"_n, udfs_global > udfs_global_table;

   //typedef uosio::multi_index< "udfsfund"_n, udfs_fund > udfs_fund_table;

   //typedef uosio::multi_index< "udfspack"_n, udfs_pack > udfs_pack_table;
   
   typedef uosio::multi_index< "udfsuser"_n, udfs_user,
                 indexed_by<"bycreatetime"_n, const_mem_fun<udfs_user, uint64_t, &udfs_user::by_create_time> >,
                 indexed_by<"byexpiration"_n, const_mem_fun<udfs_user, uint64_t, &udfs_user::by_exp_time> >,
                 indexed_by<"byspent"_n, const_mem_fun<udfs_user, uint64_t, &udfs_user::by_spent> >,
                 indexed_by<"bypacktype"_n, const_mem_fun<udfs_user, uint64_t, &udfs_user::by_pack_type> >
                > udfs_user_table;
 
   struct st_transfer {
         name        from;
         name        to;
         asset       quantity;
         std::string memo;
   };

   struct st_newaccount {
            name        from;		//创建者
            name        to;			//新账户
   };
  
   udfs_global_table 		_udfs_global;
   //udfs_fund_table		  	_udfs_fund;
   udfs_user_table			_udfs_user_list; 
   //udfs_pack_table			  _udfs_pack;

   //op   1: add or modify data;   2: delete data;	others: error
   void  set_global_var(const name& key, const uint64_t& val, const uint8_t op = 1);

   uint64_t  find_global_val(name     fieldname);	

   void update_storage();
};

}

