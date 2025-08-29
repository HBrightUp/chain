async function main() {
    // 这里的地址为前面部署的代理合约地址
    const proxyAddress = '0x0901aa37E5Dd4ae22191cCf89f44cA79EcF5B49A';

    const goverment = await ethers.getContractFactory("Goverment");
    console.log("Preparing upgrade...");
    // 升级合约
    let trx = await upgrades.upgradeProxy(proxyAddress, goverment);

    console.log("trx hash: ", trx.deployTransaction.hash);
}


main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });