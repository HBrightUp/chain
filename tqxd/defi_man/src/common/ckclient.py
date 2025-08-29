from  clickhouse_driver import Client

client = Client(host='192.168.1.27',database = 'wallet_report_prod', user='default', password='123456')
def GetCkData(sql):
    print(sql)
    data = client.execute(sql)
    print(type(data))
    return data
