package trades

import (
	"context"
	"strings"
	"time"

	"dao-exchange/apps/common"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/internal/enums"
	"dao-exchange/internal/models/order"

	"github.com/shopspring/decimal"
)

/*
 需要做的事情
 1. 统一处理需要修改状态的表具体数据
    a. 需要修改Order（状态更改） / 链上撮合，转移以及取消
	b. 需要插入历史记录
 2. 统一在一个事物中处理所有的待更新数据
*/

// OrderStatusHandler union hanlder
type OrderStatusHandler struct {
	NeedUpdateOrder   map[int][]int64 // key:status, value:id
	NeedInsertHistory []*order.NftOperationHistory
}

// NewOrderStatusHandler new obj
func NewOrderStatusHandler() *OrderStatusHandler {
	return &OrderStatusHandler{
		NeedUpdateOrder: make(map[int][]int64),
	}
}

// NftInfo base info
type NftInfo struct {
	TokenID         string
	Amount          int64
	ChainID         uint64
	ContractAddress string
}

// TradeInfo info for dex cex
type TradeInfo struct {
	TxHash         string
	From           string
	To             string
	Type           int
	Price          decimal.Decimal
	Decimls        int
	Nonce          int64
	OrderInfo      *order.NftOrder
	BlockTimestamp int64
}

// HandlerEvent handler event for status change
func (h *OrderStatusHandler) HandlerEvent(opr *dao.Dao, nftInfo NftInfo, tradeInfo TradeInfo) error {
	if tradeInfo.OrderInfo != nil {
		tradeInfo.OrderInfo.Status = sellingOrderStatus(tradeInfo.OrderInfo.Status, nftInfo, tradeInfo)
		// 准备需要更新的订单
		h.NeedUpdateOrder[tradeInfo.OrderInfo.Status] = append(h.NeedUpdateOrder[tradeInfo.OrderInfo.Status], tradeInfo.OrderInfo.ID)
	}

	// 准备需要插入的历史记录
	status := historyStatus(tradeInfo)
	h.NeedInsertHistory = append(h.NeedInsertHistory, buildHistory(status, nftInfo, tradeInfo))

	return nil
}

func sellingOrderStatus(oriStatus int, nftInfo NftInfo, tradeInfo TradeInfo) int {
	switch tradeInfo.Type {
	case common.OrderTransfer:
		// 销毁NFT
		if strings.EqualFold(tradeInfo.To, common.ZeroAddress) {

			return int(enums.OrderInValid)
		}
		if oriStatus == int(common.OrderListing) {
			// NFT 链上已转移, 必须是挂单，不能是求购，求购状态不变
			return int(enums.OrderMatched)
		}
	case common.OrderMatchFixed:
		return int(enums.OrderMatched)
	case common.OrderCancel, common.OrderCancelAll:
		return int(enums.OrderCanceled)
	}

	return oriStatus
}

func historyStatus(tradeInfo TradeInfo) int {
	switch tradeInfo.Type {
	case common.OrderTransfer:
		// MINT NFT
		if strings.EqualFold(tradeInfo.From, common.ZeroAddress) {
			return int(enums.HistoryNFTMint)
		}
		// 销毁NFT
		if strings.EqualFold(tradeInfo.To, common.ZeroAddress) {
			return int(enums.HistoryNFTBurned)
		}
		// NFT 链上已转移
		return int(enums.HistoryNFTTransfered)
	case common.OrderMatchFixed:
		return int(enums.HistoryOrderMatchFixed)
	case common.OrderCancel, common.OrderCancelAll:
		return int(enums.HistoryOrderCanceled)
	}

	return 0
}

func buildHistory(status int, nftInfo NftInfo, tradeInfo TradeInfo) *order.NftOperationHistory {
	history := &order.NftOperationHistory{
		ContractAddress: nftInfo.ContractAddress,
		TokenID:         nftInfo.TokenID,
		SellAmount:      nftInfo.Amount,
		Price:           tradeInfo.Price,
		From:            tradeInfo.From,
		To:              tradeInfo.To,
		EventType:       status,
		Hash:            tradeInfo.TxHash,
		ChainID:         int64(nftInfo.ChainID),
		CreatedAt:       time.Unix(tradeInfo.BlockTimestamp, 0),
		UpdatedAt:       time.Unix(tradeInfo.BlockTimestamp, 0),
	}
	if tradeInfo.OrderInfo != nil {
		history.ContractType = tradeInfo.OrderInfo.ContractType
		history.PaymentToken = tradeInfo.OrderInfo.Ft
		history.OrderHash = tradeInfo.OrderInfo.OrderHash
	}

	return history
}

// SyncToDB write to db
func (h *OrderStatusHandler) SyncToDB(opr dao.Dao) error {
	tx := opr.DB().Begin()
	txOpr := dao.NewDaoWithDB(tx)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancel()
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 先插入buyOrder, 然后填充MatchOrder, 在插入matchOrder
	if h.NeedInsertHistory != nil {
		if err := txOpr.Insert(ctx, h.NeedInsertHistory); err != nil {
			tx.Rollback()
			return err
		}
	}

	for k, v := range h.NeedUpdateOrder {
		if v == nil {
			continue
		}

		if err := txOpr.UpdatesByIds(ctx, &order.NftOrder{}, v, order.NftOrder{Status: k}); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// SyncToDBWithoutTx write to db
func (h *OrderStatusHandler) SyncToDBWithoutTx(ctx context.Context, opr *dao.Dao) error {
	// 先插入buyOrder, 然后填充MatchOrder, 在插入matchOrder
	if h.NeedInsertHistory != nil {
		if err := opr.Insert(ctx, h.NeedInsertHistory); err != nil {
			return err
		}
	}

	for k, v := range h.NeedUpdateOrder {
		if v == nil {
			continue
		}

		if err := opr.UpdatesByIds(ctx, &order.NftOrder{}, v, order.NftOrder{Status: k}); err != nil {
			return err
		}
	}

	return nil
}
