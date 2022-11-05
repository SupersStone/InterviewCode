package scanner

import "dao-exchange/internal/models"

// BlockScanHeight scan info
type BlockScanHeight struct {
	models.Base
	Height   uint64 `gorm:"column:height"`
	Chain    string `gorm:"column:chain"`
	TaskName string `gorm:"column:task_name"`
}
