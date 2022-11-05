package dao

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// BaseInsertDao Base dao for indexer insert data
type BaseInsertDao interface {
	Insert(ctx context.Context, db *gorm.DB, event interface{}) error
}

// BaseDaoOpr base dao opr for insert one
// at that time, it's work for indexer
type BaseDaoOpr struct {
}

// Insert update one
func (d *BaseDaoOpr) Insert(ctx context.Context, db *gorm.DB, event interface{}) error {
	tx := db.WithContext(ctx)
	if event == nil {
		return nil
	}
	res := tx.Create(event)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}

// Insert update one
func Insert(ctx context.Context, db *gorm.DB, event interface{}) error {
	tx := db.WithContext(ctx)
	if event == nil {
		return nil
	}
	res := tx.Create(event)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}
