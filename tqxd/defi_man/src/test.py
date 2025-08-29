import imp
import sys
from PyQt5.QtWidgets import QApplication, QMainWindow
from PyQt5 import QtCore, QtGui, QtWidgets,Qt
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
import json
from clickhouse_driver import result
from eth_utils import address
from eth_utils.crypto import keccak
from web3 import contract
import web3
from web3.types import Nonce
import xlrd
import xlwt
from common import dataformat
from interactive import makemarket
from contract import compoundabi
from contract import erc20abi
from contract import swapabi
from CHAINAPI import ETHAPI
from common import ckclient
from db import dbconnect
import solcx

def TestCompound():

    privkey = 'a296ebf4a6667a8fc2814c123b932194b2e06f3882971d8f3b9b71a0b55cc817'
    address = '0xDB55f4D5B6101450c8B13579442D9e6B25132A9a'
    abi_file = '../abi/arc20.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = "0x4B6b9F3695205C8468ddf9AB4025ec2A09bDfF1a"
    url = 'http://172.63.1.186:8545'
    usdt = erc20abi.ERC20(contract_abi, contract_address, url)
    usdt.PrintAllFunction()
    #usdt.Approve(privkey,'0x53a2E9148058E9Fe35490fe72ec6054abc3f6fc9', -1)

    abi_file = '../abi/cerc20_delegator.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = "0x53a2E9148058E9Fe35490fe72ec6054abc3f6fc9"
    earc_usdt = compoundabi.CErc20Delegator(contract_abi, contract_address, url)
    #earc_usdt.PrintAllFunctionEvent()
    amount = 10000000000000000000
    #earc_usdt.mint(privkey, amount)
    user_snap = earc_usdt.getAccountSnapshot(address)
    exchangeRateCurrent = earc_usdt.exchangeRateCurrent()
    brow_amount = user_snap[1] * exchangeRateCurrent/(1*10**(18+18-8))
    print(brow_amount)

    abi_file = '../abi/comptroller.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = "0xd898514F77A26df1096cF96f5024387BCd611E9f"
    comptroller = compoundabi.Comptroller(contract_abi, contract_address, url)
   # comptroller.PrintAllFunctionEvent()

    abi_file = '../abi/multi_price_oracle.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = "0xfCa85d6E3D480a94CF80Bb9E2818346AbA979A09"
    price_oracle = compoundabi.MultiPriceOracle(contract_abi, contract_address, url)
    #price_oracle.PrintAllFunctionEvent()

def TestLp():
    abi_file = "../abi/uniswap-v2-pair.json"
    contract = "0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    contract_abi = json.load(open(abi_file))
    usdt_aitd_pair = swapabi.Pair(contract_abi,contract, "http://192.168.1.165:8545")
    usdt_aitd_pair.PrintAllFunction()
    total = usdt_aitd_pair.totalSupply()
    print(total)
    abi_file = "../abi/arc20.json"
    contract = "0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    contract_abi = json.load(open(abi_file))
    usdt_token = erc20abi.ERC20(contract_abi,contract, "http://192.168.1.165:8545")
    lp_address = "0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    lp_usdt_toltal = usdt_token.GetBalance(lp_address)
    print(lp_usdt_toltal)
    contract = "0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    waitd_token = erc20abi.ERC20(contract_abi,contract, "http://192.168.1.165:8545")
    lp_address = "0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    lp_waitd_toltal = waitd_token.GetBalance(lp_address)
    print(lp_waitd_toltal)

    user_address = "0x00049ad078e1339e279f9ea343a746c72d00fa6b"
    pair_balance = usdt_aitd_pair.balanceOf(user_address)
    print(pair_balance)

    abi_file = "../abi/staking-rewards.json"
    contract = "0x9102807f101b23e23BCD9EAbcC7A36770684c735"
    contract_abi = json.load(open(abi_file))
    usdt_aitd_stake1 = swapabi.StakingRewards(contract_abi,contract, "http://192.168.1.165:8545")
    stake_balance1 = usdt_aitd_stake1.balanceOf(user_address)
    print(stake_balance1)
    usdt_aitd_stake1.PrintAllFunctionEvent()
    total = usdt_aitd_stake1.totalSupply()
    print(total)

    contract = "0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    usdt_aitd_stake2 = swapabi.StakingRewards(contract_abi,contract, "http://192.168.1.165:8545")
    stake_balance2 = usdt_aitd_stake2.balanceOf(user_address)
    print(stake_balance2)
    total = usdt_aitd_stake2.totalSupply()
    print(total)

def TestLiu():
    privkey = "0xb6223d67aeb0105c5d3261babaf64baec8b72b9ea565a374ae9a6e4f0135c018"
    fee = 0.0001
    contract = "0xa250e791dd4420e93e96b14ea4400e2dcfcfdfba"
   


def TestDecimals():
    print(dataformat.ToWei("1",0))
    print(dataformat.ToWei("1",1))
    print(dataformat.ToWei("1",6))
    print(dataformat.ToWei("1",18))


