package indexer

import "dao-exchange/internal/models"

// Token721 token 所有记录
type Token721 struct {
	models.Base
	EventDefinitionID uint64 `gorm:"column:event_definition_id;type:bigint(20) unsigned;default:0;comment:事件定义表id;NOT NULL" json:"event_definition_id"`
	BlockNumber       uint64 `gorm:"column:block_number;type:bigint(20) unsigned;default:0;comment:区块链高度;NOT NULL" json:"block_number"`
	TransactionIndex  uint64 `gorm:"column:transaction_index;type:bigint(20) unsigned;default:0;comment:交易编号;NOT NULL" json:"transaction_index"`
	LogIndex          uint64 `gorm:"column:log_index;type:bigint(20) unsigned;default:0;comment:日志编号;NOT NULL" json:"log_index"`
	BlockTimestamp    uint64 `gorm:"column:block_timestamp;type:bigint(20) unsigned;default:0;comment:区块时间戳;NOT NULL" json:"block_timestamp"`
	TokenID           string `gorm:"column:token_id;type:varchar(66);comment:erc721 token id ;NOT NULL" json:"token_id"`
	Score             string `gorm:"column:score;type:decimal(40,0);default:0;comment:用于判断先后;NOT NULL" json:"score"`
	Owner             string `gorm:"column:owner;type:varchar(42);comment:owner address;NOT NULL" json:"owner"`
}

// TableName table name
func (m *Token721) TableName() string {
	return "token721"
}
