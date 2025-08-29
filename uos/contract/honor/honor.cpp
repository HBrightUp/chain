#include "honor.hpp"

namespace honorsys {
    using namespace uosio;

    honor::honor(name receiver, name code,  uosio::datastream<const char*> ds)
     :contract(receiver, code, ds),
     _global_var(_self, _self.value),
     _honor_list(_self, _self.value),
     _ctribut_list(_self, _self.value),
     _uosc_bp_list("uosio"_n, ("uosclist"_n).value),
     _project_list(_self, _self.value)
     {
     }

    void honor::set_global_var(const name& key, const uint64_t& val, const uint8_t op) {

        uosio_assert(op == 1 || op == 2, "op 1: add or modify; 2: delete");
    
        auto global = _global_var.find(key.value);

        //delete data
        if(op == 2) {
            uosio_assert(global != _global_var.end(), "not find key");
            _global_var.erase(global);
            return ;
        }

        //add or modify data
        if(global == _global_var.end()) {
            _global_var.emplace( _self, [&] (auto& g) {
                g.key = key;
                g.val = val;
            }); 
        }
        else {
            _global_var.modify(global, _self, [&](auto& g){
                g.val = val;
            }); 
        }
    }

    bool honor::getbplist(std::vector<name>& bplist) {

        //if empty, error returns immediately 
        if(_uosc_bp_list.begin() == _uosc_bp_list.end()) {
            return false;
        }
       
        auto bp = _uosc_bp_list.rbegin();
        uint64_t time = bp->bp_valid_time;

        //get bp list from table _uosc_bp_list
        while( bp != _uosc_bp_list.rend()) {
            
            //interval time between two bp less than 2
            if(time - bp->bp_valid_time > 2) {
                break;
            }

            time = bp->bp_valid_time;
            bplist.emplace_back(bp->bpname);
            ++bp;
        }

        return true;
    }

    asset honor::get_balance( name token_contract_account, name owner, symbol_code sym_code )
    {   
        //get banlance by given account
        accounts accountstable( token_contract_account, owner.value );
        const auto& a = accountstable.get( sym_code.raw() );
        return a.balance;
    }

    ACTION honor::init() {
        
        require_auth(_self);

        name key;
        uint64_t val;

        //max number of bp
        set_global_var("maxbpnumber"_n, 15);

        //valid time of project;  
        set_global_var("overtime"_n, 3 * 24 * 3600);

        //max number of project applied by honor
        set_global_var("maxapplynum"_n, 6);

        //max credit given by bp
        set_global_var("maxcredit"_n, 100000);

        //min uos pay for registering honor account. default: 1.0000 UOS
        set_global_var("minregpay"_n, 10000);

        //max number of credit to uos once time
        set_global_var("maxconvert"_n, 10000);

        //asset of contract acccount
        symbol_code sym_code("UOS");
        auto balance = get_balance("uosio.token"_n, _self, sym_code);
        set_global_var("fund"_n, balance.amount);

        //the ratio of credit to uos
        set_global_var("ratio"_n, 1);

        //pay for register a honor account
        set_global_var("registerpay"_n, 100000);

        //grade nmae: grade + char('a'-'z')
        set_global_var("gradea"_n, 6000);        
        set_global_var("gradeb"_n, 10000);   
        set_global_var("gradec"_n, 15000);         
        set_global_var("graded"_n, 22000);
        set_global_var("gradee"_n, 30000);
        set_global_var("gradef"_n, 40000);
    }

    ACTION honor::setdata(const name& key, const uint64_t& val, const uint8_t& op) {
        require_auth(_self);
        set_global_var(key, val, op);
    }

