import logging
from eth_utils import address
from eth_utils.address import to_checksum_address
from toolz import dicttoolz
from web3 import Web3
import web3
import json
import os
import re
import time
import datetime

from web3.types import RPCEndpoint
from contract import erc20abi
from common import dataformat,setting
from eth_account.messages import (
    encode_defunct,
)

def GetAllEvent(abi):
    count = len(abi)
    events = []
    for abi_value in abi:
         abi_type = abi_value["type"]
         event = {}
         if abi_type == 'event':
             eventname =abi_value["name"]
             eventparams = "("
             input_size = len(abi_value["inputs"])
             n = 0
             for input_value in  abi_value["inputs"]:
                 n += 1
                 if n == input_size :
                    eventparams = eventparams + input_value["type"]
                 else :
                    eventparams = eventparams + input_value["type"] + ","  
             eventparams = eventparams + ")" 
             event[eventname] = eventparams
             events.append(event)
    dataformat.logger.info(events)
    return events
              

class UniswapV2Factory(object):

    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        self.contract = self.web3.eth.contract(address=contract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        dataformat.logger.info(self.contract.all_functions())
        

class UniswapV2Router02(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        self.contract = self.web3.eth.contract(address=contract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        dataformat.logger.info(self.contract.all_functions())

    def SwapAitdToUsdt(self, privkey, amount):
        #<Function swapETHForExactTokens(uint256,address[],address,uint256)>, 
        value = web3.Web3.toWei(amount, 'ether')
        amountOut = value * 10 #web3.Web3.toWei("1", 'ether')
        path = []
        swap_contract = setting.g_setting.getContract();
        waitd_address = swap_contract["swap"]["WAITD"]
        usdt_addrss = swap_contract["swap"]["USDT"]
        path.append(web3.Web3.toChecksumAddress(waitd_address))
        path.append(web3.Web3.toChecksumAddress(usdt_addrss))
        #path.append(web3.Web3.toChecksumAddress("0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"))
        #path.append(web3.Web3.toChecksumAddress("0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"))
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        t = time.time()
        deadline = int(t)  + 20 * 60
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.swapETHForExactTokens(amountOut, path, address, deadline).buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value':value })
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()

    def removeLiquidity(self, privkey,tokenA, tokenB, lp_amount):
        pk = int(privkey, 16)
        sendfrom  = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address =publicKey.to_checksum_address()
        t = time.time()
        deadline = int(t)  + 20 * 60
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        tokenA_min = web3.Web3.toWei('0', 'wei')
        tokenB_min = web3.Web3.toWei('0', 'wei')
        unsigned_tx = self.contract.functions.removeLiquidity(tokenA,tokenB, lp_amount, tokenA_min, tokenB_min, address, deadline).buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value':0 })
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()

    def SwapUsdtToAitd(self, privkey, amount, price):
        #swapExactTokensForETH(uint256,uint256,address[],address,uint256)>
        aitd_amount = web3.Web3.toWei(amount, 'ether')
        dataformat.logger.info(aitd_amount)
        ustd_amount = int(aitd_amount * price) #web3.Web3.toWei('0.01055', 'ether') #int (amountIn / price) 
        dataformat.logger.info(ustd_amount)
        path = []
        json_swap = setting.g_setting.getSwap()
        json_contract = setting.g_setting.getContract()
        swap_contract = setting.g_setting.getContract();
        waitd_address = swap_contract["swap"]["WAITD"]
        usdt_addrss = swap_contract["swap"]["USDT"]
        path.append(web3.Web3.toChecksumAddress(usdt_addrss))
        path.append(web3.Web3.toChecksumAddress(waitd_address))
        #path.append(web3.Web3.toChecksumAddress("0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"))
        #path.append(web3.Web3.toChecksumAddress("0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"))
        pk = int(privkey, 16)
        sendfrom  = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address =publicKey.to_checksum_address()
        t = time.time()
        deadline = int(t)  + 20 * 60
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.swapExactTokensForETH(ustd_amount,aitd_amount, path, address, deadline).buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value':0 })
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()
        

    def SwapAitdTokenToUsdt(self, privkey, amount, price):
        #<Function swapExactETHForTokens(uint256,address[],address,uint256)>, 
        value = web3.Web3.toWei(amount, 'wei')
        amountIn = value
        dataformat.logger.info(amountIn)
        amountOut = web3.Web3.toWei(value /1000000000000 * price, 'wei')
        dataformat.logger.info(amountOut)
        value = web3.Web3.toWei('0', 'wei')

        path = []
        #path.append(web3.Web3.toChecksumAddress("0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"))
        #path.append(web3.Web3.toChecksumAddress("0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"))
        swap_contract = setting.g_setting.getContract();
        waitd_address = swap_contract["swap"]["WAITD"]
        usdt_addrss = swap_contract["swap"]["USDT"]
        path.append(web3.Web3.toChecksumAddress(waitd_address))
        path.append(web3.Web3.toChecksumAddress(usdt_addrss))

        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()

        t = time.time()
        deadline = int(t)  + 20 * 60
        
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.swapExactTokensForTokens(amountIn, amountOut, path, address, deadline).buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value':value })
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        dataformat.logger.info(txid.hex())
        return txid.hex()

    def addLiquidity(self, token_a, token_b, privkey):
        contract_a = web3.Web3.toChecksumAddress(token_a["address"])
        desired_a =  token_a["desired"]
        min_a = token_a["min"]

        value = web3.Web3.toWei('0', 'wei')

        contract_b = web3.Web3.toChecksumAddress(token_b["address"])
        desired_b = token_b["desired"]
        min_b = token_b["min"]

        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()

        t = time.time()
        deadline = int(t)  + 20 * 60
        
        value = 0;
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.addLiquidity(contract_a, contract_b, desired_a, desired_b, min_a, min_b, address, deadline).buildTransaction({'gas':2000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value':value })
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        dataformat.logger.info(txid.hex())
        return txid.hex()

    def getSign(self, privkey, json_data):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        dataformat.logger.info(address)
        sign_data = json.dumps(json_data)
        dataformat.logger.info(sign_data)
        message = encode_defunct(text=sign_data)
        signature = sendfrom.sign_message(message)
        dataformat.logger.info('0x%x'%signature.r)
    
        dataformat.logger.info(signature)


class StakingRewardsFactory(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        self.contract = self.web3.eth.contract(address=contract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        dataformat.logger.info(self.contract.all_functions())

 

class StakingRewards(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        self.contract = self.web3.eth.contract(address=contract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        dataformat.logger.info(self.contract.all_functions())


    def earned(self, address):
        chkaddress = web3.Web3.toChecksumAddress(address)
        amount = self.contract.caller().earned(chkaddress)
        dataformat.logger.info(amount)
        return amount

    def totalSupply(self):
        total = self.contract.caller().totalSupply()
        return total

    def rewards(self, address):
        chkaddress = web3.Web3.toChecksumAddress(address)
        amount = self.contract.caller().rewards(chkaddress)
        dataformat.logger.info(amount)
        return amount

    def balanceOf(self, address):
        chkaddress = web3.Web3.toChecksumAddress(address)
        amount = self.contract.caller().balanceOf(chkaddress)
        dataformat.logger.info(amount)
        return amount

    def stake(self, amount,privkey):
         #<Function swapExactETHForTokens(uint256,address[],address,uint256)>,
        stake_amount = web3.Web3.toWei(amount, 'wei')
        dataformat.logger.info(stake_amount)
        value = web3.Web3.toWei('0', 'wei')
    
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.stake(stake_amount).buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value':value })
        signed_tx =  sendfrom.signTransaction(unsigned_tx)

        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        dataformat.logger.info(txid.hex())
        return txid.hex()


    def getReward(self,privkey):
        value = web3.Web3.toWei('0', 'ether')
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        
        unsigned_tx = self.contract.functions.getReward().buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value': value })

        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        dataformat.logger.info(txid.hex())
        return txid.hex()

    def exit(self, privkey):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        value = web3.Web3.toWei('0', 'wei')
        unsigned_tx = self.contract.functions.exit().buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce, 'value': value })

        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        dataformat.logger.info(txid.hex())
        return txid.hex()

class Pair(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        ckaddress = web3.Web3.toChecksumAddress(contract_address)
        self.contract = self.web3.eth.contract(address=ckaddress, abi =json_abi)
       
    def PrintAllFunction(self):
        #dataformat.logger.info(self.contract.all_functions())
        print(self.contract.all_functions())

    def balanceOf(self, address):
        ckaddress = web3.Web3.toChecksumAddress(address)
        amount = self.contract.caller().balanceOf(ckaddress)
        return amount

    def totalSupply(self):
        total = self.contract.caller().totalSupply() 
        return  total


    def info(self):
        name = self.contract.caller().name()
        symbol = self.contract.caller().symbol()
        decimal = self.contract.caller().decimals()
        total = self.contract.caller().totalSupply() 
        return name, symbol, decimal, total

    def Allowance(self, address, approve):
        #allowance(address,address)>    
        from_address = web3.Web3.toChecksumAddress(address)
        approve_address = web3.Web3.toChecksumAddress(approve)
        amount = self.contract.caller().allowance(from_address, approve_address)
        dataformat.logger.info(amount)
        return amount

    def changeUser(self, privkey, to_address):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        chktoaddress = web3.Web3.toChecksumAddress(to_address)
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.changeUser(address, chktoaddress).buildTransaction({'gas':600000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce})
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()
    

    def Approve(self,privkey,to_address,amount):
         #<Function approve(address,uint256)>
         banlance = 2**256 -1
         if amount > 0:
             banlance = amount
         dataformat.logger.info(banlance)
         pk = int(privkey, 16)
         sendfrom = web3.eth.Account.privateKeyToAccount(pk)
         privateKey = sendfrom._key_obj
         publicKey = privateKey.public_key
         address = publicKey.to_checksum_address()
         chktoaddress = web3.Web3.toChecksumAddress(to_address)
         gasprice = self.web3.eth.gasPrice
         nonce = self.web3.eth.getTransactionCount(address, "pending")
         unsigned_tx = self.contract.functions.approve(chktoaddress, banlance).buildTransaction({'gas':600000,
                                                      'gasPrice':gasprice, 'from': address, 'nonce':nonce})
         signed_tx =  sendfrom.signTransaction(unsigned_tx)
         txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
         return txid.hex()

    def Transfer(self, privkey, to_address, amount):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        chktoaddress = web3.Web3.toChecksumAddress(to_address)
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.transfer(chktoaddress, amount).buildTransaction({'gas':600000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce})
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()

    

def CheckTxid(txid, url):
    w3 = Web3(Web3.HTTPProvider(url))
    try:
        tx = w3.eth.getTransactionReceipt(txid)
        status = tx['status']
        if status == 0:
            return False
        if status == 1:
            return True
    except  Exception as err:
        time.sleep(3)
        dataformat.logger.info("error check txid %s"%txid)
        return CheckTxid(txid, url)
    
    return False

def Release(address, privkey, rpc_url, callbackTx):
    path = setting.g_setting.getPath()
    swap_abi = setting.g_setting.getSwap()
    abi_file = path + swap_abi["staking-rewards"] #'../abi/staking-rewards.json'    
    contract_abi  = json.load(open(abi_file))
    json_contract = setting.g_setting.getContract()
    json_contract_swap = json_contract["swap"]
    contract_address = json_contract_swap["stake-reward"]["AITD-USDT-2"] #"0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    ewards_aitd_usdt = StakingRewards(contract_abi, contract_address, rpc_url)

    exit_txid = ewards_aitd_usdt.exit(privkey)
    if CheckTxid(exit_txid, rpc_url) == False:
        dataformat.logger.error("exit txid %s failed "%exit_txid)
        print("exit txid %s failed "%exit_txid)
        return
    print("exit txid %s success "%exit_txid)
    callbackTx(address,contract_address,exit_txid, 0)
    
    contract_address = json_contract_swap["pair"]["AITD-USDT"]#"0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    abi_file =  path + swap_abi["uniswap-v2-pair"]#'../abi/uniswap-v2-pair.json'    
    contract_abi  = json.load(open(abi_file))
    pair = Pair(contract_abi, contract_address, rpc_url)
    lp_amount = pair.balanceOf(address)
    dataformat.logger.info(lp_amount)

    router_address = json_contract_swap["UniswapV2Router02"]
    allow_balance = pair.Allowance(address,router_address) #'0x3320B7E625124910BFad5CaF9DC1767205D91286')
    if allow_balance < lp_amount:
        approve_txid = pair.Approve(privkey, router_address, -1) #'0x3320B7E625124910BFad5CaF9DC1767205D91286', -1)
        if CheckTxid(approve_txid, rpc_url) == False:
            dataformat.logger.info("approve txid %s failed "%approve_txid)
            print("approve txid %s failed "%approve_txid)
            return
        dataformat.logger.info("approve txid %s success "%approve_txid)
        print("approve txid %s success "%approve_txid)

    abi_file = path + swap_abi["uniswap-v2-router0"]#'../abi/uniswap-v2-router02.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = router_address #"0x3320B7E625124910BFad5CaF9DC1767205D91286"
    uniswap_v2_router_02  = UniswapV2Router02(contract_abi, contract_address, rpc_url)
    waitd = json_contract_swap["WAITD"]#"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    usdt = json_contract_swap["USDT"]#"0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    remove_txid = uniswap_v2_router_02.removeLiquidity(privkey,waitd,usdt,lp_amount)
    if CheckTxid(remove_txid, rpc_url) == False:
        dataformat.logger.error("remove txid %s failed "%remove_txid)
        print("remove txid %s failed "%remove_txid)
        return
    print("remove txid %s success "%remove_txid)


def AddPositions(address, privkey, rpc_url, price, callbackTx):

    path = setting.g_setting.getPath()
    token_abi = setting.g_setting.getToken()
    abi_file = path + token_abi["arc20"]#'../abi/arc20.json'    
    contract_abi  = json.load(open(abi_file))

    json_contract = setting.g_setting.getContract()
    json_contract_swap = json_contract["swap"]
    contract_address = json_contract_swap["WAITD"]#"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    WAITD  = erc20abi.ERC20(contract_abi, contract_address, rpc_url)
    waitd_amount = WAITD.GetBalance(address)

    swap_abi = setting.g_setting.getSwap()
    abi_file = swap_abi["uniswap-v2-router02"]#'../abi/uniswap-v2-router02.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = json_contract_swap["UniswapV2Router02"]#"0x3320B7E625124910BFad5CaF9DC1767205D91286"
    uniswap_v2_router_02  = UniswapV2Router02(contract_abi, contract_address, rpc_url)
   
    token_waitd  = {}
    waitd_amount = WAITD.GetBalance(address)
    token_waitd["address"] = json_contract_swap["WAITD"] #"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    token_waitd["desired"] = web3.Web3.toWei(waitd_amount, 'wei')
    token_waitd["min"] = web3.Web3.toWei(waitd_amount/10*9, 'wei') 

    token_usdt  = {}
    token_usdt["address"] = json_contract_swap["USDT"]#"0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    token_usdt["desired"] = web3.Web3.toWei(waitd_amount /1000000000000 * price, 'wei')
    token_usdt["min"] = web3.Web3.toWei(waitd_amount /1000000000000 * price/10*9, 'wei')

    al_txid = uniswap_v2_router_02.addLiquidity(token_waitd, token_usdt, privkey)
    if CheckTxid(al_txid, rpc_url) == False:
        dataformat.logger.info("add liquidity txid %s failed "%al_txid)
        print("add liquidity txid %s failed "%al_txid)
        return
    dataformat.logger.info("add liquidity txid %s success "%al_txid)
    print("add liquidity txid %s success "%al_txid)
    contract_address = json_contract_swap["pair"]["AITD-USDT"]#"0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    abi_file = path + swap_abi["uniswap-v2-pair"]#'../abi/uniswap-v2-pair.json'    
    contract_abi  = json.load(open(abi_file))
    token_aitd_usdt = Pair(contract_abi, contract_address, rpc_url)
    stake_amount = token_aitd_usdt.balanceOf(address)
    dataformat.logger.info(stake_amount)
    stake_reward_2_address = json_contract_swap["stake-reward"]["AITD-USDT-2"]
    allow_balance = token_aitd_usdt.Allowance(address, stake_reward_2_address)#'0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39')
    if allow_balance < stake_amount:
        approve_txid = token_aitd_usdt.Approve(privkey, stake_reward_2_address, -1)#'0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39', -1)
        if CheckTxid(approve_txid, rpc_url) == False:
            dataformat.logger.info("approve txid %s failed "%approve_txid)
            print("approve txid %s failed "%approve_txid)
            return
        dataformat.logger.info("approve txid %s success "%approve_txid)
        print("approve txid %s success "%approve_txid)

    abi_file = path + swap_abi["staking-rewards"] #'../abi/staking-rewards.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = stake_reward_2_address #"0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    rewards_aitd_usdt = StakingRewards(contract_abi, contract_address, rpc_url)
    stake_txid = rewards_aitd_usdt.stake(stake_amount, privkey)
    if CheckTxid(stake_txid, rpc_url) == False:
            dataformat.logger.info("stake txid %s failed "%stake_txid)
            print("stake txid %s failed "%stake_txid)
            return
    dataformat.logger.info("stake txid %s success "%stake_txid)
    print("stake txid %s success "%stake_txid)


def AddPositionsReverse(address, privkey, rpc_url, price, callbackTx):
    path = setting.g_setting.getPath()
    token_abi = setting.g_setting.getToken()
    abi_file = path + token_abi["arc20"]
    #abi_file = '../abi/arc20.json'    
    contract_abi  = json.load(open(abi_file))

    json_contract = setting.g_setting.getContract()
    json_contract_swap = json_contract["swap"]
    contract_address = json_contract_swap["USDT"]
    #contract_address = "0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    USDT  = erc20abi.ERC20(contract_abi, contract_address, rpc_url)
    usdt_amount = USDT.GetBalance(address)

    swap_abi = setting.g_setting.getSwap()
    abi_file = path + swap_abi["uniswap-v2-router02"]#'../abi/uniswap-v2-router02.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = json_contract_swap["UniswapV2Router02"]#"0x3320B7E625124910BFad5CaF9DC1767205D91286"
    uniswap_v2_router_02  = UniswapV2Router02(contract_abi, contract_address, rpc_url)
   
    token_waitd  = {}
    waitd_amount = usdt_amount * 1000000000000 / price
    print(waitd_amount)
    token_waitd["address"] = json_contract_swap["WAITD"]#"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    token_waitd["desired"] = web3.Web3.toWei(waitd_amount, 'wei')
    token_waitd["min"] = web3.Web3.toWei(waitd_amount/10*9, 'wei') 

    token_usdt  = {}
    usdt_amount = USDT.GetBalance(address)
    token_usdt["address"] = json_contract_swap["USDT"]#"0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    token_usdt["desired"] = web3.Web3.toWei(usdt_amount, 'wei')
    token_usdt["min"] = web3.Web3.toWei(usdt_amount/10*9, 'wei')

    al_txid = uniswap_v2_router_02.addLiquidity(token_waitd, token_usdt, privkey)
    if CheckTxid(al_txid, rpc_url) == False:
        dataformat.logger.info("add liquidity txid %s failed "%al_txid)
        print("add liquidity txid %s failed "%al_txid)
        return
    dataformat.logger.info("add liquidity txid %s success "%al_txid)
    print("add liquidity txid %s success "%al_txid)
    contract_address = json_contract_swap["pair"]["AITD-USDT"]#"0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    abi_file = path + swap_abi["uniswap-v2-pair"] #'../abi/uniswap-v2-pair.json'    
    contract_abi  = json.load(open(abi_file))
    token_aitd_usdt = Pair(contract_abi, contract_address, rpc_url)
    stake_amount = token_aitd_usdt.balanceOf(address)
    dataformat.logger.info(stake_amount)
    stake_reward_2_address = json_contract_swap["stake-reward"]["AITD-USDT-2"]
    allow_balance = token_aitd_usdt.Allowance(address, stake_reward_2_address)#'0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39')
    if allow_balance < stake_amount:
        approve_txid = token_aitd_usdt.Approve(privkey, stake_reward_2_address, -1)#'0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39', -1)
        if CheckTxid(approve_txid, rpc_url) == False:
            dataformat.logger.info("approve txid %s failed "%approve_txid)
            print("approve txid %s failed "%approve_txid)
            return
        dataformat.logger.info("approve txid %s success "%approve_txid)
        print("approve txid %s success "%approve_txid)

    abi_file = path + swap_abi["staking-rewards"]; '../abi/staking-rewards.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = stake_reward_2_address#"0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    rewards_aitd_usdt = StakingRewards(contract_abi, contract_address, rpc_url)
    stake_txid = rewards_aitd_usdt.stake(stake_amount, privkey)
    if CheckTxid(stake_txid, rpc_url) == False:
            dataformat.logger.info("stake txid %s failed "%stake_txid)
            print("stake txid %s failed "%stake_txid)
            return
    dataformat.logger.info("stake txid %s success "%stake_txid)
    print("stake txid %s success "%stake_txid)

def Reinvest(address, privkey, rpc_url, price, callbackTx):
    path = setting.g_setting.getPath()
    swap_abi = setting.g_setting.getSwap()
    abi_file = path + swap_abi["staking-rewards"];#'../abi/staking-rewards.json'    
    contract_abi  = json.load(open(abi_file))
    json_contract = setting.g_setting.getContract()
    json_contract_swap = json_contract["swap"]

    stake_reward_2_address = json_contract_swap["stake-reward"]["AITD-USDT-2"]
    contract_address = stake_reward_2_address #"0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    rewards_aitd_usdt = StakingRewards(contract_abi, contract_address, rpc_url)
    earn_amount = rewards_aitd_usdt.earned(address)
    if earn_amount < 1000000000000:
        dataformat.logger.info("earn amount is too low %d"%earn_amount)
        print("earn amount is too low %d"%earn_amount)
        callbackTx(address,contract_address,"", 1)
        return

    reward_txid = rewards_aitd_usdt.getReward(privkey)
    if CheckTxid(reward_txid, rpc_url) == False:
        dataformat.logger.info("reward txid %s failed "%reward_txid)
        print("reward txid %s failed "%reward_txid)
        return
    
    dataformat.logger.info("reward txid %s success "%reward_txid)
    print("reward txid %s success "%reward_txid)

    token_abi = setting.g_setting.getToken()
    abi_file = path + token_abi["arc20"] #'../abi/arc20.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = json_contract_swap["WAITD"]#"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    WAITD  = erc20abi.ERC20(contract_abi, contract_address, rpc_url)
    waitd_amount = WAITD.GetBalance(address)

    abi_file = path + swap_abi["uniswap-v2-router02"]# '../abi/uniswap-v2-router02.json'    
    contract_abi  = json.load(open(abi_file))
    contract_address = json_contract_swap["UniswapV2Router02"]#"0x3320B7E625124910BFad5CaF9DC1767205D91286"
    uniswap_v2_router_02  = UniswapV2Router02(contract_abi, contract_address, rpc_url)
    swap_amount = waitd_amount /2
    swap_txid = uniswap_v2_router_02.SwapAitdTokenToUsdt(privkey, swap_amount, price)

    if CheckTxid(swap_txid, rpc_url) == False:
        dataformat.logger.info("swap txid %s failed "%swap_txid)
        print("swap txid %s failed "%swap_txid)
        return
    
    dataformat.logger.info("swap txid %s success "%swap_txid)

    token_waitd  = {}
    waitd_amount = WAITD.GetBalance(address)
    token_waitd["address"] = json_contract["WAITD"]#"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    token_waitd["desired"] = web3.Web3.toWei(waitd_amount, 'wei')
    token_waitd["min"] = web3.Web3.toWei(waitd_amount/10*9, 'wei') 

    token_usdt  = {}
    token_usdt["address"] = json_contract["USDT"] #"0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    token_usdt["desired"] = web3.Web3.toWei(waitd_amount /1000000000000 * price, 'wei')
    token_usdt["min"] = web3.Web3.toWei(waitd_amount /1000000000000 * price/10*9, 'wei')

    al_txid = uniswap_v2_router_02.addLiquidity(token_waitd, token_usdt, privkey)
    if CheckTxid(al_txid, rpc_url) == False:
        dataformat.logger.info("add liquidity txid %s failed "%al_txid)
        print("add liquidity txid %s failed "%al_txid)
        return
    dataformat.logger.info("add liquidity txid %s success "%al_txid)
    print("add liquidity txid %s success "%al_txid)
    contract_address = json_contract_swap["pair"]["AITD-USDT"]#"0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    abi_file = swap_abi["uniswap-v2-pair"]#'../abi/uniswap-v2-pair.json'    
    contract_abi  = json.load(open(abi_file))
    token_aitd_usdt = Pair(contract_abi, contract_address, rpc_url)
    stake_amount = token_aitd_usdt.balanceOf(address)
    dataformat.logger.info(stake_amount)
    allow_balance = token_aitd_usdt.Allowance(address, stake_reward_2_address)#'0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39')
    if allow_balance < stake_amount:
        approve_txid = token_aitd_usdt.Approve(privkey, stake_reward_2_address, -1)#'0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39', -1)
        if CheckTxid(approve_txid, rpc_url) == False:
            dataformat.logger.info("approve txid %s failed "%approve_txid)
            print("approve txid %s failed "%approve_txid)
            return
        dataformat.logger.info("approve txid %s success "%approve_txid)
        print("approve txid %s success "%approve_txid)

    stake_txid = rewards_aitd_usdt.stake(stake_amount, privkey)
    if CheckTxid(stake_txid, rpc_url) == False:
            dataformat.logger.info("stake txid %s failed "%stake_txid)
            print("stake txid %s failed "%stake_txid)
            return
    dataformat.logger.info("stake txid %s success "%stake_txid)
    print("stake txid %s success "%stake_txid)

def GetLpTotal(user_address):
    path = setting.g_setting.getPath()
    swap_abi = setting.g_setting.getSwap()
    abi_file = path + swap_abi["uniswap-v2-pair"]#"../abi/uniswap-v2-pair.json"
    json_contract = setting.g_setting.getContract()
    json_contract_swap = json_contract["swap"]
    contract = json_contract_swap["pair"]["AITD-USDT"]#"0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"
    contract_abi = json.load(open(abi_file))
    usdt_aitd_pair = Pair(contract_abi,contract, setting.g_setting.getRpcUrl())#"http://192.168.1.165:8545")
    total = usdt_aitd_pair.totalSupply()

    pair_address = json_contract_swap["pair"]["AITD-USDT"]#"0x1D16adac6be715F21bea6CB67F5E264a43d0b42d"

    token_abi = setting.g_setting.getToken()
    abi_file = path + token_abi["arc20"] #"../abi/arc20.json"
    contract = json_contract_swap["USDT"]#"0x848cb1a9770830da575DfD246dF2d4e38c1D40ed"
    contract_abi = json.load(open(abi_file))
    usdt_token = erc20abi.ERC20(contract_abi,contract, setting.g_setting.getRpcUrl())
    lp_usdt_toltal = usdt_token.GetBalance(pair_address)

    contract = json_contract_swap["WAITD"]#"0xEC4C225F734a614B6d6f61b5Ddf0ae96c8e85E32"
    waitd_token = erc20abi.ERC20(contract_abi,contract, setting.g_setting.getRpcUrl())#"http://192.168.1.165:8545")
    lp_waitd_toltal = waitd_token.GetBalance(pair_address)

    #user_address = "0x00049ad078e1339e279f9ea343a746c72d00fa6b"
    pair_balance = usdt_aitd_pair.balanceOf(user_address)

    abi_file = path + swap_abi["staking-rewards"]# '../abi/uniswap-v2-router02.json'    
    contract_abi  = json.load(open(abi_file))
    contract = json_contract_swap["stake-reward"]["AITD-USDT-2"]#"0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    usdt_aitd_stake2 = StakingRewards(contract_abi,contract, setting.g_setting.getRpcUrl())#"http://192.168.1.165:8545")
    stake_balance2 = usdt_aitd_stake2.balanceOf(user_address)
    stake2_total = usdt_aitd_stake2.totalSupply()
    rewards_balance2 = dataformat.FromWei(usdt_aitd_stake2.earned(user_address), 18)
    #stake1_usdt = (pair_balance + stake_balance1) * lp_usdt_toltal / total
    #stake1_aitd = (pair_balance + stake_balance1) * lp_waitd_toltal / total

    stake2_usdt = dataformat.FromWei((pair_balance + stake_balance2) * lp_usdt_toltal / total, 6)
    stake2_aitd = dataformat.FromWei((pair_balance + stake_balance2) * lp_waitd_toltal / total, 18)


    contract = json_contract_swap["stake-reward"]["AITD-USDT-3"]#"0xCf541edEF61700d454e7e0fD31aaA7BdaC7C3D39"
    usdt_aitd_stake3 = StakingRewards(contract_abi,contract, setting.g_setting.getRpcUrl())#"http://192.168.1.165:8545")
    stake_balance3 = usdt_aitd_stake3.balanceOf(user_address)
    stake3_total = usdt_aitd_stake3.totalSupply()
    rewards_balance3 = dataformat.FromWei(usdt_aitd_stake3.earned(user_address), 18)
    
    #stake1_usdt = (pair_balance + stake_balance1) * lp_usdt_toltal / total
    #stake1_aitd = (pair_balance + stake_balance1) * lp_waitd_toltal / total

    stake3_usdt = dataformat.FromWei((pair_balance + stake_balance3) * lp_usdt_toltal / total, 6)
    stake3_aitd = dataformat.FromWei((pair_balance + stake_balance3) * lp_waitd_toltal / total, 18)
    return stake2_usdt,stake2_aitd,stake_balance2,stake2_total,stake3_usdt,stake3_aitd,stake_balance3,stake3_total,rewards_balance2,rewards_balance3

    #return stake1_usdt,stake1_aitd,stake_balance1,stake1_total,stake2_usdt,stake2_aitd,stake_balance2,stake2_total
