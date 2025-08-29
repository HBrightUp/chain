#ifndef CONFIG_HPP

#define IP_MARK     	"."						//ip分隔符
#define CORE_SYMBOL 	symbol("UOS", 4)		//系统代币名
#define ZERO_ASSET  	asset(0, CORE_SYMBOL)	//0资产
#define PERIOD_SECONDS  7*24*3600				//一期的时长
#define MAX_KEY_VALUE   std::numeric_limits<uint64_t>::max()	//uint64_t最大值
//#define MAX_SUPPLY  	760000000				//基础奖励池最大金额
//#define BASE_AWARD		7600000					//节点数小于等于100时，节点基础奖励
//#define BASE_NODE_NUM	1000					//基础节点数量，100个
#define FIRST       	10000					//解锁第一级惩罚需要积分
#define SECOND      	20000  					//解锁第二级惩罚需要积分
#define THIRD			40000					//解锁第二级惩罚需要积分
#define SCORE			100000					//每期总分数
#define MULTIPLE		12						//至尊节点倍数,放大10倍存储
//#define BASE			100000					//将积分占比收益恢复到正常值
#define LOW				80						//1级惩罚扣除20%奖励
#define MID				50						//2级惩罚扣除50%奖励
#define UNION_ACCOUNT	"udfs.uos"_n			//联盟账户
#define REWARD_ACCOUNT	"udfsreceiver"_n		//奖励账户
#define PROFIT_ACCOUNT	"udfsconsumer"_n		//收益奖励账户


#endif //CONFIG_HPP

