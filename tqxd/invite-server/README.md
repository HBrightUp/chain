# data_center

## BUILD

1. 安装依赖组件
```
sudo apt install cmake gcc g++ openssl libssl-dev libboost-all-dev libcurl4 librdkafka-dev libmysqlclient-dev libgoogle-glog-dev libevent-dev 
```

2. 安装clickhouse client
```
sudo apt-get install apt-transport-https ca-certificates dirmngr
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv E0C56BD4

echo "deb https://repo.clickhouse.tech/deb/stable/ main/" | sudo tee \
    /etc/apt/sources.list.d/clickhouse.list

sudo apt-get update

sudo apt-get install clickhouse-client
```

3. 安装c++的第三方库clickhouse-cpp
详情参见 https://github.com/ClickHouse/clickhouse-cpp

4. 编译
```
mkdir bin && cd bin
cmake .. && make
```

## config配置文件
默认配置文件路径 DATA_CENTER/config/conf.json
配置文件如下所示：
```
{
    "printtoconsole":0,                 --
    "logtimestamps":1,                  --
    "logtimemicros":100,                --
	"logpath":"./log",                  --日志路径
    "daemon":true,                      --
    "rpcallowip":"127.0.0.1",           --
    "rpcservertimeout":3600,            --rpc超时时间
    "rpcport":8332,                     --
    "rpcworkqueue":16,                  --工作队列长度
    "httpthread":4,                     --处理http的线程数
    "rpcuser":"user",                   --
    "rpcpassword":"a",                  --
    "kafka":"192.168.1.108:9092",       --kafka的ip:port
	"mysql":                            --mysql配置
	{
		"url":"192.168.1.108",
		"user":"root",
		"pass":"a",
		"db":"utilutxo",
		"port":3306
    },
    "clickhouse":                       --clickhouse配置
    {
        "url":"192.168.1.108",
		"user":"default",
		"pass":"abcd4321",
		"db":"datasets",
		"port":8123
    }
}
```

