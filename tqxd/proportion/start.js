const Web3 = require("web3");
const fs = require('fs');
let config = require('./config');

let aitd_web3_rops = new Web3(new Web3.providers.HttpProvider(config.AITD_RPC_URL))

const ABI_PAIRS_FILE_PATH = "./abi/UniswapV2Pair.json"
var pairs_parsed = JSON.parse(fs.readFileSync(ABI_PAIRS_FILE_PATH));
let pairs_contract = new aitd_web3_rops.eth.Contract(pairs_parsed.abi, config.PAIRS_CONTRACT_ADDRESS);

const ABI_STAKE_FILE_PATH = "./abi/stakeReward.json"
var stake_parsed = JSON.parse(fs.readFileSync(ABI_STAKE_FILE_PATH));
let stake2_contract = new aitd_web3_rops.eth.Contract(stake_parsed.abi, config.STAKE_CONTRACT_ADDRESS2);
let stake3_contract = new aitd_web3_rops.eth.Contract(stake_parsed.abi, config.STAKE_CONTRACT_ADDRESS3);

async function getBalanceOf() {
    let sum = 0;
    for (let i = 0; i < config.ADDRESS_LIST_POOl2.length; i++) {
        const txParams = {
            to: config.STAKE_CONTRACT_ADDRESS2,
            data: stake2_contract.methods.balanceOf(config.ADDRESS_LIST_POOl2[i]).encodeABI()
        }
        let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
        let amountInt = parseInt(ret, 16);
        // if (amountInt > 0 )
        // {
        //     console.log(config.ADDRESS_LIST_POOl2[i], amountInt);
        // }
        
        sum += amountInt;
    }
    console.log(sum);
    return sum;
}

async function totalSupply() {
    const txParams = {
        to: config.PAIRS_CONTRACT_ADDRESS,
        data: pairs_contract.methods.totalSupply().encodeABI()
    }
    let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
    let amountInt = parseInt(ret, 16);
    // console.log(amountInt);
    return amountInt;
}

async function getReserves() {
    const txParams = {
        to: config.PAIRS_CONTRACT_ADDRESS,
        data: pairs_contract.methods.getReserves().encodeABI()
    }
    let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
    if (ret.length > 100) {
        let tokeA = ret.substr(0, 66);
        let tokenAInt = parseInt(tokeA, 16);
        // console.log(tokenAInt);
        let tokeB = ret.substr(ret.length - 128, 64);
        let tokenBInt = parseInt(tokeB, 16);
        return {tokenAInt, tokenBInt};
        // console.log(tokenBInt);
    } 
    // console.log(ret);
}

async function pools2() {
    let sum = 0;
    for (let i = 0; i < config.ADDRESS_LIST_POOl2.length; i++) {
        const txParams = {
            to: config.STAKE_CONTRACT_ADDRESS2,
            data: stake2_contract.methods.balanceOf(config.ADDRESS_LIST_POOl2[i]).encodeABI()
        }
        let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
        let amountInt = parseInt(ret, 16);
        
        sum += amountInt;
    }
    // console.log(sum);

    const txParams = {
        to: config.STAKE_CONTRACT_ADDRESS2,
        data: stake2_contract.methods.totalSupply().encodeABI()
    }
    let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
    let amountInt = parseInt(ret, 16);
    // console.log(amountInt);
    let ratio = sum / amountInt;

    console.log('pool2 rate: ', ratio);

    let total = await totalSupply();
    let t = await getReserves();
    let usdtAmount = (sum * t.tokenAInt) / total / 1000000;
    let aitdAmount = (sum * t.tokenBInt) / total / 1000000000000000000;
    console.log('USDT Amount: ', usdtAmount, 'AITD Amount: ', aitdAmount);
}

async function pools3() {
    let sum = 0;
    for (let i = 0; i < config.ADDRESS_LIST_POOl3.length; i++) {
        const txParams = {
            to: config.STAKE_CONTRACT_ADDRESS3,
            data: stake3_contract.methods.balanceOf(config.ADDRESS_LIST_POOl3[i]).encodeABI()
        }
        let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
        let amountInt = parseInt(ret, 16);
        sum += amountInt;
    }
    // console.log(sum);

    const txParams = {
        to: config.STAKE_CONTRACT_ADDRESS3,
        data: stake3_contract.methods.totalSupply().encodeABI()
    }
    let ret = await aitd_web3_rops.eth.call(txParams, 'latest');
    let amountInt = parseInt(ret, 16);
    // console.log(amountInt);

    let ratio = sum / amountInt;

    console.log('pool3 rate: ', ratio);

    let total = await totalSupply();
    let t = await getReserves();
    let usdtAmount = (sum * t.tokenAInt) / total / 1000000;
    let aitdAmount = (sum * t.tokenBInt) / total / 1000000000000000000;
    console.log('USDT Amount: ', usdtAmount, 'AITD Amount: ', aitdAmount);
    return ret;
}

async function start() {
    await pools2();
    await pools3();
    // await getBalanceOf();
}

start()

