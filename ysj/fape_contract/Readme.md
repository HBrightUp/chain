1.所有智能合约的部署信息可以参考<<智能合约部署信息>>；
2.大部分智能合约使用hardhat进行编译与部署，hardhat的本地使用方法可以参考以下网页：
    https://hardhat.org/hardhat-runner/docs/getting-started

3. 在安装harhdat中需要安装的第三方开库或者插件可根据编译的报错信息一步一步安装完成即可； 
4. 项目的配置文件hardhat.config.js中部署所涉及的私匙需要自行重新提供；
5. 当前目录里有战斗猿的各种智能合约，在FAPE目录，另外一个发放MEE代币的合约在MEE目录； 
6. 可升级合约的部署与升级，逻辑合约Goverment.sol文件

a） 下载项目
git clone http://git.metahori.com/blockchain/fapecontract

b）进入test目录
cd fapecontract/FAPE/testnet

c)安装hardhat工具
npm install --save-dev hardhat

d) 编译
npx hardhat compile
备注：如果中间遇到缺少js库报错，使用npm安装即可；

e) 合约部署
在 script目录下有对应的脚本，可使用命令进行部署
npx hardhat run ./scropt/deploy_v1.js

f)对于已经上线的可升级的合约，可以使用以下命令进行升级; 
npx hardhat run ./scropt/upgrade.js




