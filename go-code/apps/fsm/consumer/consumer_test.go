package consumer

import (
	"testing"

	"dao-exchange/apps/fsm/dao"
	daoMocks "dao-exchange/apps/fsm/dao/mocks"
	"dao-exchange/apps/indexer/mocks"
	"dao-exchange/config"
)

func Test_start(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}
	eventDao := dao.New(daoMocks.BuildDBCfg())
	defer eventDao.Close()

	mocks.EventDefCache()
	type args struct {
		cfg config.KafkaConf
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test for kafka",
			args: args{
				cfg: config.KafkaConf{
					Brokers:  []string{"tctest-kafka-t2-1.tc-jp1.huobiidc.com:9092", "tctest-kafka-t2-2.tc-jp1.huobiidc.com:9092", "tctest-kafka-t2-3.tc-jp1.huobiidc.com:9092"},
					Topics:   []string{"contract_event"},
					Group:    "local_indexer",
					Version:  "1.1.0",
					Assignor: "range",
					Oldest:   true,
					Verbose:  true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Start(tt.args.cfg, eventDao.DB())
		})
	}
}
