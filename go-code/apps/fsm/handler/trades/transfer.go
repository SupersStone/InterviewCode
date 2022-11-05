package trades

import (
	"context"
	"strconv"
	"strings"

	"dao-exchange/apps/common"
	fsmdao "dao-exchange/apps/fsm/dao"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/enums"
	"dao-exchange/internal/models/indexer"
	"dao-exchange/internal/models/order"
	"dao-exchange/pkg/myerr"

	"gorm.io/gorm"
)

// TransferEvent just impl actions interface
type TransferEvent struct {
	fsmOpr *fsmdao.Dao
	Opr    *dao.Dao
}

// NewTransferEvent new obj
func NewTransferEvent(db *gorm.DB) *TransferEvent {
	return &TransferEvent{
		fsmOpr: fsmdao.NewDaoWithDB(db),
		Opr:    dao.NewDaoWithDB(db),
	}
}

// HandleEvent status
func (e *TransferEvent) HandleEvent(data interface{}, eventDef *indexer.EventDefinition) (*OrderStatusHandler, error) {
	parseredEvent := data.(*indexer.Erc721TransferEvent)
	chainID, err := strconv.ParseUint(eventDef.ChainID, 10, 64)
	if err != nil {
		return nil, common.ErrNotSupportChainID
	}

	nftInfo := NftInfo{
		TokenID:         parseredEvent.TokenID,
		ChainID:         chainID,
		ContractAddress: eventDef.ContractAddress,
	}

	tradeInfo := TradeInfo{
		TxHash:         parseredEvent.TransactionHash,
		From:           parseredEvent.From,
		To:             parseredEvent.To,
		Type:           common.OrderTransfer,
		BlockTimestamp: int64(parseredEvent.BlockTimestamp),
	}

	orderStatusHandler := NewOrderStatusHandler()
	if strings.EqualFold(parseredEvent.From, common.ZeroAddress) {
		err = orderStatusHandler.HandlerEvent(e.Opr, nftInfo, tradeInfo)
		return orderStatusHandler, err
	}

	// tokenID := ethCommon.HexToHash(transferEvent.TokenID).Hex()
	ctx := context.Background()
	query := order.NftOrder{
		Maker:   parseredEvent.From,
		ChainID: int64(chainID),
		NftID:   parseredEvent.TokenID,
		Status:  int(enums.OrderListing),
	}
	order, err := e.fsmOpr.GetOrderInfo(ctx, query)
	if err != nil {
		if err == myerr.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	payToken := cache.LoadPayTokenWithContract(order.Ft, chainID)
	tradeInfo.OrderInfo = order
	tradeInfo.Decimls = int(payToken.Decimal)
	err = orderStatusHandler.HandlerEvent(e.Opr, nftInfo, tradeInfo)
	return orderStatusHandler, err
}
