#include <gtest/gtest.h>
#include "database/db_clickhouse.h"
#include "config.h"
void InitConfig()
{
    std::string filePath = "../conf/conf.json";
    configManager *cf = configManager::GetInstance();
    cf->init_config(filePath);
    const std::string sql1 = "Drop TABLE IF EXISTS datasets.A";
    const std::string sql2 = "Drop TABLE IF EXISTS datasets.B";
    const std::string sql3 = "Drop TABLE IF EXISTS datasets.C";
    g_db_clickhouse->openDB();
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql1));
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql2));
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql3));
}
TEST(dbClickHouse, openDB)
{
    InitConfig();
    EXPECT_NO_THROW(g_db_clickhouse->openDB());
}

TEST(dbClickHouse, openDB_Throw)
{
    configManager *cf = configManager::GetInstance();
    std::string host = cf->m_sets.m_Clickhouse.host;
    cf->m_sets.m_Clickhouse.host = "12.3.5.5";
    //EXPECT_FALSE(g_db_clickhouse->openDB());
    cf->m_sets.m_Clickhouse.host = host;
}

TEST(dbClickHouse, execute_createTable)
{
    const std::string sql1 = "CREATE TABLE IF NOT EXISTS datasets.A(`A` INT , `B` INT)ENGINE = TinyLog()";
    const std::string sql2 = "CREATE TABLE IF NOT EXISTS datasets.B(`HEIGHT` UInt64)ENGINE = Log";
    const std::string sql3 = "CREATE TABLE IF NOT EXISTS datasets.C(`TXID` String , `AMOUNT` UInt32 , `ADDRESS` String)ENGINE = MergeTree() PRIMARY KEY(`TXID`) ORDER BY (`TXID`, `AMOUNT`) SETTINGS index_granularity = 8192";
    g_db_clickhouse->openDB();
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql1));
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql2));
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql3));
}
TEST(dbClickHouse, execute_Throw)
{
    const std::string sql1 = "INSERT INTO datasets.blocks (height) VALUES (11)";
    EXPECT_ANY_THROW(g_db_clickhouse->execute(sql1));
}

TEST(dbClickHouse, insertDB)
{
    auto aColumn = std::make_shared<ColumnInt32>();
    auto bColumn = std::make_shared<ColumnInt32>();
    auto cColumn = std::make_shared<ColumnUInt64>();
    for (int i = 0; i < 10000; i++)
    {
        aColumn->Append(i);
        bColumn->Append(i + 10001);
        cColumn->Append(i + 1.234F);
    }
    Block bk;
    bk.AppendColumn("A", aColumn);
    bk.AppendColumn("B", bColumn);
    EXPECT_NO_THROW(g_db_clickhouse->insertDB("datasets.A" , bk));
    Block bk1;
    bk1.AppendColumn("HEIGHT", cColumn);
    EXPECT_NO_THROW(g_db_clickhouse->insertDB("datasets.B", bk1));
}
TEST(dbClickHouse, insertDB_Throw)
{
    auto aColumn = std::make_shared<ColumnInt64>();
    auto bColumn = std::make_shared<ColumnInt64>();
    for (int i = 10001; i < 20000; i++)
    {
        aColumn->Append(i);
        if (i < 15000)
            bColumn->Append(i + 20001);
    }
    Block bk;
    bk.AppendColumn("A", aColumn);
    // C++ exception : all columns in block must have same count of rows
    EXPECT_THROW(bk.AppendColumn("B", bColumn), std::exception);
    // DB::Exception : Cannot convert: Int64 to Int32
    EXPECT_THROW(g_db_clickhouse->insertDB("datasets.A", bk), clickhouse::ServerException);
}
TEST(dbClickHouse, execute_Select)
{
    std::string strQuery = "SELECT A,B FROM datasets.A";
    Block bk;
    g_db_clickhouse->openDB();
    g_db_clickhouse->getData(strQuery , bk);
    int colCount = bk.GetColumnCount();
    int rowCount = bk.GetRowCount();
    std::cout<< "colCount:"<< colCount << "rowCount:"<< rowCount<< std::endl;    
    EXPECT_TRUE(colCount > 0);
    EXPECT_TRUE(rowCount > 0);
    EXPECT_EQ(bk[0]->As<ColumnInt32>()->At(167) , 167);
    EXPECT_EQ(bk[1]->As<ColumnInt32>()->At(167) , 167+10001);
}
TEST(dbClickHouse, execute_Select_Throw)
{
    std::string strQuery = "SELECT A,B,C FROM datasets.A";
    Block bk;
    g_db_clickhouse->openDB();
    EXPECT_ANY_THROW(g_db_clickhouse->getData(strQuery , bk));
}
TEST(dbClickHouse, execute_DropTable)
{
    const std::string sql1 = "Drop TABLE IF EXISTS datasets.A";
    const std::string sql2 = "Drop TABLE IF EXISTS datasets.B";
    const std::string sql3 = "Drop TABLE IF EXISTS datasets.C";
    g_db_clickhouse->openDB();
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql1));
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql2));
    EXPECT_NO_THROW(g_db_clickhouse->execute(sql3));
    delete g_db_clickhouse;
}