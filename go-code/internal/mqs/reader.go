package mqs

import (
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// Reader represents a kafka-go consumer group consumer
type Reader struct {
	*kafka.Reader
}

// NewReader new consumer
func NewReader(cfg KafkaCfg, accessKeyID, secretAccessKey, region string) *Reader {
	dialer := newDialer(newCredentialMechanism(accessKeyID, secretAccessKey, region))
	go keepDialer(accessKeyID, secretAccessKey, region, dialer)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     cfg.Brokers,
		GroupID:     cfg.GroupID,
		GroupTopics: cfg.GroupTopics,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		Dialer:      dialer,
	})
	return &Reader{
		Reader: r,
	}
}

// keepDialer keep living
func keepDialer(accessKeyID, secretAccessKey, region string, dialer *kafka.Dialer) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			dialer = newDialer(newCredentialMechanism(accessKeyID, secretAccessKey, region))
		}
	}
}

// Close close kafka
func (r *Reader) Close() {
	r.Close()
}
