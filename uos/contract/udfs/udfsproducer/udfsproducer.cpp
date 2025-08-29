#include "udfsproducer.hpp"

using namespace uosio;

namespace udfsuossys {
	udfsproducer::udfsproducer(name receiver, name code, datastream<const char*> ds):
	contract(receiver, code, ds),
	_udfsglobal(receiver, receiver.value),
	_udfsnode(receiver, receiver.value),
	_udfsreward(receiver, receiver.value)
	{ }

	//IP convert to uint
	uint32_t udfsproducer::ipToInt(std::string strIp)
	{
		uosio_assert(strIp.size() <= 15, "this is not an IP");
		uint32_t nRet = 0;
		std::string strTmp;
		std::size_t pos;

		for(int i = 0; i < 3; ++i){
			pos = strIp.find(IP_MARK, 0);
			uosio_assert(pos != std::string::npos && pos <= 3, "this is not an IP");
			strTmp = strIp.substr(0, pos);
			strIp = strIp.substr(pos+1);
			uosio_assert(std::stoi(strTmp) >= 0 && std::stoi(strTmp) <= 255, "this is not an ip");
			nRet += std::stoi(strTmp) << ((3-i)*8);
			
		}
		uosio_assert(strIp.size() <= 3, "this is not an IP");
		uosio_assert(std::stoi(strTmp) >= 0 && std::stoi(strIp) <= 255, "this is not an ip");
		nRet += std::stoi(strIp);
		uosio_assert(isPublic(nRet), "this is not a public IP");
		return nRet;
	}

	//uint convert to IP
	std::string udfsproducer::intToIp(uint32_t num){  
		std::string strIP = "";  
		for (int i = 0; i < 4; ++i){  
			uint32_t tmp = (num >> ((3-i)*8)) & 0xFF;  

			strIP += std::to_string(tmp);
			if (i < 3){
				strIP += IP_MARK;
			}
		}  

		return strIP;  
	} 

	//IP filter：
	//local loop network address--127.0.0.1(2130706433)-127.255.255.254(2147483646);
	//private network address--10.0.0.0(167772160)-10.255.255.255(184549375)、192.168.0.0(3232235520)-192.168.255.255(3232301055)、172.16.0.0(2886729728)-172.16.255.255(2886795263);
	//specific network address--0.0.0.0(0)、169.254.0.1(2851995649)-169.254.255.255(2852061183)
	bool udfsproducer::isPublic(uint32_t num){
		if (num == 0)	return false;
		if (num >= 2130706433 && num <= 2147483646)	return false;
		if (num >= 167772160  && num <= 184549375 )	return false;
		if (num >= 3232235520 && num <= 3232301055)	return false;
		if (num >= 2886729728 && num <= 2886795263)	return false;
		if (num >= 2851995649 && num <= 2852061183)	return false;
		return true;
		
	}

	//phase convert to scope
	std::string udfsproducer::phaseToStr(uint64_t phase){
		std::string strScope = std::to_string(phase);
		uosio_assert(strScope.size() <= 12, "phase error, length should less than 12 charactor");
		for (int i = 0; i < strScope.size(); ++i){
			strScope[i] += 49;			//convert every character，0--a,1--b, and so on
		}
		return strScope;
	}
	
	void udfsproducer::sendInline(name to, asset quantity, std::string memo){
		action(
			permission_level{_self, "active"_n},
			"uosio.token"_n, "transfer"_n,
			std::make_tuple(_self, to, quantity, memo) 	//std::string("transfer")
		).send();
	}
	
	void udfsproducer::setKey(name key, uint64_t value, bool del){
		auto itr = _udfsglobal.find(key.value);

		if (del) {
			uosio_assert(itr != _udfsglobal.end(), "the key doesn't exist");
			_udfsglobal.erase(itr);
			return;
		}

		if (itr == _udfsglobal.end()) {
			_udfsglobal.emplace(_self, [&](auto& one){
				one.key = key;
				one.val = value;
			});
		} else {
			_udfsglobal.modify(itr, name(0), [&](auto& one){
				one.val = value;
			});
		}
	}

