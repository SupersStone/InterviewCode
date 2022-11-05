package worship

import (
	"context"

	"dao-exchange/apps/common"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/models/indexer"
	"dao-exchange/pkg/myerr"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// NFTWorship just impl actions interface
type NFTWorship struct {
	Opr *dao.Dao
}

// NewNFTOwner new obj
func NewNFTWorship(db *gorm.DB) *NFTWorship {
	return &NFTWorship{
		Opr: dao.NewDaoWithDB(db),
	}
}

// HandleEvent status
func (e *NFTWorship) HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*StatusHandler, error) {
	statusHandler := &StatusHandler{}
	parseredEvent, ok := data.(*indexer.TokenWorship)
	if !ok {
		return nil, errors.New("handle_event data is not tranfer event")
	}

	ctx := context.Background()
	score := common.CalcScore(parseredEvent.BlockNumber, parseredEvent.TransactionIndex, parseredEvent.LogIndex)
	worshipStatus, err := e.Opr.QueryWorshipStatus(ctx, eventDef.ChainID, eventDef.ContractAddress, parseredEvent.TokenID)
	if err != nil {
		if err != myerr.ErrRecordNotFound {
			return nil, err
		}

		statusHandler.NeedInsertStatus = &indexer.TokenWorshipStatus{
			ChainID:         eventDef.ChainID,
			ContractAddress: eventDef.ContractAddress,
			Score:           score,
			EndTime:         parseredEvent.ReleaseTimestamp,
			StartTime:       parseredEvent.BlockTimestamp,
			TokenID:         parseredEvent.TokenID,
			Worship:         parseredEvent.Type,
		}
		return statusHandler, nil
	}
	worshipStatus.Score = score
	worshipStatus.Worship = parseredEvent.Type
	statusHandler.NeedUpdateStatus = worshipStatus

	return statusHandler, nil
}
