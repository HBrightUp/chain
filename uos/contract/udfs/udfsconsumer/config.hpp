#ifndef CONFIG_HPP

//#define MAX_KEY_VALUE       std::numeric_limits<uint64_t>::max()	//uint64_t最大值
#define CORE_SYMBOL 		symbol("UOS", 4)				//系统代币名
#define ZERO_ASSET  		asset(0, CORE_SYMBOL)			//0资产
#define SUPER_STORAGE   	2*1024*1024*1024*1024			//2T
#define SUPREME_STORAGE 	5*SUPER_STORAGE					//10T
#define UNION				"udfs.uos"_n					//联盟账户
#define PRODUCER_ACCOUNT	"udfsproducer"_n				//生产账户


#endif //CONFIG_HPP
