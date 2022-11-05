package worship

import (
	"context"

	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/models/indexer"
)

type Actions interface {
	HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*StatusHandler, error)
}

// StatusHandler union hanlder
type StatusHandler struct {
	NeedUpdateStatus *indexer.TokenWorshipStatus
	NeedInsertStatus *indexer.TokenWorshipStatus
}

// SyncToDB sync to db
func (h *StatusHandler) SyncToDB(ctx context.Context, opr *dao.Dao) error {
	if h.NeedUpdateStatus != nil {
		if err := opr.UpdateWorshipStatusByID(ctx, *h.NeedUpdateStatus, h.NeedUpdateStatus.Worship); err != nil {
			return err
		}
	}

	if h.NeedInsertStatus != nil {
		return opr.Insert(ctx, h.NeedInsertStatus)
	}

	return nil
}
