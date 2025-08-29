#include "db_mysql.h"
#include <utility>
#include <iostream>
#include <config.h>
#include <util.h>
DBMysql::DBMysql()
{
}

DBMysql::~DBMysql()
{
}

bool DBMysql::openDB()
{
    configManager *cf = configManager::GetInstance();
    if (mysql_init(&mysql_) == NULL)
    {
        return false;
    }
    std::cout<< "host:" << cf->m_sets.m_Mysql.host << std::endl;
    if (!mysql_real_connect(&mysql_, cf->m_sets.m_Mysql.host.c_str(), cf->m_sets.m_Mysql.user.c_str(),
                cf->m_sets.m_Mysql.pass.c_str(), cf->m_sets.m_Mysql.dbName.c_str(),
                cf->m_sets.m_Mysql.port, NULL, 0))
    {
        std::string errorStr= mysql_error(&mysql_);
        return error("open DB FAIL: %s", errorStr);
    }

    return true;
}

bool DBMysql::refreshDB(const std::string& sql)
{

    int ret = mysql_real_query(&mysql_, sql.c_str(), strlen(sql.c_str()));
    if (ret != 0 && mysql_errno(&mysql_) != 1062)
    {
        return false;
    }

}

bool DBMysql::getData(const std::string& sql,  std::map<int, DataType> col_type ,json& json_data) 
{

    int ret = mysql_real_query(&mysql_, sql.c_str(), strlen(sql.c_str()));
    if (ret != 0)
    {
        return error("exec DB FAIL: %s", sql);
    }

    MYSQL_RES *result = mysql_store_result(&mysql_);
    size_t num_rows = mysql_num_rows(result);

    //LOG(INFO) << "data size: " << num_rows;
    int col_size = col_type.size();

    for (size_t i = 0; i < num_rows; ++i)
    {
        MYSQL_ROW row = mysql_fetch_row(result);
        json row_data = json::array();

        for (size_t j = 0; j < col_size; j++)
        {
            std::string sql_data = row[j];
            DataType type  = col_type[j];	
            if ( INT == type )
            {
                int real_data = std::stoi(sql_data);
                row_data.push_back(real_data);
            }
            else if ( DOUBLE == type )
            {
                double real_data = std::stof(sql_data);
                row_data.push_back(real_data);
            }
            else if ( STRING == type )
            {
                row_data.push_back(sql_data);
            }
        }
        json_data.push_back(row_data);
    }
    mysql_free_result(result);

    return true;
}

void DBMysql::batchRefreshDB(const std::vector<std::string>& vect_sql)
{
    mysql_query(&mysql_,"START TRANSACTION");
    for(uint32_t i = 0; i < vect_sql.size(); i ++)
    {
        refreshDB(vect_sql.at(i));
    }
    mysql_query(&mysql_,"COMMIT");
}

void DBMysql::closeDB()
{
    mysql_close(&mysql_);
}
