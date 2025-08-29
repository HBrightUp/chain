#include <gtest/gtest.h>
#include "messagequeue/kafka_dataprocess.h"
#include "config.h"
#include "db_kafkaproduce.h"
#include "json.hpp"
#include <thread>
#include "database/db_clickhouse.h"
void initDB(std::string sql1 , std::string sql2){
    g_db_clickhouse->openDB();
    g_db_clickhouse->execute(sql1);
    g_db_clickhouse->execute(sql2);
}
void getData1(int i , std::string& s){
    s = "";
    nlohmann::json js;
    js["UUID"] = i;
    js["Attribute_Id"] = i+1;
    js["Attribute_Name"] = "xzgtest";
    js["Attribute_type"] = i;
    js["value_type"] = 4;
    js["value_data"] = i;
    s = js.dump();
}

void thread_routine1(void* ss){
    ProducerKafka* pKa = (ProducerKafka*)ss;
    std::string topic = "ckAttribute";
    std::string groupID = "0";
    configManager *cf = configManager::GetInstance();
    pKa->init_kafka(0, const_cast<char*>(cf->m_sets.m_Kafka.c_str()), const_cast<char*>(topic.c_str()));
    std::string tmp;
    for(int i = 0 ; i < 1000000 ; i++){
        getData1(i , tmp);
        pKa->push_data_to_kafka(tmp.c_str() , tmp.size());    
    }
}
void getData2(int i , std::string& s){
    s = "";
    nlohmann::json js;
    js["UUID"] = i;
    js["Action_ID"] = i+1;
    js["Action_Name"] = "xzgtest";
    js["Action_Input_Type"] = i;
    js["Action_Output_Type"] = i;
    js["value_type"] = 4;
    js["value_data"] = i;
    s = js.dump();
    if(i ==19)
        std::cout<< s << std::endl;
}
void thread_routine2(void* ss){
    ProducerKafka* pKa = (ProducerKafka*)ss;
    std::string topic = "ckAction";
    std::string groupID = "0";
    configManager *cf = configManager::GetInstance();
    pKa->init_kafka(0, const_cast<char*>(cf->m_sets.m_Kafka.c_str()), const_cast<char*>(topic.c_str()));
    std::string tmp;
    for(int i = 0 ; i < 1000000 ; i++){
        getData2(i , tmp);
        pKa->push_data_to_kafka(tmp.c_str() , tmp.size());    
    }
}

void getData3(int i , std::string& s){
    s = "";
    nlohmann::json js;
    js["UUID"] = i;
    js["Business_ID"] = i+1;
    js["Business_Name"] = "b_n";
    js["Logic_ID"] = i+2;
    js["Logic_Name"] = "l_n";
    js["Fund_ID"] = i+3;
    js["Fund_Name"] = "f_n";
    js["Account_ID"] = i+4;
    js["Account_Name"] = "a_n";
    js["Transaction_ID"] = i+5;
    js["Transaction_Name"] = "t_n";
    js["value_type"] = 4;
    js["value_data"] = i;
    s = js.dump();
}
void thread_routine3(void* ss){
    ProducerKafka* pKa = (ProducerKafka*)ss;
    std::string topic = "ckBusiness";
    std::string groupID = "0";
    configManager *cf = configManager::GetInstance();
    pKa->init_kafka(0, const_cast<char*>(cf->m_sets.m_Kafka.c_str()), const_cast<char*>(topic.c_str()));
    std::string tmp;
    for(int i = 0 ; i < 1000000 ; i++){
        getData3(i , tmp);
        pKa->push_data_to_kafka(tmp.c_str() , tmp.size());    
    }
}

