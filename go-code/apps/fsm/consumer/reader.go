package consumer

import (
	"context"
	"dao-exchange/apps/common"
	"dao-exchange/apps/fsm/handler"
	"dao-exchange/config"
	"dao-exchange/internal/mqs"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Reader represents a kafka-go consumer group consumer
type Reader struct {
	*mqs.Reader
	db      *gorm.DB
	actions map[string]*handler.MsgAction
}

// NewReader new consumer
func NewReader(cfg mqs.KafkaCfg, cred config.AWSCredential, db *gorm.DB, actions map[string]*handler.MsgAction) *Reader {
	return &Reader{
		Reader:  mqs.NewReader(cfg, cred.AccessKeyID, cred.SecretAccessKey, cred.Region),
		db:      db,
		actions: actions,
	}
}

// Start received msg
func (r *Reader) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Millisecond)
	for {
		select {
		case <-ticker.C:
			m, err := r.FetchMessage(ctx)
			if err != nil {
				break
			}

			resp, err := handler.ConsumeMsg(m.Value, r.actions)
			if err != nil && (errors.Is(err, common.ErrNotFoundEventDef) || errors.Is(err, common.ErrNotSupportEvent) || errors.Is(err, common.ErrUnMarshalFail)) {
				logrus.Warnf("EventHandler: not match event err = %s", err.Error())
				if err := r.CommitMessages(ctx, m); err != nil {
					logrus.Warnf("EventHandler: CommitMessages fail")
				}

				continue
			}

			if err := handler.SyncToDB(r.db, resp); err != nil {
				logrus.Warn("ConsumeClaim Process DB Error")
				continue
			}

			if err := r.CommitMessages(ctx, m); err != nil {
				continue
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		case <-ctx.Done():
			r.Close()
			ticker.Stop()
		}
	}

}
