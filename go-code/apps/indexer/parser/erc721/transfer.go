package erc721

import (
	"encoding/json"
	"log"
	"math/big"
	"strconv"
	"strings"

	"dao-exchange/apps/indexer/model"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TransferParse parse nft match order event
type TransferParse struct {
	contractAbi abi.ABI
}

// NewTransferParse new obj
func NewTransferParse(abiStr string) *TransferParse {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	return &TransferParse{
		contractAbi: contractAbi,
	}
}

// Parse parser message
func (p *TransferParse) Parse(eventMsg []byte) (interface{}, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(eventMsg, eventLog); err != nil {
		return nil, err
	}

	var transferEvent model.TokenTransferEvent
	transferEvent.From = common.HexToAddress(eventLog.Log.Topics[1])
	transferEvent.To = common.HexToAddress(eventLog.Log.Topics[2])
	dataBytes := common.Hex2Bytes(eventLog.Log.Topics[3][2:])
	transferEvent.TokenId = new(big.Int).SetBytes(dataBytes)

	return convertModel(eventLog, transferEvent), nil
}

func convertModel(event *model.EventModel, transferEvent model.TokenTransferEvent) *indexer.Erc721TransferEvent {
	res := &indexer.Erc721TransferEvent{}
	res.EventDefinitionID = cache.LoadEventDef(event.Log.Address, event.Log.Topics[0], strconv.FormatInt(event.ChainID, 10)).ID
	res.BlockHash = event.Log.BlockHash
	res.TransactionHash = event.Log.TransactionHash
	res.BlockNumber = hexutil.MustDecodeUint64(event.Log.BlockNumber)
	res.TransactionIndex = hexutil.MustDecodeUint64(event.Log.TransactionIndex)
	res.LogIndex = hexutil.MustDecodeUint64(event.Log.LogIndex)
	res.BlockTimestamp = uint64(event.BlockTimestamp)

	res.From = strings.ToLower(transferEvent.From.Hex())
	res.To = strings.ToLower(transferEvent.To.Hex())
	res.TokenID = strings.ToLower(common.BigToHash(transferEvent.TokenId).Hex())

	return res
}