TEST(kaDataProcess , init_Attribute)
{
    // init config
    std::string filePath = "../conf/conf.json";
    configManager *cf = configManager::GetInstance();
    cf->init_config(filePath);
    // std::string sql1 = "drop table IF EXISTS datasets.user_attribute";
    // std::string sql2 = "CREATE TABLE IF NOT EXISTS datasets.user_attribute(`UUID` 				UInt64,`Attribute_Id` 		UInt64,`Attribute_Name` 	String,`Attribute_type` 	UInt8,`Value_bit` 		Nullable(Int8),`Value_ubit` 		Nullable(UInt8),`Value_int64`	 	Nullable(Int64),`Value_uint64` 		Nullable(UInt64),`Value_double` 		Nullable(Float64),`Value_string` 		Nullable(String),`Value_date` 		Nullable(DateTime))ENGINE = MergeTree ORDER BY UUID SETTINGS index_granularity = 8192";
    // initDB(sql1 , sql2);
    // KafkaDataProcess kaProcess;
    // ProducerKafka pKa;
    // std::thread threadProd(thread_routine1, &pKa);
    //kaProcess.init("ckAttribute");
}

TEST(kaDataProcess , init_Action)
{
    // configManager *cf = configManager::GetInstance();
    // std::string sql1 = "drop table IF EXISTS datasets.user_action";
    // std::string sql2 = "CREATE TABLE IF NOT EXISTS datasets.user_action ( `UUID` UInt64, `Action_ID` UInt64, `Action_Name` String, `Action_Input_Type` UInt8, `Action_Output_Type` UInt8, `Value_bit` Nullable(Int8), `Value_ubit` Nullable(UInt8), `Value_int64` Nullable(Int64), `Value_uint64` Nullable(UInt64), `Value_double` Nullable(Float64), `Value_string` Nullable(String), `Value_date` Nullable(DateTime) ) ENGINE = MergeTree ORDER BY UUID SETTINGS index_granularity = 8192";
    // initDB(sql1 , sql2);
    // KafkaDataProcess kaProcess;
    // ProducerKafka pKa;
    // std::thread threadProd(thread_routine2, &pKa);
    // kaProcess.init("ckAction");
    // threadProd.join();
}

TEST(kaDataProcess , init_Business)
{
    // configManager *cf = configManager::GetInstance();
    // std::string sql1 = "drop table IF EXISTS datasets.business";
    // std::string sql2 = "CREATE TABLE IF NOT EXISTS datasets.business ( `UUID` UInt64, `Business_ID` UInt64, `Business_Name` String, `Logic_ID` UInt64, `Logic_Name` String, `Fund_ID` UInt64, `Fund_Name` String, `Account_ID` UInt64, `Account_Name` String, `Transaction_ID` UInt64, `Transaction_Name` String, `Value_bit` Nullable(Int8), `Value_ubit` Nullable(UInt8), `Value_int64` Nullable(Int64), `Value_uint64` Nullable(UInt64), `Value_double` Nullable(Float64), `Value_string` Nullable(String), `Value_date` Nullable(DateTime) ) ENGINE = MergeTree ORDER BY UUID SETTINGS index_granularity = 8192";
    // initDB(sql1 , sql2);
    // KafkaDataProcess kaProcess;
    // ProducerKafka pKa;
    // std::thread threadProd(thread_routine3, &pKa);
    // kaProcess.init("ckBusiness");
    // threadProd.join();
}



TEST(kaDataProcess , Init_KafkaProduce){
    // configManager *cf = configManager::GetInstance();
    // KafkaDataProcess kaProcess1;
    // kaProcess1.m_topic = "ckAttribute";

    // KafkaDataProcess kaProcess2;
    // kaProcess2.m_topic = "ck_Action";

    // KafkaDataProcess kaProcess3;
    // kaProcess3.m_topic = "ck_Business";

    // std::thread t1(thread_routine1 , &kaProcess1);

    // std::thread t2(thread_routine2 , &kaProcess2);

    // std::thread t3(thread_routine3 , &kaProcess3);

    // t1.join();
    // t2.join();
    // t3.join();



    // nlohmann::json js = "[1600135406.28479, 3, \"BTC\", \"deposit\", 1.23450000]"_json;
    // for(auto &el:js)
    //     std::cout<< el<< std::endl;
    // std::cout<< js[0] << "---" << js[1] << std::endl;


}

