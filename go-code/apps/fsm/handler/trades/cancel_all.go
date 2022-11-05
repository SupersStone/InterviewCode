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

// CancelAllEvent handler
type CancelAllEvent struct {
	Opr    *dao.Dao
	fsmOpr *fsmdao.Dao
}

// NewCancelAllEvent new obj
func NewCancelAllEvent(db *gorm.DB) *CancelAllEvent {
	return &CancelAllEvent{
		Opr:    dao.NewDaoWithDB(db),
		fsmOpr: fsmdao.NewDaoWithDB(db),
	}
}

// HandleEvent event
func (e *CancelAllEvent) HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*OrderStatusHandler, error) {
	parseredEvent := data.(*indexer.NftexAllOrderCanceledEvent)
	ctx := context.Background()
	chainID, err := strconv.ParseUint(eventDef.ChainID, 10, 64)
	if err != nil {
		return nil, common.ErrNotSupportChainID
	}

	query := order.NftOrder{
		Maker:   parseredEvent.Maker,
		ChainID: int64(chainID),
		Status:  int(common.OrderListing),
	}

	orders, err := e.fsmOpr.ListOrders(ctx, query)
	if err != nil {
		return nil, errors.Wrapf(err, "get sell info maker:%s, chainID:%d", parseredEvent.Maker, chainID)
	}

	orderStatusHandler := NewOrderStatusHandler()
	for _, info := range orders {
		orderStatusHandler.NeedUpdateOrder[common.OrderCancelAll] = append(orderStatusHandler.NeedUpdateOrder[common.OrderCancelAll], info.ID)
		nftInfo := NftInfo{
			TokenID: info.NftID,
			ChainID: chainID,
		}
		tradeInfo := TradeInfo{
			TxHash:         parseredEvent.TransactionHash,
			From:           parseredEvent.Maker,
			Type:           common.OrderCancelAll,
			OrderInfo:      &info,
			BlockTimestamp: int64(parseredEvent.BlockTimestamp),
		}

		orderStatusHandler.NeedInsertHistory = append(orderStatusHandler.NeedInsertHistory, buildCanlcelAllHistory(nftInfo, tradeInfo))
	}

	return orderStatusHandler, nil
}

func buildCanlcelAllHistory(nftInfo NftInfo, tradeInfo TradeInfo) *order.NftOperationHistory {
	return buildHistory(common.OrderCancelAll, nftInfo, tradeInfo)
}

func (e *CancelAllEvent) syncToDB(opr *dao.Dao, datas []*OrderStatusHandler) error {
	tx := opr.DB().Begin()
	txOpr := dao.NewDaoWithDB(tx)
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, handler := range datas {
		if err := handler.SyncToDB(*txOpr); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
