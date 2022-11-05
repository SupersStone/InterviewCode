package worship

import (
	"encoding/json"
	"log"
	"math/big"
	"strings"

	"dao-exchange/apps/indexer/model"
	"dao-exchange/apps/indexer/parser"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// RedeemedParse parse nft offered event
type RedeemedParse struct {
	contractAbi abi.ABI
}

// NewRedeemedParse new obj
func NewRedeemedParse(abiStr string) *RedeemedParse {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	return &RedeemedParse{
		contractAbi: contractAbi,
	}
}

// Parse parse message
func (p *RedeemedParse) Parse(eventMsg []byte) (interface{}, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(eventMsg, eventLog); err != nil {
		return nil, err
	}

	var event model.TokenRedeemedEvent
	event.Nft = common.HexToAddress(eventLog.Log.Topics[1])
	dataBytes := common.Hex2Bytes(eventLog.Log.Topics[2][2:])
	event.TokenId = new(big.Int).SetBytes(dataBytes)
	event.Redeemer = common.HexToAddress(eventLog.Log.Topics[3])

	return p.convertModel(eventLog, &event), nil
}

func (p *RedeemedParse) convertModel(eventLog *model.EventModel, event *model.TokenRedeemedEvent) *indexer.TokenWorship {
	offeredEvent := &indexer.TokenWorship{}
	offeredEvent.EventBase = *parser.FillBaseModel(eventLog)
	offeredEvent.NftAddress = strings.ToLower(event.Nft.Hex())
	offeredEvent.Redeemer = strings.ToLower(event.Redeemer.Hex())
	offeredEvent.TokenID = strings.ToLower(common.BigToHash(event.TokenId).Hex())
	offeredEvent.Type = 2

	return offeredEvent
}
