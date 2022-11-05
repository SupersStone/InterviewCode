package dao

import (
	"context"
	"reflect"

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

// Insert insert one
func (d *Dao) Insert(ctx context.Context, info interface{}) error {
	tx := d.db.WithContext(ctx)
	res := tx.Create(info)
	if res.Error != nil {
		return res.Error
	}

	infoLen := 1
	v := reflect.ValueOf(info)
	if v.Kind() == reflect.Slice {
		infoLen = v.Len()
	}

	if res.RowsAffected != int64(infoLen) {
		return errors.Errorf("rows affected must %d, but got: %d", infoLen, res.RowsAffected)
	}

	return nil
}

// UpdateOne update one
func (d *Dao) UpdateOne(ctx context.Context, wheres interface{}, update interface{}) (err error) {
	tx := d.db.WithContext(ctx)
	res := tx.Where(wheres).Updates(update)
	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return
}

// UpdatesByIds update many
func (d *Dao) UpdatesByIds(ctx context.Context, models interface{}, ids []int64, update interface{}) (err error) {
	tx := d.db.WithContext(ctx)
	res := tx.Model(models).Where("id IN ?", ids).Updates(update)
	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected != int64(len(ids)) {
		return errors.Errorf("rows affected must %d, but got: %d", len(ids), res.RowsAffected)
	}

	return
}
