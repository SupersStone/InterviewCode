package indexer

// EventBase event common struct
type EventBase struct {
	EventDefinitionID uint64 `gorm:"column:event_definition_id;type:bigint(20) unsigned;default:0;comment:事件定义表id;NOT NULL" json:"event_definition_id"`
	BlockHash         string `gorm:"column:block_hash;type:varchar(66);comment:区块hash;NOT NULL" json:"block_hash"`
	TransactionHash   string `gorm:"column:transaction_hash;type:varchar(66);comment:交易hash;NOT NULL" json:"transaction_hash"`
	BlockNumber       uint64 `gorm:"column:block_number;type:bigint(20) unsigned;default:0;comment:区块链高度;NOT NULL" json:"block_number"`
	TransactionIndex  uint64 `gorm:"column:transaction_index;type:bigint(20) unsigned;default:0;comment:交易编号;NOT NULL" json:"transaction_index"`
	LogIndex          uint64 `gorm:"column:log_index;type:bigint(20) unsigned;default:0;comment:日志编号;NOT NULL" json:"log_index"`
	BlockTimestamp    uint64 `gorm:"column:block_timestamp;type:bigint(20) unsigned;default:0;comment:区块时间戳;NOT NULL" json:"block_timestamp"`
}
