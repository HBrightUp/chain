#!/usr/bin/python3

from os import rename
from re import S, search
import sys

from eth_utils import address
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
from  contract import erc20abi
from  common import ckclient, setting, dataformat

class Record(QDialog):
    filter_result = []
    search_type =None
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("record")
        loadUi(ui_file, self)
        self.btn_dump_json.clicked.connect(self.onDumpJson)
        self.btn_dump_excel.clicked.connect(self.onDumpExcel)
        
        self.btn_search_address.clicked.connect(self.onSearchAddress)
        self.btn_search_contract.clicked.connect(self.onSearchContract)
        self.btn_balance.clicked.connect(self.onBalance)
        self.btn_balance_contract.clicked.connect(self.onBalanceContract)
        self.search_type = 1
        self.tableView.doubleClicked.connect(self.onSearch)
    
    
    def onSearch(self, index):
        dataformat.logger.info("on search address")
        row = index.row()
        col = index.column()
        ck_data = self.tableView.model().item(row, col).text()
        headers = []
        
        if self.search_type == 1:
            headers.append("交易")
            headers.append("发送者")
            headers.append("接收者")
            headers.append("交易金额")
            headers.append("时间")
            sql = "SELECT tx_hash,`from`, `to`, amount,  toDateTime(`timestamp`) from block_transaction WHERE "
            if col ==0 :
                sql = sql + "tx_hash='%s'"%(ck_data)
            if col ==1 :
                sql = sql + "`from`='%s'"%(ck_data)
            if col == 2:
                sql = sql + "`to`='%s'"%(ck_data)
            self.showData(sql, headers)
   
    def checkConfig(self, file_type):
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            print("Not select file !")
            return open_file,False
        
        return open_file,True

    def onBalance(self):
        file_type = 'excel files(*.xls)'
        execl_file,ret = self.checkConfig(file_type)
        if  ret == False :
            return 

        #json_data  = json.load(open(json_file))
        result = dataformat.GetExcelData(execl_file)
        cols = len(result)
        sql_aitd = "SELECT transaction_hash ,`from`, `to`, amount,  toDateTime(`timestamp`),block_number from block_number_transaction WHERE "
        sql_usdt = "SELECT transaction_hash ,param0 as `from`,    param1 as `to`, param2 as `amount`, toDateTime(blocktime), block_number from event_log WHERE "
        model = QtGui.QStandardItemModel()
        
        #self.tableView.setModel(model)
        iter = 1
        headers = []
        headers.append("交易")
        headers.append("发送者")
        headers.append("接收者")
        headers.append("交易金额")
        headers.append("时间")
        headers.append("块高度")
        model.setRowCount(10000)
        model.setHorizontalHeaderLabels(headers)
        self.filter_result.clear()
        rpc = ETHAPI(setting.g_setting.getRpcUrl())
        contract_address = "0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
        for i in range(cols):
            send = result[i][0]
            print(send)
            aitd = result[i][1]
            ustd = result[i][2]
            recieve = result[i][3]
      
            sql = sql_aitd + "`from` = LOWER('%s') and block_number >2936769 and toUInt64(amount) > 0 ORDER BY block_number ;"%(send)
            list_result = ckclient.GetCkData(sql)
            row_count = len(list_result)
            col_count = len(headers)
            
            for k in range(row_count):
                raw_value = list_result[k]
                filter_value = []
                for j in range(col_count):
                    if j == 3 :
                        print(raw_value[0])
                        transaction = rpc.getTransaction(str(raw_value[0]))
                        balance = transaction["value"]
                        amount = dataformat.FromWei(balance, 18)
                        model.setItem(iter, j, QtGui.QStandardItem(str(amount)))
                        filter_value.append(str(amount))
                    else :
                        model.setItem(iter, j, QtGui.QStandardItem(str(raw_value[j])))
                        filter_value.append(str(raw_value[j]))
                iter = iter + 1
                self.filter_result.append(filter_value)

            sql = sql_usdt + "address = LOWER('%s') AND (`param0` = LOWER('%s') OR param1 =LOWER('%s')) AND event_name = 'Transfer' AND block_number >2936769 ORDER BY  block_number;"%(contract_address, send,send) 
            list_result = ckclient.GetCkData(sql)
            row_count = len(list_result)
            col_count = len(headers)
            
            for k in range(row_count):
                raw_value = list_result[k]
                filter_value = []
                for j in range(col_count):
                    if j == 3 :
                        balance = int(raw_value[j])
                        amount = dataformat.FromWei(balance, 18)
                        model.setItem(iter, j, QtGui.QStandardItem(str(amount)))
                        filter_value.append(str(amount))
                    else :
                        model.setItem(iter, j, QtGui.QStandardItem(str(raw_value[j])))
                        filter_value.append(str(raw_value[j]))
                iter = iter + 1
                self.filter_result.append(filter_value)
        self.tableView.setModel(model)
        self.tableView.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeToContents)



    def onBalanceContract(self):
        file_type = 'excel files(*.xls)'
        execl_file,ret = self.checkConfig(file_type)
        if  ret == False :
            return 

        #json_data  = json.load(open(json_file))
        result = dataformat.GetExcelData(execl_file)
        cols = len(result)
        contract_address = self.lineEdit_contract_balance.text()
        sql_usdt = "SELECT transaction_hash ,param0 as `from`,    param1 as `to`, param2 as `amount`, toDateTime(blocktime), block_number from event_log WHERE "
        model = QtGui.QStandardItemModel()
        #self.tableView.setModel(model)
        iter = 0
        headers = []
        headers.append("交易")
        headers.append("发送者")
        headers.append("接收者")
        headers.append("交易金额")
        headers.append("时间")
        headers.append("块高度")
        model.setRowCount(10000)
        self.filter_result.clear()
        for i in range(cols):
            send = result[i][0]
            print(send)
            aitd = result[i][1]
            ustd = result[i][2]
            recieve = result[i][3]

            sql = sql_usdt + "address = LOWER('%s') AND (`param0` = LOWER('%s') OR param1 =LOWER('%s')) AND event_name = 'Transfer' AND block_number >2936769 ORDER BY  block_number;"%(contract_address, send,send) 
            list_result = ckclient.GetCkData(sql)
            row_count = len(list_result)
            col_count = len(headers)
            
            for k in range(row_count):
                raw_value = list_result[k]
                filter_value = []
                for j in range(col_count):
                    if j == 3 :
                        balance = int(raw_value[j])
                        amount = dataformat.FromWei(balance, 6)
                        model.setItem(iter, j, QtGui.QStandardItem(str(amount)))
                        filter_value.append(str(amount))
                    else :
                        model.setItem(iter, j, QtGui.QStandardItem(str(raw_value[j])))
                        filter_value.append(str(raw_value[j]))
                iter = iter + 1
                self.filter_result.append(filter_value)
        
        self.tableView.setModel(model)
        self.tableView.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeToContents)

    def showData(self,sql, headers):
        self.filter_result.clear()
        list_result = ckclient.GetCkData(sql)
        
        model = QtGui.QStandardItemModel()
        self.tableView.setModel(model)
        model.setHorizontalHeaderLabels(headers)
        row_count = len(list_result)
        col_count = len(headers)
        model.setRowCount(row_count)
        for i in range(row_count):
            raw_value = list_result[i]
            filter_value = []
            for j in range(col_count):
                model.setItem(i, j, QtGui.QStandardItem(str(raw_value[j])))
                filter_value.append(str(raw_value[j]))
            self.filter_result.append(filter_value)
        
        self.tableView.setModel(model)
        self.tableView.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeToContents)

    def onSearchAddress(self):
        dataformat.logger.info("on search address")
        address = self.lineEdit_address.text();
        sql = "SELECT tx_hash,`from`, `to`, amount,  toDateTime(`timestamp`) from block_transaction WHERE  `from` = LOWER('%s') OR `to` = LOWER('%s');"%(address,address)
        headers = []
        headers.append("交易")
        headers.append("发送者")
        headers.append("接收者")
        headers.append("交易金额")
        headers.append("时间")
        self.showData(sql,headers)

    def onSearchContract(self):
        dataformat.logger.info("on search contract")
        address = self.lineEdit_address.text();
        contract = self.lineEdit_contract.text();
        sql = "SELECT tx_hash,`from`, `to`, amount,  toDateTime(`timestamp`) from block_transaction WHERE  `from` = LOWER('%s') OR `to` = LOWER('%s');"%(address,address)
        headers = []
        headers.append("交易")
        headers.append("发送者")
        headers.append("接收者")
        headers.append("交易金额")
        headers.append("时间")
        self.showData(sql,headers)

    def getDumpFileName(self,suffix):
        address = self.lineEdit_address.text();
        contract =self.lineEdit_contract.text()
        ret_name = ''
        if address != '':
            ret_name = address + "_"
        if contract != '':
            ret_name = ret_name + contract + "_"
        ret_name = ret_name + suffix
        return ret_name
    
    def onDumpJson(self):
        dataformat.logger.info("on dump json")
        ret_result = self.filter_result
       
        file_type = 'json files(*.json)'
        filename = self.getDumpFileName("record.json")
        save_file = QFileDialog.getSaveFileName(self, '保存文件',filename, file_type)
        
        with open(save_file[0], "w") as f:
            json.dump(ret_result, f)
        
    def onDumpExcel(self):
        dataformat.logger.info("dump excel")
        ret_result = self.filter_result
        workbook = xlwt.Workbook(encoding = 'utf-8')
        worksheet = workbook.add_sheet("record")
        row_size = len(ret_result)
        for i in range(row_size):
            row_value = ret_result[i]
            col_size = len(row_value)
            for j in range(col_size):
                worksheet.write(i,j, row_value[j])

        file_type = 'excel files(*.xls)'
        filename = self.getDumpFileName("record.xls")
        save_file = QFileDialog.getSaveFileName(self, '保存文件',filename, file_type)
        workbook.save(save_file[0])



  