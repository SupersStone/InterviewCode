package indexer

import "dao-exchange/internal/models"

// NftexAllOrderCanceledEvent   取消全部订单事件
type NftexAllOrderCanceledEvent struct {
	models.Base
	EventBase
	Maker string `gorm:"column:maker;type:varchar(42);comment:订单maker address;NOT NULL" json:"maker"`
	Nonce int64  `gorm:"column:nonce;type:int(11);comment:dex合约用户nonce;NOT NULL" json:"nonce"`
}

// TableName table name
func (m *NftexAllOrderCanceledEvent) TableName() string {
	return "nftex_all_order_canceled_event"
}
