// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const fs = require('fs');

const ABI_GOVERMETNT_PATH = "./artifacts/contracts/goverment.sol/Goverment.json";

async function deployAllContract() {

 
  //deploy Goverment
  const govermentFactory = await hre.ethers.getContractFactory("Goverment");
  const TransparentUpgradeableProxy = await upgrades.deployProxy(govermentFactory, { initializer: 'initialize' });
  console.log("TransparentUpgradeableProxy address:", TransparentUpgradeableProxy.address);



  const TransparentUpgradeableProxy_address = TransparentUpgradeableProxy.address;

  
  sleep(20000);


  let goverment = await upgrades.erc1967.getImplementationAddress(TransparentUpgradeableProxy_address);
  console.log("goverment address: ", goverment);

  let adminProxy = await upgrades.erc1967.getAdminAddress(TransparentUpgradeableProxy_address);
  console.log("adminProxy address: ", adminProxy);

  console.log("Deploy all contract finished.");



}

function sleep(delay) {
  var start = (new Date()).getTime();
  while ((new Date()).getTime() - start < delay) {
      // 使用  continue 实现；
      continue; 
  }
}

async function main() {
  console.log("start running ...");

  await deployAllContract();

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
