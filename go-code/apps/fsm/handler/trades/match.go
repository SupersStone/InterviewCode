package trades

import (
	"context"
	"strconv"

	"dao-exchange/apps/common"
	fsmdao "dao-exchange/apps/fsm/dao"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/enums"
	"dao-exchange/internal/models/indexer"
	"dao-exchange/internal/models/order"

	"gorm.io/gorm"
)

// FixedPriceMatchEvent just impl actions interface
type FixedPriceMatchEvent struct {
	Opr    *dao.Dao
	fsmOpr *fsmdao.Dao
}

// NewFixedPriceMatchEvent new obj
func NewFixedPriceMatchEvent(db *gorm.DB) *FixedPriceMatchEvent {
	return &FixedPriceMatchEvent{
		Opr:    dao.NewDaoWithDB(db),
		fsmOpr: fsmdao.NewDaoWithDB(db),
	}
}

// HandleEvent event
func (e *FixedPriceMatchEvent) HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*OrderStatusHandler, error) {
	parseredEvent := data.(*indexer.FixedPriceMatchedEvent)
	chainID, err := strconv.ParseUint(eventDef.ChainID, 10, 64)
	if err != nil {
		return nil, common.ErrNotSupportChainID
	}

	payToken := cache.LoadPayTokenWithContract(parseredEvent.Ft, chainID)
	if payToken == nil {
		return nil, common.ErrNotSupportPayToken
	}

	ctx := context.Background()
	query := order.NftOrder{
		NftID:     parseredEvent.NftID,
		OrderHash: parseredEvent.OrderHash,
		ChainID:   int64(chainID),
		Status:    enums.OrderListing,
	}

	order, err := e.fsmOpr.GetOrderInfo(ctx, query)
	if err != nil {
		return nil, err
	}

	orderStatusHandler := NewOrderStatusHandler()
	err = orderStatusHandler.HandlerEvent(e.Opr,
		NftInfo{
			TokenID:         parseredEvent.NftID,
			ChainID:         chainID,
			ContractAddress: parseredEvent.Nft,
		},
		TradeInfo{
			TxHash:         parseredEvent.TransactionHash,
			From:           parseredEvent.Maker,
			To:             parseredEvent.Taker,
			Type:           common.OrderMatchFixed,
			BlockTimestamp: int64(parseredEvent.BlockTimestamp),
			OrderInfo:      order,
			Decimls:        int(payToken.Decimal),
		},
	)
	return orderStatusHandler, err
}
