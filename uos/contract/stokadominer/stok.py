#! /usr/bin/env python3

import os
import sys
import time
import configparser
import subprocess
import json


class global_info:
    rpclist = set()

    def __init__(self):
        self.rpclist.add("https://testrpc1.uosio.org:20580")
        self.rpclist.add("http://testrpc2.uosio.org:9000")
        self.rpclist.add("http://testrpc2.uosio.org:9007")




def run_special_command(cmd):
    return subprocess.getoutput(cmd)

def get_info(rpc):
    cmd = "cluos -u " + rpc + " get info"
    ###cmd = "curl --request POST --url " + rpc + "/v1/chain/get_info"
    result = run_special_command(cmd)
    return result

def get_head_block_num():
    good_rpc = get_good_rpc()
    getinfo = get_info(good_rpc)
    head_block_num = 0
    if getinfo is not None:
        try:
            json_getinfo = json.loads(getinfo)
            head_block_num = json_getinfo.get("head_block_num")
        except Exception as err:
            pass
    return head_block_num

def get_good_rpc():
    g = global_info()
    good_rpc = "rpc"
    for rpc in g.rpclist:
        time.sleep(0.1)
        getinfo = get_info(rpc)
        if getinfo is not None:
            try:
                json_getinfo = json.loads(getinfo)
                last_irreversible_block_num = json_getinfo.get("head_block_num")
                if last_irreversible_block_num:
                    good_rpc = rpc
                    break
                else:
                    continue
            except Exception as err:
                pass
    if good_rpc == "rpc":
        get_good_rpc()
    return good_rpc

def set_block(head_block_num):
    good_rpc = get_good_rpc()
    ###cluos  -u https://testrpc1.uosio.org:20580 push action stokadominer issue '["100"]'  -p stokadominer@active

    cmd = "cluos  -u " + good_rpc + " push action stokadominer issue " + "'[" + '"' + str(head_block_num) + '"'+ "]' " +"-p stokadominer@active"
    print(cmd)
    result = run_special_command(cmd)
    print(result)




def main():
    print("start program...")
    while True:
        time.sleep(0.1)
        block_num = get_head_block_num()
        if block_num > 0:
            set_block(block_num)







if __name__ == '__main__':
    main()