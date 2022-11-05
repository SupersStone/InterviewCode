package contracts

import "dao-exchange/internal/models"

// SystemMetadataInfo 链系统信息表
type SystemMetadataInfo struct {
	models.Base
	ChainType      int    `gorm:"column:f_chain_type;type:tinyint(4);default:0;comment:链类型 0(EVM兼容链) 1(Solana) 2(Polkadot);NOT NULL" json:"chain_type"`
	ChainName      string `gorm:"column:f_chain_name;type:varchar(60);comment:链名称 Polygon Ethereum Solana Polkadot;NOT NULL" json:"chain_name"`
	ChainID        uint64 `gorm:"column:f_chain_id;type:int(11);default:0;comment:chain ID;NOT NULL" json:"chain_id"`
	DexContract    string `gorm:"column:f_dex_contract;type:varchar(256);comment:dex 合约地址;NOT NULL" json:"dex_contract"`
	NftContract    string `gorm:"column:f_nft_contract;type:varchar(256);comment:nft 合约地址;NOT NULL" json:"nft_contract"`
	DexContractAbi string `gorm:"column:f_dex_contract_abi;type:text;comment:dex 合约ABI信息" json:"dex_contract_abi"`
	NftContractAbi string `gorm:"column:f_nft_contract_abi;type:text;comment:nft 合约ABI信息" json:"nft_contract_abi"`
	DexType        int    `gorm:"column:f_dex_type;type:tinyint(4);default:1;comment:DEX类型 1(fixed_price) 2(auction)" json:"dex_type"`
	ExpirationTime int64  `gorm:"column:f_expiration_time;type:bigint(20);default:2592000;comment:默认订单过期时间，不指定默认三十天;NOT NULL" json:"expiration_time"`
}

// TableName table name
func (m *SystemMetadataInfo) TableName() string {
	return "system_metadata_info"
}
