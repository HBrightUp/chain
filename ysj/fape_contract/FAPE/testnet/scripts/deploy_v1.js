// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require("hardhat");
const fs = require('fs');

const dapp = "0x1602eBb54fF8c513006085B9Ce2d6C21EBccc7CE";
const ABI_GOVERMETNT_PATH = "./artifacts/contracts/goverment.sol/Goverment.json";

async function deployAllContract() {

  //deploy TokenFAPE
  const FAPEFactory = await hre.ethers.getContractFactory("TokenFAPE");
  const tokenFAPE = await FAPEFactory.deploy();
  await tokenFAPE.deployed();
  console.log("TokenFAPE address:", tokenFAPE.address);


//deploy  TicketNFT
  const ticketNFTFactory = await hre.ethers.getContractFactory("TicketNFT");
  const ticketNFT = await ticketNFTFactory.deploy();
  await ticketNFT.deployed();
  console.log("ticketNFT address:", ticketNFT.address);

  //deploy Goverment
  const govermentFactory = await hre.ethers.getContractFactory("Goverment");
  const TransparentUpgradeableProxy = await upgrades.deployProxy(govermentFactory, { initializer: 'initialize' });
  console.log("TransparentUpgradeableProxy address:", TransparentUpgradeableProxy.address);

  //const tokenUSDT_address = usdt.address;
    //const tokenFAPE_address = "0xb7F36019ce25E047833b8087CfB9d73Eea36C33c";
  const tokenFAPE_address = tokenFAPE.address;
  const tokenUSDT_address = "0x5C2346C83d156B74a4366E3F8192ed609EA90f51";
  const ticketNFT_address = ticketNFT.address;
  const TransparentUpgradeableProxy_address = TransparentUpgradeableProxy.address;

  
  sleep(20000);


  let goverment = await upgrades.erc1967.getImplementationAddress(TransparentUpgradeableProxy_address);
  console.log("goverment address: ", goverment);

  let adminProxy = await upgrades.erc1967.getAdminAddress(TransparentUpgradeableProxy_address);
  console.log("adminProxy address: ", adminProxy);

  console.log("Deploy all contract finished.");

  //get instance of Goverment
  const provider = new ethers.getDefaultProvider();
  const goverment_parsed = JSON.parse(fs.readFileSync(ABI_GOVERMETNT_PATH));
  let gover = new ethers.Contract(TransparentUpgradeableProxy_address, goverment_parsed.abi, provider);
  console.log(`gover instance : ${gover.address}`);

  const [owner] = await ethers.getSigners();

  //gover init
  trx = await gover.connect(owner).setTokenFAPE(tokenFAPE_address);
  console.log("goverment setTokenFAPE hash: ", trx.hash);

  trx = await gover.connect(owner).setTicketNFT(ticketNFT_address);
  console.log("goverment setTicketNFT hash: ", trx.hash);

  //ticketNFT init
  trx = await ticketNFT.connect(owner).setProxy(TransparentUpgradeableProxy_address);
  console.log("ticketNFT setProxy hash: ", trx.hash);

  console.log("Init all contract finished.");

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
