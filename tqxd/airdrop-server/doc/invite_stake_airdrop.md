# airdrop服务文档



## 设置空投数量的上限下限

### URL

```
127.0.0.1:8080/invite_stake_airdrop/set_limit
```



### 请求方式

```
POSt
```



### 请求参数

| 参数名称 | 参数类型 | 是否必须 | 参数描述 |
| -------- | -------- | -------- | -------- |
| upper    | string   | 是       | 空投上限 |
| lower    | string   | 是       | 空投下限 |



### 请求示例(Postman)

```
请求方式：POST  

URL: 127.0.0.1:8080/invite_stake_airdrop/set_limit

参数(选择Body, 再选择raw)：
{
    "upper": "400000000.0",
    "lower": "2.1"
}
```



### 返回参数

| 参数名称 | 参数类型 | 参数描述          |
| -------- | -------- | ----------------- |
| message  | string   | 消息描述          |
| status   | int      | 0: 失败， 1: 成功 |



### 响应示例

```
{
    "message": "success",
    "status": 1
}
```



### 异常示例

```
{
    "message": "Key: 'AirdropLimit.Upper' Error:Field validation for 'Upper' failed on the 'required' tag",
    "status": 0
}
```



## 获取空投数量的上限下限

### URL

```
127.0.0.1:8080/invite_stake_airdrop/get_limit
```



### 请求方式

```
GET
```



### 请求示例(Postman)

```
请求方式：GET 

URL: 127.0.0.1:8080/invite_stake_airdrop/get_limit

无参数

```



### 返回参数

| 参数名称 | 参数类型 | 参数描述          |
| -------- | -------- | ----------------- |
| message  | string   | 消息描述          |
| status   | int      | 0: 失败， 1: 成功 |
| lower    | string   | 下限              |
| upper    | string   | 上限              |



### 响应示例

```
{
    "lower": "2.1",
    "message": "success",
    "status": 1,
    "upper": "40000000000000000000"
}
```



## 空投

### URL

```
127.0.0.1:8080/invite_stake_airdrop/airdrop
```



### 请求方式

```
POST
```



### 请求参数

| 参数名称 | 参数类型 | 是否必须 | 参数描述 |
| -------- | -------- | -------- | -------- |
| address  | string   | 是       | 空投地址 |
| value    | string   | 是       | 空投数量 |



### 请求示例(Postman)

```
请求方式：POST  

URL: 127.0.0.1:8080/invite_stake_airdrop/airdrop

参数(选择Body, 再选择raw)：
{
    "address" : "0xee31e38007D819E00a386fa4308F42c13871D55D",
    "value" : "20.1"
}
```



### 返回参数

| 参数名称 | 参数类型 | 参数描述          |
| -------- | -------- | ----------------- |
| message  | string   | 消息描述          |
| status   | int      | 0: 失败， 1: 成功 |
| hash     | string   | hash              |



### 响应示例

```
{
    "hash": "0xf23095d611d62200e26ed696c269e7f66574d1ad2630e8ec8dcfa105cd46e18d",
    "message": "success",
    "status": 1
}
```



### 异常示例

```
{
    "message": "invalid address",
    "status": 0
}
```



## 获取空投状态

### URL

```
127.0.0.1:8080/invite_stake_airdrop/get_tx_status
```



### 请求方式

```
GET
```



### 请求参数

| 参数名称 | 参数类型 | 是否必须 | 参数描述 |
| -------- | -------- | -------- | -------- |
| hash     | string   | 是       | tx hash  |



### 请求示例(Postman)

```
请求方式：POST  

URL: 127.0.0.1:8080/invite_stake_airdrop/get_tx_status

参数(选择Body, 再选择raw)：
{
    "hash" : "0xf23095d611d62200e26ed696c269e7f66574d1ad2630e8ec8dcfa105cd46e18d"
}
```



### 返回参数

| 参数名称 | 参数类型 | 参数描述          |
| -------- | -------- | ----------------- |
| message  | string   | 消息描述          |
| status   | int      | 0: 失败， 1: 成功 |
| hash     | string   | hash              |



### 响应示例

```
{
    "hash": "0xf23095d611d62200e26ed696c269e7f66574d1ad2630e8ec8dcfa105cd46e18d",
    "message": "success",
    "status": 1
}
```



### 异常示例

```
{
    "message": "invalid hash",
    "status": 0
}
```