	void udfsproducer::addValue(name key, uint64_t plus, bool negative){ 
		auto itr = _udfsglobal.find(key.value);

		if (itr == _udfsglobal.end()) {
			uosio_assert(false, "can not find the key");
			uosio_assert(plus >= 0, "emplace, value can't be negative");
			_udfsglobal.emplace(_self, [&](auto& one){
				one.key = key;
				one.val = plus;
			});
		} else {
			_udfsglobal.modify(itr, name(0), [&](auto& one){
				if (negative == true) {
					uosio_assert(one.val >= plus, "modify, numerical downward overflow");
					one.val -= plus;
				} else {
					uosio_assert((MAX_KEY_VALUE - plus) >= one.val, "modify, numerical upward overflow");
					one.val += plus;
				}
			});
		}
	}

/*
	void udfsproducer::checkFreeze(){
		auto freezeitr = _udfsglobal.find("freeze"_n.value);
		if (freezeitr != _udfsglobal.end())
			uosio_assert(freezeitr->val == 0, "the contract if freezed");
	}
*/
	void udfsproducer::init(){					//OK
		require_auth2(_self.value, "udfs"_n.value);
		auto itr = _udfsglobal.begin();
		uosio_assert(itr == _udfsglobal.end(), "the contract already initalized");
		setKey("phase"_n, 0);						//period
		setKey("startat"_n, now());					//timestamp of the period start at
		setKey("duration"_n, PERIOD_SECONDS); 		//duration of every period
		setKey("supernum"_n, 0);					//number of super nodes
		setKey("supremenum"_n, 0);					//number of supreme nodes
		//setKey("freezed"_n, 0);					//freeze contract
		setKey("multiple"_n, MULTIPLE);				//times of supreme nodes(x10)
		setKey("first"_n, FIRST);					//how many points eliminate the first penalty level
		setKey("second"_n, SECOND);
		setKey("third"_n, THIRD);
		setKey("low"_n, LOW);						//only 80% of the reward will be awarded when the node at the first penalty level
		setKey("mid"_n, MID);						//only 50% of the reward will be awarded when the node at the second penalty level, the third penalty level get zero

		setKey("phasenum"_n, 0);					//number of period
		setKey("utreward"_n, 0);					//all the reward with nodes which buy the token by UT
		setKey("ethreward"_n, 0);					//all the reward with nodes which buy the token by ETH
	}

	void udfsproducer::loginnode(name owner, std::string strip, uint8_t level){			//OK
		require_auth(owner);
		uosio_assert(level == 1, "unknow node type");

		auto nodeitr = _udfsnode.find(owner.value);

		uint32_t ip = ipToInt(strip);
		auto node_index = _udfsnode.get_index<"byip"_n>();
		auto ipitr = node_index.find(ip);
		uosio_assert(ipitr == node_index.end(), "this ip already be bounded");

		udfstoken_t _udfstoken(UNION_ACCOUNT, UNION_ACCOUNT.value);
		auto token_index = _udfstoken.get_index<"owner"_n>();
		auto tokenitr = token_index.find(owner.value);
		uosio_assert(tokenitr != token_index.end() && tokenitr->owner == owner.value, "this user have no token");

		if (nodeitr == _udfsnode.end()) {
			_udfsnode.emplace(_self, [&](auto &a){
				a.owner = owner;
				a.ip = ip;
				a.login_at = tokenitr->apply_time;
				//a.login_at = now();
				a.level = level;
				a.coin_type = tokenitr->coin_type;
				a.amount = tokenitr->amount;
				a.status = 1;
			});	
		} else {
			uosio_assert(nodeitr->status == 0, "this node already logined");
        	uosio_assert(!isPublic(nodeitr->ip), "this node already bonunded ip");
			_udfsnode.modify(nodeitr, name(0), [&](auto &a){
				a.ip = ip;
				//a.login_at = now();
				a.level = level;
				a.status = 1;
			});
		}

		if (level == 1) {
			addValue("supernum"_n, 1);
		} else {
			addValue("supremenum"_n, 1);
		}
		
	}
	
