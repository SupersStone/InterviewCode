package dex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"dao-exchange/apps/indexer/contracts/nft"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestMatchOrderParse_parse(t *testing.T) {
	mocksEventDefCache()
	type fields struct {
		abiStr string
	}
	type args struct {
		eventMsg    json.RawMessage
		contractAbi abi.ABI
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *indexer.FixedPriceMatchedEvent
		wantErr bool
	}{
		{
			name: "test order fixed match log",
			fields: fields{
				abiStr: nft.NftMetaData.ABI,
			},
			args: args{
				eventMsg:    readMockMatchOrderData(),
				contractAbi: abi.ABI{},
			},
			want: &indexer.FixedPriceMatchedEvent{
				EventBase: indexer.EventBase{
					EventDefinitionID: 1,
					BlockHash:         "0x3af9e8391251c9c5b233e06b5c4f03acbfaf951d1ae4f087701fb16b20805c94",
					TransactionHash:   "0x642bc63d3e54bfb31fc707188ed89d3379fa29484f3dfd10cacd32ff938e6a48",
					BlockNumber:       17398401,
					TransactionIndex:  0,
					LogIndex:          3,
					BlockTimestamp:    1659085959,
				},
				Maker:            "0x12fe20f224062bcce4ab2a30714571b77143525f",
				Taker:            "0x1cb3a2c6eef8d08f9a090dff0319f689199a471e",
				OrderHash:        "0x980751fdcc18a4df76f8f09124427bb3623027024e515703275c5090e97ab845",
				OrderBytes:       "0x2ff099ab253a425343254468bb9bcca9e79695c6edb1cf046f6b673e93cff6fc00000000000000000000000012fe20f224062bcce4ab2a30714571b77143525f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012fe20f224062bcce4ab2a30714571b77143525f0000000000000000000000000000000000000000000000000000000005f5e10000000000000000000000000000000000000000000000000000000000633afdaa00000000000000000000000000000000000000000000000000000000633afde6000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000011cd312adf18ed14743056b942f85a1e3cee5ef232cb3bd24235d47fa91cf016b",
				RoyaltyRecipient: "0x12fe20f224062bcce4ab2a30714571b77143525f",
				RoyaltyRate:      "100000000",
				StartTime:        1664810410,
				ExpireTime:       1664810470,
				MakerNonce:       "0",
				TakerGetNft:      1,
				Nft:              "0xe6e466a0bef75002545df0c68f61479bce970db6",
				Ft:               "0x0000000000000000000000000000000000000000",
				NftID:            "0x000000000000000000000000000000000000000000000000000000000000007b",
				NftAmount:        "0",
				FtAmount:         "2000000000000000000",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewMatchOrder(tt.fields.abiStr)
			event, err := p.Parse(tt.args.eventMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchOrderParse.parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(event, tt.want) {
				t.Errorf("MatchOrderParse.parse() real = %v, want %v", event, tt.want)
			}
		})
	}
}

func readMockMatchOrderData() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("../../mocks/match_order.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened match_order.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func mocksEventDefCache() {
	defs := []indexer.EventDefinition{}
	def := indexer.EventDefinition{
		Base: models.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ContractAddress: "0x37bbbec6cd25bb128887b39b1d9c922cddbd0987",
		EventSignature:  "0xf98e5acebea57eaddbc56ef62e1ad0c1fb3f8b3b06e16d83ef8e4ad578bdad52",
		ContractType:    "d1verse-dex",
		EventName:       "FixedPriceOrderMatched",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)

	def = indexer.EventDefinition{
		Base: models.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ContractAddress: "0x37bbbec6cd25bb128887b39b1d9c922cddbd0987",
		EventSignature:  "0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985",
		ContractType:    "d1verse-dex",
		EventName:       "OrderCancelled",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)

	def = indexer.EventDefinition{
		Base: models.Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		ContractAddress: "0x37bbbec6cd25bb128887b39b1d9c922cddbd0987",
		EventSignature:  "0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936888",
		ContractType:    "d1verse-dex",
		EventName:       "AllOrdersCancelled",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)

	cache.NewLocalEventCache().InsertCache(defs)
}
