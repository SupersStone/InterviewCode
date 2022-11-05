package dao

import (
	"context"
	"reflect"

	"dao-exchange/internal/models/order"
	"dao-exchange/pkg/myerr"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// OrderDao interface of order dao
type OrderDao interface {
	Query() (order.NftOrder, error)
}

// Query query order with TokenID and chainID where order status is listing.
func (d *Dao) Query(ctx context.Context, tokenID string, chainID uint64) (*order.NftOrder, error) {
	cond := make(map[string]interface{}, 0)
	cond["f_token_id"] = tokenID
	cond["f_chain_id"] = chainID
	cond["f_status"] = 0

	order := order.NftOrder{}
	tx := d.db.WithContext(ctx)
	res := tx.Model(&order).Take(&order, cond)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, myerr.ErrRecordNotFound
		}

		return nil, res.Error
	}

	return &order, nil
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
func (d *Dao) UpdatesByIds(ctx context.Context, models interface{}, ids []uint64, update interface{}) (err error) {
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

// Update update one
func (d *Dao) Update(ctx context.Context, orderHash string, updateMap map[string]interface{}) (err error) {
	tx := d.db.WithContext(ctx)
	res := tx.Model(&order.NftOrder{}).Where("f_order_hash = ? and f_status = 0", orderHash).Updates(updateMap)
	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return
}

// GetOrderInfo query order with TokenID and chainID where order status is listing.
func (d *Dao) GetOrderInfo(ctx context.Context, query order.NftOrder) (*order.NftOrder, error) {
	order := order.NftOrder{}
	tx := d.db.WithContext(ctx)
	res := tx.Where(&query).Take(&order)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, myerr.ErrRecordNotFound
		}

		return nil, res.Error
	}

	return &order, nil
}

// ListOrders 查询用户正在挂单的订单信息
func (d *Dao) ListOrders(ctx context.Context, query order.NftOrder) ([]order.NftOrder, error) {
	tx := d.db.WithContext(ctx)
	result := []order.NftOrder{}
	res := tx.Where(&query).Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	return result, nil
}
