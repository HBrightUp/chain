const {ethers} = require("ethers");
const provider = new ethers.getDefaultProvider()

async function main() {

let num = 1000;

for(let i = 0; i < num; i++) {
    let a = await provider.getStorageAt(
        //"0xa41397521877fbAE76DEcfa010Defceb4dF215fE",
        "0xB82964e00b1ABFD6dC71A1Cea80AAf31817FF966",
        i
    )
    console.log(a)
  }
}



main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
