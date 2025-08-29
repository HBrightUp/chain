#!/usr/bin/python3

from logging import Formatter
from os import rename
from re import S
import sys
from PyQt5.QtSql import QSqlQuery, QSqlTableModel
from PyQt5.QtWidgets import QApplication, QMainWindow, QDialog
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from eth_account import account
from eth_account.account import Account
from web3 import contract, eth
import web3
from CHAINAPI import ETHAPI, checkPwd 
from db import dbconnect
import xlwt
import ast
from  common import dataformat,setting
from  contract import erc20abi, swapabi


class Balance(QDialog):
    db_account = None
    filter_result = []
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("balance")
        loadUi(ui_file, self)
    
        self.btn_dump_json.clicked.connect(self.onDumpJson)
        self.btn_dump_excel.clicked.connect(self.onDumpExcel)
        self.btn_search.clicked.connect(self.onSearch)
        self.btn_add.clicked.connect(self.onAdd)
        self.btn_del.clicked.connect(self.onDel)
        self.db_account = dbconnect.DBSqlite(setting.g_setting.getLocalDB())
        #self.db_account = dbconnect.DBSqlite('../db/account.db')
        sql = 'CREATE TABLE coin (symbol char(50) NOT NULL, asset_name CHAR(256) NOT NULL, asset_decimal int not null, contract_address TEXT NOT NULL, total TEXT NOT NULL);'
        self.db_account.createTable(sql)
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(200, 200, 300, 25)
        self.step = 0
        self.pbar.setRange(0,100)
        self.pbar.hide()
        self.timer = QBasicTimer()
  
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

    def onAdd(self):
        dataformat.logger.info("on trigger add")
        contract_address = self.lineEdit_contract.text()
        json_token = setting.g_setting.getToken()
        #ui_file = setting.g_setting.getPath() + json_token["arc20"]
        json_file = setting.g_setting.getPath() + json_token["arc20"]    
        json_abi = json.load(open(json_file))
        token = erc20abi.ERC20(json_abi, contract_address, setting.g_setting.getRpcUrl())
        
        coin,symbol,decimal,total = token.info()
        if coin == "" :
             QMessageBox.information(self, '导入', '合约错误！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
             return
        sql = "INSERT INTO coin VALUES"
        sql_value = " ('%s','%s','%d','%s','%s')"%(symbol, coin, decimal, contract_address, total)
        sql_exec = sql + sql_value
        ret = self.db_account.insertData(sql_exec)
        if ret == True:
            QMessageBox.information(self, '导入', '导入代币成功！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        else :
            QMessageBox.information(self, '导入', '导入代币失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)

    def onDel(self):
        dataformat.logger.info("on trigger del")
        contract_address = self.lineEdit_contract.text()
        sql = "DELETE FROM coin WHERE contract_address = '%s' ;"%(contract_address)
        ret = self.db_account.delData(sql)
        if ret == True:
            QMessageBox.information(self, '删除', '删除代币成功！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        else :
            QMessageBox.information(self, '删除', '删除代币失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        

    def getBalance(self, address, contract):
        json_token = setting.g_setting.getToken()
        json_file = setting.g_setting.getPath() + json_token["arc20"] 
        #json_file = '../abi/arc20.json'    
        json_abi  = json.load(open(json_file))
        token = erc20abi.ERC20(json_abi, contract, setting.g_setting.getRpcUrl())
        amount = token.GetBalance(address)
        decimals = token.decimals()
        balance = dataformat.FromWei(amount,decimals)
        print(balance)
        return balance


    def onSearch(self):
        self.pbar.show()
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(100, self)
        sys = self.lineEdit_sys.text();
        user = self.lineEdit_user.text();
        coin = self.lineEdit_coin.text();
        logic = self.lineEdit_logic.text();
        url  = setting.g_setting.rpc_url
        api = ETHAPI(url)
        sys = "%" + sys +"%"
        user = "%" + user +"%"
        coin = "%" + coin +"%"
        logic = "%" + logic +"%"
        
        sql_data = "select distinct address from account where  sys like '%s' and user like '%s' and coin like '%s'and logic like '%s'"%(sys,user,coin,logic)
        accounts = self.db_account.getData(sql_data)
        account_count  = len(accounts)
        sql_data = "select symbol, contract_address from coin;"
        assets = self.db_account.getData(sql_data)
        asset_count = len(assets)
        step = 100 / (account_count +1)

        model = QtGui.QStandardItemModel()
        self.tableView.setModel(model)
        model.setHorizontalHeaderItem(0, QtGui.QStandardItem("地址"))
        model.setHorizontalHeaderItem(1, QtGui.QStandardItem("AITD"))
        model.setRowCount(asset_count)
        self.filter_result.clear()
        list_tmp = []
        list_tmp.append("address")
        list_tmp.append("AITD")
        for j in range(asset_count):
            coin_name = assets[j][0]
            print(coin_name)
            model.setHorizontalHeaderItem(j+2,QtGui.QStandardItem(coin_name))
            list_tmp.append(coin_name)
        
        #return stake1_usdt,stake1_aitd,stake_balance1,stake1_total,stake2_usdt,stake2_aitd,stake_balance2,stake2_total
        '''list_tmp.append("stake1_usdt")
        list_tmp.append("stake1_aitd")
        list_tmp.append("stake1_balance")
        list_tmp.append("stake1_total")'''
        list_tmp.append("stake2_usdt")
        list_tmp.append("stake2_aitd")
        list_tmp.append("stake2_balance")
        list_tmp.append("stake2_total")

        list_tmp.append("stake3_usdt")
        list_tmp.append("stake3_aitd")
        list_tmp.append("stake3_balance")
        list_tmp.append("stake3_total")
        list_tmp.append("stake2_rewards")
        list_tmp.append("stake3_rewards")
        '''model.setHorizontalHeaderItem(asset_count+2 + 1, QtGui.QStandardItem("stake1_usdt"))
        model.setHorizontalHeaderItem(asset_count+2 + 2 , QtGui.QStandardItem("stake1_aitd"))
        model.setHorizontalHeaderItem(asset_count+2 + 3, QtGui.QStandardItem("stake_balance1"))
        model.setHorizontalHeaderItem(asset_count+2 + 4, QtGui.QStandardItem("stake1_total"))
        '''
        model.setHorizontalHeaderItem(asset_count+2 + 0, QtGui.QStandardItem("stake2_usdt"))
        model.setHorizontalHeaderItem(asset_count+2 + 1, QtGui.QStandardItem("stake2_aitd"))
        model.setHorizontalHeaderItem(asset_count+2 + 2, QtGui.QStandardItem("stake_balance2"))
        model.setHorizontalHeaderItem(asset_count+2 + 3, QtGui.QStandardItem("stake2_total"))

        model.setHorizontalHeaderItem(asset_count+2 + 4, QtGui.QStandardItem("stake3_usdt"))
        model.setHorizontalHeaderItem(asset_count+2 + 5, QtGui.QStandardItem("stake3_aitd"))
        model.setHorizontalHeaderItem(asset_count+2 + 6, QtGui.QStandardItem("stake_balance3"))
        model.setHorizontalHeaderItem(asset_count+2 + 7, QtGui.QStandardItem("stake3_total"))

        model.setHorizontalHeaderItem(asset_count+2 + 8, QtGui.QStandardItem("stake2_rewards"))
        model.setHorizontalHeaderItem(asset_count+2 + 9, QtGui.QStandardItem("stake3_rewards"))

        self.filter_result.append(list_tmp)

        for index in range(account_count):
            address = accounts[index][0]
            model.setItem(index, 0, QtGui.QStandardItem(address))
            list_add =  []
            list_add.append(address)
            balance = api.getBalance(address.strip())
            list_add.append(balance)
            model.setItem(index, 1, QtGui.QStandardItem(str(balance)))
            for j in range(asset_count):
                contract_address = assets[j][1]
                amount = self.getBalance(address.strip(), contract_address.strip())
                list_add.append(amount)
                model.setItem(index, j+2, QtGui.QStandardItem(str(amount)))
            stake2_usdt,stake2_aitd,stake_balance2,stake2_total,stake3_usdt,stake3_aitd,stake_balance3,stake3_total,stake2_rewards,stake3_rewards =swapabi.GetLpTotal(address)
            '''model.setItem(index, asset_count+2 + 1, QtGui.QStandardItem(str(stake1_usdt)))
            model.setItem(index, asset_count+2 + 2 , QtGui.QStandardItem(str(stake1_aitd)))
            model.setItem(index, asset_count+2 + 3, QtGui.QStandardItem(str(stake_balance1)))
            model.setItem(index, asset_count+2 + 4, QtGui.QStandardItem(str(stake1_total)))'''
            model.setItem(index, asset_count+2 + 0, QtGui.QStandardItem(str(stake2_usdt)))
            model.setItem(index, asset_count+2 + 1, QtGui.QStandardItem(str(stake2_aitd)))
            model.setItem(index, asset_count+2 + 2, QtGui.QStandardItem(str(stake_balance2)))
            model.setItem(index, asset_count+2 + 3, QtGui.QStandardItem(str(stake2_total)))

            model.setItem(index, asset_count+2 + 4, QtGui.QStandardItem(str(stake3_usdt)))
            model.setItem(index, asset_count+2 + 5, QtGui.QStandardItem(str(stake3_aitd)))
            model.setItem(index, asset_count+2 + 6, QtGui.QStandardItem(str(stake_balance3)))
            model.setItem(index, asset_count+2 + 7, QtGui.QStandardItem(str(stake3_total)))
            
            model.setItem(index, asset_count+2 + 8, QtGui.QStandardItem(str(stake2_rewards)))
            model.setItem(index, asset_count+2 + 9, QtGui.QStandardItem(str(stake3_rewards)))
        
            list_add.append(stake2_usdt)
            list_add.append(stake2_aitd)
            list_add.append(stake_balance2)
            list_add.append(stake2_total)
        
            list_add.append(stake3_usdt)
            list_add.append(stake3_aitd)
            list_add.append(stake_balance3)
            list_add.append(stake3_total)

            list_add.append(stake2_rewards)
            list_add.append(stake3_rewards)
    
            self.filter_result.append(list_add)
            self.step = step*index
            self.pbar.setValue(step*index)
        self.tableView.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeToContents)
        #self.showData(sql_data)
        self.step =100
        self.pbar.setValue(100)
        self.pbar.hide()

    
    def getSearchSQL(self): 
        sys_content = self.lineEdit_sys_1.text();
        sys_content = '%'+ sys_content +'%'
        user_content = self.lineEdit_user_1.text();
        user_content = '%'+ user_content +'%'
        coin = self.lineEdit_coin_1.text();
        coin = '%'+ coin +'%'
        logic_content = self.lineEdit_logic_1.text();
        logic_content = '%'+ logic_content +'%'
        sql_value = "select * from account where  sys like '%s' and user like '%s' and coin like '%s'and logic like '%s'"%(sys_content,user_content,coin,logic_content)
        return sql_value
  

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
        ret_result = self.filter_result
       
        file_type = 'json files(*.json)'
        filename = self.getDumpFileName("balance.json")
        save_file = QFileDialog.getSaveFileName(self, '保存文件',filename, file_type)
        
        with open(save_file[0], "w") as f:
            json.dump(ret_result, f)
        
    def onDumpExcel(self):
        dataformat.logger.info("dump excel")
        ret_result = self.filter_result
        workbook = xlwt.Workbook(encoding = 'utf-8')
        worksheet = workbook.add_sheet("balance")
        row_size = len(ret_result)
        for i in range(row_size):
            row_value = ret_result[i]
            col_size = len(row_value)
            for j in range(col_size):
                worksheet.write(i,j, row_value[j])

        file_type = 'excel files(*.xls)'
        filename = self.getDumpFileName("balance.xls")
        save_file = QFileDialog.getSaveFileName(self, '保存文件',filename, file_type)
        workbook.save(save_file[0])



  