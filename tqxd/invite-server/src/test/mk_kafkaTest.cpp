#include <gtest/gtest.h>
#include "config.h"
#include "messagequeue/mk_dataprocess.h"
#include <mysql/mysql.h>
#include <json.hpp>
#include <vector>
#include <string>
TEST(mk_kafkaTest, insert)
{
    std::string str_order = "{\"event\": 2, \"order\": {\"id\": 6512, \"market\": \"AITDUSDT\", \"source\": \"3\", \"type\": 1, \"side\": 2, \"user\": 675572, \"ctime\": 1611208328.3611839, \"mtime\": 1611208351.885426, \"price\": \"1e-8\", \"amount\": \"1500000000\", \"taker_fee\": \"0\", \"maker_fee\": \"0\", \"left\": \"1082310241.75\", \"deal_stock\": \"417689758.25\", \"deal_money\": \"4.1768975825\", \"deal_fee\": \"0e-12\"}, \"stock\": \"AITD\", \"money\": \"USDT\"}";
    std::string str_deal = "[1611202403.7112429, \"AITDUSDT\", 5110, 5160, 755577, 1460415, \"1e-8\", \"69031320.91000000\", \"0.06903132091000000000\", \"6903132.091000000000\", 2, 244, \"AITD\", \"USDT\"]";
    std::vector<std::string> vecStr , vecStr1;
    std::string orderTopic = "orders" ;
    std::string dealTopic = "deals" ;
    mkDataProcess::msgDispose(orderTopic , str_order , vecStr);

    mkDataProcess::msgDispose(dealTopic , str_deal , vecStr1);
    for (auto i : vecStr)
        std::cout<< i << std::endl;
        for (auto i : vecStr1)
        std::cout<< i << std::endl;
}
