require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
require('@openzeppelin/hardhat-upgrades');

// This is a sample Hardhat task. To learn how to create your own go to
// https://hardhat.org/guides/create-task.html
task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners();

  for (const account of accounts) {
    console.log(account.address);
  }
});

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */

 module.exports = {
  solidity: {
    version: '0.8.4',
    settings: {
        optimizer: {
            enabled: true,
            runs: 9999,
        },
    },
  },

  defaultNetwork: "polygonMainnet",
  networks: {
    hardhat: {
    },

    bsctestnet: {
      url: "https://data-seed-prebsc-1-s3.binance.org:8545",
      network_id: 0x61,
      accounts: ["0xa992b126e10478be3a5b175c996b329107710a1a8fe3e57839c58e82f2251d1c"]
    },
    polygontestnet: {
      url: "https://matic-testnet-archive-rpc.bwarelabs.com",
      accounts: ["0xa992b126e10478be3a5b175c996b329107710a1a8fe3e57839c58e82f2251d1c"],
    },
    polygonMainnet: {
      url: "https://polygon-rpc.com",
      accounts: ["0xa992b126e10478be3a5b175c996b329107710a1a8fe3e57839c58e82f2251d1c"],
    },
  },
  etherscan: {
    apiKey: {
      polygon: ""
    }
  },

};