	void udfsproducer::updatenode(name owner, int8_t status, bool del){			//OK
		require_auth2(_self.value, "udfs"_n.value);
		uosio_assert(status != 0, "can not modify node status to 0");
		auto node = _udfsnode.find(owner.value);
		uosio_assert(node != _udfsnode.end(), "the node isn't in node table");
		uosio_assert(node->status != 0, "the node already logouted");
		del = false;
		if (del) {
			_udfsnode.erase(node);
			
		} else {
			uosio_assert(status >= -4 && status <= 2, "unknow status");
			_udfsnode.modify(node, name(0), [&](auto& one){
				one.status = status;
			});
		}
	}

	void udfsproducer::logoutnode(name owner){					//OK
		require_auth(owner);
		auto node = _udfsnode.find(owner.value);
		uosio_assert(node != _udfsnode.end(), "the node isn't in node table");
		uosio_assert(node->status != 0, "the node already logouted");
		
		//TODO:每期仅允许注销一次

		if (node->level == 1) {
			addValue("supernum"_n, 1, true);
		} else if (node->level == 2) {
			addValue("supremenum"_n, 1, true);
		} else {
			uosio_assert(false, "unknow node type");
		}

		udfstoken_t _udfstoken(UNION_ACCOUNT, UNION_ACCOUNT.value);
		auto token_index = _udfstoken.get_index<"owner"_n>();
		auto tokenitr = token_index.find(owner.value);
		if (tokenitr == token_index.end() || tokenitr->owner != owner.value){
			_udfsnode.erase(node);
		}

		_udfsnode.modify(node, name(0), [&](auto &one){
			one.ip = 0;
			one.level = 0;
			one.status = 0;
		});

	}

	void udfsproducer::unbindnode(name owner){
		require_auth2(_self.value, "udfs"_n.value);
		auto node = _udfsnode.find(owner.value);
		uosio_assert(node != _udfsnode.end(), "the node isn't in node table");
		uosio_assert(node->status != 0, "the node already logouted");

		if (node->level == 1) {
			addValue("supernum"_n, 1, true);
		} else if (node->level == 2) {
			addValue("supremenum"_n, 1, true);
		} else {
			uosio_assert(false, "unknow node type");
		}

		udfstoken_t _udfstoken(UNION_ACCOUNT, UNION_ACCOUNT.value);
		auto token_index = _udfstoken.get_index<"owner"_n>();
		auto tokenitr = token_index.find(owner.value);
		if (tokenitr == token_index.end() || tokenitr->owner != owner.value){
			_udfsnode.erase(node);
		}

		_udfsnode.modify(node, name(0), [&](auto &one){
			one.ip = 0;
			one.level = 0;
			one.status = 0;
		});
	}

	void udfsproducer::punishnode(name owner){				//OK
		require_auth2(_self.value, "udfs"_n.value);
		auto node = _udfsnode.find(owner.value);
		uosio_assert(node != _udfsnode.end(), "the node isn't in node table");
		auto super = _udfsglobal.get("supernum"_n.value);
		auto supreme = _udfsglobal.get("supremenum"_n.value);
		auto multi = _udfsglobal.get("multiple"_n.value);
		auto first = _udfsglobal.get("first"_n.value);
		auto second = _udfsglobal.get("second"_n.value);
		auto third = _udfsglobal.get("third"_n.value);
		auto node_num = super.val * 10 + supreme.val * multi.val;		//nodenum放大了10倍，因此newscore也需要放大10倍

		if(node_num <= 0)	node_num = 10;	
		uint64_t newscore = node->punish >= 2 ? SCORE * third.val * 10 / node_num : (node->punish == 1 ? SCORE * second.val * 10 / node_num : SCORE * first.val * 10 / node_num);
		//uint8_t exp = node->punish >= 2 ? THIRD : (node->punish == 1 ? SECOND : FIRST);
		
		_udfsnode.modify(node, name(0), [&](auto &one){
			if (one.punish <= 2)
				one.punish ++;
			if (one.score < newscore)
				one.score += newscore;
		});	
		
	}

