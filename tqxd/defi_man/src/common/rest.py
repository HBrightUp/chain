#!/usr/bin/python3
import requests
proxy  = '192.168.1.34:1080'
proxies = {
    "http": "http://%(proxy)s/" % {'proxy':proxy},
    "https": "https://%(proxy)s/" % {'proxy':proxy}
}
r = requests.get(url='https://gate.kickex.com/api/v1/market/pairs')
print(r)