def GetAddLp():
    address_file = "../accountliu.xls"
    accounts =  dataformat.GetExcelData(address_file)
    count = len(accounts)
    filter_results = []
    for i in range(count):
        address = accounts[i][0]
        sql = "SELECT param1, param2, transaction_hash from  event_log el  WHERE  el.transaction_hash  in (SELECT  transaction_hash from  block_number_transaction WHERE  `from` = LOWER('%s')) and event_name  ='Mint';"%(address)
        amount = ckclient.GetCkData(sql)
        amount_count = len(amount) 
        for j in range(amount_count):
            usdt = dataformat.FromWei(str(amount[j][0]), 6)
            aitd = dataformat.FromWei(str(amount[j][1]), 18)
            txid = str(amount[j][2])
            result_value = []
            result_value.append(address)
            result_value.append(usdt)
            result_value.append(aitd)
            result_value.append(txid)
            sql_time = "SELECT toDateTime(`timestamp`)  from block_number_transaction bnt WHERE transaction_hash ='%s'"%(txid)
            tt = ckclient.GetCkData(sql_time)
            result_value.append(str(tt[0][0]))
            filter_results.append(result_value)

    print("get data")
    workbook = xlwt.Workbook(encoding = 'utf-8')
    worksheet = workbook.add_sheet("record")
    row_size = len(filter_results)
    for i in range(row_size):
        row_value = filter_results[i]
        col_size = len(row_value)
        for j in range (5):
            worksheet.write(i,j, row_value[j])
    filename = "add_lp_record.xls"
    workbook.save(filename)

def GetReward():
    address_file = "../accountliu.xls"
    accounts =  dataformat.GetExcelData(address_file)
    count = len(accounts)
    filter_results = []
    for i in range(count):
        address = accounts[i][0]
        sql = "SELECT param1, transaction_hash from  event_log el  WHERE  el.transaction_hash  in (SELECT  transaction_hash from  block_number_transaction WHERE `from` = LOWER('%s')) and event_name  ='RewardPaid';"%(address)
        amount = ckclient.GetCkData(sql)
        amount_count = len(amount) 
        for j in range(amount_count):
            aitd = dataformat.FromWei(str(amount[j][0]), 18)
            txid = str(amount[j][1])
            result_value = []
            result_value.append(address)
            result_value.append(aitd)
            result_value.append(txid)
            sql_time = "SELECT toDateTime(`timestamp`)  from block_number_transaction bnt WHERE transaction_hash ='%s'"%(txid)
            tt = ckclient.GetCkData(sql_time)
            result_value.append(str(tt[0][0]))
            filter_results.append(result_value)

    print("get data")
    workbook = xlwt.Workbook(encoding = 'utf-8')
    worksheet = workbook.add_sheet("record")
    row_size = len(filter_results)
    for i in range(row_size):
        row_value = filter_results[i]
        col_size = len(row_value)
        for j in range (4):
            worksheet.write(i,j, row_value[j])
    filename = "withdraw_reward.xls"
    workbook.save(filename)

def compileSolc():
    version  = solcx.get_installable_solc_versions()
    print(version)
    #solcx.install_solc(version="latest", show_progress=False, solcx_binary_path=None)
    print(solcx.get_solc_version())

def testRlp():
    from web3 import Web3, IPCProvider
    from web3.middleware import geth_poa_middleware
    import rlp

    endpoint = 'http://192.168.1.11:18545'
    w3 = Web3(Web3.HTTPProvider(endpoint))
    w3.eth.defaultAccount = Web3.toChecksumAddress('0x1C7EAD88395528bFF24012d810Fa49FfDBD5dDDC')
    Nonce  = 0 
    rlp.encode
    contract_address = keccak(rlp([w3.eth.defaultAccount, Nonce]))
    print(contract_address.hex())


def testVerify():
    from web3 import Web3, IPCProvider
    from web3.middleware import geth_poa_middleware
    import rlp
    endpoint = 'http://172.63.1.44:8545'
    w3 = Web3(Web3.HTTPProvider(endpoint))
    file = "./crer.csv"
    result = dataformat.GetCsvData(file)
    count = len(result)

    for i in range(count):
        order = result[i][0]
        txid = result[i][1] 
        print(i)
        try:
            tx = w3.eth.getTransaction(txid)
            print(tx)
        except web3.exceptions.TransactionNotFound:
            print(txid)
            print("not found")
            pass
        else:
            print(order)
            print("found")
        

def testMysql():
    db_mysql = dbconnect.DBMysql("192.168.1.10",3306,"root","123456","yoho_moniter_test")
    sql = "SELECT * FROM token;"
    result = db_mysql.getData(sql)
    count = len(result)
    for i in range(count):
        row = result[i]
        print(i)
        print(row[0])
        #print(row[1])
        #print(row[2])
        #print(row[3])
        #print(row[4])
 
def testBacktrade():
    import backtrader as bt
    cerebro = bt.Cerebro();
    print("init asset: %.2f" % cerebro.broker.getvalue())
    cerebro.run();
    print("end asset: %.2f" % cerebro.broker.getvalue())

def testCoinmarket():
    # encoding:utf-8
    from requests import Request,Session
    from requests.exceptions import ConnectionError, Timeout, TooManyRedirects
    
    url = 'https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest'
    parameters = {
    'start':'1',
    'limit':'5000',
    'convert':'USD'
    }   
    headers = {
    'Accepts': 'application/json',
    'X-CMC_PRO_API_KEY': 'dde49e15-3db6-4c80-8d9b-2d3d6edb4870',
    }

    session = Session()
    session.headers.update(headers)

    try:
        response = session.get(url, params=parameters)
        data = json.loads(response.text)
        print(json.dumps(data, sort_keys=False, indent=2))
    except (ConnectionError, Timeout, TooManyRedirects) as e:
        print(e)



if __name__ == '__main__':
    #app = QApplication(sys.argv)
    #TestDecimals()
    #api = ETHAPI('http://192.168.1.165:8545')
    #balance = api.getBalance('0x331873335Cb7a391340fbC77aFc521f42598F5d3')
    #accout = accountnew.AccounNew()
    #accout_new.show()
    #sys.exit(app.exec_())
    #GetReward()
    #compileSolc()
    #testRlp()
    #testVerify()
    #testMysql()
    #testBacktrade()
    testCoinmarket()

    