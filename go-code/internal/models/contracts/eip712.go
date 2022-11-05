package contracts

import "dao-exchange/internal/models"

// DappDomain EIP712 Domain表
type DappDomain struct {
	models.Base
	Name              string `gorm:"column:f_name;type:varchar(40);comment:名字;NOT NULL" json:"name"`
	Version           string `gorm:"column:f_version;type:varchar(20);comment:版本;NOT NULL" json:"version"`
	ChainID           int    `gorm:"column:f_chain_id;type:int(11);default:0;comment:链ID;NOT NULL" json:"chain_id"`
	VerifyingContract string `gorm:"column:f_verifying_contract;type:varchar(256);comment:验证合约地址;NOT NULL" json:"verifying_contract"`
}

// TableName table name
func (m *DappDomain) TableName() string {
	return "dapp_domain"
}
