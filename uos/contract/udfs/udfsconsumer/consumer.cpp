
#include "consumer.hpp"
#include <cmath>

namespace udfs_consumer {

  using namespace uosio;

	uint8_t from_hex( char c ) {
		if( c >= '0' && c <= '9' )
			return c - '0';
		if( c >= 'a' && c <= 'f' )
			return c - 'a' + 10;
		if( c >= 'A' && c <= 'F' )
			return c - 'A' + 10;
		uosio_assert(false, "Invalid hex character ");
		return 0;
	}
  
	size_t from_hex( const std::string& hex_str, char* out_data, size_t out_data_len ) {
		std::string::const_iterator i = hex_str.begin();
		uint8_t* out_pos = (uint8_t*)out_data;
		uint8_t* out_end = out_pos + out_data_len;
		while( i != hex_str.end() && out_end != out_pos ) {
			*out_pos = from_hex( *i ) << 4;   
			++i;
			if( i != hex_str.end() )  {
				*out_pos |= from_hex( *i );
				++i;
			}
			++out_pos;
		}
		return out_pos - (uint8_t*)out_data;
	}
  
	checksum256 hex_to_sha256(const std::string& hex_str) {
		uosio_assert(hex_str.length() == 64, "invalid sha256");
		checksum256 checksum;
		from_hex(hex_str,(char*) checksum.data(), 32 );
		return checksum;
	}

	std::string& trim(std::string &s) {
  		if (s.empty()) {
        	return s;
    	}
			
    	s.erase(0,s.find_first_not_of(" "));
		s.erase(s.find_last_not_of(" ") + 1);
		return s;
	}
  
	/***
	 *重载localtime()
	 *时间戳转换成日期（北京时间，东八区）
	***/
	int localtime(time_t timestamp, struct tm &ret_tm, int time_zone = 8) {
		static const int days_every_month[12] = { 31,28,31,30,31,30,31,31,30,31,30,31 };	//平年各月份天数
		static const bool four_year[4] = { false, false, true, false }; 	//1970开始，判断是否为闰年，4年为周期
		static const bool four_hundred_year[4] = { true, false, false, false }; //百年不闰，四百年再闰
	
		//int time_zone = 8;		//多8小时
		timestamp += time_zone * 3600;
		//struct tm ret_tm;
	
		ret_tm.tm_isdst = 0;		//夏时令，这个东西没用过，默认是0，不过以防万一，还是重置下
	
		//秒
		ret_tm.tm_sec = (int)(timestamp % 60);
		int tmp = (int)(timestamp / 60);
	
		//分
		ret_tm.tm_min = tmp % 60;
		tmp = tmp / 60;
	
		//时
		ret_tm.tm_hour = tmp % 24;
		tmp = tmp / 24;
	
		/**
		  *1970.1.1到2000.1.1，一共历时10957天
		  *4年1461天，1461*24 = 35064小时
		  *100年36524或36525天
		  *400年146097天
		**/
		int four_year_count = 0;			//四年计数
		int four_hundred_year_count = 0;	//四百年计数
		int left_days = 0;					//剩余天数
		int leave_years = 0;				//过去年数
	
		if (tmp > 10957) {
			four_hundred_year_count = (tmp - 10957) / 146097;		//过去多少个400年，相对于2000.1.1而言
			left_days = (tmp - 10957) % 146097;
			int i = 0;			//此处i表示过了多少个百年，例如2100年1月1日，正好过100年,i == 1
			for (; i < 4; ++i) {
				int hundred_year_days = four_hundred_year[i] ? 36525 : 36524;
				if (left_days < hundred_year_days)		break;
				left_days -= hundred_year_days;
			}
			if (i >= 1 && left_days <= 58)	i--;			//当日期小于等于2月28日时，意味着这个百年并未多算一天，i--
			tmp += (four_hundred_year_count * 3 + i);		//将多算的天数补上
		}
	
		four_year_count = tmp / 1461;		//过去多少个4年
		left_days = tmp % 1461;
		int i = 0;				//i表示以4年为周期计数完之后，还过了几年
		for (; i < 4; ++i) {
			int year_days = four_year[i] ? 366 : 365;
			if (left_days < year_days)		break;
			left_days -= year_days;
		}
		//年
		leave_years = (four_year_count << 2) + 1970 + i;
		ret_tm.tm_year = leave_years;
	
		int j = 0;
		for (; j < 12; ++j) {
			int month_days = days_every_month[j];
			if (i == 2 && j == 1) {
				month_days++;
			}
	
			if (left_days < month_days) {
				break;
			}
			left_days -= month_days;
		}
		//月
		ret_tm.tm_mon = j + 1;
		//日
		ret_tm.tm_mday = left_days + 1;
	 
		return 0;
	}

