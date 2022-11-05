package parser

import (
	"strconv"

	"dao-exchange/apps/indexer/model"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Parser parse event
type Parser interface {
	Parse(eventMsg []byte) (interface{}, error)
}

// FillBaseModel fill base model
func FillBaseModel(event *model.EventModel) *indexer.EventBase {
	res := &indexer.EventBase{}
	res.EventDefinitionID = cache.LoadEventDef(event.Log.Address, event.Log.Topics[0], strconv.FormatInt(event.ChainID, 10)).ID
	res.BlockHash = event.Log.BlockHash
	res.TransactionHash = event.Log.TransactionHash
	res.BlockNumber = hexutil.MustDecodeUint64(event.Log.BlockNumber)
	res.TransactionIndex = hexutil.MustDecodeUint64(event.Log.TransactionIndex)
	res.LogIndex = hexutil.MustDecodeUint64(event.Log.LogIndex)
	res.BlockTimestamp = uint64(event.BlockTimestamp)
	return res
}
