

// todo : 线下构建默克尔树证明
const fs = require('fs');
const BN = require('bignumber.js');
const { MerkleTree } = require('merkletreejs');
const keccak256 = require('keccak256');
const ethers = require('ethers');

// 读取数据
let logs = fs.readFileSync('../bayc-transfer.csv').toString().split("\n");

logs = logs.map(item => {
    return item.split(",");
})

console.log(logs);

// 筛选数据
let address = [];
// logs.map(item =>{
//     // item[3] = token id
//     // itme[2] = owner
//     address[item[3]] = item[2];
// })
// 去掉表头数据
for ( let item of logs){
    if (item[2] && item[2] !== 'to_address'){
        address[item[3]] = item[2];
    }
}

console.log(address);



// 奖励

let AMOUNT = new BN(400).times(new BN(10).pow(18)).toString(10);
// 去重数据
let users = [];
address.map((item) => {
    users[item] = AMOUNT;
});

console.log(users);


// 数据筛选完毕，生成默克尔树 npm install merkletreejs  npm install keccak256
// 开始生成默克尔树
let addressLeaf = [];

// 第一步： 叶子节点数据
let leafs = [];
for ( let addressItem in users){
    let leaf = ethers.utils.solidityKeccak256(['address','uint256'], [addressItem,users[addressItem]]);
    leafs.push(leaf);
    addressLeaf[addressItem] = leaf;
}

// 第二步，生成树根
let tree = new MerkleTree(leafs, keccak256, {sort:true});

console.log("树根是：", tree.getHexRoot());


// 第三步；生成每个叶子的proof
let proofs = [];
leafs.map(item => {
    console.log("proof", tree.getHexProof(item));
    proofs[item] =  tree.getHexProof(item);
})


// 第四步： 将数据保存到本地
let result = [];
for (let i in addressLeaf){
    result.push([i,AMOUNT,proofs[addressLeaf[i]]]);
}

console.log(result);


// 写入本地
fs.writeFileSync('whitelist.json',JSON.stringify(result))