    ACTION honor::reward(const name project) {

        auto idx = _project_list.find(project.value);
        uosio_assert(idx != _project_list.end(), "project not exist");

        //whether current project is overtime
        auto overtime = _global_var.find(("overtime"_n).value);
        uosio_assert(overtime != _global_var.end(), "cannot find overtime field");

        if(idx->update_time  + overtime->val <  now()) {
            _project_list.modify(idx, idx->owner, [&](auto& p){
                p.state = state_overtime;
            }); 

            return ;
        }

        uosio_assert(idx->state == state_ongoing, "not a ongoing project");

        require_auth(idx->owner);

        //get the total count of all bp assessed
        uint64_t total_credit = 0;

        for(auto bpcredit : idx->credit_list) {
            total_credit += bpcredit.credit;
        }

        std::vector<name> bplist;
        uosio_assert(getbplist(bplist), "get bplist falied");

        //only 2/3 bp given the assess, the honor account can get the award
        //uosio_assert(static_cast<double>(idx->credit_list.size()) / bplist.size() >= 2.0 / 3, "must 2/3 bp assessed project" );
        uosio_assert(idx->credit_list.size()  >=    bplist.size() * 2 / 3, "must 2/3 bp assessed project" );

        auto honor = _honor_list.find( idx->owner.value);
        uosio_assert( honor != _honor_list.end(), "receiver not a honor account");
        
        auto credit = total_credit / bplist.size();

        _honor_list.modify(honor, honor->account, [&](auto& h) {
            h.credit += credit;
        });

         _project_list.modify(idx, idx->owner, [&](auto& p){
            p.credit = credit;
            p.state = state_finished;
        }); 

    }

    //user action

    ACTION honor::upgrade(const name& user, uint64_t credit) {
        require_auth(user);

        std::string base = "grade";
        std::string dest;
        uint64_t total_grade = 1;

        //find the scope of grade
        auto begin = _global_var.find(("gradea"_n).value);
        uosio_assert(begin != _global_var.end(), "not find gradea field");
        auto end = begin;

        for (char ch = 'b'; ch <= 'z'; ++ch) {
            dest = base + ch;
            ++end;
            if(end->key == uosio::name(dest)) { 
                ++total_grade;
            }
            else {
                break;
            }
        }

        print("begin: ", begin->val, " end: ", end->val, "total_grade: ", total_grade);
        
        auto honor = _honor_list.find( user.value);
        uosio_assert( honor != _honor_list.end(), "user not a  honor account");

        //check grade
        uosio_assert(honor->vip < total_grade + 1, "full grade");
        uosio_assert(honor->credit >= credit, "not enough credit");

        auto vip = honor->vip;
        uint8_t pos = 1;
        auto ct = credit;
        auto last_vip = vip;
        auto used_credit = 0;

        //honor account upgrade 
        for(auto idx = begin; idx != end; ++idx ) {
            auto require_credit = idx->val;

            if(vip > pos) {
                ++pos;
                continue;
            }
            
            if( ct >= require_credit) {
                ++last_vip;
                ct = ct - require_credit;
                used_credit += require_credit;
            } else {
                break;
            }
        }

        if(last_vip > vip) {
            _honor_list.modify( honor, honor->account, [&](auto& h) {
                h.vip = last_vip;
                h.credit -= used_credit;
            });
        }
    }

    ACTION honor::convertuos(const name& user, uint64_t credit) {

        require_auth(user);

        auto honor = _honor_list.find( user.value);

        //check honor account 
        uosio_assert( honor != _honor_list.end(), "user is not a honor account");

        //check credit os user
        uosio_assert(honor->credit >= credit, "not enough credit");

        //check time of convert
        uosio_assert(now() - honor->last_convert >=  seconds_one_day, "only one time convert credit to uos every day");

        //check max convert  from credit to uos
        auto maxconvert = _global_var.find( ("maxconvert"_n).value);
        uosio_assert(maxconvert != _global_var.end(), "not find maxconvert field");
        uosio_assert(credit <= maxconvert->val, "convert credit to uos to many");

        //calculate valid credit
        auto ratio = _global_var.find( ("ratio"_n).value);
        uosio_assert(ratio != _global_var.end(), "not find ratio field");
        auto rate = ratio->val;

        auto amount = credit / rate;
        uosio_assert(amount > 0, "not enough credit to convert");

        auto used_credit = amount * rate;
        auto real_amount = amount * 10000;
        
        auto fund = _global_var.find( ("fund"_n).value);
        uosio_assert(fund != _global_var.end(), "not find fund field");
        uosio_assert(fund->val >= real_amount, "not enought fund");

        _honor_list.modify( honor, honor->account, [&](auto& h) {
            h.credit -= used_credit;
            h.last_convert = now();
        });
        
        auto fund_update = fund->val - real_amount;
        set_global_var("fund"_n, fund_update, 1);

        symbol sym("UOS", 4);
        asset crt(real_amount, sym);
        
        //transfer asset from contract account to user
        auto action_data = make_tuple(_self, user, crt, std::string("convertuos"));
        action(permission_level{_self, "active"_n}, "uosio.token"_n, "transfer"_n, action_data).send();
    }

    ACTION honor::apply(name project, name proposer, std::string reason) {
        
        require_auth(proposer);

        auto pro = _project_list.find(project.value);
        auto honor = _honor_list.find(proposer.value);

        //check para
        uosio_assert(reason.length() <= 256 && reason.length() > 0, "invalid reason");  
        uosio_assert(honor != _honor_list.end(), "proposer not a honor account");
        uosio_assert(pro == _project_list.end(), "project existed");

        uint32_t counts_apply = 0;
        auto apply =  _project_list.get_index<"owner"_n>();
        auto begin = apply.lower_bound(proposer.value);
        auto end = apply.upper_bound(proposer.value);

        for(auto pos = begin; pos != end; ++pos) {
            if(pos->owner == proposer && pos->state == state_ongoing) {
                 ++counts_apply;
            }
        }

        //a honor account apply project must less than seven at the same time
        auto global = _global_var.find(("maxapplynum"_n).value);
        uosio_assert(global != _global_var.end(), "cannot find maxapplynum field");
        uosio_assert(counts_apply < global->val, "apply too many projects");

        _project_list.emplace( proposer, [&] (auto& p) {
            p.project = project;
            p.owner = proposer;
            p.credit = 0;
            p.update_time = now();
            p.reason = reason;
            p.state = state_ongoing;
        });
    }

    ACTION honor::assess(const name bp, const name project, const uint64_t credit) {

        require_auth(bp); 
        
        auto pro = _project_list.find(project.value);
        uosio_assert(pro != _project_list.end(), "project not exist");

        //check state of project
        uosio_assert(pro->state == "ongoing", "state of project not ongoing");

        //check valid time of project
        auto global = _global_var.find(("overtime"_n).value);
        uosio_assert(global != _global_var.end(), "cannot find overtime field");
        uosio_assert(pro->update_time  + global->val >  now(), "assess a project of overtime");

        //check credit
        global = _global_var.find(("maxcredit"_n).value);
        uosio_assert(global != _global_var.end(), "cannot find maxcredit field");
        uosio_assert(credit <= global->val, "credit given too large");

        uosio_assert(check_authority(bp), "account not have authority to assess");

        struct bpcredit bpcrt{bp, credit};

        //check max apply number once time
        global = _global_var.find(("maxbpnumber"_n).value);
        uosio_assert(global != _global_var.end(), "cannot find maxbpnumber field");

        _project_list.modify( pro, _self, [&](auto& p) {
            if(p.credit_list.size() < global->val) {

                //bp not permit assessing repeatedly
                for(auto assessed : p.credit_list) {
                    if(assessed.bp == bp) {
                        uosio_assert(false, "current bp has been given its credit");
                    }
                }

                p.credit_list.emplace_back(bpcrt);
            }
            else {
                uosio_assert(false, "assessed bp  must less than max bp number");
            }
        });
    }

    bool honor::check_authority(const name bp) {
        std::vector<name> bplist;
        uosio_assert(getbplist(bplist), "get bplist falied");

        //check the bp name is in a list of authority
        bool is_find = false;

        for(auto p : bplist) {
            if(p == bp) {
                is_find = true;
                break;
            }
        }

        return is_find;
    }

    ACTION honor::rmproject(const name project) {
        require_auth(_self);

        auto pro = _project_list.find(project.value);
        uosio_assert(pro != _project_list.end(), "project not exist");

        _project_list.erase(pro);
    }

    ACTION honor::update() {
         
        auto current_time = now();
        auto overtime = _global_var.find(("overtime"_n).value);
        uosio_assert(overtime != _global_var.end(), "cannot find overtime field");

        //get current bp list
        std::vector<name> bplist;
        uosio_assert(getbplist(bplist), "get bplist falied");

        uint64_t credit,total_credit;

        //update the state of all project
        for(auto project = _project_list.begin(); project != _project_list.end(); ++project) {
            if(project->update_time + overtime->val < current_time && project->state == state_ongoing) {

                //check 2/3 dpos
                if(project->credit_list.size() >= bplist.size() * 2 / 3) {

                    //calc all credit of current project
                    total_credit = 0;
                    for(auto bpcredit : project->credit_list) {
                        total_credit += bpcredit.credit;
                    }

                    //check  honor account
                    auto honor = _honor_list.find( project->owner.value);
                    uosio_assert( honor != _honor_list.end(), "receiver not a honor account");
                    
                    auto credit = total_credit / bplist.size();

                    _honor_list.modify(honor, _self, [&](auto& h) {
                        h.credit += credit;
                    });

                    _project_list.modify(project, _self, [&](auto& p){
                        p.credit = credit;
                        p.state = state_finished;
                    }); 
                }
                else {
                    //no 2/3 dpos project, the state is setting to overtime
                    _project_list.modify( project, _self, [&](auto& p) {
                        if(p.update_time + overtime->val < current_time && p.state == state_ongoing) {
                            p.state = state_overtime;
                        }
                    });
                }
            }    
        }
     }

     ACTION honor::transfer(const name& from, const name& to, const asset& quantity, const std::string& memo) {

        symbol sym("UOS", 4);
        uosio_assert(quantity.symbol == sym, "transfer must system coin");
        uosio_assert(quantity.amount > 0, "asset must positive");

        if(to == _self) {
            //register honor account
            if (memo == "register") {
                auto register_pay = _global_var.find( ("registerpay"_n).value);
                uosio_assert(register_pay != _global_var.end(), "not find registerpay field");

                auto honor = _honor_list.find( from.value);

                uosio_assert( honor == _honor_list.end(), "account is a honor account already");
                uosio_assert(quantity.amount == register_pay->val, "error transfer of register honor ");

                _honor_list.emplace( _self, [&](auto& h) {
                    h.account = from;
                    h.vip = 1;
                    h.credit = 0;
                    h.last_convert = 0;
                });
            }

            auto register_pay = _global_var.find( ("registerpay"_n).value);
            auto fund = _global_var.find( ("fund"_n).value);
            
            if(register_pay != _global_var.end() &&  fund != _global_var.end()) {
                set_global_var("fund"_n, fund->val + quantity.amount);

                auto ctribut = _ctribut_list.find( from.value);
                if( quantity.amount > register_pay->val) {
                    if(ctribut  == _ctribut_list.end()) {
                        _ctribut_list.emplace( _self, [&](auto& c) {
                            c.contributor = from;
                            c.quantity = quantity;
                        });
                    }
                    else {
                        _ctribut_list.modify( ctribut, _self, [&](auto& c) {
                            c.quantity += quantity;
                        });
                    }    
                }
            }
        }
        else if(from == _self) {
            if(memo != "convertuos") {
                 uosio_assert(false, "invalid transfer.");
            } 
        }
     }
}

#define UOSIO_DISPATCH_HONOR( TYPE, MEMBERS ) \
extern "C" { \
   void apply( uint64_t receiver, uint64_t code, uint64_t action ) { \
      auto self = receiver; \
      if( code == self || code == ("uosio.token"_n).value) { \
      	 if( action == ("transfer"_n).value ){ \
           if(code != ("uosio.token"_n).value) \
                return ; \
            uosio::execute_action(uosio::name(receiver), uosio::name(code), &honorsys::honor::transfer); \
      	 } \
         switch( action ) { \
            UOSIO_DISPATCH_HELPER( TYPE, MEMBERS ) \
         } \
      } \
   } \
}

UOSIO_DISPATCH_HONOR(honorsys::honor, (init)(setdata)(reward) (upgrade)(convertuos) (apply) (assess) (rmproject)(update) )






















