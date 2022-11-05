// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
const hre = require('hardhat')

async function main() {
    // deploy address
    const deployer = await hre.ethers.getSigners();

    console.log("Deploying contracts with the account:", deployer[0].address);

    console.log("Account balance:", (await deployer[0].getBalance()).toString());


    // deploy erc1155 token contracts
    const InterviewTokenInstance = await hre.ethers.getContractFactory('InterviewErc1155Token');
    const Name = "Token";
    const Symbol = "Token";
    const initNotRevealedUri = "https://hei she bei jing tu"

    const InterviewTokenInstanceDeployer = await InterviewTokenInstance.deploy(initNotRevealedUri, Name, Symbol);
    await InterviewTokenInstanceDeployer.deployed();
    console.log('Interview Token address: ' + InterviewTokenInstanceDeployer.address);


}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })