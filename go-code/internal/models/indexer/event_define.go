package indexer

import (
	"dao-exchange/internal/models"
	"fmt"
)

// EventDefinition 事件定义
type EventDefinition struct {
	models.Base
	ContractAddress string `gorm:"column:contract_address;type:varchar(42);comment:合约地址;NOT NULL" json:"contract_address"`
	EventSignature  string `gorm:"column:event_signature;type:varchar(66);comment:事件签名;NOT NULL" json:"event_signature"`
	ContractType    string `gorm:"column:contract_type;type:varchar(20);comment:合约类型;NOT NULL" json:"contract_type"`
	EventName       string `gorm:"column:event_name;type:varchar(30);comment:事件方法名;NOT NULL" json:"event_name"`
	ChainID         string `gorm:"column:chain_id;type:varchar(20);comment:chain id;NOT NULL" json:"chain_id"`
	ChainName       string `gorm:"column:chain_name;type:varchar(50);comment:链名;NOT NULL" json:"chain_name"`
}

// TableName table name
func (m *EventDefinition) TableName() string {
	return "event_definition"
}

// CacheKey cache key
func (m *EventDefinition) CacheKey() string {
	return fmt.Sprintf("%s-%s", m.ContractType, m.EventName)
}
