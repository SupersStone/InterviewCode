package indexer

import "dao-exchange/internal/models"

// TokenWorshipStatus token 所有记录
type TokenWorshipStatus struct {
	models.Base
	ChainID         string `gorm:"column:chain_id;type:varchar(20);comment:chain id;NOT NULL" json:"chain_id"`
	ContractAddress string `gorm:"column:contract_address;type:varchar(42);comment:合约地址;NOT NULL" json:"contract_address"`
	Score           string `gorm:"column:score;type:decimal(40,0);default:0;comment:用于判断先后;NOT NULL" json:"score"`
	EndTime         uint64 `gorm:"column:end_time;type:bigint(20) unsigned;default:0;comment:日志编号;NOT NULL" json:"end_time"`
	StartTime       uint64 `gorm:"column:start_time;type:bigint(20) unsigned;default:0;comment:区块时间戳;NOT NULL" json:"start_time"`
	TokenID         string `gorm:"column:token_id;type:varchar(66);comment:erc721 token id ;NOT NULL" json:"token_id"`
	Worship         int8   `gorm:"column:worship;type:tinyint(4);comment:worship;NOT NULL" json:"worship"`
}

// TableName table name
func (m *TokenWorshipStatus) TableName() string {
	return "token_worship_status"
}
