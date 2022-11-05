# 生成ERC20的事件代码
solc --abi  erc20.sol --overwrite -o ./
abigen --abi ERC20.abi --pkg erc20 --out erc20.go


# 生成NFT DEX 的事件代码
solc --abi  order_event.sol --overwrite -o ./
abigen --abi NFTOrderEvent.abi --pkg nft --out order_event.go

# 生成NFT Worship 的事件代码
solc --abi  worship.sol --overwrite -o ./
abigen --abi NFTWorshipEvent.abi --pkg nft --out worship.go