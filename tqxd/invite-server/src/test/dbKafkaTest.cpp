#include <gtest/gtest.h>
#include "database/db_kafka_consume.h"
#include "database/db_kafka_produce.h"
#include "config.h"
#include <thread> 
void readKafkaData(std::string topic, std::string msg, void* args)
{
    std::cout<< msg << std::endl;
}
void thread_routine(void* ss){
    // DBKafkaConsume *ka = (DBKafkaConsume*)ss;
    // sleep(2);
    // DBKafkaProduce proKafa;
    // std::string brokers = "192.168.1.66:9092";
    // std::string topic = "ckAttribute";
    // std::string groupID = "0";
    // proKafa.init_kafka(0 , brokers, topic);
    // std::stringstream tmp;
    // for(int i = 0 ; i < 10 ; i++){
    //     tmp.str("");
    //     tmp << "{\"UUID\":1234,\"Attribute_Id\":1235,\"Attribute_Name\":\"att_Name\",\"Attribute_Type\":1,\"value_type\":2,\"value_data\":1234}";
    //     std::string str = tmp.str();
    //     proKafa.push_data_to_kafka(str);    
    // }
    // sleep(3);
    // ka->stop();
    // sleep(3);
}
TEST(dbKafkaTest, Init_Kafka)
{
    // configManager *cf = configManager::GetInstance();
    // std::string filePath = "../conf/conf.json";
    // cf->init_config(filePath);
    // DBKafkaConsume *dbK = new DBKafkaConsume;
    // std::thread threadProd(thread_routine, dbK);

    // std::string topic = "test";
    // std::string groupID = "0";
    // EXPECT_NO_THROW(dbK->initKafka(readKafkaData, dbK));
    // dbK->consume(1000);
    // delete dbK;
    // dbK = nullptr;
    // threadProd.join();
}
TEST(dbKafkaTest, Init_Kafka_Error)
{
    // std::string topic = "test";
    // std::string groupID = "0";
    // std::string brokers = "192.168.1.1:9092";
    // DBKafka dbK;
    // dbK.initKafka(brokers , topic , groupID , RdKafka::Topic::OFFSET_END ,readKafkaData);
}