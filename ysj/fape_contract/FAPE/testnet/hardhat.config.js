require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
require('@openzeppelin/hardhat-upgrades');

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

  defaultNetwork: "polygontestnet",
  networks: {
    hardhatLocal: {
      url: "http://192.168.1.26:9090/",
      network_id: 31337,
      accounts: ["0xa992b126e10478be3a5b175c996b329107710a1a8fe3e57839c58e82f2251d1c"]
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

  },

};
