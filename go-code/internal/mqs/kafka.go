package mqs

import (
	"encoding/hex"
	"time"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// ProducerCfg config
type ProducerCfg struct {
	Brokers        []string `json:"brokers"`
	Topic          string   `json:"topic"`
	Version        string   `json:"version"`
	FlushFrequency int      `json:"flush_frequency"`
	FlushMessages  int      `json:"flush_messages"`
}

// InitProducer init kafka producer
func InitProducer(conf ProducerCfg) (sarama.AsyncProducer, error) {
	kafkaConf := sarama.NewConfig()
	kafkaConf.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConf.Producer.Partitioner = sarama.NewHashPartitioner
	kafkaConf.Producer.Return.Errors = true
	// Flush batches every 4ms
	kafkaConf.Producer.Flush.Frequency = 4 * time.Millisecond
	kafkaConf.Producer.Compression = sarama.CompressionSnappy
	// 最大消息大小: 4M
	kafkaConf.Producer.MaxMessageBytes = 4 * 1024 * 1024

	if conf.FlushFrequency > 0 {
		kafkaConf.Producer.Flush.Frequency = time.Duration(conf.FlushFrequency) * time.Millisecond
	}

	kafkaConf.Producer.Flush.Messages = 16
	if conf.FlushMessages > 0 {
		kafkaConf.Producer.Flush.Messages = conf.FlushMessages
	}

	producer, err := sarama.NewAsyncProducer(conf.Brokers, kafkaConf)
	if err != nil {
		return nil, errors.Wrapf(err, "init kafka producer failed, config: %+v", conf)
	}

	return producer, nil
}

// SaramaProducer kafka producer
type SaramaProducer struct {
	sarama.AsyncProducer
	topic string
}

// NewAsyncProducer new obj
func NewAsyncProducer(topic string, asyncProducer sarama.AsyncProducer) *SaramaProducer {
	return &SaramaProducer{
		AsyncProducer: asyncProducer,
		topic:         topic,
	}
}

// PublishMsg msg
func (p *SaramaProducer) PublishMsg(key []byte, data []byte) {
	logrus.Infof("push msg: %s from contract_address: %s, to topic: %s", data, hex.EncodeToString(key), p.topic)
	msg := sarama.ProducerMessage{
		Topic: p.topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(data),
	}
	p.Input() <- &msg
}
