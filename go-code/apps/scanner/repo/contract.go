package repo

import (
	"context"
	"dao-exchange/internal/models/scanner"
	"dao-exchange/pkg/myerr"

	"gorm.io/gorm"
)

// QueryContracts query
func (d *Dao) QueryContracts(chain string, chainID int) ([]*scanner.BlockScanContract, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tx := d.db.WithContext(ctx)
	contracts := []*scanner.BlockScanContract{}
	if res := tx.Where(&scanner.BlockScanContract{
		Chain:   chain,
		ChainID: chainID,
	}).Find(&contracts); res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, myerr.ErrRecordNotFound
		}
		return nil, res.Error
	}

	if len(contracts) == 0 {
		return nil, myerr.ErrRecordNotFound
	}

	return contracts, nil
}

// UpdateContractsStatus update to db
func (d *Dao) UpdateContractsStatus(endHeight uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	tx := d.db.WithContext(ctx)
	res := tx.Model(&scanner.BlockScanContract{}).Where("f_end_height != 0").Updates(map[string]interface{}{"f_end_height": endHeight})
	return res.Error
}
