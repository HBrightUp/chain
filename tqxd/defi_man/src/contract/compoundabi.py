import logging
from eth_utils import address
from eth_utils.address import to_checksum_address
from web3 import Web3
import web3
import json
import os
import re
import time
import datetime

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
    print(events)
    return events
              

class Comptroller(object):

    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        ckcontract_address = web3.Web3.toChecksumAddress(contract_address)
        self.contract = self.web3.eth.contract(address=ckcontract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        print(self.contract.all_functions())
        

class CErc20Delegator(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        self.contract = self.web3.eth.contract(address=contract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        print(self.contract.all_functions())

    def mint(self, privkey, amount):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        print(address)
        #chktoaddress = web3.Web3.toChecksumAddress(to_address)
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.mint( amount).buildTransaction({'gas':1000000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce})
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        print(txid.hex())
        return txid.hex()
    
    def getAccountSnapshot(self, address):
        chkaddress = web3.Web3.toChecksumAddress(address)
        result = self.contract.caller().getAccountSnapshot(chkaddress)
        print(result)
        return result

    def exchangeRateCurrent(self):
        result = self.contract.caller().exchangeRateCurrent()
        print(result)
        return result

    

class MultiPriceOracle(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        self.contract = self.web3.eth.contract(address=contract_address, abi =json_abi)
       
    def PrintAllFunctionEvent(self):
        print(self.contract.all_functions())
