package erc721

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"time"

	"dao-exchange/apps/indexer/contracts/nft/token"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func init() {
	mocksEventDefCache()
}

func TestTransferParse_parse(t *testing.T) {
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
		want    *indexer.Erc721TransferEvent
		wantErr bool
	}{
		{
			name: "test token transfer log",
			fields: fields{
				abiStr: token.TokenMetaData.ABI,
			},
			args: args{
				eventMsg:    readMockTokenTransferData(),
				contractAbi: abi.ABI{},
			},
			want: &indexer.Erc721TransferEvent{
				EventBase: indexer.EventBase{
					EventDefinitionID: 1,
					BlockHash:         "0x92a44fa63c96f02836dab0ecd7dc6a4060ec1f7d43fdb8e75bf23193e535a85d",
					TransactionHash:   "0x99ff9d5c430603b96d764b3c55b15b8ac8a9c3e52c3c0a82a9321720669f6505",
					BlockNumber:       17915279,
					TransactionIndex:  1,
					LogIndex:          4,
					BlockTimestamp:    1660636595,
				},
				From:    "0x085601fd1a9e72d05010e60e171a6984d1c3c0f2",
				To:      "0x8b5b5cfa27bddead123e3e05e20f55162dfba64f",
				TokenID: "0x085601fd1a9e72d05010e60e171a6984d1c3c0f200003b9aca2d000000001ee3",
			},
			wantErr: false,
		},

		{
			name: "test token mint log",
			fields: fields{
				abiStr: token.TokenMetaData.ABI,
			},
			args: args{
				eventMsg:    readMockTokenTransferData(),
				contractAbi: abi.ABI{},
			},
			want: &indexer.Erc721TransferEvent{
				EventBase: indexer.EventBase{
					EventDefinitionID: 1,
					BlockHash:         "0x92a44fa63c96f02836dab0ecd7dc6a4060ec1f7d43fdb8e75bf23193e535a85d",
					TransactionHash:   "0x99ff9d5c430603b96d764b3c55b15b8ac8a9c3e52c3c0a82a9321720669f6505",
					BlockNumber:       17915279,
					TransactionIndex:  1,
					LogIndex:          4,
					BlockTimestamp:    1660636595,
				},
				From:    "0x085601fd1a9e72d05010e60e171a6984d1c3c0f2",
				To:      "0x8b5b5cfa27bddead123e3e05e20f55162dfba64f",
				TokenID: "0x085601fd1a9e72d05010e60e171a6984d1c3c0f200003b9aca2d000000001ee3",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewTransferParse(tt.fields.abiStr)
			event, err := p.Parse(tt.args.eventMsg)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferParse.parse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(event, tt.want) {
				t.Errorf("TransferParse.parse() real = %v, want %v", event, tt.want)
			}
		})
	}
}

func readMockTokenTransferData() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("../../mocks/token_transfer.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened token_transfer.json")
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
		ContractAddress: "0x83be785152a581e8f801b85e67cec3820a614d40",
		EventSignature:  "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		ContractType:    "erc721",
		EventName:       "Transfer",
		ChainID:         "80001",
		ChainName:       "PloygonTest",
	}
	defs = append(defs, def)
	cache.NewLocalEventCache().InsertCache(defs)
}
