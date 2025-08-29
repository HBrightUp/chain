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

def createContract(privkey,abi,bytecode):
    w3 = Web3(Web3.HTTPProvider("http://127.0.0.1:8545"))
    pk = int(privkey, 16)
    sendfrom = web3.eth.Account.privateKeyToAccount(pk)
    privateKey = sendfrom._key_obj
    publicKey = privateKey.public_key
    address = publicKey.to_checksum_address()

    gasprice = w3.eth.gasPrice
    nonce = w3.eth.getTransactionCount(address, "pending")
    contract = w3.eth.contract(abi=abi, bytecode=bytecode)
    unsigned_tx = contract.constructor().buildTransaction({'gas':10000000,'gasPrice':gasprice, 'from': address, 'nonce':nonce})
    print(unsigned_tx)
    signed_tx =  sendfrom.signTransaction(unsigned_tx)
    print(signed_tx)
    txid = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
    print(txid.hex())



if __name__ == '__main__':
    abi_file = 'build/UnitrollerAdminStorage.abi'    
    contract_abi  = json.load(open(abi_file))
    bytecode_file = 'build/UnitrollerAdminStorage.bin'
    fd = open(bytecode_file)
    bytecode = '0x' + fd.read()
    
    