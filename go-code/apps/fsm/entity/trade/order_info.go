package trade

import (
	"dao-exchange/internal/models"

	"github.com/shopspring/decimal"
)

// OrderInfo 订单状态表
type OrderInfo struct {
	models.Base
	OrderHash    string          `gorm:"column:order_hash;type:varchar(256);comment:挂单的哈希;NOT NULL" json:"order_hash"`
	Creator      string          `gorm:"column:creator;type:varchar(66);default:0;" json:"address"`
	TokenID      string          `gorm:"column:token_id;type:varchar(256);default:0;" json:"token_id"`
	OrderType    int             `gorm:"column:order_type;type:tinyint(4);default:0;" json:"order_type"`
	PaymentType  int             `gorm:"column:payment_type;type:tinyint(4);default:0;" json:"payment_type"`
	PaymentToken string          `gorm:"column:payment_token;type:varchar(66);comment:支付token合约地址;NOT NULL" json:"payment_token"`
	ChainID      uint64          `gorm:"column:chain_id;type:int(11);default:0;" json:"chain_id"`
	Price        decimal.Decimal `gorm:"column:price;type:decimal(40,20);default:0;" json:"price"`
	Status       int             `gorm:"column:status;type:tinyint(4);default:0;" json:"status"`
}

// TableName table name
func (m *OrderInfo) TableName() string {
	// TODO change name
	return "order_info"
}

// NFTOrderHistory 订单历史记录表
type NFTOrderHistory struct {
	models.Base
	TxHash       string          `gorm:"column:tx_hash;type:varchar(256);comment:交易哈希;NOT NULL" json:"tx_hash"`
	From         string          `gorm:"column:from;type:varchar(66);default:0;" json:"from"`
	To           string          `gorm:"column:to;type:varchar(66);default:0;" json:"to"`
	Amount       decimal.Decimal `gorm:"column:amount;type:decimal(40,0);default:0;" json:"amount"`
	TokenID      string          `gorm:"column:token_id;type:varchar(256);default:0;" json:"token_id"`
	EventType    int             `gorm:"column:order_type;type:tinyint(4);default:0;" json:"event_type"`
	PaymentToken string          `gorm:"column:payment_token;type:varchar(66);comment:支付token合约地址;NOT NULL" json:"payment_token"`
	ChainID      uint64          `gorm:"column:chain_id;type:int(11);default:0;" json:"chain_id"`
	Price        decimal.Decimal `gorm:"column:price;type:decimal(40,20);default:0;" json:"price"`
	Status       int             `gorm:"column:status;type:tinyint(4);default:0;" json:"status"`
	Timestamp    int64           `gorm:"column:timestamp;type:bigint(20);default:0;" json:"timestamp"`
}

// TableName table name
func (m *NFTOrderHistory) TableName() string {
	// TODO change name
	return "nft_order_history"
}
