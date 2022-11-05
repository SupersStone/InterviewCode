package indexer

import (
	"dao-exchange/internal/models"
)

// FixedPriceMatchedEvent   固定价格订单成交事件(新商城使用)
type FixedPriceMatchedEvent struct {
	models.Base
	EventBase
	Maker            string `gorm:"column:maker;type:varchar(42);comment:订单maker address;NOT NULL" json:"maker"`
	Taker            string `gorm:"column:taker;type:varchar(42);comment:订单taker address;NOT NULL" json:"taker"`
	OrderHash        string `gorm:"column:order_hash;type:varchar(66);comment:make订单hash;NOT NULL" json:"make_order_digest"`
	OrderBytes       string `gorm:"column:order_bytes;type:text;comment:make订单eip712标准订单内容" json:"make_order_bytes"`
	RoyaltyRecipient string `gorm:"column:royalty_recipient;type:varchar(42);comment:版税接收人;NOT NULL" json:"royalty_recipient"`
	RoyaltyRate      string `gorm:"column:royalty_rate;type:varchar(66);comment:合约版税;NOT NULL" json:"royalty_rate"`
	StartTime        int64  `gorm:"column:start_time;type:bigint(20);default:0;comment:订单开始时间戳;NOT NULL" json:"start_time"`
	ExpireTime       int64  `gorm:"column:expire_time;type:bigint(20);default:0;comment:订单过期时间戳;NOT NULL" json:"expire_time"`
	MakerNonce       string `gorm:"column:maker_nonce;type:varchar(66);comment:dex合约用户nonce;NOT NULL" json:"maker_nonce"`
	TakerGetNft      int    `gorm:"column:taker_get_nft;type:tinyint(2);default:1;comment:是否是taker获取的nft,1-> 是, 0-> 否;NOT NULL" json:"taker_get_nft"`
	Nft              string `gorm:"column:nft;type:varchar(42);comment:nft 地址;NOT NULL" json:"nft"`
	Ft               string `gorm:"column:ft;type:varchar(42);comment:erc20 合约地址;NOT NULL" json:"ft"`
	NftID            string `gorm:"column:nft_id;type:varchar(66);default:0;comment:nft id;NOT NULL" json:"nft_id"`
	NftAmount        string `gorm:"column:nft_amount;type:varchar(66);comment:nft数量" json:"nft_amount"`
	FtAmount         string `gorm:"column:ft_amount;type:varchar(66);comment:付款erc20的数量" json:"ft_amount"`
}

// TableName table name
func (m *FixedPriceMatchedEvent) TableName() string {
	return "fixed_price_matched_event"
}
