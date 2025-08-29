#!/usr/bin/
import sys 
import json
#client = Client(host='192.168.1.27',database = 'wallet_report_prod', user='default', password='123456')

class Setting(object):
    rpc_url = "http://192.168.1.165:8545"
    program_path = "./"
    local_db = "account.db"
    ck_host = "192.168.1.27"
    ck_db = "wallet_report_prod"
    ck_user = "default"
    ck_pswd = "123456"
    json_ui = None
    json_swap = None
    json_token = None
    json_compound = None
    json_contract = None
    json_yoho = None
    def __init__(self):
        self.rpc_url = "http://192.168.1.165:8545"
        self.program_path = sys.path[0]

    def setRpcUrl(self, rpc_url):
        self.rpc_url = rpc_url
    
    def setPath(self, path):
        self.program_path = path

    def getRpcUrl(self):
        return self.rpc_url
    
    def getPath(self):
        return self.program_path

    def setLocalDB(self, db_file):
        self.local_db = db_file

    def getLocalDB(self):
        return self.local_db

    def setClickHouse(self, host, db, user,pswd):
        self.ck_host = host
        self.ck_db = db
        self.ck_user = user
        self.ck_pswd = pswd
    
    def getCkHost(self):
        return self.ck_host
    
    def getCkDB(self):
        return self.ck_db

    def getCkUser(self):
        return self.ck_user

    def getCkPswd(self):
        return self.ck_pswd

    def setUi(self, json_ui):
        self.json_ui = json_ui

    def getUi(self, ui_name):
        try:
            self.json_ui[ui_name]
        except KeyError:
            return ""
        return self.json_ui[ui_name]

    def setCompound(self, json_compound):
        self.json_compound = json_compound

    def getCompound(self):
        return self.json_compound

    def setSwap(self, json_swap):
        self.json_swap = json_swap

    def getSwap(self):
        return self.json_swap

    def setYoho(self, json_yoho):
        self.json_yoho = json_yoho

    def getYoho(self):
        return self.json_yoho

    def setToken(self, json_token):
        self.json_token = json_token

    def getToken(self):
        return self.json_token

    def setContract(self, json_contract):
        self.json_contract = json_contract

    def getContract(self):
        return self.json_contract

    def readConfig(self,json_config):
        try:
            self.setPath(json_config["path"])
            self.setLocalDB(json_config["db"]["sqlite"])
            self.setRpcUrl(json_config["rpc"]["AITD"])
            json_ck = json_config["db"]["clickhouse"]
            self.setClickHouse(json_ck["host"], json_ck["database"], json_ck["user"], json_ck["password"])
            json_ui  = json_config["ui"]
            self.setUi(json_ui)
            json_compound = json_config["abi"]["compound"]
            self.setCompound(json_compound)
            json_swap = json_config["abi"]["swap"]
            self.setSwap(json_swap)
            print(json_swap)
            json_token = json_config["abi"]["token"]
            print(json_token);
            self.setToken(json_token)
            json_yoho = json_config["abi"]["yoho"]
            self.setYoho(json_yoho)
            json_contract = json_config["contract"]
            self.setContract(json_contract)
        except KeyError:
            print("read config fail")

        
g_setting = Setting()