	void udfsproducer::setbonus(name owner, uint64_t phase, uint8_t level, uint64_t score, asset award){			//OK
		require_auth2(_self.value, "udfs"_n.value);
		std::string scope = phaseToStr(phase);
		udfsbonus_t _udfsbonus(_self, name(scope).value);
		
		uosio_assert(score <= SCORE, "score can not more than 100000");
		uosio_assert(level == 1, "node leve error, should be 1");
		uosio_assert(award > ZERO_ASSET, "must use system coin, and amount must more than 0");
		auto reward = _udfsreward.find(phase);
		uosio_assert(reward != _udfsreward.end(), "can not find the reward");
		uosio_assert(reward->status != 2, "can not insert record in settlement");	//状态为2，则表示本期已经开始进行结算

		auto node = _udfsnode.find(owner.value);
		uosio_assert(is_account(owner), "the account is not exist");
		uosio_assert(node != _udfsnode.end(), "the node isn't in node table");
		
		auto bonus = _udfsbonus.find(owner.value);
		uosio_assert(bonus == _udfsbonus.end(), "the node is already in bonus table");

		auto multitr = _udfsglobal.get("multiple"_n.value);
		uint8_t multiple = level == 2 ? multitr.val : 10;		//基础奖励倍数，至尊节点是超级节点1.2倍

		if (reward != _udfsreward.end()){
			_udfsreward.modify(reward, name(0), [&](auto &one){
				one.nodes += multiple;
				one.score += score;
				one.bill += award;
			});
		} 
		
		_udfsbonus.emplace(_self, [&](auto &one){
			one.owner = owner;  	
			one.ip = node->ip;
			one.level = level; 
			one.score = score;
			one.award = award;
		});	

	}

	void udfsproducer::clearbonus(uint64_t phase, uint32_t num){			//OK
		require_auth2(_self.value, "udfs"_n.value);

		std::string scope = phaseToStr(phase);
		udfsbonus_t _udfsbonus(_self, name(scope).value);
		auto bonus = _udfsbonus.begin();

		auto reward = _udfsreward.find(phase);
		uosio_assert(reward != _udfsreward.end(), "unknow period");

		if (reward->status == 0)
			_udfsreward.modify(reward, name(0), [&](auto &one){
				one.status = 1;
			});

		for (uint32_t i = 0; i < num; ++i){
			if (bonus != _udfsbonus.end())
				bonus = _udfsbonus.erase(bonus);
			
			if(bonus == _udfsbonus.end()){
				_udfsreward.modify(reward, name(0), [&](auto &one){
					one.nodes = 0;
					one.score = 0;
					one.bill = ZERO_ASSET;
				});
				break;
			}	
		}
	}
	
