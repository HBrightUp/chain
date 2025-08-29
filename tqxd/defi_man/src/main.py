#!/usr/bin/python3

from logging import log
import sys
from PyQt5.QtWidgets import QApplication, QMainWindow
from PyQt5 import QtGui
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from interactive import account, coinview, token, config, record, makemarket, airdrop, yoho
from common import dataformat, setting
from CHAINAPI import ETHAPI, CheckAccount
from pathlib import Path
#app = QApplication(sys.argv)
class Manager(QMainWindow):
    account = None
    coin_view = None
    makemarket = None
    token = None
    airdrop = None
    rpcdlg = None
    record = None
    init = False
    yoho = None
    def __init__(self):
        super().__init__()
        '''from PyQt5.uic import loadUi
        loadUi("../ui/mainwindow.ui", self)'''
        self.resize(800,600)
        self.setWindowTitle('钱包管理')
        
        reg = QRegExp('^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$')
        validator = QtGui.QRegExpValidator(self)
        validator.setRegExp(reg)

        #module        
        menubar = self.menuBar()
        menubar.setNativeMenuBar(False)
        fileMenu = menubar.addMenu('File')
        icon = QtGui.QIcon()
        fileMenu.setTitle("文件")
        #file 
        actionopen = QAction(icon, 'open', self)
        actionopen.setText('打开')
        actionopen.triggered.connect(self.onOpen)
        fileMenu.addAction(actionopen)

        actionclose = QAction(icon, 'close', self)
        actionclose.setText('关闭')
        actionclose.triggered.connect(self.onClose)
        fileMenu.addAction(actionclose)

        #manager
        managerMenu = menubar.addMenu('Manager')
        managerMenu.setTitle("管理")
        
        #self.actiontoken.triggered.connect(self.onToken)
        actiontoken = QAction(icon, 'token', self)
        actiontoken.setText('代币')
        actiontoken.triggered.connect(self.onToken)
        managerMenu.addAction(actiontoken)

        #self.actionbridge.triggered.connect(self.onBridge)
        actionbridge = QAction(icon, 'bridge', self)
        actionbridge.setText('桥')
        actionbridge.triggered.connect(self.onBridge)
        managerMenu.addAction(actionbridge)
        
        #self.actionairdrop.triggered.connect(self.onAirdrop)
        actionairdrop = QAction(icon, 'airdrop', self)
        actionairdrop.setText('空投')
        actionairdrop.triggered.connect(self.onAirdrop)
        managerMenu.addAction(actionairdrop)

        #self.actionquantization.triggered.connect(self.onQuantization)
        actionquantization = QAction(icon, 'quantization', self)
        actionquantization.setText('量化')
        actionquantization.triggered.connect(self.onQuantization)
        managerMenu.addAction(actionquantization)

        #self.actionaccount.triggered.connect(self.onAccount)
        actionaccount = QAction(icon, 'account', self)
        actionaccount.setText('账户')
        actionaccount.triggered.connect(self.onAccount)
        managerMenu.addAction(actionaccount)

        #self.actionborrow_lend.triggered.connect(self.onBorrowLend)
        actionborrow_lend = QAction(icon, 'borrow_lend', self)
        actionborrow_lend.setText('借贷')
        actionborrow_lend.triggered.connect(self.onBorrowLend)
        managerMenu.addAction(actionborrow_lend)

        #self.actionborrow_lend.triggered.connect(self.onBorrowLend)
        actionyoho = QAction(icon, 'yoho', self)
        actionyoho.setText('YOHO')
        actionyoho.triggered.connect(self.onYoho)
        managerMenu.addAction(actionyoho)
     
        #view
        managerView = menubar.addMenu('View')
        managerView.setTitle("查询")

        #self.actionasset.triggered.connect(self.onAsset)
        actionasset = QAction(icon, 'asset', self)
        actionasset.setText('资金')
        actionasset.triggered.connect(self.onAsset)
        managerView.addAction(actionasset)
        #self.actionrecord.triggered.connect(self.onRecord)
        actionrecord = QAction(icon, 'record', self)
        actionrecord.setText('记录')
        actionrecord.triggered.connect(self.onRecord)
        managerView.addAction(actionrecord)

        #operation
        lbl_src_addr = QLabel('label', self)
        lbl_src_addr.setText("源地址:")
        lbl_src_addr.move(15, 35)
        self.lineEdit_src_address = QLineEdit(self)
        self.lineEdit_src_address.move(60, 35)
        self.lineEdit_src_address.setFixedWidth(300)
        self.lineEdit_src_address.setFixedHeight(25)

        lbl_src_priv = QLabel('label', self)
        lbl_src_priv.setText("源私钥:")
        lbl_src_priv.move(15, 75)
        self.lineEdit_src_privkey = QLineEdit(self)
        self.lineEdit_src_privkey.move(60, 75)
        self.lineEdit_src_privkey.setFixedWidth(300)
        self.lineEdit_src_privkey.setFixedHeight(25)

        lbl_dst_addr = QLabel('label', self)
        lbl_dst_addr.setText("目标地址:")
        lbl_dst_addr.move(400, 35)
        self.lineEdit_dst_address = QLineEdit(self)
        self.lineEdit_dst_address.move(460, 35)
        self.lineEdit_dst_address.setFixedWidth(300)
        self.lineEdit_dst_address.setFixedHeight(25)

        lbl_dst_priv = QLabel('label', self)
        lbl_dst_priv.setText("目标私钥:")
        lbl_dst_priv.move(400, 75)
        self.lineEdit_dst_privkey = QLineEdit(self)
        self.lineEdit_dst_privkey.move(460, 75)
        self.lineEdit_dst_privkey.setFixedWidth(300)
        self.lineEdit_dst_privkey.setFixedHeight(25)

        lbl_amount_addr = QLabel('label', self)
        lbl_amount_addr.setText("金额:")
        lbl_amount_addr.move(15, 110)
        self.lineEdit_amount = QLineEdit(self)
        self.lineEdit_amount.move(50, 110)
        self.lineEdit_amount.setFixedWidth(150)
        self.lineEdit_amount.setFixedHeight(25)
        
        lbl_fee_addr = QLabel('label', self)
        lbl_fee_addr.setText("手续费:")
        lbl_fee_addr.move(240, 110)
        self.lineEdit_fee = QLineEdit(self)
        self.lineEdit_fee.move(290, 110)
        self.lineEdit_fee.setFixedWidth(150)
        self.lineEdit_fee.setFixedHeight(25)
 
        self.btn_transfer = QPushButton('转账', self)
        self.btn_transfer.clicked.connect(self.onTransfer)
        self.btn_transfer.move(460,108)        

        #self.initUI()
    def checkInput(self):
        from_address = self.lineEdit_src_address.text()
        from_privkey = self.lineEdit_src_privkey.text()
        ret = CheckAccount(from_address.strip(),from_privkey.strip())
        if ret == False:
            dataformat.logger.error("check address %s failed!"%(from_address))
            return False

        to_address = self.lineEdit_dst_address.text()
        to_privkey = self.lineEdit_dst_privkey.text()
        ret = CheckAccount(to_address.strip(),to_privkey.strip())
        if ret == False:
            dataformat.logger.error("check address %s failed!"%(to_address))
            return False

    def onYoho(self):
        dataformat.logger.info('operation yoho')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 
        self.yoho.show()
        
    def onTransfer(self):
        dataformat.logger.info('operation transfer')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 
        if self.checkInput() == False:
            QMessageBox.warning(self, '转账', '地址和私钥不匹配',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 
        from_privkey = self.lineEdit_src_privkey.text()
        amount = self.lineEdit_amount.text();
        fee = self.lineEdit_fee.text()
        aitd = ETHAPI(setting.g_setting.getRpcUrl())
        to_address = self.lineEdit_dst_address.text()
        result = aitd.offlineSign(from_privkey,to_address,amount,fee)
        dataformat.logger.info(result)
        QMessageBox.warning(self, '转账', '转账完成',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)

    def checkUIFile(self):
        ret = False
        base_path = setting.g_setting.getPath();
        ui_file = base_path + setting.g_setting.getUi("account")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("addressimport")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;
        
        ui_file = base_path + setting.g_setting.getUi("airdrop")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("balance")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("coinview")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("configdialog")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("dialogpassword")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("makemarket")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("record")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("rpcdialog")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("token")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;

        ui_file = base_path + setting.g_setting.getUi("yoho")
        if Path(ui_file).is_file() == False:
            QMessageBox.warning(self, '找不到文件', ui_file,  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return ret;
        ret = True
        return ret;
 

    def onOpen(self):
        dataformat.logger.info('action open triggered')
        file_type = 'json files(*.json)'
        open_file, _ = QFileDialog.getOpenFileName(self, '选择文件','', file_type)
        if open_file == "":
            dataformat.logger.info("Not select file !")
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 
        json_config  = json.load(open(open_file))
        setting.g_setting.readConfig(json_config)
        dataformat.logger.info(json_config)
        if self.checkUIFile() == False:
            dataformat.logger.info("no ui file!")
            return 

        self.init = True
        self.account = account.Account()
        self.coin_view = coinview.CoinView()
        self.makemarket = makemarket.MakeMarket()
        self.token = token.Token()
        self.airdrop = airdrop.Airdrop()
        self.rpcdlg = config.RpcDlg()
        self.record = record.Record()
        self.yoho = yoho.YohoContral()

    def onClose(self):
        dataformat.logger.info('action close triggered')

    def onToken(self):
        dataformat.logger.info('action token triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 
        self.token.show()
        
    def onBridge(self):
        dataformat.logger.info('action Bridge triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 

    def onAirdrop(self):
        dataformat.logger.info('action airdrop triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        self.airdrop.show()
    
    def onQuantization(self):
        dataformat.logger.info('action quantization triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        self.makemarket.show()

    def onAccount(self):
        dataformat.logger.info('action account triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        self.account.show()

    def onBorrowLend(self):
        dataformat.logger.info('action borrow lend triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return 

    def onAsset(self):
        dataformat.logger.info('action asset triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        self.coin_view.exec()

    def onRecord(self):
        dataformat.logger.info('action record triggered')
        if self.init == False:
            QMessageBox.warning(self, '初始化', '请打开配置文件',  QMessageBox.Yes | QMessageBox.Cancel, QMessageBox.Yes)
            return
        self.record.show()

if __name__ == '__main__':
    app = QApplication(sys.argv)
    dataformat.InitLog()
    dataformat.logger.debug('begin sys')
    ex = Manager()
    ex.show(); 
    dataformat.logger.debug('manager show')
    sys.exit(app.exec_())
