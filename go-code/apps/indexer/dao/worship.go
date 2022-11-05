package dao

import (
	"context"

	"dao-exchange/internal/models/indexer"
	"dao-exchange/pkg/myerr"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// QueryNFT query order with TokenID and chainID where order status is listing.
func (d *Dao) QueryWorshipStatus(ctx context.Context, chainID string, contractAddr, tokenID string) (*indexer.TokenWorshipStatus, error) {
	token := &indexer.TokenWorshipStatus{}
	tx := d.db.WithContext(ctx)
	res := tx.Where(indexer.TokenWorshipStatus{
		ChainID:         chainID,
		ContractAddress: contractAddr,
		TokenID:         tokenID,
	}).Take(token)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, myerr.ErrRecordNotFound
		}

		return nil, res.Error
	}

	return token, nil
}

// UpdateNFTOwnerByID update one
func (d *Dao) UpdateWorshipStatusByID(ctx context.Context, wheres indexer.TokenWorshipStatus, status int8) (err error) {
	tx := d.db.WithContext(ctx)
	res := tx.Where("id = ? and score < ?", wheres.ID, wheres.Score).
		Updates(&indexer.TokenWorshipStatus{
			Worship: status,
		})
	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return
}
