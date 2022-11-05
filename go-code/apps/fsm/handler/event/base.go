package event

import (
	"context"
	"fmt"

	"dao-exchange/apps/fsm/dao"

	"gorm.io/gorm"
)

// Base just impl actions interface
type Base struct {
	eventName    string
	contractType ContractType
	eventDao     dao.BaseInsertDao
}

// EventType empty handler for reveiced msg
func (h *Base) EventType() ContractType {
	return DAODex
}

// EventName empty handler for reveiced msg
func (h *Base) EventName() string {
	return h.eventName
}

// EventKey empty handler for reveiced msg
func (h *Base) EventKey() string {
	return fmt.Sprintf("%s-%s", h.contractType, h.eventName)
}

// InsertDB write to db
func (h *Base) InsertDB(db *gorm.DB, data interface{}) error {
	return h.eventDao.Insert(context.Background(), db, data)
}
