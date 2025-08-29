# Basic Sample Hardhat Project

This project demonstrates a basic Hardhat use case. It comes with a sample contract, a test for that contract, a sample script that deploys that contract, and an example of a task implementation, which simply lists the available accounts.

Try running some of the following tasks:

```shell
npx hardhat accounts
npx hardhat compile
npx hardhat clean
npx hardhat test
npx hardhat node
node scripts/sample-script.js
npx hardhat help
```


/****************************  项目说明 *********************************

//airDrop使用说明
1. 获取无聊猿的 10000 个用户地址
npx ts-node ./getAPEAddress.ts  > APEAddress.md

2.空投脚本 
修改程序的运行参数后运行以下脚本
npx ts-node ./airdrop.ts > dropresult.md

//contract使用说明
1. 部署合约到本地
npx hardhat run ./scripts/deploy_v2.js --network hardhat

2.测试脚本