	consumer::consumer(name receiver, name code, uosio::datastream<const char*> ds)
	:contract(receiver, code, ds),
	_udfs_global(_self, _self.value),
	//_udfs_fund(_self, _self.value),
	_udfs_user_list(_self, _self.value)
	{
	}

   
	ACTION consumer::init() {
		require_auth(_self);
		auto itr = _udfs_global.begin();
		uosio_assert(itr == _udfs_global.end(), "the contract already initalized");
		//set_global_var("udfsprice"_n, 10000);       //  1M 多少钱
		set_global_var("outfund"_n, 0); 
		set_global_var("infund"_n, 0); 
		set_global_var("usernum"_n, 0);
		set_global_var("totalstorage"_n, 0);
		set_global_var("usedstorage"_n, 0);
		set_global_var("allotstorage"_n, 0);
		// recommendfee  

		update_storage();		//更新总存储
	}
 
	//find the name  of value 
	uint64_t  consumer::find_global_val(name  fieldname )
	{
		//check max convert  from credit to uos
		auto fielditer = _udfs_global.find( fieldname.value);
		uosio_assert(fielditer != _udfs_global.end(), "not find fielditer ");
		//uosio_assert(credit <= fielditer->val, "convert credit to uos to many");
		return fielditer->val;
	}	
  
	// op 1 add ; op  2 delete
	void consumer::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {
		uosio_assert(op == 1 || op == 2, "op 1: add or modify; 2: delete");
  		auto fielditer = _udfs_global.find(key.value);
		print( "consumer::set_global_var: ", fielditer->val,  "\n");	
		//delete data
		if(op == 2) {
			uosio_assert(fielditer != _udfs_global.end(), "not find key");
			_udfs_global.erase(fielditer);
			return ;
		}
  
		//add or modify data
		if(fielditer == _udfs_global.end()) {
			_udfs_global.emplace( _self, [&] (auto& g) {
				g.key = key;
				g.val = val;
			}); 
		}
		else {
			_udfs_global.modify(fielditer, _self, [&](auto& g){
				g.val = val;
			}); 
		}
	}

	void consumer::update_storage(){
		udfs_global_table _producer_global(PRODUCER_ACCOUNT, PRODUCER_ACCOUNT.value);
		auto super = _producer_global.get("supernum"_n.value);
		auto supreme = _producer_global.get("supremenum"_n.value);
		auto totalitr = _udfs_global.find("totalstorage"_n.value);
		if (totalitr != _udfs_global.end())
	 		_udfs_global.modify(totalitr, name(0), [&](auto &a){
	 			a.val = super.val * SUPER_STORAGE + supreme.val * SUPREME_STORAGE;
			});
	}
  
