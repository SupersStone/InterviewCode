package indexer

import (
	"dao-exchange/internal/models"
)

// TokenWorship token 所有记录
type TokenWorship struct {
	models.Base
	EventBase
	ReleaseTimestamp uint64 `gorm:"column:release_timestamp;type:bigint(20) unsigned;default:0;comment:区块时间戳;NOT NULL" json:"release_timestamp"`
	NftAddress       string `gorm:"column:nft_address;type:varchar(66);comment:nft address ;NOT NULL" json:"nft_address"`
	TokenID          string `gorm:"column:token_id;type:varchar(66);comment:erc721 token id ;NOT NULL" json:"token_id"`
	Votary           string `gorm:"column:votary;type:varchar(66);comment:votary;NOT NULL" json:"votary"`
	Redeemer         string `gorm:"column:redeemer;type:varchar(42);comment:owner address;NOT NULL" json:"redeemer"`
	Type             int8   `gorm:"column:type;type:tinyint(4);comment:type 1 offer 2 redeeme;NOT NULL" json:"type"`
}

// TableName table name
func (m *TokenWorship) TableName() string {
	return "token_worship_history"
}
