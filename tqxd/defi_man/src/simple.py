#!/usr/bin/python3

import sys
from PyQt5.QtWidgets import QApplication, QMainWindow
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
import xlrd
from common import dataformat
from interactive import makemarket
from common import setting

if __name__ == '__main__':
    app = QApplication(sys.argv)
    config_file = "./config/config.json"
    json_config  = json.load(open(config_file))
    setting.g_setting.readConfig(json_config)
    dataformat.InitLog()
    dataformat.logger.debug('begin sys')
    
    ex = makemarket.MakeMarket()
    ex.show(); 
    dataformat.logger.debug('manager show')
    #accout = accountnew.AccounNew()
    #accout_new.show()
    sys.exit(app.exec_())
