package dao

import (
	"context"
	"dao-exchange/internal/models/indexer"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// EventDefDao interface of order dao
type EventDefDao struct {
}

// QueryAll update one
func (d *EventDefDao) QueryAll(ctx context.Context, db *gorm.DB) ([]indexer.EventDefinition, error) {
	tx := db.WithContext(ctx)
	events := []indexer.EventDefinition{}
	res := tx.Model(&indexer.EventDefinition{}).Find(&events)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected < 1 {
		return nil, errors.Errorf("rows affected must bigger than 1, but got: %d", res.RowsAffected)
	}

	return events, nil
}

// Insert update one
func (d *EventDefDao) Insert(ctx context.Context, db *gorm.DB, eventDef *indexer.EventDefinition) error {
	tx := db.WithContext(ctx)
	res := tx.Create(eventDef)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected != 1 {
		return errors.Errorf("rows affected must 1, but got: %d", res.RowsAffected)
	}

	return nil
}
