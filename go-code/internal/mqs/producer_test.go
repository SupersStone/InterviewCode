package mqs

import (
	"testing"
)

func Test_test(t *testing.T) {
	if testing.Short() {
		t.Skip("short模式下会跳过该测试用例")
	}

	type args struct {
		cfg             KafkaCfg
		secretAccessID  string
		secretAccessKey string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				cfg: KafkaCfg{
					Brokers: []string{
						"b-3.msknftdao.15jz9z.c18.kafka.us-east-1.amazonaws.com:9098",
						"b-2.msknftdao.15jz9z.c18.kafka.us-east-1.amazonaws.com:9098",
						"b-1.msknftdao.15jz9z.c18.kafka.us-east-1.amazonaws.com:9098",
					},
					Topic: "contract_event_log",
				},
				secretAccessID:  "AKIAVJM3MCTDQN3BVBWX",
				secretAccessKey: "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",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test(tt.args.cfg, tt.args.secretAccessID, tt.args.secretAccessKey)
		})
	}
}
