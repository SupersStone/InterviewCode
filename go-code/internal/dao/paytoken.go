package dao

import (
	"context"
	"dao-exchange/internal/models/contracts"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// PayTokenDao interface of order dao
type PayTokenDao struct {
}

// QueryAll update one
func (d *PayTokenDao) QueryAll(ctx context.Context, db *gorm.DB) ([]contracts.NftChainToken, error) {
	tx := db.WithContext(ctx)
	events := []contracts.NftChainToken{}
	res := tx.Model(&contracts.NftChainToken{}).Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

// Insert update one
func (d *PayTokenDao) Insert(ctx context.Context, db *gorm.DB, payToken *contracts.NftChainToken) error {
	tx := db.WithContext(ctx)
	res := tx.Create(payToken)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}
