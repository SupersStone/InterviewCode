package contracts

import "dao-exchange/internal/models"

// NftChainToken  支持的block chain 和代币信息
type NftChainToken struct {
	models.Base
	ChainID      int64  `gorm:"column:chain_id" db:"chain_id" json:"chain_id" form:"chain_id"`                     //  支持公链的id, 默认以太坊主网
	Name         string `gorm:"column:name" db:"name" json:"name" form:"name"`                                     //  支持公链名称
	TokenAddress string `gorm:"column:token_address" db:"token_address" json:"token_address" form:"token_address"` //  支付的token地址
	TokenName    string `gorm:"column:token_name" db:"token_name" json:"token_name" form:"token_name"`             //  支付的token名字
	TokenSymbol  string `gorm:"column:token_symbol" db:"token_symbol" json:"token_symbol" form:"token_symbol"`     //  支付的token简称
	Decimal      int64  `gorm:"column:decimal" db:"decimal" json:"decimal" form:"decimal"`                         //  支付合约代币精度
	Deleted      int64  `gorm:"column:deleted" db:"deleted" json:"deleted" form:"deleted"`                         //  逻辑删除,  0:未删除, 1:已删除
}

// TableName table name
func (NftChainToken) TableName() string {
	return "nft_chain_token"
}
