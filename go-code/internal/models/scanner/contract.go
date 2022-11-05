package scanner

import "dao-exchange/internal/models"

// BlockScanContract care contract info
type BlockScanContract struct {
	models.Base
	Name        string `json:"name" gorm:"column:name"`
	Address     string `json:"address" gorm:"column:address"`
	Chain       string `json:"chain" gorm:"column:chain"`
	ChainID     int    `json:"chain_id" gorm:"column:chain_id"`
	StartHeight uint64 `json:"start_height" gorm:"column:start_height"`
	EndHeight   uint64 `json:"end_height" gorm:"column:end_height"`
	Status      int8   `json:"status" gorm:"column:status"`
	Abi         string `json:"abi" gorm:"column:abi"`
}
