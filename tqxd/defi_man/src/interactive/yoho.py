
#!/usr/bin/python3

import decimal
import enum
from logging import FATAL, setLogRecordFactory
from os import rename
from re import I, S
import sys
from PyQt5.QtSql import QSqlQuery, QSqlTableModel
from PyQt5.QtWidgets import QApplication, QMainWindow, QDialog
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from eth_utils import address
from web3.types import TxData
from CHAINAPI import ETHAPI, CheckAccount
from contract import yohoabi
from db import dbconnect
import xlwt
import ast
from common import dataformat, setting
from  enum import Enum
import csv
import time


class TopMenu(Enum):
    FILE = 0
    SETTING = 1
    MAN = 2

class SecondMenu(Enum):
    RELATION = 0
    BATCH = 1
    MINT = 2
    OPENWHITELIST = 10
    CLOSEWHITELITE = 11
    TOKENFEERATE = 12
    REWARDSRATE = 13
    ADDWHITELIST = 14
    OPENHOLD = 15
    CLOSEHOLD = 16
    LOWERLIMIT =17
    TOKENMAN = 20
    LEVELMAN = 21
    REWARDMAN =22

class YohoContral(QMainWindow):
    top_menu = TopMenu.FILE
    second_menu = SecondMenu.RELATION
    token = None
    level = None
    stake_rewards = None
    invite_rewards = None

    def __init__(self):
        super().__init__()
        from PyQt5.uic import loadUi
        ui_file = setting.g_setting.getPath() + setting.g_setting.getUi("yoho")
        loadUi(ui_file, self)
        reg = QRegExp('^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)
   
        self.lineEdit_transfer_fee.setValidator(validator) 
        self.lineEdit_rewards_fee.setValidator(validator) 
     
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(200, 200, 200, 25)
        self.step = 0
        self.pbar.setRange(0,100)
        self.pbar.hide()
        self.timer = QBasicTimer()

        self.resetShow()
        self.actionimport_relation.triggered.connect(self.importRelation)
        self.actionbatch_airdrop.triggered.connect(self.batchAirdrop)
        self.actionimport_mint.triggered.connect(self.importMint)

        self.actionopen_whitelist.triggered.connect(self.openWhitelist)
        self.actionclose_whitelist.triggered.connect(self.closeWhitelist)
        self.actiontransfer_fee.triggered.connect(self.transferFee)
        self.actionadd_whitelist.triggered.connect(self.addWhitelist)
        self.action_change_admin_token.triggered.connect(self.changeAdminToken)
        self.actionchange_admin_level.triggered.connect(self.changeAdminLevel)
        self.actionchange_admin_rewards.triggered.connect(self.changeAdminRewards)

        self.btc_ok.clicked.connect(self.onOK)
        self.initContract()
        self.pbar = QProgressBar(self)
        self.pbar.setGeometry(200, 200, 300, 25)
        self.step = 0
        self.pbar.setRange(0,1000)
        self.pbar.hide()
        self.timer = QBasicTimer()
  
    def timerEvent(self, e):
        if self.step >= 1000:
            self.step = 0
            self.pbar.setValue(self.step)
            self.timer.stop()
            return  
        self.step = self.step+1
        self.pbar.setValue(self.step)   

    def initContract(self):
        contracts = setting.g_setting.getContract()
        json_yoho = setting.g_setting.getYoho()
        path = setting.g_setting.getPath()

        token_address = contracts["yoho"]["Token"]
        stake_rewards_address = contracts["yoho"]["StakeRewards"]
        level_address = contracts["yoho"]["Level"]
        invite_rewards_address = contracts["yoho"]["InviteRewards"]
        
        token_abi = json.load(open(path+json_yoho["token"]))
        stake_rewards_abi = json.load(open(path+json_yoho["stake-rewards"]))
        level_abi = json.load(open(path+json_yoho["level"]))
        invite_rewards_abi = json.load(open(path+json_yoho["invite-rewards"]))
 
        self.token = yohoabi.Yoho(token_abi, token_address, setting.g_setting.getRpcUrl())
        self.level = yohoabi.Level(level_abi, level_address, setting.g_setting.getRpcUrl())
        self.stake_rewards = yohoabi.StakeRewards(stake_rewards_abi, stake_rewards_address, setting.g_setting.getRpcUrl())
        self.invite_rewards = yohoabi.InviteRewards(invite_rewards_abi, invite_rewards_address, setting.g_setting.getRpcUrl())

    def onOK(self):
        if self.checkAddress() == False:
            return
        privkey = self.lineEdit_privkey_man.text()
        # 
        txid = ""
        if self.second_menu == SecondMenu.TOKENFEERATE:
            rate = self.lineEdit_transfer_fee.text()
            txid = self.token.setBasisPointsRate(privkey, rate)
        elif self.second_menu == SecondMenu.REWARDSRATE:
            rate = self.lineEdit_rewards_fee.text()
            txid = self.stake_rewards.setLevelRewardRate(privkey, rate)
        elif self.second_menu == SecondMenu.ADDWHITELIST:
            address = self.lineEdit_address_whitelist.text()
            txid = self.token.addWhiteList(privkey,address)
        elif self.second_menu == SecondMenu.LOWERLIMIT:
            rate = self.lineEdit_rewards_fee.text()
            txid = self.stake_rewards.setLevelRewardRate(privkey, rate)
        elif self.second_menu == SecondMenu.TOKENMAN:
            address = self.lineEdit_address_man_new.text()
            txid = self.token.transferOwnership(privkey, address)
        elif self.second_menu == SecondMenu.LEVELMAN:
            address = self.lineEdit_address_man_new.text()
            txid = self.level.transferOwnership(privkey, address)
        elif self.second_menu == SecondMenu.REWARDMAN:
            address = self.lineEdit_address_man_new.text()
            txid = self.stake_rewards.updateAdmin(privkey, address)
        QMessageBox.information(self, '转账', txid,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
        
    def changeAdminRewards(self):
        self.top_menu = TopMenu.MAN
        self.second_menu = SecondMenu.REWARDMAN
        self.resetShow()
     
    def changeAdminLevel(self):
        self.top_menu = TopMenu.MAN
        self.second_menu = SecondMenu.LEVELMAN
        self.resetShow()
      
    def changeAdminToken(self):
        self.top_menu = TopMenu.MAN
        self.second_menu = SecondMenu.TOKENMAN
        self.resetShow()
      
    def addWhitelist(self):
        self.top_menu = TopMenu.SETTING
        self.second_menu = SecondMenu.ADDWHITELIST
        self.resetShow()
        
    def transferFee(self):
        self.top_menu = TopMenu.SETTING
        self.second_menu = SecondMenu.TOKENFEERATE
        self.resetShow()
      

    def openWhitelist(self):
        self.top_menu = TopMenu.SETTING
        self.second_menu = SecondMenu.OPENWHITELIST
        self.resetShow()
        dataformat.logger.info("on open whitelist")
        if self.checkAddress() == False:
            return
        privkey = self.lineEdit_privkey_man.text() 
        self.token.resetLimiteTransfer(privkey,True)

    def closeWhitelist(self):
        self.top_menu = TopMenu.SETTING
        self.second_menu = SecondMenu.CLOSEWHITELITE
        self.resetShow()
        dataformat.logger.info("on close whitelist")
        if self.checkAddress() == False:
            return
        privkey = self.lineEdit_privkey_man.text() 
        self.token.resetLimiteTransfer(privkey, False)
    
    def batchAirdrop(self):
        self.top_menu = TopMenu.FILE
        self.second_menu = SecondMenu.BATCH
        self.resetShow()
        dataformat.logger.info("on batch airdrop")
        if self.checkAddress() == False:
            return 
        file_type = 'csv files(*.csv)'
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            QMessageBox.information(self, '文件', '请选择文件！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        result = dataformat.GetCsvData(open_file)
        privkey = self.lineEdit_privkey_man.text()
        self.pbar.show()
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(1000, self)
        count = len(result)
        step = 1000 / (count +1)

        for i in range(count):
            to = result[i][0]
            amount = result[i][1] 
            amount_wei = dataformat.ToWei(amount,18)
            txid = self.token.transfer(privkey,to, amount_wei)
            print(txid)
            print(i)
            time.sleep(0.01)
            self.step = step*i
            self.pbar.setValue(step*i)
        self.step =1000
        self.pbar.setValue(1000)
        self.pbar.hide()

        
    def importMint(self):
        self.top_menu = TopMenu.FILE
        self.second_menu = SecondMenu.RELATION
        self.resetShow()
        dataformat.logger.info("on inmport relation")
        if self.checkAddress() == False:
                return 
    
        file_type = 'csv files(*.csv)'
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            QMessageBox.information(self, '文件', '请选择文件！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        result = dataformat.GetCsvData(open_file)
        privkey = self.lineEdit_privkey_man.text()
  
        self.pbar.show()
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(1000, self)
        count = len(result)
        step = 1000 / (count +1)
        
        for i in range(count):
            lower = result[i][0] 
            upper = self.level.getUpper(lower)
            if upper == "0x0000000000000000000000000000000000000000":
                print("lower has not upper : ", lower)
                continue
                
            amount = self.token.GetBalance(lower)
            if amount == 0:
                print("has not money : ", lower)
                continue 
                
            
            stake_upper = self.invite_rewards.lowerToUpperStake(lower)
            if stake_upper == "0x0000000000000000000000000000000000000000":
                print(stake_upper) 
                txid = self.invite_rewards.importStakeRelation(privkey, lower)
                print(txid)
            print(i)

            #txid = self.invite_rewards.importStakeRelation(privkey, lower)
            #print(txid)
            #print(i)
            #time.sleep(0.005)
            
            #print(org_upper)
            self.step = step*i
            self.pbar.setValue(step*i)
        self.step =1000
        self.pbar.setValue(1000)
        self.pbar.hide()

    def importRelation(self):
        self.top_menu = TopMenu.FILE
        self.second_menu = SecondMenu.MINT
        self.resetShow()
        dataformat.logger.info("on inmport mint")
        if self.checkAddress() == False:
                return 
    
        file_type = 'csv files(*.csv)'
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            QMessageBox.information(self, '文件', '请选择文件！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        result = dataformat.GetCsvData(open_file)
        privkey = self.lineEdit_privkey_man.text()
  
        self.pbar.show()
        if self.timer.isActive():
            self.timer.stop()
        else:
            self.timer.start(1000, self)
        count = len(result)
        step = 1000 / (count +1)
        
        for i in range(count):
            upper = result[i][0]
            lower = result[i][1] 
            org_upper = self.level.getUpper(lower)
            if org_upper == "0x0000000000000000000000000000000000000000":
                txid = "0x12345678910123212452123124512345541"#self.level.importRelation(privkey,upper, lower)
                print(txid)
            print(i)
            #time.sleep(0.005)
            
            #print(org_upper)
            self.step = step*i
            self.pbar.setValue(step*i)
        self.step =1000
        self.pbar.setValue(1000)
        self.pbar.hide()


    def resetShow(self):
        self.lineEdit_transfer_fee.setDisabled(True)
        self.lineEdit_rewards_fee.setDisabled(True)
        self.lineEdit_address_whitelist.setDisabled(True)
        self.lineEdit_lower_limit.setDisabled(True)
        self.lineEdit_rewards_fee.setDisabled(True)
        self.lineEdit_address_man_new.setDisabled(True)
        self.lineEdit_privkey_man_new.setDisabled(True)
        
        if self.second_menu == SecondMenu.TOKENFEERATE:
            self.lineEdit_transfer_fee.setDisabled(False)
        elif self.second_menu == SecondMenu.REWARDSRATE:
            self.lineEdit_rewards_fee.setDisabled(False)
        elif self.second_menu == SecondMenu.ADDWHITELIST:
            self.lineEdit_address_whitelist.setDisabled(False)
        elif self.second_menu == SecondMenu.LOWERLIMIT:
            self.lineEdit_lower_limit.setDisabled(False)
        elif self.second_menu == SecondMenu.LOWERLIMIT:
            self.lineEdit_rewards_fee.setDisabled(False)
        elif self.second_menu == SecondMenu.TOKENMAN or self.second_menu == SecondMenu.LEVELMAN or self.second_menu == SecondMenu.REWARDMAN:
            self.lineEdit_address_man_new.setDisabled(False)
            self.lineEdit_privkey_man_new.setDisabled(False)
      
  
    def checkAddress(self):
        man_address = self.lineEdit_address_man.text()
        man_privkey = self.lineEdit_privkey_man.text()
        
        ret = CheckAccount(man_address,man_privkey)
        if ret == False:
            QMessageBox.information(self, '地址', '管理地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False
        if self.top_menu == TopMenu.FILE and self.second_menu == SecondMenu.RELATION:
            ret = self.level.checkAdmin(man_address)
        if ret == False:
            QMessageBox.information(self, '地址', '管理员地址布不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return False

        if self.top_menu == TopMenu.MAN:
            man_address_new = self.lineEdit_address_man_new.text()
            man_privkey_new = self.lineEdit_privkey_man_new.text()
            ret = CheckAccount(man_address_new,man_privkey_new)
            if ret == False:
                QMessageBox.information(self, '地址', '新管理地址于私钥不匹配！',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
                return False
        return ret

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



