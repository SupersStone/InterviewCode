package handler

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"dao-exchange/apps/common"
	"dao-exchange/apps/fsm/handler/event"
	"dao-exchange/apps/fsm/handler/token"
	"dao-exchange/apps/fsm/handler/trades"
	"dao-exchange/apps/fsm/handler/worship"
	worshiphandler "dao-exchange/apps/fsm/handler/worship"
	"dao-exchange/apps/indexer/dao"
	"dao-exchange/apps/indexer/model"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var defaultHandlerActions = make(map[string]*MsgAction)

// MsghandlerResult msg
type MsghandlerResult struct {
	event      interface{}
	trade      *trades.OrderStatusHandler
	token      *token.StatusHandler
	worship    *worship.StatusHandler
	processLog *indexer.ProcessLog
}

// MsgAction msg
type MsgAction struct {
	TradeAction   trades.Actions
	EventAction   event.Actions
	TokenAction   token.Actions
	WorshipAction worshiphandler.Actions
}

// InitTestAction init handler actions
func InitTestAction(db *gorm.DB) map[string]*MsgAction {
	actions := make(map[string]*MsgAction)
	matchEvent := event.NewMatchEvent()
	actions[matchEvent.EventKey()] = &MsgAction{
		EventAction: matchEvent,
	}
	transferEvent := event.NewTransferEvent()
	actions[transferEvent.EventKey()] = &MsgAction{
		EventAction: transferEvent,
		TokenAction: token.NewNFTOwner(db),
	}
	cancelOrder := event.NewOrderCancelledEvent()
	actions[cancelOrder.EventKey()] = &MsgAction{
		EventAction: cancelOrder,
	}
	defaultHandlerActions = actions
	return actions
}

// InitAction init handler actions
func InitAction(db *gorm.DB) map[string]*MsgAction {
	actions := make(map[string]*MsgAction)
	matchEvent := event.NewMatchEvent()
	actions[matchEvent.EventKey()] = &MsgAction{
		TradeAction: trades.NewFixedPriceMatchEvent(db),
		EventAction: matchEvent,
	}

	transferEvent := event.NewTransferEvent()
	actions[transferEvent.EventKey()] = &MsgAction{
		TradeAction: trades.NewTransferEvent(db),
		EventAction: transferEvent,
		TokenAction: token.NewNFTOwner(db),
	}

	worshipOfferedEvent := event.NewWorshipOfferedEvent()
	actions[worshipOfferedEvent.EventKey()] = &MsgAction{
		EventAction:   worshipOfferedEvent,
		WorshipAction: worship.NewNFTWorship(db),
	}

	worshipRedeemedEvent := event.NewWorshipRedeemedEvent()
	actions[worshipRedeemedEvent.EventKey()] = &MsgAction{
		EventAction:   worshipRedeemedEvent,
		WorshipAction: worship.NewNFTWorship(db),
	}

	cancelOrder := event.NewOrderCancelledEvent()
	actions[cancelOrder.EventKey()] = &MsgAction{
		TradeAction: trades.NewCancelEvent(db),
		EventAction: cancelOrder,
	}

	cancelAllOrder := event.NewOrderCancellAllEvent()
	actions[cancelAllOrder.EventKey()] = &MsgAction{
		TradeAction: trades.NewCancelAllEvent(db),
		EventAction: cancelOrder,
	}
	defaultHandlerActions = actions
	return actions
}

// ConsumeMsg consumer message from kafka
func ConsumeMsg(message []byte, hanlderActions map[string]*MsgAction) (*MsghandlerResult, error) {
	if hanlderActions == nil {
		hanlderActions = defaultHandlerActions
	}

	eventLog := &model.EventModel{}
	var err error
	if err := json.Unmarshal(message, eventLog); err != nil {
		return nil, common.ErrUnMarshalFail
	}

	eventDef := cache.LoadEventDef(eventLog.GetContractAddr(), eventLog.GetFunctionSign(), strconv.FormatInt(eventLog.ChainID, 10))
	if eventDef == nil {
		return nil, common.ErrNotFoundEventDef
	}

	action, ok := hanlderActions[eventDef.CacheKey()]
	if !ok {
		return nil, common.ErrNotSupportEvent
	}

	var processLogs = &indexer.ProcessLog{
		EventBase: indexer.EventBase{
			EventDefinitionID: eventDef.ID,
			BlockHash:         eventLog.Log.BlockHash,
			TransactionHash:   eventLog.Log.TransactionHash,
			BlockNumber:       hexutil.MustDecodeUint64(eventLog.Log.BlockNumber),
			TransactionIndex:  hexutil.MustDecodeUint64(eventLog.Log.TransactionIndex),
			LogIndex:          hexutil.MustDecodeUint64(eventLog.Log.LogIndex),
			BlockTimestamp:    uint64(eventLog.BlockTimestamp),
		},
		LogJSON: string(message),
	}

	resp := &MsghandlerResult{}
	resp.event, err = action.EventAction.HandlerMsg(message)
	if err != nil {
		processLogs.Status = common.Fail
		processLogs.Msg = err.Error()
		logrus.Errorf("EventHandler: err = %s", err.Error())
	} else {
		if action.TradeAction != nil {
			resp.trade, err = action.TradeAction.HandleEvent(resp.event, eventDef)
			if err != nil {
				processLogs.Status = common.Fail
				processLogs.Msg = err.Error()
				logrus.Errorf("TradeAction EventHandler: err = %s", err.Error())
			} else {
				processLogs.Status = common.Success
			}
		}
		if action.TokenAction != nil {
			resp.token, err = action.TokenAction.HandleEvent(resp.event, eventDef)
			if err != nil {
				processLogs.Status = common.Fail
				processLogs.Msg = err.Error()
				logrus.Errorf("TokenAction EventHandler: err = %s", err.Error())
			} else {
				processLogs.Status = common.Success
			}
		}
		if action.WorshipAction != nil {
			resp.worship, err = action.WorshipAction.HandleEvent(resp.event, eventDef)
			if err != nil {
				processLogs.Status = common.Fail
				processLogs.Msg = err.Error()
				logrus.Errorf("WorshipAction EventHandler: err = %s", err.Error())
			} else {
				processLogs.Status = common.Success
			}
		}
	}

	resp.processLog = processLogs
	return resp, nil
}

// SyncToDB to db
func SyncToDB(db *gorm.DB, data *MsghandlerResult) error {
	tx := db.Begin()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() {
		cancel()
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := dao.Insert(ctx, tx, data.processLog); err != nil {
		tx.Rollback()
		return err
	}

	if err := dao.Insert(ctx, tx, data.event); err != nil {
		tx.Rollback()
		return err
	}

	if data.trade != nil {
		if err := data.trade.SyncToDBWithoutTx(ctx, dao.NewDaoWithDB(tx)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if data.token != nil {
		if err := data.token.SyncToDB(ctx, dao.NewDaoWithDB(tx)); err != nil {
			tx.Rollback()
			return err
		}
	}

	if data.worship != nil {
		if err := data.worship.SyncToDB(ctx, dao.NewDaoWithDB(tx)); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
