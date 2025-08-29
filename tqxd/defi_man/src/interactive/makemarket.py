
#!/usr/bin/python3

from os import SEEK_END, rename
import sys
from time import sleep
from PyQt5.QtSql import QSqlQuery, QSqlTableModel
from PyQt5.QtWidgets import QApplication, QMainWindow, QDialog
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from eth_account.account import Account
from CHAINAPI import ETHAPI, checkPwd 
from db import dbconnect
import xlwt
import ast
from common import dataformat,setting
from contract import swapabi
from interactive import config, balance
class MakeMarket(QMainWindow):
    db_account = None
    dlg_settting = None
    dlg_asset = None
    dlg_balance = None
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("makemarket")
        loadUi(ui_file, self)
        #loadUi("../ui/makemarket.ui", self)
        self.dlg_balance = balance.Balance()
        self.btn_reinvest.clicked.connect(self.onReinvest)
        self.btn_dump_json.clicked.connect(self.onDumpJson)
        self.btn_dump_excel.clicked.connect(self.onDumpExcel)
        self.btn_search.clicked.connect(self.onSearch)
        self.btn_release.clicked.connect(self.onRelease)
        self.btn_add.clicked.connect(self.onAdd)
        self.btn_add_reverse.clicked.connect(self.onAddReverse)
        self.db_account = dbconnect.DBSqlite(setting.g_setting.getLocalDB())
        #self.db_account = dbconnect.DBSqlite('../db/account.db')
        sql = 'CREATE TABLE account ( sys TEXT NOT NULL, coin char(50) NOT NULL, user CHAR(50) NOT NULL, logic TEXT NOT NULL, address CHAR(100) NOT NULL, privekey TEXT NOT NULL);'
        self.db_account.createTable(sql)
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(350, 150, 300, 30)
        self.step = 0
        self.pbar.setRange(0,100)
        self.pbar.hide()
        self.timer = QBasicTimer()
        self.dlg_settting = config.SettingDlg();
        self.lineEdit_url.setText(setting.g_setting.getRpcUrl())

        reg = QRegExp('^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)
        self.lineEdit_price.setValidator(validator)
        
        #self.actionimport.triggered.connect(self.onSetting)
        self.actionopen_json.triggered.connect(self.onImportJson)
        self.actionopen_excel.triggered.connect(self.onImportExcel)
        self.actionasset.triggered.connect(self.onAsset)

    def callbackTx(self,from_address, to, txid, status):
        dataformat.logger.info("address %s, to %s, txid %s"%(from_address, to, txid))

    def onAsset(self):
        dataformat.logger.info("on asset")
        self.dlg_balance.exec()

    def onAdd(self):
        dataformat.logger.info("on add")
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(100, self)

        self.onSearch()
        url = self.lineEdit_url.text()
        str_price = self.lineEdit_price.text()
        if str_price == "":
            dataformat.logger.warn("price not float!")
            self.pbar.hide()
            return

        price  = float(str_price)
        if price < 0:
            dataformat.logger.warn("price not setting!")
            self.pbar.hide()
            return 
        if url == "":
            dataformat.logger.warn("url not setting!")
            self.pbar.hide()
            return 
        accounts = self.getAccountInfo()
        counts = len(accounts)
 
        self.pbar.show()
        step = 100 / (counts +1)
        dataformat.logger.info("step %d"%(step))
        for index in range(counts):
            address = accounts[index][0]
            privkey = accounts[index][1]
            print("address:%s"%address)
            swapabi.AddPositions(address,privkey,url, price, self.callbackTx)
            self.step = step*index
            self.pbar.setValue(step*index)

        self.step =100
        self.pbar.setValue(100)
        self.pbar.hide()

    def onAddReverse(self):
        dataformat.logger.info("on add reverse")
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(100, self)

        self.onSearch()
        url = self.lineEdit_url.text()
        str_price = self.lineEdit_price.text()
        if str_price == "":
            dataformat.logger.warn("price not float!")
            self.pbar.hide()
            return

        price  = float(str_price)
        if price < 0:
            dataformat.logger.warn("price not setting!")
            self.pbar.hide()
            return 
        if url == "":
            dataformat.logger.warn("url not setting!")
            self.pbar.hide()
            return 
        accounts = self.getAccountInfo()
        counts = len(accounts)
 
        self.pbar.show()
        step = 100 / (counts +1)
        dataformat.logger.info("step %d"%(step))
        for index in range(counts):
            address = accounts[index][0]
            privkey = accounts[index][1]
            print("address:%s"%address)
            swapabi.AddPositionsReverse(address,privkey,url, price, self.callbackTx)
            self.step = step*index
            self.pbar.setValue(step*index)

        self.step =100
        self.pbar.setValue(100)
        self.pbar.hide()


    def onSetting(self):
        dataformat.logger.info("on setting")
        self.dlg_settting.exec()
        version = self.dlg_settting.getVersion()
        sys = self.dlg_settting.getSys()
        coin = self.dlg_settting.getCoin()
        logic = self.dlg_settting.getLogic()
        user = self.dlg_settting.getUser()
    
    def onRelease(self):
        dataformat.logger.info("on release")
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(100, self)

        self.onSearch()
        url = self.lineEdit_url.text()
        if url == "":
            dataformat.logger.warn("url not setting!")
            self.pbar.hide()
            return 
        accounts = self.getAccountInfo()
        counts = len(accounts)
 
        self.pbar.show()
        step = 100 / (counts +1)
        dataformat.logger.info("step %d"%(step))
        for index in range(counts):
            address = accounts[index][0]
            privkey = accounts[index][1]
            print("address:%s"%address)
            swapabi.Release(address,privkey,url, self.callbackTx)
            self.step = step*index
            self.pbar.setValue(step*index)

        self.step =100
        self.pbar.setValue(100)
        self.pbar.hide()


    def checkConfig(self, file_type):
        self.dlg_settting.exec()
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            dataformat.logger.warn("Not select file !")
            return open_file,False
        
        str_version = self.dlg_settting.getVersion()
        version = 0
        if str_version == '':
            dataformat.logger.warn("Not input version !")
            return open_file,False

        version = int(str_version)
        if version < 1 or version > 3:
            dataformat.logger.warn("version error !")
            return open_file,False    
        return open_file,True


    def checkVersion(self, version):
        sys = self.dlg_settting.getSys()
        coin = self.dlg_settting.getCoin()
        logic = self.dlg_settting.getLogic()
        user = self.dlg_settting.getUser()
        if version == 1 and (sys == '' or coin == '' or logic == '' or user == ''):
            dataformat.logger.warn("system coin logic user must input in version 1!")
            return False
        return True

    def importDB(self, version, dict_accounts, is_excel):
        sys = self.dlg_settting.getSys()
        coin = self.dlg_settting.getCoin()
        logic = self.dlg_settting.getLogic()
        user = self.dlg_settting.getUser()
        count = len(dict_accounts)
        sql = "INSERT INTO account VALUES"
        address = ""
        privkey = ""
        if version == 1 :
            for i in range(count):
                dict_data = dict_accounts[i]
                if is_excel == False:
                    address = dict_data["address"]
                    privkey = dict_data["privkey"]
                else :
                    address = dict_data[0]
                    privkey = dict_data[1]
                sql_value = " ('%s','%s','%s','%s','%s',\"%s\")"%(sys,coin,user,logic,address,privkey)
                sql_exec = sql + sql_value
                self.db_account.insertData(sql_exec)
        
        if version == 3 :
            for i in range(count):
                dict_data = dict_accounts[i]
                sql_value = " ('%s','%s','%s','%s','%s',\"%s\")"%(dict_data[0],dict_data[1],dict_data[2],dict_data[3],dict_data[4],dict_data[5])
                sql_exec = sql + sql_value
                self.db_account.insertData(sql_exec)


    def onImportJson(self):
        file_type = 'json files(*.json)'
        json_file,ret = self.checkConfig(file_type)
        if  ret == False :
            return 
        version = int(self.dlg_settting.getVersion())
        ret = self.checkVersion(version)
        if ret == False:
            return 
        
        json_accounts  = json.load(open(json_file))
        self.importDB(version, json_accounts, False)


    def onImportExcel(self):
        file_type = 'excel files(*.xls)'
        execl_file,ret = self.checkConfig(file_type)
        if  ret == False :
            return 
        version = int(self.dlg_settting.getVersion())
        ret = self.checkVersion(version)
        if ret == False:
            return 
        #json_data  = json.load(open(json_file))
        result = dataformat.GetExcelData(execl_file)
        self.importDB(version, result, True)


    def timerEvent(self, e):
        if self.step >= 100:
            self.step = 0
            self.pbar.setValue(self.step)
            self.timer.stop()
            return  
        self.step = self.step+1
        self.pbar.setValue(self.step)   


    def showData(self,sql):
        model = QSqlTableModel()
        query = QSqlQuery(sql, self.db_account.getDB())
        model.setEditStrategy(QSqlTableModel.OnFieldChange)
        model.setQuery(query)
        model.submitAll()
        self.tableView.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeToContents)#Stretch)
        self.tableView.setModel(model)
        self.tableView.show()
        model.select()

    def getAccountInfo(self):
        sys = self.lineEdit_sys.text();
        user = self.lineEdit_user.text();
        coin = self.lineEdit_coin.text();
        logic = self.lineEdit_logic.text();
  
        sys = '%'+ sys +'%'
        user = '%'+ user +'%'
        coin = '%'+ coin +'%'
        logic = '%'+ logic +'%'
        sql = "select address, privekey from account where  sys like '%s' and user like '%s' and coin like '%s'and logic like '%s'"%(sys,user,coin,logic)
       
        ret_result = self.db_account.getData(sql)
        return ret_result


    def onReinvest(self):
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(100, self)

        self.onSearch()
        url = self.lineEdit_url.text()
        str_price = self.lineEdit_price.text()
        if str_price == "":
            dataformat.logger.warn("price not float!")
            self.pbar.hide()
            return

        price  = float(str_price)
        if price < 0:
            dataformat.logger.warn("price not setting!")
            self.pbar.hide()
            return 
        if url == "":
            dataformat.logger.warn("url not setting!")
            self.pbar.hide()
            return 
        accounts = self.getAccountInfo()
        counts = len(accounts)
 
        self.pbar.show()
        step = 100 / (counts +1)
        dataformat.logger.info("step %d"%(step))
        for index in range(counts):
            address = accounts[index][0]
            privkey = accounts[index][1]
            print("address:%s"%address)
            swapabi.Reinvest(address,privkey,url, price, self.callbackTx)
            self.step = step*index
            self.pbar.setValue(step*index)

        self.step =100
        self.pbar.setValue(100)
        self.pbar.hide()
       

    def getSearchSQL(self):
        sys = self.lineEdit_sys.text();
        sys = '%'+ sys +'%'
        user = self.lineEdit_user.text();
        user = '%'+ user +'%'
        coin = self.lineEdit_coin.text();
        coin = '%'+ coin +'%'
        logic = self.lineEdit_logic.text();
        logic = '%'+ logic +'%'
        sql_value = "select * from account where  sys like '%s' and user like '%s' and coin like '%s'and logic like '%s'"%(sys,user,coin,logic)
        return sql_value
    
    def onSearch(self):
        sql = self.getSearchSQL();
        self.showData(sql)

    def getDumpFileName(self,suffix):
        sys = self.lineEdit_sys.text();
        user = self.lineEdit_user.text();
        coin = self.lineEdit_coin.text();
        logic = self.lineEdit_logic.text();
        ret_name = ''
        if sys != '':
            ret_name = sys + "_"
        if user != '':
            ret_name = ret_name + user + "_"
        if coin != '':
            ret_name = ret_name + coin + "_"
        if logic != '':
            ret_name = ret_name + logic + "_"
        
        ret_name = ret_name + suffix
        return ret_name
    
    def onDumpJson(self):
        print("dump json")
        sql = self.getSearchSQL();
        self.showData(sql)
        ret_result = self.db_account.getData(sql)
        filename = self.getDumpFileName("account.json")
        with open(filename, "w") as f:
            json.dump(ret_result, f)
        
    def onDumpExcel(self):
        print("dump excel")
        sql = self.getSearchSQL();
        self.showData(sql)
        ret_result = self.db_account.getData(sql)
        workbook = xlwt.Workbook(encoding = 'utf-8')
        worksheet = workbook.add_sheet("account")
        row_size = len(ret_result)
        for i in range(row_size):
            row_value = ret_result[i]
            col_size = len(row_value)
            for j in range(col_size):
                worksheet.write(i,j, row_value[j])
        filename = self.getDumpFileName("account.xls")
        workbook.save(filename)
        
