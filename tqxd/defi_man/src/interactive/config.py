#!/usr/bin/python3

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
import xlrd 
from CHAINAPI import ETHAPI, checkPwd
from  db import dbconnect
import xlwt
import ast
from common import dataformat, setting
from contract import erc20abi

class RpcDlg(QDialog):
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("rpcdialog")
        loadUi(ui_file, self)
        #loadUi("../ui/rpcdialog.ui", self)    
        self.lineEdit_AITD.setText(setting.g_setting.getRpcUrl()) 
        self.buttonBox.accepted.connect(self.onAccept)
        self.buttonBox.rejected.connect(self.onReject)
        
    def onAccept(self):
        self.hide()
        setting.g_setting.setRpcUrl(self.lineEdit_AITD.text())  

    def onReject(self):
        self.hide()

class SettingDlg(QDialog):
    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("configdialog")
        loadUi(ui_file, self)
        #loadUi("../ui/configdialog.ui", self)
        reg = QRegExp('[0-9]+$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)
        self.lineEdit_version.setValidator(validator) 
        self.buttonBox.accepted.connect(self.onAccept)
        self.buttonBox.rejected.connect(self.onReject)
        
    def onAccept(self):
        self.hide()

    def onReject(self):
        self.hide()

    def getSys(self):
        sys_content = self.lineEdit_sys.text()
        return sys_content

    def getLogic(self):
        logic = self.lineEdit_logic.text()
        return logic
    
    def getUser(self):
        user = self.lineEdit_user.text()
        return user
    
    def getCoin(self):
        coin = self.lineEdit_coin.text()
        return coin

    def getVersion(self):
        version = self.lineEdit_version.text()
        return version