	void udfsproducer::paybonus(name owner, uint64_t phase){			//OK,考虑实时修改奖池金额
		require_auth2(_self.value, "udfs"_n.value);
		
		std::string scope = phaseToStr(phase);
		udfsbonus_t _udfsbonus(_self, name(scope).value);
		auto node = _udfsnode.find(owner.value);
		uosio_assert(node != _udfsnode.end(), "can not find the node");
		auto bonus = _udfsbonus.find(owner.value);
		auto reward = _udfsreward.find(phase);
		uosio_assert(reward != _udfsreward.end(), "can not find the record in udfsrewad");

		if (_udfsbonus.begin() == _udfsbonus.end()){
			_udfsreward.erase(reward);
			return;
		}
		uosio_assert(bonus != _udfsbonus.end(), "can not find the record in udfsbonus");
		
		if (reward->status != 2){

			uosio_assert(reward->bill <= reward->profit, "reward error, bill should less than profit");		//核对奖池总额
			_udfsreward.modify(reward, name(0), [&](auto &one){
				one.status = 2;
			});
		}

		auto multitr = _udfsglobal.get("multiple"_n.value);
		uint8_t multiple = bonus->level == 2 ? multitr.val : 10;		//基础奖励倍数，至尊节点是超级节点1.2倍,此处放大了10倍
		asset contribution = ZERO_ASSET;
		
		contribution = bonus->award;//reward->bill * bonus->score / SCORE;			//贡献奖励

		asset award = contribution;

		//uosio_assert(award <= bonus->award, "award error");
		//award = bonus->award;
		asset newcontribution = contribution;
	
		//惩罚节点
		if (node->punish >= 3) { 
			newcontribution = ZERO_ASSET;
		} else if (node->punish == 2) {	
			auto mid = _udfsglobal.get("mid"_n.value);
			newcontribution = contribution * mid.val / 100;
		} else if (node->punish ==1 ) {
			auto low = _udfsglobal.get("low"_n.value);
			newcontribution = contribution * low.val / 100;
		}

		award = newcontribution;
		if (award > ZERO_ASSET) {
			std::string memo = "phase:" + std::to_string(phase) + 
							   ",punish:" + std::to_string(node->punish) + 
						   	   ",score:" + std::to_string(bonus->score);
			sendInline(owner, award, memo);
		}

		asset back = contribution - newcontribution;
		if (back > ZERO_ASSET)
			sendInline(REWARD_ACCOUNT, back, "punish asset to profit");

		if (node->coin_type == 1) {
			addValue("utreward"_n, newcontribution.amount);
			addValue("utreward"_n, back.amount, true);
		} else if (node->coin_type == 2){
			addValue("ethreward"_n, newcontribution.amount);
			addValue("ethreward"_n, back.amount, true);
		} else {
			uosio_assert(false, "unknow coin type");
		}
			
		_udfsnode.modify(node, name(0), [&](auto &one){
			one.award += award;
		});
		
		if (node->score != 0){
			if (node->score <= bonus->score) {
				_udfsnode.modify(node, name(0), [&](auto &one){
					one.score = 0;
					one.punish = 0;
				});
			} else {
				_udfsnode.modify(node, name(0), [&](auto &one){
					one.score -= bonus->score;
				});
			}
		}

		bonus = _udfsbonus.erase(bonus);
		if (_udfsbonus.begin() == _udfsbonus.end()){
			_udfsreward.erase(reward);
		}
	}

	void udfsproducer::updatekey(name key, uint64_t value, bool del){			//OK
		require_auth2(_self.value, "udfs"_n.value);
		setKey(key, value, del);		
	}
	
	void udfsproducer::setperiod(uint64_t phase){				//OK
		require_auth2(_self.value, "udfs"_n.value);
		uosio_assert(std::to_string(phase).size() <= 12, "phase error, length should not more than 12 charactor");
		auto thephase = _udfsglobal.find("phase"_n.value);
		uosio_assert(thephase != _udfsglobal.end(), "global error, can't find phase param");
		uosio_assert(thephase->val < phase, "new phase must larger than current phase");
		auto duration = _udfsglobal.get("duration"_n.value);
		auto start = _udfsglobal.find("startat"_n.value);
		uosio_assert((now() - start->val) >= duration.val, "the period is not reach expiration");	//TODO:可以稍微改小间隔
		auto phasenum = _udfsglobal.find("phasenum"_n.value);

		if (thephase->val != phase){
			_udfsglobal.modify(thephase, name(0), [&](auto &one){
				one.val = phase;
			});

			_udfsglobal.modify(start, name(0), [&](auto &one){
				one.val = now();
			});

			_udfsglobal.modify(phasenum, name(0), [&](auto &one){
				one.val ++;
			});

			_udfsreward.emplace(_self, [&](auto &one){
				one.phase = phase;
			});
		}
	}

