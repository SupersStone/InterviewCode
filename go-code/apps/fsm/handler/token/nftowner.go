package token

import (
	"context"
	"strings"

	"dao-exchange/apps/common"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/models/indexer"
	"dao-exchange/pkg/myerr"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// NFTOwner just impl actions interface
type NFTOwner struct {
	Opr *dao.Dao
}

// NewNFTOwner new obj
func NewNFTOwner(db *gorm.DB) *NFTOwner {
	return &NFTOwner{
		Opr: dao.NewDaoWithDB(db),
	}
}

// HandleEvent status
func (e *NFTOwner) HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*StatusHandler, error) {
	statusHandler := &StatusHandler{}
	parseredEvent, ok := data.(*indexer.Erc721TransferEvent)
	if !ok {
		return nil, errors.New("handle_event data is not tranfer event")
	}

	score := common.CalcScore(parseredEvent.BlockNumber, parseredEvent.TransactionIndex, parseredEvent.LogIndex)
	// 处理 Mint
	if strings.EqualFold(parseredEvent.From, common.ZeroAddress) {
		statusHandler.NeedInsertNFTOwner = &indexer.Token721{
			EventDefinitionID: eventDef.ID,
			BlockNumber:       parseredEvent.BlockNumber,
			TransactionIndex:  parseredEvent.TransactionIndex,
			LogIndex:          parseredEvent.LogIndex,
			BlockTimestamp:    parseredEvent.BlockTimestamp,
			TokenID:           parseredEvent.TokenID,
			Owner:             parseredEvent.To,
			Score:             score,
		}

		return statusHandler, nil
	}

	ctx := context.Background()
	tokenInfo, err := e.Opr.QueryNFT(ctx, eventDef.ID, parseredEvent.TokenID)
	if err != nil {
		if err != myerr.ErrRecordNotFound {
			return nil, err
		}

		statusHandler.NeedInsertNFTOwner = &indexer.Token721{
			EventDefinitionID: eventDef.ID,
			BlockNumber:       parseredEvent.BlockNumber,
			TransactionIndex:  parseredEvent.TransactionIndex,
			LogIndex:          parseredEvent.LogIndex,
			BlockTimestamp:    parseredEvent.BlockTimestamp,
			TokenID:           parseredEvent.TokenID,
			Owner:             parseredEvent.To,
			Score:             score,
		}

		return statusHandler, nil
	}

	tokenInfo.Score = score
	statusHandler.NeedUpdateNFTOwner = tokenInfo

	return statusHandler, nil
}
