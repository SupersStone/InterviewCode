package dao

import (
	"context"
	"dao-exchange/internal/models/indexer"

	"gorm.io/gorm"
)

// WorshipDao interface of worship dao
type WorshipDao struct {
}

// NewWorshipDao new
func NewWorshipDao() *WorshipDao {
	return &WorshipDao{}
}

// QueryAll update one
func (d *WorshipDao) QueryAll(db *gorm.DB, tokenID string) (*indexer.TokenWorshipStatus, error) {
	tx := db.WithContext(context.Background())
	nftWorship := &indexer.TokenWorshipStatus{}
	res := tx.Model(&indexer.TokenWorshipStatus{TokenID: tokenID}).Take(nftWorship)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, res.Error
	}

	return nftWorship, nil
}
