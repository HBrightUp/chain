import sys
from PyQt5.QtCore import *
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *
from PyQt5.QtSql import QSqlDatabase,QSqlQuery
import pymysql
from pymysql.err import NotSupportedError

class DBSqlite(object):
    db_file ='tmp.db'
    db_database = None
    def __init__(self, filename):
        self.db_file = filename
        if QSqlDatabase.contains('qt_sql_default_connection'):
            self.db_database = QSqlDatabase.database('qt_sql_default_connection')
        else :
            self.db_database = QSqlDatabase.addDatabase('QSQLITE')
        
        self.db_database.setDatabaseName(filename)
        if not self.db_database.open():
            QMessageBox.critical(None, ("无法打开数据库"),(filename),QMessageBox.Cancel)
            return
    def createTable(self, sql):
        query = QSqlQuery()
        query.exec_(sql)

    def getData(self, sql):
        query = QSqlQuery();
        ret_result = []
        if query.exec(sql):
            col_count = query.record().count();
            while query.next():
                raw_value = []
                for i in range(col_count):
                    raw_value.append(query.value(i))
                ret_result.append(raw_value)
        return ret_result

    def getDB(self):
        return self.db_database

    def insertData(self, sql):
         query = QSqlQuery()
         ret = query.exec_(sql)
         return ret

    def delData(self, sql):
        query = QSqlQuery()
        ret = query.exec_(sql)
        return ret


    def closeDB(self):
        self.db_database.close()


class DBMysql(object):
    host = "127.0.0.1"
    port = 3306
    user = "root"
    pwd = "123456"
    dbname = "test"
    conn = None
    def __init__(self, host, port, user, pwd, dbname):
        self.host = host
        self.port = port
        self.user = user
        self.pwd = pwd
        self.dbname = dbname

    def open(self):
        try:
            self.conn = pymysql.connect(host=self.host,port=self.port,user=self.user,passwd=self.pwd)
            self.conn.select_db(self.dbname)
        except pymysql.err.OperationalError:
            print("can not open db")
            return False
        else:
            return True


    def getData(self,sql):
        result = []
        if self.open() == False:
            print("Open db fail")
            return result
        cur =  self.conn.cursor()
        cur.execute(sql)
        result = cur.fetchall()      
        return result


        