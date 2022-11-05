package dao

import (
	"context"

	"dao-exchange/internal/models/contracts"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// SystemMetadataDao interface of order dao
type SystemMetadataDao struct {
}

// QueryAll update one
func (d *SystemMetadataDao) QueryAll(ctx context.Context, db *gorm.DB) ([]contracts.SystemMetadataInfo, error) {
	tx := db.WithContext(ctx)
	metadatas := []contracts.SystemMetadataInfo{}
	res := tx.Model(&contracts.SystemMetadataInfo{}).Find(&metadatas)
	if res.Error != nil {
		return nil, res.Error
	}

	return metadatas, nil
}

// Insert update one
func (d *SystemMetadataDao) Insert(ctx context.Context, db *gorm.DB, metadata *contracts.SystemMetadataInfo) error {
	tx := db.WithContext(ctx)
	res := tx.Create(metadata)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}
