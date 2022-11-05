package trades

import (
	"testing"

	"dao-exchange/apps/indexer/dao"
	"dao-exchange/apps/indexer/dao/mocks"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/indexer"
)

func TestTransferEvent_handle(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}

	eventDao := dao.New(mocks.BuildDBCfg())
	defer eventDao.Close()

	initCache(*eventDao)

	type args struct {
		transferEvent *indexer.Erc721TransferEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test mint",
			args: args{
				transferEvent: &indexer.Erc721TransferEvent{
					EventBase: indexer.EventBase{
						EventDefinitionID: 4,
						BlockHash:         "0x3204313ce6ddd36649d3d6282035ecb76c8b2068f43e78d94d29489facabeb57",
						TransactionHash:   "0xc24286df253edc7d13df5e152e06db8503e066309c16f9e299ea525687de3940",
						BlockNumber:       18577196,
						TransactionIndex:  0,
						LogIndex:          0,
						BlockTimestamp:    1662622346,
					},
					From:    "0x0000000000000000000000000000000000000000",
					To:      "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd",
					TokenID: "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd00003b9aca3a000000000030",
				},
			},
			wantErr: false,
		},
		{
			name: "test transfer",
			args: args{
				transferEvent: &indexer.Erc721TransferEvent{
					EventBase: indexer.EventBase{
						EventDefinitionID: 4,
						BlockHash:         "0x3204313ce6ddd36649d3d6282035ecb76c8b2068f43e78d94d29489facabeb57",
						TransactionHash:   "0xc24286df253edc7d13df5e152e06db8503e066309c16f9e299ea525687de3940",
						BlockNumber:       18577196,
						TransactionIndex:  0,
						LogIndex:          0,
						BlockTimestamp:    1662622346,
					},
					From:    "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd",
					To:      "0xa7af7e741a6997d09d73d0f8d0719e590211248a",
					TokenID: "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd00003b9aca3a000000000030",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &TransferEvent{
				Opr: eventDao,
			}
			eventDef := cache.LoadEventDefWithID(tt.args.transferEvent.EventDefinitionID)
			if _, err := e.HandleEvent(tt.args.transferEvent, eventDef); (err != nil) != tt.wantErr {
				t.Errorf("TransferEvent.handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func initCache(opr dao.Dao) {
	cache.NewLocalEventCache().UpdateCache(opr.DB())
	cache.NewPayTokenCache().UpdateCache(opr.DB())
}
