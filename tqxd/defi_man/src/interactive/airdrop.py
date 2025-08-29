
#!/usr/bin/python3

import decimal
from logging import FATAL, setLogRecordFactory
from os import rename
import sys
from PyQt5.QtSql import QSqlQuery, QSqlTableModel
from PyQt5.QtWidgets import QApplication, QMainWindow, QDialog
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from web3.types import TxData
from CHAINAPI import ETHAPI, CheckAccount
from contract import erc20abi
from db import dbconnect
import xlwt
import ast
from common import dataformat,setting

class Airdrop(QMainWindow):
    db_account = None
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("airdrop")
        loadUi(ui_file, self)
         
        self.db_account = dbconnect.DBSqlite(setting.g_setting.getLocalDB())
        reg = QRegExp('^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)
        sql = 'CREATE TABLE airdrop (from_address char(200) NOT NULL, txid CHAR(256) NOT NULL, to_address char(200) NOT NULL, amount CHAR(256) NOT NULL, contract  CHAR(200));'
        self.db_account.createTable(sql)

        self.lineEdit_amount.setValidator(validator) 
        self.lineEdit_fee.setValidator(validator) 
     
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(200, 200, 200, 25)
        self.step = 0
        self.pbar.setRange(0,100)
        self.pbar.hide()
        self.timer = QBasicTimer()

        self.btn_std.clicked.connect(self.onAirdrop)
        self.btn_std_batch.clicked.connect(self.onAirdropBatch)
        self.btn_contract_batch.clicked.connect(self.onAirdropContractBatch)
        self.btn_verify.clicked.connect(self.onVerify)


    def txRecord(self, from_address, to, amount, txid, contract):
        dataformat.logger.info("%s send %s to %s, record %s, contract %s"%(from_address ,amount, to, txid, contract))
        sql = "INSERT INTO airdrop VALUES"
        sql_value = " ('%s','%s','%s','%s','%s')"%(from_address, txid, to, amount,contract)
        sql_exec = sql + sql_value
        self.db_account.insertData(sql_exec)      
  
  
    def onAirdrop(self):
        dataformat.logger.info("on airdrop")
        if self.checkAddress() == False:
            return 
        api = ETHAPI(setting.g_setting.getRpcUrl())
        amount = self.lineEdit_amount.text()
        from_address = self.lineEdit_address.text()
        privkey = self.lineEdit_privkey.text()
        to_address = self.lineEdit_recieve.text()
        fee = self.lineEdit_amount.text()
        txid = api.offlineSign(privkey, to_address, amount, fee)
        self.txRecord(from_address, to_address, amount, txid)

    def onAirdropBatch(self):
        dataformat.logger.info("on airdrop batch")
        if self.checkAddress() == False:
                return 
        file_type = 'excel files(*.xls)'
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            QMessageBox.information(self, '文件', '请选择文件！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        result = dataformat.GetExcelData(open_file)
        from_address = self.lineEdit_address.text()
        privkey = self.lineEdit_privkey.text()
        count = len(result)
        api = ETHAPI(setting.g_setting.getRpcUrl())
        nonce = api.getNonce(from_address)
        fee = self.lineEdit_fee.text()
        for i in range(count):
            to_address = result[i][0]
            amount = result[i][1]
           
            txid = api.offlineSignNonce(privkey, to_address, amount, fee, nonce + i)
            print(to_address)
            print(amount)
            print(txid)
            print('\n')
            self.txRecord(from_address, to_address, amount, txid, "sys")

    def onVerify(self):
        address = self.lineEdit_address.text()
        if address == "":
            sql = "SELECT * FROM airdrop;"
            self.showData(sql)
        else:
            sql = "SELECT * FROM airdrop WHERE  from_address='%s';"%(address)
            self.showData(sql)

    def onAirdropContractBatch(self):
        dataformat.logger.info("on contract airdrop batch")
        if self.checkAddress() == False:
                return 
        file_type = 'excel files(*.xls)'
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            QMessageBox.information(self, '文件', '请选择文件！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        result = dataformat.GetExcelData(open_file)
        from_address = self.lineEdit_address.text()
        privkey = self.lineEdit_privkey.text()
        count = len(result)
        api = ETHAPI(setting.g_setting.getRpcUrl())
        nonce = api.getNonce(from_address)
        fee = self.lineEdit_fee.text()
        path = setting.g_setting.getPath()
        token_abi = setting.g_setting.getToken()
        for i in range(count):
            to_address = result[i][0]
            amount = result[i][1]
            contract = result[i][2]
            decimal = result[i][3]
            print(token_abi["arc20"])
            abi_file = path+token_abi["arc20"]
            contract_abi  = json.load(open(abi_file))
            token =  erc20abi.ERC20(contract_abi,contract, setting.g_setting.getRpcUrl())
            if decimal != token.decimals():
                QMessageBox.information(self, '精度', '精度不一致停止转账！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
                return;

            txid = api.offlineSignERC20Nonce(privkey, to_address, contract, amount, fee, decimal, nonce + i)
            print(to_address)
            print(amount)
            print(contract)
            print(decimal)
            print(txid)
            print('\n')
            self.txRecord(from_address, to_address, amount, txid, contract)

    def checkAddress(self):
        from_address = self.lineEdit_address.text()
        from_privkey = self.lineEdit_privkey.text()
        ret = CheckAccount(from_address,from_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '发送地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False

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



