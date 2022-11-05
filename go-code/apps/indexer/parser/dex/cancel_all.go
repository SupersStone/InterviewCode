package dex

import (
	"dao-exchange/apps/indexer/model"
	"dao-exchange/apps/indexer/parser"
	"dao-exchange/internal/models/indexer"
	"encoding/json"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// CancelAllOrder parse nft match order event
type CancelAllOrder struct {
	contractAbi abi.ABI
}

// NewCancelAllOrder new obj
func NewCancelAllOrder(abiStr string) *CancelAllOrder {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	return &CancelAllOrder{
		contractAbi: contractAbi,
	}
}

// Parse parse message
func (p *CancelAllOrder) Parse(eventMsg []byte) (interface{}, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(eventMsg, eventLog); err != nil {
		return nil, err
	}

	var cancelEvent model.CancelAllOrderEvent
	cancelEvent.Offerer = common.HexToAddress(eventLog.Log.Topics[1])
	dataBytes := common.FromHex(eventLog.Log.Data)
	err := p.contractAbi.UnpackIntoInterface(&cancelEvent, "AllOrdersCancelled", dataBytes)
	if err != nil {
		return nil, err
	}

	return p.convertModel(eventLog, &cancelEvent), nil
}

func (p *CancelAllOrder) convertModel(event *model.EventModel, cancelEvent *model.CancelAllOrderEvent) *indexer.NftexAllOrderCanceledEvent {
	eventModel := &indexer.NftexAllOrderCanceledEvent{}
	eventModel.EventBase = *parser.FillBaseModel(event)
	eventModel.Maker = strings.ToLower(cancelEvent.Offerer.Hex())
	eventModel.Nonce = cancelEvent.IncreasedNonce.Int64()

	return eventModel
}
