package indexer

import (
	"context"
	"encoding/json"
	"strconv"

	"dao-exchange/apps/common"
	"dao-exchange/apps/indexer/model"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/dao"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
)

// Parser parse msg to model
type Parser struct {
	db *gorm.DB
}

// New obj
func New(db *gorm.DB) *Parser {
	return &Parser{
		db: db,
	}
}

// HandlerMsg handler msg
func (p *Parser) HandlerMsg(message []byte) error {
	processLog, err := hanldeMsg(message)
	if err != nil {
		return err
	}

	return dao.Insert(context.Background(), p.db, processLog)
}

// ConsumeMsg consumer message from kafka
func hanldeMsg(message []byte) (*indexer.ProcessLog, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(message, eventLog); err != nil {
		return nil, err
	}

	eventDef := cache.LoadEventDef(eventLog.GetContractAddr(), eventLog.GetFunctionSign(), strconv.FormatInt(eventLog.ChainID, 10))
	if eventDef == nil {
		return nil, common.ErrNotFoundEventDef
	}

	processLog := &indexer.ProcessLog{
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

	return processLog, nil
}
