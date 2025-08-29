const { expect } = require("chai");
const { ethers } = require("hardhat");
const fs = require('fs');

const TransparentUpgradeableProxy_address = "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707";
const ABI_GOVERMETNT_PATH = "./artifacts/contracts/goverment.sol/Goverment.json";
const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000";

async function getGovermentInstance() {
  let provider = new ethers.getDefaultProvider();
  const goverment_parsed = JSON.parse(fs.readFileSync(ABI_GOVERMETNT_PATH));
  let gover = new ethers.Contract(TransparentUpgradeableProxy_address, goverment_parsed.abi, provider);
  console.log(gover);
  return gover;
}



describe("FAPE project test", function() {
  it("contract init test", async function() {
    /*
    let goverment = await upgrades.erc1967.getImplementationAddress(TransparentUpgradeableProxy_address);
    console.log("goverment address: ", goverment);

    let adminProxy = await upgrades.erc1967.getAdminAddress(TransparentUpgradeableProxy_address);
    console.log("adminProxy address: ", adminProxy);
  
    let provider = new ethers.getDefaultProvider();
    const goverment_parsed = JSON.parse(fs.readFileSync(ABI_GOVERMETNT_PATH));
    let gover = new ethers.Contract(TransparentUpgradeableProxy_address, goverment_parsed.abi, provider);
    console.log(`gover instance : ${gover.address}`);
    */

    let gover = await getGovermentInstance();
    const [owner] = await ethers.getSigners();
    console.log(await gover.connect(owner).tokenALC())

    //console.log(gover);
    //expect(await gover.connect(owner).tokenALC()).to.be.not.equal("0x0");
    //expect(await gover.connect(owner).ticketNFT()).to.be.not.equal("0x0");
    //expect(await gover.connect(owner).dapp()).to.be.not.equal("0x0");
    //expect(await gover.connect(owner).tokenALC()).to.be.not.equal("0x0");
    //expect(await gover.connect(owner).getStorageGover()).to.be.not.equal("0x0");

    //expect(await gover.tokenALC()).to.be.not.equal("0x0");

  });

});



describe("bindSuper test", function() {

  
  it("Should bind root user node at first time.", async function() {
    console.log("bindSuper test.");
/*
    let gover =  await getGovermentInstance();
    let dapp = await  ethers.provider.getSigner(1);
    
    const givenUserAddress = "0x049d4aC96A4AC24fAFD9C574951D0ffB9A594170";
    const givenSuperiorAddress = "0xA1bDf8C186B49Db485Ce3B2B0aD6eb9E8b8D1d23";
    expect(await gover.connect(dapp).getSuperior(givenUserAddress)).to.be.equal(ZERO_ADDRESS);
    let trx = await gover.connect(dapp).bindSuperior(  givenUserAddress, givenSuperiorAddress);
    console.log("bindSuperior hash: ", trx.hash);
    expect(await gover.connect(dapp).getSuperior(givenUserAddress)).to.equal(givenSuperiorAddress);

    */
  });

});