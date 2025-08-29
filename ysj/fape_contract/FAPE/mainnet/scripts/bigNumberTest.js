async function main() {
    const hre = require("hardhat");
    const fs = require('fs');
    
    
    const ABI_STORAGEGOVER_PATH = "./artifacts/contracts/libraries/storageGover.sol/storageGover.json";
    const STORAGEADDRESS_ADDRESS = "0xf0593f5C308D2C4b908F66a53BB6152Ba99450df";

    //get instance of Goverment
    const provider = new ethers.getDefaultProvider();
    const storageGover_parsed = JSON.parse(fs.readFileSync(ABI_STORAGEGOVER_PATH));
    const storageGover = new ethers.Contract(STORAGEADDRESS_ADDRESS, storageGover_parsed.abi, provider);
    console.log(`storageGover instance : ${storageGover.address}`);

    const [owner] = await ethers.getSigners();
    trx = await storageGover.connect(owner).setAdmissionTicket(200000000000000000000);
    console.log("storageGover setAdmissionTicket hash: ", trx.hash);


}


main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });