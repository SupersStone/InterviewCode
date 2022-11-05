package token

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
	NeedUpdateNFTOwner *indexer.Token721
	NeedInsertNFTOwner *indexer.Token721
}

// SyncToDB sync to db
func (h *StatusHandler) SyncToDB(ctx context.Context, opr *dao.Dao) error {
	if h.NeedUpdateNFTOwner != nil {
		if err := opr.UpdateNFTOwnerByID(ctx, *h.NeedUpdateNFTOwner, h.NeedUpdateNFTOwner.Owner); err != nil {
			return err
		}
	}
	if h.NeedInsertNFTOwner != nil {
		return opr.Insert(ctx, h.NeedInsertNFTOwner)
	}

	return nil
}
