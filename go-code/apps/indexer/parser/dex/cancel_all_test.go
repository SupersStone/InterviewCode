package dex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"dao-exchange/apps/indexer/contracts/nft"
	"dao-exchange/internal/models/indexer"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestCancelAllOrder_parse(t *testing.T) {
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
		want    *indexer.NftexAllOrderCanceledEvent
		wantErr bool
	}{
		{
			name: "test order fixed match log",
			fields: fields{
				abiStr: nft.NftMetaData.ABI,
			},
			args: args{
				eventMsg:    []byte{},
				contractAbi: abi.ABI{},
			},
			want: &indexer.NftexAllOrderCanceledEvent{
				EventBase: indexer.EventBase{
					EventDefinitionID: 1,
					BlockHash:         "0x2759df2da4c14bca01d8307ba5acab37e27489a6fbceaff44d102f91813f42b3",
					TransactionHash:   "0xdee8ee781d846e4f438ee6735cde0c7f5816e07cac19b7b2559495627addbce8",
					BlockNumber:       18346467,
					TransactionIndex:  0,
					LogIndex:          0,
					BlockTimestamp:    1661930159,
				},
				Maker: "0x8bd0a3f1b798a824cf384f2f502056eaa93c9b23",
				Nonce: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewCancelAllOrder(tt.fields.abiStr)
			msg := readMockCancelAllOrderData()
			event, err := p.Parse(msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CancelOrder.parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(event, tt.want) {
				t.Errorf("CancelOrder.parse() real = %v, want %v", event, tt.want)
			}
		})
	}
}

func readMockCancelAllOrderData() []byte {
	// Open our jsonFile
	jsonFile, err := os.Open("../../mocks/orders/cancel_all.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened cancel_all.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