	/*void udfsproducer::cleannodes(name owner){
		require_auth(_self);
		
		auto nodeitr = _udfsnode.begin();
		udfstoken_t _udfstoken(UNION_ACCOUNT, UNION_ACCOUNT.value);
		auto token_index = _udfstoken.get_index<"owner"_n>();
		auto tokenitr = token_index.begin();
		
		for (uint32_t i = 0; i < num; i++){
			tokenitr++;
		}
		uint32_t index = 0;
		for(;tokenitr != token_index.end() && index < 100;){
		//if(tokenitr != token_index.end()){
			nodeitr = _udfsnode.find(tokenitr->owner);
			if (nodeitr == _udfsnode.end()){
				_udfsnode.emplace(_self, [&](auto &a){
					a.owner = name(tokenitr->owner);
					a.ip = 0;
					a.login_at = tokenitr->apply_time;
					a.level = 0;
					a.coin_type = tokenitr->coin_type;
					a.amount = tokenitr->amount;
				});
			}
			tokenitr++;
			index ++;
		}
		
		//uint32_t index = 0;
		for(; nodeitr != _udfsnode.end() &&  index < num; index++){
			tokenitr = token_index.find(nodeitr->owner.value);
			if (tokenitr == token_index.end() || tokenitr->owner != nodeitr->owner.value){
				_udfsnode.erase(nodeitr);
			}
			nodeitr ++;
		}
		
		auto nodeitr = _udfsnode.find(owner.value);
		udfstoken_t _udfstoken(UNION_ACCOUNT, UNION_ACCOUNT.value);
		auto token_index = _udfstoken.get_index<"owner"_n>();
		auto tokenitr = token_index.find(owner.value);

		if (tokenitr == token_index.end() || tokenitr->owner != nodeitr->owner.value){
                _udfsnode.erase(nodeitr);
        }
				

	}*/

	void udfsproducer::clear(uint32_t num){				//OK
		require_auth2(_self.value, "udfs"_n.value);
		int i = 0;
		auto itr = _udfsglobal.begin();
		for(; itr != _udfsglobal.end() && i < num; ++i) 
			itr = _udfsglobal.erase(itr);

		auto itr1 = _udfsnode.begin();
		for(; itr1 != _udfsnode.end() && i < num; ++i)
			itr1 = _udfsnode.erase(itr1);

		auto itr2 = _udfsreward.begin();
		for(; itr2 != _udfsreward.end() && i < num; ++i){
			auto phase =  itr2->phase;
			std::string scope = phaseToStr(phase);
			udfsbonus_t _udfsbonus(_self, name(scope).value);
			auto bonusitr = _udfsbonus.begin();
			for (; i < num; ++i){
				if (bonusitr != _udfsbonus.end()) {
					bonusitr = _udfsbonus.erase(bonusitr);
				} else {
					break;
				}
			}
			
			if (bonusitr == _udfsbonus.end())
				itr2 = _udfsreward.erase(itr2);
		}        
	}