	ACTION consumer::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
		require_auth(_self);
		set_global_var(key, val, op);
		update_storage();
	}

	void consumer::transfer(const name& from, const name& to, const asset& quantity, const std::string& memo) 
	{  	
		require_auth(from);
  		print( "consumer::transfer storage  from: ", from, " memo: ", memo, " \n ");
		uosio_assert(quantity.symbol == CORE_SYMBOL, "transfer must system coin");
		uosio_assert(quantity.amount > 0, "asset must positive");
	  
		if(to == _self) {
		//register uid account
		std::string str = memo;
		auto pos = str.find(";", 0);
		if (pos == std::string::npos)	return;
	   
		std::string strBuy = str.substr(0, pos);
		std::string strPack = str.substr(pos+1);
		trim(strBuy);
		trim(strPack);
		uint32_t pack = (uint32_t)stoi(strPack);
		if (strBuy != "buy")		return;

		update_storage();
	   
		auto priceiter = _udfs_global.find( ("udfsprice"_n).value);
		uosio_assert(priceiter != _udfs_global.end(), "not find uidfee field"); 		
		print("consumer::transfer udfsprice : ", priceiter->val, "\n ");
		uosio_assert(priceiter->val >= 1, "consumer  price error");

		int64_t  buystorage = quantity.amount / priceiter->val; 
		uosio_assert(buystorage > 0, "consumer storage too less");
		auto nowtime= now();	

		auto  owneriter = _udfs_user_list.find(from.value);
 
		if(_udfs_user_list.end() == owneriter)
		{
			print( "consumer::transfer  emplace  \n " );
			_udfs_user_list.emplace(_self, [&](auto &a){
				a.user = from;
				a.create_time = nowtime;
				a.exp_time = nowtime;
				a.storage = buystorage;
				a.used = 0;
				a.spentmoney = quantity.amount;
			});
   			print( "consumer::transfer  emplace   ", quantity.amount, "  ", buystorage, "  \n  " );

			auto usernumitr = _udfs_global.find("usernum"_n.value);
			uosio_assert(usernumitr != _udfs_global.end(), "can not find usernum");
			_udfs_global.modify(usernumitr, name(0), [&](auto &g){
				g.val++;
			});
		}else{
			print( "consumer::transfer  modify   ", owneriter->user, "  \n " ); 
			_udfs_user_list.modify(owneriter, name(0), [&](auto &a){
				a.exp_time = nowtime;
				a.storage += buystorage;
				a.spentmoney += quantity.amount;
	 		});
		}

		name fieldname("infund");
		auto fielditer = _udfs_global.find(fieldname.value);
		_udfs_global.modify(fielditer, _self, [&](auto& g){
			g.val += quantity.amount;
		}); 
		auto total = _udfs_global.get("totalstorage"_n.value);
		auto allotitr = _udfs_global.find("allotstorage"_n.value);
		uosio_assert(total.val - allotitr->val >= buystorage, "there are not enough space to sell");
		_udfs_global.modify(allotitr, name(0), [&](auto &g){
			g.val += buystorage;
		});

		/*  modify by tanke 2019.8.6 begin  */
		/*
		struct tm ret_tm;
		localtime(nowtime, ret_tm);
		uint64_t phase = ret_tm.tm_year * 10000 + ret_tm.tm_mon * 100 + ret_tm.tm_mday;
		auto funditr = _udfs_fund.find(phase);
		if (funditr == _udfs_fund.end()){
			_udfs_fund.emplace(_self, [&](auto &a){
				a.phase = phase;
				a.infund = quantity;
			});

		} else {
			_udfs_fund.modify(funditr, name(0), [&](auto &a){
				a.infund += quantity;
			});
		}
		*/
		/*  modify by tanke 2019.8.6 end  */ 	 		  
		}
		else if(from == _self) { // 给udfs 基金
	 	
			if( memo != "udfsfund" ) {
				uosio_assert(false, "consumer:: invalid transfer must spend to udfsfund.");
			} 
			name  fieldname("outfund");
			auto  fielditer = _udfs_global.find(fieldname.value);
			_udfs_global.modify(fielditer, _self, [&](auto& g){
				g.val += quantity.amount;
			}); 
		}
	} // consumer::transfer

	ACTION consumer::setstorage(const name &user, const int64_t &storage){		//OK
		require_auth(_self);
		uosio_assert(user != _self, "this contract account can not be allocated space");
		uosio_assert(is_account(user), "known account");
		uosio_assert(storage != 0, "storage can not be zero");

		update_storage();
	
		auto useritr = _udfs_user_list.find(user.value);
		if (storage > 0){
			if (useritr == _udfs_user_list.end()){
				_udfs_user_list.emplace(_self, [&](auto &a){
					a.user = user;
					a.create_time = now();
					a.storage = storage;
					a.gived = storage;
				});

				auto usernumitr = _udfs_global.find("usernum"_n.value);
				uosio_assert(usernumitr != _udfs_global.end(), "can not find usernum");
				_udfs_global.modify(usernumitr, name(0), [&](auto &g){
					g.val++;
				});
			} else {
				_udfs_user_list.modify(useritr, name(0), [&](auto &a){
					a.storage += storage;
					a.gived += storage;
				});
			}
		
			auto allotitr = _udfs_global.find("allotstorage"_n.value);
			auto total = _udfs_global.get("totalstorage"_n.value);
			uosio_assert(allotitr != _udfs_global.end(), "can not find totalstorage");
		
			_udfs_global.modify(allotitr, name(0), [&](auto &g){
				uosio_assert(total.val - g.val >= storage, "there are not enough space for give");
				g.val += storage;
			});
		}else{
			uosio_assert(useritr != _udfs_user_list.end(), "can not find user");
			uint64_t newstorage = abs(storage);
			if (useritr->gived < newstorage)
				newstorage = useritr->gived;
		
			_udfs_user_list.modify(useritr, name(0), [&](auto &a){
				a.gived -= newstorage;
				a.storage -= newstorage;
			});
			
			auto allotitr = _udfs_global.find("allotstorage"_n.value);
			uosio_assert(allotitr != _udfs_global.end(), "can not find usedstorage");
			_udfs_global.modify(allotitr, name(0), [&](auto &g){
				g.val -= newstorage;
			});
		}
		//uosio_assert(useditr->val <= totalitr->val, "no space to gived");
	}

	ACTION consumer::setused(const name &user, const int64_t &storage){			//OK
		require_auth(_self);
		uosio_assert(storage != 0, "storage can not be zero");
		auto useritr = _udfs_user_list.find(user.value);
		uosio_assert(useritr != _udfs_user_list.end(), "can not find the user");
		auto useditr = _udfs_global.find("usedstorage"_n.value);
		uosio_assert(useditr != _udfs_global.end(), "can not find usedstorage");
		if (storage > 0){
			_udfs_user_list.modify(useritr, name(0), [&](auto &a){
				uosio_assert(a.storage >= (a.used + storage), "too less storage space to store the file");
				a.used += storage;
			});

			_udfs_global.modify(useditr, name(0), [&](auto &g){
				g.val += storage;
			});
		}else{
			uint64_t newstorage = abs(storage);
			_udfs_user_list.modify(useritr, name(0), [&](auto &a){
				uosio_assert(a.used >= newstorage, "too much storage space to release");
				a.used -= newstorage;
			});

			_udfs_global.modify(useditr, name(0), [&](auto &g){
				g.val -= newstorage;
			});
		}
	
	}

	ACTION consumer::clear(const uint64_t &num){
		require_auth(_self);
		auto globalitr = _udfs_global.begin();
		uint64_t i = 0;
		for(; globalitr != _udfs_global.end() && i < num; ++i){
			globalitr = _udfs_global.erase(globalitr);
		}
	
		auto useritr = _udfs_user_list.begin();
		for(; useritr != _udfs_user_list.end() && i < num; ++i){
			useritr = _udfs_user_list.erase(useritr);
		}
		/*
		auto funditr = _udfs_fund.begin();
		for(;funditr != _udfs_fund.end() && i < num; ++i){
			funditr = _udfs_fund.erase(funditr);
		}
		*/
	}
  
	void consumer::apply(uint64_t receiver, uint64_t code, uint64_t action) {
		auto &thiscontract = *this;
		if (action == ( "transfer"_n ).value && code == ( "uosio.token"_n ).value ) {
			auto transfer_data = unpack_action_data<st_transfer>();
			//transfer(transfer_data.from, transfer_data.to, transfer_data.quantity, transfer_data.memo);
			return;
		}
		uosio_assert( code == receiver, "consumer must self");  
	}     
} // udfs_consumer

extern "C" {
	[[noreturn]] void apply(uint64_t receiver, uint64_t code, uint64_t action) {
		udfs_consumer::consumer udfsc(uosio::name(receiver),uosio::name(code),uosio::datastream<const char*>(nullptr, 0) );
		udfsc.apply(receiver, code, action);

		auto self = receiver;
		if (code == self){
			switch (action) {
				UOSIO_DISPATCH_HELPER( udfs_consumer::consumer,  (init)(setdata)(setstorage)(setused)(clear) );
			} 
		}
		uosio_exit(0);
	}
}
