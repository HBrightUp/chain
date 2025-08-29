import  ethers  from 'ethers';
import  BigNumber  from 'bignumber.js';
const fs = require('fs');
const common = require('./common');

async function main() {

  const provider = new ethers.providers.JsonRpcProvider(common.AITD_MAINNET_RPC);
  const invite_parsed = JSON.parse(fs.readFileSync(common.ABI_INVITE_FILE_PATH));
  const inviteStaket = new ethers.Contract(common.INVITE_STAKE_CONTRACT_ADDRESS, invite_parsed.abi, provider);
  console.log(`inviteStaket deployed to: ${inviteStaket.address}`);

  let compAddress = new Set<string>();
  getCompangAddress(compAddress);
  console.log("Amount of company account: ", compAddress.size);

  const userOfLength = await inviteStaket.userLength();
  console.log("Amount of account in inviteStake contract: ", userOfLength.toString());

  const totalAmount = await inviteStaket.totalAmount();

  let userStakeTotal:BigNumber = new BigNumber(0);
  let userInfo:[Boolean, string, string, string, string, string];
  let compStakeMap = new Map();

  let curItem:number = 0;
  let curAddress:string;

  for (var addr of compAddress) {
    curAddress = addr.toString().trim();
    console.log("addr: ", curAddress, "  on processing: ", ++curItem);

    if (!ethers.utils.isAddress(curAddress)) {
      console.log("Warning(involid address): ", curAddress);
      continue;
    }

    let userId:number = await inviteStaket.userOfPid(curAddress);
    if (userId  == 0) {
      continue;
    }

    userInfo = await inviteStaket.userInfoSet(userId - 1); 
    compStakeMap.set(curAddress, userInfo[1]); 
    userStakeTotal =  userStakeTotal.plus(new BigNumber(userInfo[1].toString()));
  }

  const compPrecent = (userStakeTotal.div(new BigNumber(totalAmount.toString())  ).toNumber() * 100).toString() + `%`;

  console.log("Stake amount of company: ", userStakeTotal.toString()); 
  console.log("Stake amount of all user: ", totalAmount.toString());
  console.log(`Stake precenage of company: ${compPrecent}`);

  let reportHear = (new Date(Date.parse(new Date().toString()))).toString();
  reportHear += `\nCompany aitd stake information: ${userStakeTotal.toString()}`
  reportHear += `\nAll aitd stake information: ${totalAmount} `
  reportHear +=  "\nThe precentage of Company: " + compPrecent;
  reportHear += "\nAll company address amount: " + compStakeMap.size.toString();
  reportHear += "\n\nStake informat of all company address as follows: \n";

  fs.writeFileSync(common.REPORT_FILE_PATH, reportHear, err => {
    if (err) {
        console.log("save information failed!");
      return 
    }
  })
  
  let userStakeInfo:string;
  for( let [ key, value ] of compStakeMap ) {
    userStakeInfo = key.toString() +" : " + value.toString() + "\n";
    fs.writeFileSync(common.REPORT_FILE_PATH, userStakeInfo, { flag: 'a+' }, err => {
      if (err) {
        console.log("save userStakeInfo failed!");
        return 
      }
    })
  }

  console.log("The statistics of company stake information finished, details for the report file(stakeReport.md) on current directory.")
}

async function getCompangAddress(companyAddress: Set<string>) {
  const data = fs.readFileSync(common.ADDRESS_COMPANG_PATH, 'UTF-8');
    const lines = data.split(/\r?\n/);
    lines.forEach((line) => {
      companyAddress.add(line);
    });
    return ;
}

main( )
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
});
