//#include "config.hpp"
#include <cstring>
#include <string>
#include <uosiolib/uosio.hpp>
#include <uosiolib/asset.hpp>
#include <uosiolib/print.hpp>


using namespace uosio;

static constexpr uint64_t MAX_KEY_VALUE = std::numeric_limits<uint64_t>::max();
#define CORE_SYMBOL 	symbol("UOS", 4)
#define UDFS_ACCOUNT	"udfsconsumer"_n

CONTRACT udisk : public contract {
	public:
		using contract::contract;
		udisk(name receiver, name code, datastream<const char*> ds)
			:contract(receiver, code, ds),
			_udisk_global(_self, _self.value),
			_udisk_user(_self, _self.value)
			//_udisk_pack(_self, _self.value)
			{}

		ACTION init(){
			require_auth(_self);
			auto itr = _udisk_global.begin();
			uosio_assert(itr == _udisk_global.end(), "the contract already initialized");
			//setKey("price"_n, 10000);
			setKey("usernum"_n, 0);
			setKey("infund"_n, 0);
			setKey("outfund"_n, 0);
			setKey("totalstorage"_n, 0);
	  		setKey("usedstorage"_n, 0);
			setKey("allotstorage"_n, 0);

			update_storage();
		}

		ACTION setstorage(const name &user, const int64_t &storage/*, const uint64_t expiration*/){		//
			require_auth(_self);
			uosio_assert(user != _self, "this contract account can not be allocated space");
			uosio_assert(is_account(user), "known account");
			uosio_assert(storage != 0, "storage can not be zero");
		
			update_storage();
			
			auto useritr = _udisk_user.find(user.value);
			if (storage > 0){
				if (useritr == _udisk_user.end()){
					_udisk_user.emplace(_self, [&](auto &a){
						a.user = user;
						a.create_time = now();
						a.storage = storage;
						a.gived = storage;
					});

					addValue("usernum"_n, 1);
				} else {
					_udisk_user.modify(useritr, name(0), [&](auto &a){
						a.storage += storage;
						a.gived += storage;
					});
				}
				
				auto allotitr = _udisk_global.find("allotstorage"_n.value);
				auto total = _udisk_global.get("totalstorage"_n.value);
				uosio_assert(allotitr != _udisk_global.end(), "can not find allotstorage");
				
				_udisk_global.modify(allotitr, name(0), [&](auto &g){
					uosio_assert(total.val - g.val >= storage, "there are not enough space for give");
					g.val += storage;
				});
/*
				udisk_give_table _udisk_gived(_self, user.value);		
				if (expiration != 0)			//expiration为入参
					uosio_assert(expiration > now(), "expiration time can not less than now");
				auto giveitr = _udisk_gived.find(expiration);
				if (giveitr == _udisk_gived.end()){
					_udisk_gived.emplace(_self, [&](auto &a){
						a.exp_time = expiration;
						a.storage = storage;
					})
				} else {
					_udisk_gived.emplace(giveitr, name(0), [&](auto &a){
						a.storage += storage;
					})
				}
*/
			}else{
				uosio_assert(useritr != _udisk_user.end(), "can not find user");
				uint64_t newstorage = abs(storage);
				if (useritr->gived < newstorage)
					newstorage = useritr->gived;
				
				_udisk_user.modify(useritr, name(0), [&](auto &a){
					a.gived -= newstorage;
					a.storage -= newstorage;
				});

				addValue("allotstorage"_n, newstorage, true);
				
/*
				udisk_give_table _udisk_gived(_self, user.value);
				uint64_t expiration = now();
				auto giveitr = _udisk_gived.find(expiration);
				uosio_assert(giveitr == _udisk_gived.end(), "can not find record");
				uosio_assert(storage == giveitr->storage, "storage error")
				_udisk_gived.erase(giveitr);
*/			
			}
		}

		ACTION setused(const name &user, const int64_t &storage){			//OK
			require_auth(_self);
			//uosio_assert(user != _self, "this contract account can not be allocated space");
			uosio_assert(storage != 0, "storage can not be zero");
			auto useritr = _udisk_user.find(user.value);
			uosio_assert(useritr != _udisk_user.end(), "can not find the user");
			//auto useditr = _udisk_global.find("usedstorage"_n.value);
			//uosio_assert(useditr != _udisk_global.end(), "can not find usedstorage");
			if (storage > 0){
				_udisk_user.modify(useritr, name(0), [&](auto &a){
					uosio_assert(a.storage >= (a.used + storage), "too less storage space to store the file");
					a.used += storage;
				});

				addValue("usedstorage"_n, storage);
			}else{
				uint64_t newstorage = abs(storage);
				_udisk_user.modify(useritr, name(0), [&](auto &a){
					uosio_assert(a.used >= newstorage, "too much storage space to release");
					a.used -= newstorage;
				});

				addValue("usedstorage"_n, newstorage, true);
			}	
			
		}
		
		ACTION setdata(name key, uint64_t value, bool del){
			require_auth(_self);
			setKey(key, value, del);
		}

		ACTION clear(uint64_t num){
			require_auth(_self);
			auto itr = _udisk_global.begin();
			uint64_t i = 0;
			for(; itr != _udisk_global.end() && i < num; ++i){
				itr = _udisk_global.erase(itr);
			}

			auto itr1 = _udisk_user.begin();
			for(; itr1 != _udisk_user.end() && i < num; ++i){
				itr1 = _udisk_user.erase(itr1);
			}
		}

		void transfer(name from, name to, asset quantity, std::string memo){		//有待需求明确
			if(from == _self){
				return;
			}
			uosio_assert(  to == _self , "must transfer to this contract");
			uosio_assert(  quantity.symbol == CORE_SYMBOL, "must use system coin");

			trim(memo);

			if (memo == "buy"){
				auto price = _udisk_global.get("price"_n.value);
				uint64_t storage = getStorage(quantity, price.val);
				auto useritr = _udisk_user.find(from.value);
				if (useritr == _udisk_user.end()){
					_udisk_user.emplace(_self, [&](auto &a){
						a.user = from;
						a.storage = storage;
					});
				}else{
					_udisk_user.modify(useritr, name(0), [&](auto &a){
						a.storage += storage;
					});
				}
			}
			
		}

		void apply(uint64_t receiver, uint64_t code, uint64_t action){
			auto &thiscontract = *this;
			if (action == ( "transfer"_n ).value && code == ( "uosio.token"_n ).value ) {
				auto transfer_data = unpack_action_data<st_transfer>();
				//transfer(transfer_data.from, transfer_data.to, transfer_data.quantity, transfer_data.memo);
				return;
			}

			if (code != get_self().value) return;
			switch (action) {
				UOSIO_DISPATCH_HELPER( udisk, (init)(setstorage)(setused)(clear) );
			}; 
		}

	private:
		TABLE udisk_global {
			name     key;
			uint64_t val;

			uint64_t primary_key() const { return key.value; }
			UOSLIB_SERIALIZE(udisk_global, (key)(val) )
		};

		TABLE udisk_user { 
      		name      user;          	 //  消费者
      		uint64_t  create_time;       //  注册时间
      		uint64_t  exp_time;			 //  到期时间
      		uint8_t	  pack_type;		 //  套餐包类型
      		uint64_t  storage;           //  总的存储大小（购买+赠送） 
      		uint64_t  used;              //  已使用的存储大小，有增减
      		uint64_t  gived;			 //  赠送的存储
	  		uint64_t  spentmoney;        //  累加,充值金额

      		uint64_t  primary_key() const { return user.value; }
      		uint64_t  by_create_time()const { return create_time; }
	  		uint64_t  by_exp_time()const { return exp_time; }
	  		uint64_t  by_spent() const { return spentmoney; }
	  		uint64_t  by_pack_type() const {return (uint64_t)pack_type;}
	  		UOSLIB_SERIALIZE(udisk_user, (user)(create_time)(exp_time)(pack_type)(storage)(used)(gived)(spentmoney))
  		};
/*
		TABLE udisk_gived {
			uint64_t exp_time;
			uint64_t storage;
			
			uint64_t primary_key() const { return exp_time; }
			UOSLIB_SERIALIZE(udisk_gived, (exp_time)(storage))
		};

		TABLE udisk_pack{
			uint64_t id;
			uint64_t storage;
			uint64_t duration;
			asset	 quantity;

			uint64_t primary_key() const { return id; }
			uint64_t get_storage() const { return storage; }
			UOSLIB_SERIALIZE(udisk_pack, (id)(storage)(duration)(quantity));
		};
*/
		struct st_transfer{
			name		from;
			name		to;
			asset		quantity;
			std::string	memo;
		};

		typedef multi_index<"udiskglobal"_n, udisk_global> 		udisk_global_table;
		typedef multi_index<"udiskuser"_n, udisk_user>			udisk_user_table;
		typedef multi_index<"udfsuser"_n, udisk_user>			udfs_user_table;
		//typedef multi_index<"udiskgived"_n, udisk_gived>		udisk_give_table;
		//typedef multi_index<"udiskpack"_n, udisk_pack>			udisk_pack_table;

		udisk_global_table _udisk_global;
		udisk_user_table  _udisk_user;
		//udisk_pack_table  _udisk_pack;

		inline std::string& trim(std::string &s) {
        	if (s.empty()) {
               return s;
        	}
			
        	s.erase(0,s.find_first_not_of(" "));
        	s.erase(s.find_last_not_of(" ") + 1);
        	return s;
		}

		inline uint64_t getStorage(asset quantity, uint64_t price){
			return quantity.amount / price;
		}

		void setKey(name key, uint64_t value, bool del = false){
			auto itr = _udisk_global.find(key.value);

			if (del) {
				uosio_assert(itr != _udisk_global.end(), "the key doesn't exist");
				_udisk_global.erase(itr);
				return;
			}

			if (itr == _udisk_global.end()) {
				_udisk_global.emplace(_self, [&](auto& a){
					a.key = key;
					a.val = value;
				});
			} else {
				_udisk_global.modify(itr, name(0), [&](auto& a){
					a.val = value;
				});
			}
		}

		void addValue(name key, uint64_t plus, bool negative = false){ 
			auto itr = _udisk_global.find(key.value);
			uosio_assert(itr != _udisk_global.end(), "can not find the key");
			
			_udisk_global.modify(itr, name(0), [&](auto& one){
					if (negative == true) {
						uosio_assert(one.val >= plus, "modify, numerical downward overflow");
						one.val -= plus;
					} else {
						uosio_assert((MAX_KEY_VALUE - plus) >= one.val, "modify, numerical upward overflow");
						one.val += plus;
					}
			});
		}

		void update_storage(){
	  		udfs_user_table _udfs_user(UDFS_ACCOUNT, UDFS_ACCOUNT.value); 
			auto itr = _udfs_user.find(_self.value);
	  		auto totalitr = _udisk_global.find("totalstorage"_n.value);
	  		if (itr != _udfs_user.end()){
	 	  		_udisk_global.modify(totalitr, name(0), [&](auto &a){
	 		  		a.val = itr->storage;
		  		});
	  		}
  		}
		
};

extern "C" {
	[[noreturn]] void apply(uint64_t receiver, uint64_t code, uint64_t action) {
		udisk disk( name(receiver), name(code), datastream<const char*>(nullptr, 0) );
		disk.apply(receiver, code, action);
		uosio_exit(0);
	}
}

