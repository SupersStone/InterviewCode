package dex

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"dao-exchange/apps/indexer/model"
	"dao-exchange/apps/indexer/parser"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MatchOrderParse parse nft match order event
type MatchOrderParse struct {
	contractAbi abi.ABI
}

// NewMatchOrder new obj
func NewMatchOrder(abiStr string) *MatchOrderParse {
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	return &MatchOrderParse{
		contractAbi: contractAbi,
	}
}

// Parse parse message
func (p *MatchOrderParse) Parse(eventMsg []byte) (interface{}, error) {
	eventLog := &model.EventModel{}
	if err := json.Unmarshal(eventMsg, eventLog); err != nil {
		return nil, err
	}

	var matchOrder model.FixedPriceOrderMatch
	matchOrder.Maker = common.HexToAddress(eventLog.Log.Topics[1])
	matchOrder.Taker = common.HexToAddress(eventLog.Log.Topics[2])
	matchOrder.OrderHash = common.HexToHash(eventLog.Log.Topics[3])

	dataBytes, err := common.ParseHexOrString(eventLog.Log.Data)
	if err != nil {
		return nil, err
	}

	err = p.contractAbi.UnpackIntoInterface(&matchOrder, "FixedPriceOrderMatched", dataBytes)
	if err != nil {
		log.Fatal(err)
	}

	var tokenStruct model.TokenBytesInfo
	err = p.contractAbi.UnpackIntoInterface(&tokenStruct, "AssetBytesInfo", matchOrder.AssetsBytes)
	if err != nil {
		log.Fatal(err)
	}

	var makerOrder model.OrderBytesInfo
	err = p.contractAbi.UnpackIntoInterface(&makerOrder, "OrderBytesInfo", matchOrder.OrderBytes)
	if err != nil {
		log.Fatal(err)
	}

	return p.convertModel(eventLog, &matchOrder, &makerOrder, &tokenStruct), nil
}

func (p *MatchOrderParse) convertModel(event *model.EventModel, orderMatch *model.FixedPriceOrderMatch, makerOrder *model.OrderBytesInfo, tokenInfo *model.TokenBytesInfo) *indexer.FixedPriceMatchedEvent {
	matchEvent := &indexer.FixedPriceMatchedEvent{}
	matchEvent.EventBase = *parser.FillBaseModel(event)
	matchEvent.Maker = strings.ToLower(orderMatch.Maker.Hex())
	matchEvent.Taker = strings.ToLower(orderMatch.Taker.Hex())
	matchEvent.OrderHash = strings.ToLower(orderMatch.OrderHash.Hex())
	matchEvent.OrderBytes = strings.ToLower(hexutil.Encode(orderMatch.OrderBytes))
	matchEvent.RoyaltyRecipient = strings.ToLower(makerOrder.RoyaltyRecipient.Hex())
	matchEvent.RoyaltyRate = makerOrder.RoyaltyRate.String()
	matchEvent.StartTime = int64(makerOrder.StartAt)
	matchEvent.ExpireTime = int64(makerOrder.ExpireAt)
	matchEvent.MakerNonce = strconv.FormatUint(makerOrder.MakerNonce, 10)
	if makerOrder.TakerGetNft {
		matchEvent.TakerGetNft = 1
	} else {
		matchEvent.TakerGetNft = 0
	}

	matchEvent.Nft = strings.ToLower(tokenInfo.Nft.Hex())
	matchEvent.Ft = strings.ToLower(tokenInfo.Ft.Hex())
	matchEvent.NftID = strings.ToLower(common.BigToHash(tokenInfo.NftId).Hex())
	matchEvent.NftAmount = tokenInfo.NftAmount.String()
	matchEvent.FtAmount = tokenInfo.FtAmount.String()

	return matchEvent
}
