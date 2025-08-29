// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const fs = require('fs');

async function deployStoneContract() {

  //deploy  TicketNFT
  const stoneNFTFactory = await hre.ethers.getContractFactory("StoneNFT");
  const stoneNFT = await stoneNFTFactory.deploy();
  await stoneNFT.deployed();
  console.log("stoneNFT address:", stoneNFT.address);

  console.log("Deploy stone contract finished.");


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

  await deployStoneContract();

}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
