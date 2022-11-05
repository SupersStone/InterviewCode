package trades

import (
	"testing"

	"dao-exchange/apps/indexer/dao"
	"dao-exchange/apps/indexer/dao/mocks"
	"dao-exchange/internal/cache"
	"dao-exchange/internal/models/indexer"
)

func TestCancelEvent_Handle(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}

	eventDao := dao.New(mocks.BuildDBCfg())
	defer eventDao.Close()

	initCache(*eventDao)

	type fields struct {
		Opr dao.Dao
	}
	type args struct {
		cancelEvent *indexer.NftexOrderCanceledEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test 1",
			args: args{
				cancelEvent: &indexer.NftexOrderCanceledEvent{
					EventBase: indexer.EventBase{
						EventDefinitionID: 2,
						BlockHash:         "0x6a8ee4d5aa8496317bb9765315f483bb89853d2965c68fba04155def6620ac58",
						TransactionHash:   "0x790a68c7f3130e8b9881f65d7eb16c59a2a7ffa70180cb91fc76e2fe3f1d4090",
						BlockNumber:       18604562,
						TransactionIndex:  0,
						LogIndex:          0,
						BlockTimestamp:    1662704444,
					},
					Maker:     "0xa7af7e741a6997d09d73d0f8d0719e590211248a",
					OrderHash: "0x14a79b5db8881de5bf6d458f49b83101c614b93531b3331e2939acde8e2e6aa1",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &CancelEvent{
				Opr: eventDao,
			}
			eventDef := cache.LoadEventDefWithID(tt.args.cancelEvent.EventDefinitionID)
			if _, err := e.HandleEvent(tt.args.cancelEvent, eventDef); (err != nil) != tt.wantErr {
				t.Errorf("CancelEvent.Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
