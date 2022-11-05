package dex

import (
	"dao-exchange/apps/indexer/model"
	"dao-exchange/internal/models/indexer"
	"encoding/json"
	"log"
	"strings"

	"dao-exchange/apps/indexer/parser"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// CancelOrder parse nft match order event
type CancelOrder struct {
	contractAbi abi.ABI
}

// NewCancelOrder new obj
func NewCancelOrder(abiStr string) *CancelOrder {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	return &CancelOrder{
		contractAbi: contractAbi,
	}
}

// Parse parse message
func (p *CancelOrder) Parse(eventMsg []byte) (interface{}, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(eventMsg, eventLog); err != nil {
		return nil, err
	}

	var cancelEvent model.CancelOrderEvent
	cancelEvent.Maker = common.HexToAddress(eventLog.Log.Topics[1])
	cancelEvent.OrderHash = common.HexToHash(eventLog.Log.Topics[2])

	return p.convertModel(eventLog, &cancelEvent), nil
}

func (p *CancelOrder) convertModel(event *model.EventModel, cancelEvent *model.CancelOrderEvent) *indexer.NftexOrderCanceledEvent {
	eventModel := &indexer.NftexOrderCanceledEvent{}
	eventModel.EventBase = *parser.FillBaseModel(event)
	eventModel.Maker = strings.ToLower(cancelEvent.Maker.Hex())
	eventModel.OrderHash = strings.ToLower(cancelEvent.OrderHash.Hex())

	return eventModel
}
