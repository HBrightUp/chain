#include <gtest/gtest.h>
#include "config.h"
#include "database/db_kafka_produce.h"
#include <mysql/mysql.h>
#include <json.hpp>
TEST(DataInsertTest, insert)
{
    std::string filePath = "../conf/conf.json";
    configManager *cf = configManager::GetInstance();
    cf->init_config(filePath);

    // get data  from mysql
    MYSQL xzgsql_;
    if (mysql_init(&xzgsql_) == NULL)
    {
        return ;
    }
    if (!mysql_real_connect(&xzgsql_, "172.31.18.204", "jys",
                            "jys@Y582ALuwN", "hurong_web", 3306, NULL, 0))
    {
        std::string errorStr = mysql_error(&xzgsql_);
        return;
    }
    // select id , uid , coin , ops_type , utime , amount FROM sgp_otc.otc_eth_transaction;
    std::string sql = "select id , uid , coin , ops_type , utime , amount , business_id FROM sgp_otc.otc_eth_transaction";
    int ret = mysql_real_query(&xzgsql_, sql.c_str(), strlen(sql.c_str()));
    if (ret != 0)
    {
        std::cout<<"exec DB FAIL:" << sql << std::endl;
        return;
    }

    MYSQL_RES *result = mysql_store_result(&xzgsql_);
    size_t num_rows = mysql_num_rows(result);
    std::string logicID, logicName, fondID, fondName, accountName, businessID, businessName, amount;

    businessID = businessName = "OTC_ETH_TRANS";
    uint64_t accountID, updateTime, uuid;
    // kafka init
    DBKafkaProduce *cDBProduceBus = nullptr;
    cDBProduceBus = new DBKafkaProduce;
    cDBProduceBus->init_kafka(0, cf->m_sets.m_Kafka.c_str(), KA_TOPIC_BUSINESS);

    for (size_t i = 0; i < num_rows; ++i)
    {
        MYSQL_ROW row = mysql_fetch_row(result);
        uuid = std::stol(row[0]);
        accountID = std::stol(row[1]);
        fondID = fondName = accountName = row[2];
        int ops_type = std::stoi(row[3]);
        logicID = logicName = ops_type == 1 ? "outps" : "inps";
        updateTime = std::stol(row[4]);
        amount = row[5];

        nlohmann::json newData ;
        newData.push_back(uuid);
        newData.push_back(businessID);
        newData.push_back(businessName);
        newData.push_back(businessName);
        newData.push_back(logicID);
        newData.push_back(logicName);
        newData.push_back(fondID);
        newData.push_back(fondName);
        if(i%3 != 1)
            newData.push_back(accountID);
        newData.push_back(row[6]);
        newData.push_back("transName");
        newData.push_back(amount);
        newData.push_back(updateTime);
        // 将数据组织为json
        std::string data = newData.dump();
        std::cout<< "i:" << i << "---" << data<< std::endl;
        try
        {
            cDBProduceBus->push_data_to_kafka(data.c_str(), data.size());
        }
        catch (std::exception &e)
        {
            std::stringstream errorMsg;
            errorMsg << "push data to kafka failed!error msg is " << e.what();
            throw std::invalid_argument(errorMsg.str());
        }
        catch (...)
        {
            std::stringstream errorMsg;
            errorMsg << "push data to kafka failed!";
            throw std::invalid_argument(errorMsg.str());
        }
    }
    cDBProduceBus->destroy();
    mysql_free_result(result);

    return ;
}