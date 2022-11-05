package repo

import (
	"context"
	"dao-exchange/internal/models/scanner"

	"github.com/pkg/errors"
)

// QueryCurrentHeightWithTaskName query current height
func (d *Dao) QueryCurrentHeightWithTaskName(taskName string) (height uint64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	result := &scanner.BlockScanHeight{}
	tx := d.db.WithContext(ctx)
	if res := tx.Where(&scanner.BlockScanHeight{
		TaskName: taskName,
	}).Take(&result); res.Error != nil {
		return 0, res.Error
	}

	return result.Height, nil
}

// Insert insert to db
func (d *Dao) Insert(data *scanner.BlockScanHeight) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tx := d.db.WithContext(ctx)
	res := tx.Create(data)
	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return res.Error
}

// UpdateBlockHeight update to db
func (d *Dao) UpdateBlockHeight(taskName, chain string, update *scanner.BlockScanHeight) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tx := d.db.WithContext(ctx)
	res := tx.Where(&scanner.BlockScanHeight{
		Chain:    chain,
		TaskName: taskName,
	}).Updates(update)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}
