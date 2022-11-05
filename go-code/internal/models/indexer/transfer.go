package indexer

import (
	"dao-exchange/internal/models"
)

// Erc721TransferEvent erc721 transfer事件
type Erc721TransferEvent struct {
	models.Base
	EventBase
	From    string `gorm:"column:from;type:varchar(42);comment:from address;NOT NULL" json:"from"`
	To      string `gorm:"column:to;type:varchar(42);comment:to address;NOT NULL" json:"to"`
	TokenID string `gorm:"column:token_id;type:varchar(66);comment:token id;NOT NULL" json:"token_id"`
}

// TableName table name
func (m *Erc721TransferEvent) TableName() string {
	return "erc721_transfer_event"
}
