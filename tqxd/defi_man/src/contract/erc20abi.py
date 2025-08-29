
import logging
from eth_utils.address import to_checksum_address
from web3 import Web3
import web3
import json
import os
import re
import time
import datetime
from common import dataformat
"""
    An alternative Contract API.

    This call:

    > contract.caller({'from': eth.accounts[1], 'gas': 100000, ...}).add(2, 3)
    is equivalent to this call in the classic contract:
    > contract.functions.add(2, 3).call({'from': eth.accounts[1], 'gas': 100000, ...})

    Other options for invoking this class include:

    > contract.caller.add(2, 3)

    or

    > contract.caller().add(2, 3)

    or

    > contract.caller(transaction={'from': eth.accounts[1], 'gas': 100000, ...}).add(2, 3)
 """
class ERC20(object):
    
    def __init__(self, json_abi, contract_address, endpoint):
        if endpoint.find('http://') == -1:
                endpoint = 'http://' + endpoint
        self.web3 = Web3(Web3.HTTPProvider(endpoint))
        ckaddress = web3.Web3.toChecksumAddress(contract_address)
        self.contract = self.web3.eth.contract(address=ckaddress, abi =json_abi)
       
    def PrintAllFunction(self):
        print(self.contract.all_functions())

    def GetBalance(self, address):
        ckaddress = web3.Web3.toChecksumAddress(address)
        amount = self.contract.caller().balanceOf(ckaddress)
        return amount

    def info(self):
        try:
            name = self.contract.caller().name()
            symbol = self.contract.caller().symbol()
            decimal = self.contract.caller().decimals()
            total = self.contract.caller().totalSupply() 
            return name, symbol, decimal, total
        except ValueError as err:
            dataformat.logger.error(str(err))
            return "", "", "", ""
        except Exception as err:
            dataformat.logger.error(str(err))
            return "", "", "", ""
        except :
            dataformat.logger.error("unknow error")
            return "", "", "", ""
    def decimals(self):
        decimal = self.contract.caller().decimals()
        return decimal
        
    def Allowance(self, address, approve):    
        from_address = web3.Web3.toChecksumAddress(address)
        approve_address = web3.Web3.toChecksumAddress(approve)
        amount = self.contract.caller().allowance(from_address, approve_address)
        print(amount)
        return amount

    def Approve(self,privkey,to_address,amount):
         #<Function approve(address,uint256)>
         banlance = 2**256 -1
         if amount > 0:
             banlance = amount
         print(banlance)
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

    def transferOwnership(self, privkey, to_address):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        chktoaddress = web3.Web3.toChecksumAddress(to_address)
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.transferOwnership(chktoaddress).buildTransaction({'gas':600000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce})
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()

    def mint(self, privkey, amount):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.mint(amount).buildTransaction({'gas':600000,
                                                        'gasPrice':gasprice, 'from': address, 'nonce':nonce})
        signed_tx =  sendfrom.signTransaction(unsigned_tx)
        txid = self.web3.eth.sendRawTransaction(signed_tx.rawTransaction)
        return txid.hex()


    def burn(self, privkey, amount):
        pk = int(privkey, 16)
        sendfrom = web3.eth.Account.privateKeyToAccount(pk)
        privateKey = sendfrom._key_obj
        publicKey = privateKey.public_key
        address = publicKey.to_checksum_address()
        
        gasprice = self.web3.eth.gasPrice
        nonce = self.web3.eth.getTransactionCount(address, "pending")
        unsigned_tx = self.contract.functions.burn(amount).buildTransaction({'gas':600000,
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
 