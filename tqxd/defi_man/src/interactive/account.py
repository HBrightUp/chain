#!/usr/bin/python3

from os import rename
import sys
from PyQt5.QtSql import QSqlQuery, QSqlTableModel
from PyQt5.QtWidgets import  QMainWindow
from PyQt5 import QtGui
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from CHAINAPI import ETHAPI, CheckAccount
from db import dbconnect
import xlwt
import ast
from common import dataformat,setting
from interactive import config,balance

class AddressImport(QMainWindow):
    db_account = None
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("addressimport")
        loadUi(ui_file, self)
        self.btn_import.clicked.connect(self.onImport)
    
    def setDB(self, db):
        self.db_account = db
    
    def checkInput(self):
        sys = self.lineEdit_sys.text()
        if sys == "":
            QMessageBox.warning(self, '输入', '系统不能为空！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False
        
        coin = self.lineEdit_coin.text()
        if coin == "":
            QMessageBox.warning(self, '输入', '币种不能为空！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False

        logic = self.lineEdit_logic.text()
        if logic == "":
            QMessageBox.warning(self, '输入', '业务不能为空！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False
        user = self.lineEdit_userc.text()
        if user == "":
            QMessageBox.warning(self, '输入', '用户不能为空！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False
        return True


    def onImport(self):
        address = self.lineEdit_address.text();
        privkey = self.lineEdit_privkey.text();
        ret = CheckAccount(address.strip(), privkey.strip())
        if ret == False:
            QMessageBox.warning(self, '导入', '地址和私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 
        ret = self.checkInput()
        if ret == False:
            return 

        sys = self.lineEdit_sys.text();
        user = self.lineEdit_user.text();
        coin = self.lineEdit_coin.text();
        logic = self.lineEdit_logic.text();

        sql = "INSERT INTO account VALUES"
        sql_value = " ('%s','%s','%s','%s','%s',\"%s\")"%(sys, coin, user, logic,address,privkey)
        sql_exec = sql + sql_value
        ret = self.db_account.insertData(sql_exec)
        if ret == True:
            QMessageBox.information(self, '导入', '导入成功！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        else :
            QMessageBox.information(self, '导入', '导入失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        

class Account(QMainWindow):
    db_account = None
    dlg_settting = None
    dlg_asset = None
    dlg_import = None
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("account")
        loadUi(ui_file, self)
        reg = QRegExp('[0-9]+$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)
        self.lineEdit_quantity.setValidator(validator) 
        self.btn_create.clicked.connect(self.onCreateAccount)
        self.btn_dump_json.clicked.connect(self.onDumpJson)
        self.btn_dump_excel.clicked.connect(self.onDumpExcel)
        self.btn_search.clicked.connect(self.onSearch)
        
        self.db_account = dbconnect.DBSqlite(setting.g_setting.getLocalDB())
        sql = 'CREATE TABLE account ( sys TEXT NOT NULL, coin char(50) NOT NULL, user CHAR(50) NOT NULL, logic TEXT NOT NULL, address CHAR(100) NOT NULL, privekey TEXT NOT NULL);'
        self.db_account.createTable(sql)
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(400, 200, 300, 25)
        self.step = 0
        self.pbar.setRange(0,100)
        self.pbar.hide()
        self.timer = QBasicTimer()
        self.dlg_settting = config.SettingDlg();
        self.dlg_import = AddressImport()
        self.dlg_asset = balance.Balance()
        
        self.actionimport.triggered.connect(self.onImport)
        self.actionopen_json.triggered.connect(self.onImportJson)
        self.actionopen_excel.triggered.connect(self.onImportExcel)
        self.actionasset.triggered.connect(self.onAsset)

    def onImport(self):
        dataformat.logger.info("on triggerd  import address")
        self.dlg_import.setDB(self.db_account)
        self.dlg_import.show()
        dataformat.logger.info("import finish")

    def onAsset(self):
        dataformat.logger.info("on triggerd asset")
        self.dlg_asset.show()
        
    def onSetting(self):
        dataformat.logger.info("on triggerd asset")
        self.dlg_settting.exec()
       
    def checkConfig(self, file_type):
        self.dlg_settting.exec()
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            print("Not select file !")
            return open_file,False
        
        str_version = self.dlg_settting.getVersion()
        version = 0
        if str_version == '':
            print("Not input version !")
            return open_file,False

        version = int(str_version)
        if version < 1 or version > 3:
            print("version error !")
            return open_file,False    
        return open_file,True

    def checkVersion(self, version):
        sys = self.dlg_settting.getSys()
        coin = self.dlg_settting.getCoin()
        logic = self.dlg_settting.getLogic()
        user = self.dlg_settting.getUser()
        if version == 1 and (sys == '' or coin == '' or logic == '' or user == ''):
            print("system coin logic user must input in version 1!")
            QMessageBox.information(self, '导入', '导入版本1的账户必须输入用户，业务，币种！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False
        return True


    def importDB(self, version, dict_accounts):
        sys = self.dlg_settting.getSys()
        coin = self.dlg_settting.getCoin()
        logic = self.dlg_settting.getLogic()
        user = self.dlg_settting.getUser()
        count = len(dict_accounts)
        sql = "INSERT INTO account VALUES"
        ret = False
        if version == 1 :
            for i in range(count):
                dict_data = dict_accounts[i]
                address = dict_data["address"]
                privkey = dict_data["privkey"]
                sql_value = " ('%s','%s','%s','%s','%s',\"%s\")"%(sys,coin,user,logic,address,privkey)
                sql_exec = sql + sql_value
                ret = self.db_account.insertData(sql_exec)
        
        if version == 3 :
            for i in range(count):
                dict_data = dict_accounts[i]
                sql_value = " ('%s','%s','%s','%s','%s',\"%s\")"%(dict_data[0],dict_data[1],dict_data[2],dict_data[3],dict_data[4],dict_data[5])
                sql_exec = sql + sql_value
                ret = self.db_account.insertData(sql_exec)
        
        if ret == True:
            QMessageBox.information(self, '导入', '导入成功！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        else :
            QMessageBox.information(self, '导入', '导入失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)

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
        self.importDB(version, json_accounts)


    def onImportExcel(self):
        file_type = 'excel files(*.xls)'
        execl_file,ret = self.checkConfig(file_type)
        if  ret == False :
            return 
        version = int(self.dlg_settting.getVersion())
        ret = self.checkVersion(version)
        if ret == False:
            return 
        result = dataformat.GetExcelData(execl_file)
        self.importDB(version, result)


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


    def onCreateAccount(self):
        dataformat.logger.info("on create acount")
        self.pbar.show()
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(100, self)

        sys_content = self.lineEdit_sys.text();    
        user_content = self.lineEdit_user.text();
        
        coin = self.lineEdit_coin.text();
        logic_content = self.lineEdit_logic.text();
        if (sys_content == '' or user_content == '' or coin == '' or logic_content == ''):
            dataformat.logger.warn("创建私钥基础信息不能为空")
            QMessageBox.information(self, '生成', '创建私钥基础信息不能为空！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)  
            self.pbar.hide()
            return 

        if (self.lineEdit_quantity.text() == '' ):
            dataformat.logger.warn("需要指定生成数量")
            QMessageBox.information(self, '生成', '请指定生成数量',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)  
            self.pbar.hide()
            return 

        quantity = int(self.lineEdit_quantity.text());
        api = ETHAPI(setting.g_setting.getRpcUrl())
        sql = "INSERT INTO account VALUES"
        step = 100 / (quantity +1)
        for index in range(quantity):
            ret = api.createAccount()
            dict_ret = json.loads(ret)
            dict_data = dict_ret['data']
            address = dict_data['address']
            privkey = dict_data['privateKey']           
            sql_value = " ('%s','%s','%s','%s','%s',\"%s\")"%(sys_content,coin,user_content,logic_content,address,privkey)
            sql_exec = sql + sql_value
            self.db_account.insertData(sql_exec)
            self.step = step*index
            self.pbar.setValue(step*index)
        sql_data = "select * from account where  sys like '%s' and user like '%s' and coin like '%s'and logic like '%s'"%(sys_content,user_content,coin,logic_content)
        self.showData(sql_data)
        self.step =100
        self.pbar.setValue(100)
        self.pbar.hide()
        QMessageBox.information(self, '生成', '账户生成完成！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)

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
        dataformat.logger.info("on dump json")
        file_type = 'json files(*.json)'
        sql = self.getSearchSQL();
        self.showData(sql)

        ret_result = self.db_account.getData(sql)
        filename = self.getDumpFileName("account.json")
        save_file = QFileDialog.getSaveFileName(self, '保存文件',filename, file_type)
        
        with open(save_file[0], "w") as f:
            json.dump(ret_result, f)
        QMessageBox.information(self, '导出', '导出完成',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        
    def onDumpExcel(self):
        dataformat.logger.info("dump excel")
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
        file_type = 'excel files(*.xls)'

        filename = self.getDumpFileName("account.xls")
        save_file = QFileDialog.getSaveFileName(self, '保存文件',filename, file_type)
        workbook.save(save_file[0])
        QMessageBox.information(self, '导出', '导出完成',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
   

