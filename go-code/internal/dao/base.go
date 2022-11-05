package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Insert update one
func Insert(ctx context.Context, db *gorm.DB, data interface{}) error {
	tx := db.WithContext(ctx)
	if data == nil {
		return nil
	}
	res := tx.Create(data)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}
