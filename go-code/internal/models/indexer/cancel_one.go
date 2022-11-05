package indexer

import "dao-exchange/internal/models"

// NftexOrderCanceledEvent   取消订单事件
type NftexOrderCanceledEvent struct {
	models.Base
	EventBase
	Maker     string `gorm:"column:maker;type:varchar(42);comment:订单maker address;NOT NULL" json:"maker"`
	OrderHash string `gorm:"column:order_hash;type:varchar(66);comment:订单hash;NOT NULL" json:"order_hash"`
}

// TableName table name
func (m *NftexOrderCanceledEvent) TableName() string {
	return "nftex_order_canceled_event"
}
