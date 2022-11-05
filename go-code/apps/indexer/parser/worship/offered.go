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

// OfferedParse parse nft offered event
type OfferedParse struct {
	contractAbi abi.ABI
}

// NewOfferedParse new obj
func NewOfferedParse(abiStr string) *OfferedParse {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	return &OfferedParse{
		contractAbi: contractAbi,
	}
}

// Parse parse message
func (p *OfferedParse) Parse(eventMsg []byte) (interface{}, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(eventMsg, eventLog); err != nil {
		return nil, err
	}

	var event model.TokenOfferedEvent
	event.Nft = common.HexToAddress(eventLog.Log.Topics[1])
	dataBytes := common.Hex2Bytes(eventLog.Log.Topics[2][2:])
	event.TokenId = new(big.Int).SetBytes(dataBytes)
	event.Votary = common.HexToAddress(eventLog.Log.Topics[3])
	dataBytes, err := common.ParseHexOrString(eventLog.Log.Data)
	if err != nil {
		return nil, err
	}

	err = p.contractAbi.UnpackIntoInterface(&event, "TokenOffered", dataBytes)
	if err != nil {
		log.Fatal(err)
	}

	return p.convertModel(eventLog, &event), nil
}

func (p *OfferedParse) convertModel(eventLog *model.EventModel, event *model.TokenOfferedEvent) *indexer.TokenWorship {
	offeredEvent := &indexer.TokenWorship{}
	offeredEvent.EventBase = *parser.FillBaseModel(eventLog)
	offeredEvent.NftAddress = strings.ToLower(event.Nft.Hex())
	offeredEvent.Redeemer = strings.ToLower(event.Redeemer.Hex())
	offeredEvent.Votary = strings.ToLower(event.Votary.Hex())
	offeredEvent.TokenID = strings.ToLower(common.BigToHash(event.TokenId).Hex())
	offeredEvent.ReleaseTimestamp = event.ReleaseTimestamp.Uint64()
	offeredEvent.Type = 1

	return offeredEvent
}
