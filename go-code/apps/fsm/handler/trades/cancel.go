package trades

import (
	"context"
	"strconv"

	"dao-exchange/apps/common"
	fsmdao "dao-exchange/apps/fsm/dao"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/models/indexer"
	"dao-exchange/internal/models/order"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// CancelEvent handler
type CancelEvent struct {
	Opr    *dao.Dao
	fsmOpr *fsmdao.Dao
}

// NewCancelEvent new obj
func NewCancelEvent(db *gorm.DB) *CancelEvent {
	return &CancelEvent{
		Opr:    dao.NewDaoWithDB(db),
		fsmOpr: fsmdao.NewDaoWithDB(db),
	}
}

// HandleEvent 当取消订单的时候，出现有人已经付款（币币账户，并且撮合单已经记录了数据，需要将撮合状态修改为失败
func (e *CancelEvent) HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*OrderStatusHandler, error) {
	parseredEvent := data.(*indexer.NftexOrderCanceledEvent)
	ctx := context.Background()
	chainID, err := strconv.ParseUint(eventDef.ChainID, 10, 64)
	if err != nil {
		return nil, common.ErrNotSupportChainID
	}

	query := order.NftOrder{
		OrderHash: parseredEvent.OrderHash,
		ChainID:   int64(chainID),
		Status:    int(common.OrderListing),
	}

	order, err := e.fsmOpr.GetOrderInfo(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "get sell info orderHash:%s, chainID:%d", parseredEvent.OrderHash, chainID)
	}

	orderStatusHandler := NewOrderStatusHandler()
	err = orderStatusHandler.HandlerEvent(e.Opr,
		NftInfo{
			TokenID: order.NftID,
			ChainID: chainID,
		},
		TradeInfo{
			TxHash:         parseredEvent.TransactionHash,
			From:           parseredEvent.Maker,
			Type:           common.OrderCancel,
			OrderInfo:      order,
			BlockTimestamp: int64(parseredEvent.BlockTimestamp),
		},
	)
	return orderStatusHandler, err
}
