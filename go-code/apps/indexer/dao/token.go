package dao

import (
	"context"

	"dao-exchange/internal/models/indexer"
	"dao-exchange/pkg/myerr"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// QueryNFT query order with TokenID and chainID where order status is listing.
func (d *Dao) QueryNFT(ctx context.Context, eventID uint64, tokenID string) (*indexer.Token721, error) {
	tokenInfo := &indexer.Token721{}
	tx := d.db.WithContext(ctx)
	res := tx.Where(indexer.Token721{
		EventDefinitionID: eventID,
		TokenID:           tokenID,
	}).Take(tokenInfo)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, myerr.ErrRecordNotFound
		}

		return nil, res.Error
	}

	return tokenInfo, nil
}

// UpdateNFTOwnerByID update one
func (d *Dao) UpdateNFTOwnerByID(ctx context.Context, wheres indexer.Token721, owner string) (err error) {
	tx := d.db.WithContext(ctx)
	res := tx.Where("id = ? and score < ?", wheres.ID, wheres.Score).Updates(&indexer.Token721{
		Owner: owner,
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
