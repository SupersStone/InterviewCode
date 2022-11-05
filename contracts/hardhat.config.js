/** @type import('hardhat/config').HardhatUserConfig */

require('dotenv').config()

require('@openzeppelin/hardhat-upgrades')
require('@nomiclabs/hardhat-etherscan')
require('@nomiclabs/hardhat-waffle')
require('hardhat-gas-reporter')
require('solidity-coverage')
require('hardhat-contract-sizer')
require('hardhat-abi-exporter')


task('accounts', 'Prints the list of accounts', async (taskArgs, hre) => {
  const accounts = await hre.ethers.getSigners()

  for (const account of accounts) {
    console.log(account.address)
  }
})

// You need to export an object to set up your config
// Go to https://hardhat.org/config/ to learn more

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
module.exports = {
  solidity: {
    version: '0.8.13',
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      },
      outputSelection: {
        '*': {
          '*': ['storageLayout']
        }
      }
    }
  },
  defaultNetwork: "bsctest",
  networks: {
    hardhat: {
      initialBaseFeePerGas: 0 // workaround from https://github.com/sc-forks/solidity-coverage/issues/652#issuecomment-896330136 . Remove when that issue is closed.
    },
    bsctest: {
      url: "https://rpc.ankr.com/bsc_testnet_chapel",
      accounts:  [`b4ba724c1537e47e884570f4aafeef56a78ec1388af1fae2388a68f53f86f2c2`]
    },
    bsc: {
      url: process.env.BSC_MAIN_URL,
      accounts: process.env.PRIVATE_KEY !== undefined ? [process.env.PRIVATE_KEY] : []
    }
  },
  gasReporter: {
    enabled: process.env.REPORT_GAS !== undefined,
    currency: 'USD',
    gasPrice: 200,
    showTimeSpent: true,
    coinmarketcap: process.env.COINMARKETCAP_API
  },
  etherscan: {
    apiKey: process.env.BSCSCAN_API_KEY  // bsc
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: false
  },
  abiExporter: {
    path: './abi/',
    clear: true,
    flat: true,
    only: [],
    spacing: 2,
    pretty: false
  }
}