	void udfsproducer::transfer(name from, name to, asset quantity, std::string memo){			//OK
		if(from == _self){
			//require_auth(_self);
			return;
		}
		uosio_assert(  to == _self , "must transfer to this contract");
		uosio_assert(  quantity.symbol == CORE_SYMBOL, "must use system coin");
		
		auto pos = memo.find(";", 0);
		std::string strType = memo.substr(0,pos);
		//uosio_assert(strType == "profit", "type error");
		/*std::string strPhase = memo.substr(pos+1);
		std::string::size_type sz = 0;
		uint64_t thephase = (uint64_t)stoull( strPhase, &sz, 10 );

		auto reward = _udfsreward.find(thephase);
		uosio_assert(reward != _udfsreward.end(), "can not find the phase");

		uosio_assert(from == REWARD_ACCOUNT, "reward account error");
		_udfsreward.modify(reward, name(0), [&](auto &one){
				one.profit += quantity;
			});
		*/	
		
		if (strType == "transfer"){
			/*uosio_assert(from == SUPPLY_ACCOUNT, "supply account error");
			auto maxsup = _udfsglobal.get("maxsup"_n.value);
			uosio_assert((quantity.amount + reward->supply.amount) <= maxsup.val, "supply is too much");
			
			_udfsreward.modify(reward, name(0), [&](auto &one){
				one.supply += quantity;
			});
			*/
		} else if (strType == "profit"){
			std::string strPhase = memo.substr(pos+1);
        	std::string::size_type sz = 0;
        	uint64_t thephase = (uint64_t)stoull( strPhase, &sz, 10 );

			auto reward = _udfsreward.find(thephase);
			uosio_assert(reward != _udfsreward.end(), "can not find the phase");
			uosio_assert(from == REWARD_ACCOUNT, "profit account error");
			_udfsreward.modify(reward, name(0), [&](auto &one){
				one.profit += quantity;
			});
			
		} else {
			uosio_assert(false, "type error");
		}
		
	}

	void udfsproducer::nodevoter(uint64_t owner){
		auto nodeitr = _udfsnode.find(owner);
		uosio_assert(is_account(name(owner)), "the account is not exist");
		udfstoken_t _udfstoken(UNION_ACCOUNT, UNION_ACCOUNT.value);
		auto token_index = _udfstoken.get_index<"owner"_n>();

		auto itr = token_index.find(owner);
		uosio_assert(itr != token_index.end() && itr->owner == owner, "token error");

		if (nodeitr == _udfsnode.end()){
			_udfsnode.emplace(_self, [&](auto &a){
				a.owner = name(owner);
				a.ip = 0;
				a.login_at = itr->apply_time;
				a.level = 0;
				a.coin_type = itr->coin_type;
				a.amount = itr->amount;
			});
		} else {
			_udfsnode.modify(nodeitr, name(0), [&](auto &a){
				a.coin_type = itr->coin_type;
				a.amount = itr->amount;
				a.login_at = itr->apply_time;
			});
		}
	}

	void udfsproducer::apply(uint64_t receiver, uint64_t code, uint64_t action) {
		auto &thiscontract = *this;
		if (action == ( "transfer"_n ).value && code == ( "uosio.token"_n ).value ) {
			auto transfer_data = unpack_action_data<st_transfer>();
			transfer(transfer_data.from, transfer_data.to, transfer_data.quantity, transfer_data.memo);
			return;
		}
		
		if (code == (UNION_ACCOUNT.value)){
			uint64_t owner;
			if (action == "voter"_n.value)
				owner = unpack_action_data<st_nodevoter>().owner;
			else if (action == "addnode"_n.value)
				owner = unpack_action_data<st_addnode>().owner;
			else 
				return;
			nodevoter(owner);
			return;
		}

		if (code != get_self().value) return;
		switch (action) {
			UOSIO_DISPATCH_HELPER( udfsproducer, 
				(init)(loginnode)(updatenode)(logoutnode)(unbindnode)(punishnode)(setbonus)(clearbonus)(paybonus)(updatekey)(setperiod)(clear)/*(cleannodes)(addnode)*/ );
		}; 
	}
	
}

extern "C" {
	[[noreturn]] void apply(uint64_t receiver, uint64_t code, uint64_t action) {
		udfsuossys::udfsproducer udfs( name(receiver), name(code), datastream<const char*>(nullptr, 0) );
		udfs.apply(receiver, code, action);
		uosio_exit(0);
	}
}


