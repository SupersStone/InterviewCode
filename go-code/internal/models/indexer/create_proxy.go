package indexer

import "dao-exchange/internal/models"

// NftexProxyEvent 代理创建event
type NftexProxyEvent struct {
	models.Base
	EventBase
	UserAddress  string `gorm:"column:user_address;type:varchar(66);comment:用户地址;NOT NULL" json:"user_address"`
	ProxyAddress string `gorm:"column:proxy_address;type:varchar(66);comment:代理地址;NOT NULL" json:"proxy_address"`
}

// TableName table name
func (m *NftexProxyEvent) TableName() string {
	return "nftex_proxy_event"
}
