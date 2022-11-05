package trades

import (
	"testing"

	"dao-exchange/apps/indexer/dao"
	"dao-exchange/apps/indexer/dao/mocks"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/indexer"
)

func TestFixedPriceMatchEvent_Handle(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}

	eventDao := dao.New(mocks.BuildDBCfg())
	defer eventDao.Close()

	initCache(*eventDao)

	type args struct {
		matchEvent *indexer.FixedPriceMatchedEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				matchEvent: &indexer.FixedPriceMatchedEvent{
					EventBase: indexer.EventBase{
						EventDefinitionID: 5,
						BlockHash:         "0x3204313ce6ddd36649d3d6282035ecb76c8b2068f43e78d94d29489facabeb57",
						TransactionHash:   "0xc24286df253edc7d13df5e152e06db8503e066309c16f9e299ea525687de3940",
						BlockNumber:       18577196,
						TransactionIndex:  0,
						LogIndex:          3,
						BlockTimestamp:    1662622346,
					},
					Maker:            "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd",
					Taker:            "0xa7af7e741a6997d09d73d0f8d0719e590211248a",
					OrderHash:        "0xbd76f8646e6328695346d199655a042d468aa41d842e03198585f77b667b20be",
					OrderBytes:       "0xf864ff4b08b0229d60f2daf2b5c71d60922e1604d52cfbee43d5b0101c1b3fb3000000000000000000000000e8f5b7da307c6c0c1a963eb6443810249920d2dd0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e8f5b7da307c6c0c1a963eb6443810249920d2dd000000000000000000000000e8f5b7da307c6c0c1a963eb6443810249920d2dd0000000000000000000000000000000000000000000000000000000002faf08000000000000000000000000000000000000000000000000000000000631894c60000000000000000000000000000000000000000000000000000000064c3b5e9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001a95e3a69572d8935dbe7c8e297118a664ccb8049ee6f0720cdc8a1ef03e8892d",
					RoyaltyRecipient: "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd",
					RoyaltyRate:      "50000000",
					StartTime:        1662555334,
					ExpireTime:       1690547689,
					MakerNonce:       "0x0",
					TakerGetNft:      1,
					Nft:              "0xe48538159e84aa92d0d9b92e4d789d4fa91208c8",
					Ft:               "0x29e18f4ae8026e3e03561e3a6252bba93286a9e8",
					NftID:            "0xe8f5b7da307c6c0c1a963eb6443810249920d2dd00003b9aca3a000000000030",
					NftAmount:        "0",
					FtAmount:         "20000000000000000000",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &FixedPriceMatchEvent{
				Opr: eventDao,
			}
			eventDef := cache.LoadEventDefWithID(tt.args.matchEvent.EventDefinitionID)
			if _, err := e.HandleEvent(tt.args.matchEvent, eventDef); (err != nil) != tt.wantErr {
				t.Errorf("FixedPriceMatchEvent.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
