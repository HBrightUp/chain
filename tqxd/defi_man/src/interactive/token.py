
#!/usr/bin/python3

from logging import setLogRecordFactory
from os import rename
import sys
from PyQt5.QtSql import QSqlQuery, QSqlTableModel
from PyQt5.QtWidgets import QApplication, QMainWindow, QDialog
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from eth_account.account import Account
from web3.types import TxData
from CHAINAPI import ETHAPI, checkPwd, CheckAccount
from db import dbconnect
import xlwt
import ast
from common import dataformat,setting
from contract import erc20abi
from interactive import config

class Token(QMainWindow):
    db_account = None
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("token")
        loadUi(ui_file, self)
        #loadUi("../ui/token.ui", self)
        self.db_account = dbconnect.DBSqlite(setting.g_setting.getLocalDB())
        #self.db_account = dbconnect.DBSqlite('../db/account.db')
        reg = QRegExp('^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)

        self.lineEdit_amount.setValidator(validator) 
        self.lineEdit_fee.setValidator(validator) 
     
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(200, 200, 200, 25)
        self.step = 0
        self.pbar.setRange(0,100)
        self.pbar.hide()
        self.timer = QBasicTimer()

        self.actionopen.triggered.connect(self.onOpen)
        self.actionclose.triggered.connect(self.onClose)
        self.actionpandect.triggered.connect(self.onPandect)
        self.actioninfo.triggered.connect(self.onInfo)
        self.actionturnover.triggered.connect(self.onTurnover)
        self.actionapprove.triggered.connect(self.onApprove)
        self.actiontransfer.triggered.connect(self.onTransfer)
        self.actionburn.triggered.connect(self.onBurn)
        self.actionmint.triggered.connect(self.onMint)
    
    def getToken(self):
        contract_address = self.lineEdit_contract.text()
        if contract_address == "":
            dataformat.logger.warning("error contranct address")
            return None, False
        path = setting.g_setting.getPath()
        token_abi = setting.g_setting.getToken()
        abi_file = path + token_abi["arc20"]#'../abi/arc20.json'    
        contract_abi  = json.load(open(abi_file))
        token = erc20abi.ERC20(contract_abi, contract_address, setting.g_setting.getRpcUrl())
        return token,True
        
  
    def onOpen(self):
        dataformat.logger.info("trigger open")
  
    def onClose(self):
        dataformat.logger.info("trigger close")
    
    def onPandect(self):
        dataformat.logger.info("trigger pandect")
        sql = "SELECT * FROM coin"
        self.showData(sql)

    def onInfo(self):
        dataformat.logger.info("trigger info")
        contract_address = self.lineEdit_contract.text()
        sql = "SELECT * FROM coin WHERE contract_address = \'%s\'"%(contract_address)
        self.showData(sql)

    def checkAddress(self):
        from_address = self.lineEdit_src_address.text()
        from_privkey = self.lineEdit_src_privkey.text()
        ret = CheckAccount(from_address,from_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '源地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False

        to_address = self.lineEdit_dst_address.text()
        to_privkey = self.lineEdit_dst_privkey.text()
        ret = CheckAccount(to_address,to_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '源地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False
        return True

    def onTurnover(self):
        dataformat.logger.info("trigger trun over")
        token, ret = self.getToken()
        if ret == False:
            QMessageBox.information(self, '合约', '合约地址失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        ret = self.checkAddress()
        if ret == False:
            dataformat.logger.error("address input error")
            return

        from_privkey = self.lineEdit_src_privkey.text()
        to_address = self.lineEdit_dst_address.text()
        txid = token.transferOwnership(from_privkey,to_address)
        dataformat.logger.info(txid)
        QMessageBox.information(self, '权限转移完成', txid,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)

  
    def onApprove(self):
        dataformat.logger.info("trigger approve")
        from_address = self.lineEdit_src_address.text()
        from_privkey = self.lineEdit_src_privkey.text()

        ret = CheckAccount(from_address,from_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '源地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        
        token, ret = self.getToken()
        if ret == False:
            QMessageBox.information(self, '合约', '合约地址失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        #from_privkey = self.lineEdit_src_privkey.text()
        to_address = self.lineEdit_dst_address.text()
        txid = token.Approve(from_privkey, to_address, -1)
        dataformat.logger.info(txid)


    def onTransfer(self):
        dataformat.logger.info("trigger transfer")
        token, ret = self.getToken()
        if ret == False:
            QMessageBox.information(self, '合约', '合约地址失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        ret = self.checkAddress()
        if ret == False:
            dataformat.logger.error("address input error")
            return

        from_privkey = self.lineEdit_src_privkey.text()
        to_address = self.lineEdit_dst_address.text()
        balance = self.lineEdit_amount.text()
        amount = dataformat.ToWei(balance, token.decimals());
        txid = token.Transfer(from_privkey, to_address, amount)
        dataformat.logger.info(txid)
        QMessageBox.information(self, '转账完成', txid,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        

    def onBurn(self):
        dataformat.logger.info("trigger burn")
        from_address = self.lineEdit_src_address.text()
        from_privkey = self.lineEdit_src_privkey.text()
        print(from_address)
        print(from_privkey)
        ret = CheckAccount(from_address,from_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '源地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        
        token, ret = self.getToken()
        if ret == False:
            QMessageBox.information(self, '合约', '合约地址失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        amount = self.lineEdit_amount.text()
        balance = self.lineEdit_amount.text()
        amount = dataformat.ToWei(balance, token.decimals());
        txid = token.burn(from_privkey, amount)
        dataformat.logger.info(txid)
        QMessageBox.information(self, '销毁完成', txid,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)

    def onMint(self):
        dataformat.logger.info("trigger mint")
        from_address = self.lineEdit_src_address.text()
        from_privkey = self.lineEdit_src_privkey.text()
        ret = CheckAccount(from_address,from_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '源地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        token, ret = self.getToken()
        if ret == False:
            QMessageBox.information(self, '合约', '合约地址失败！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        balance = self.lineEdit_amount.text()
        amount = dataformat.ToWei(balance, token.decimals());
        txid = token.mint(from_privkey, amount)
        dataformat.logger.info(txid)
        QMessageBox.information(self, '增发完成', txid,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        

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



