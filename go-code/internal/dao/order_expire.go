package dao

import (
	"context"
	"dao-exchange/apis/model"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// PayTokenDao interface of order dao
type OrderDao struct {
}

// QueryAll update one
func (d *OrderDao) QueryAllExpireOrder(ctx context.Context, db *gorm.DB) ([]model.NftOrder, error) {
	tx := db.WithContext(ctx)
	events := []model.NftOrder{}
	// 满足条件的查询出来,查询出当前所有挂单的数据
	res := tx.Where(map[string]interface{}{"status": 1, "taker_get_nft": 1}).Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

// Update one
func (d *OrderDao) UpdateExpireOrder(ctx context.Context, db *gorm.DB, order model.NftOrder, id uint64) error {
	tx := db.WithContext(ctx)
	res := tx.Where(&model.NftOrder{Id: id}).Updates(order)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}
	return nil
}

// Insert update one
func (d *OrderDao) InsertHistoryOrder(ctx context.Context, db *gorm.DB, order *model.NftOperationHistory) error {
	tx := db.WithContext(ctx)
	res := tx.Create(order)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}
	fmt.Println("成功插入历史数据")

	return nil
}
