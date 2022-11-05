package mqs

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"time"

	sigv4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/aws_msk_iam_v2"
	"github.com/segmentio/kafka-go/sasl/scram"
	"github.com/sirupsen/logrus"
)

// KafkaCfg config
type KafkaCfg struct {
	Brokers     []string `json:"brokers"`
	Topic       string   `json:"topic"`
	GroupTopics []string `json:"group_topics"`
	GroupID     string   `json:"group_id"`
}

// Producer producer
type Producer struct {
	*kafka.Writer
	cfg             KafkaCfg
	accessKeyID     string
	secretAccessKey string
}

// NewProducer new obj
func NewProducer(cfg KafkaCfg) *Producer {
	return &Producer{
		Writer: newKafkaWriter(cfg),
		cfg:    cfg,
	}
}

// NewProducerWithCerd new obj
func NewProducerWithCerd(cfg KafkaCfg, accessKeyID, secretAccessKey string) *Producer {
	writer := newKafkaWriterIAM(cfg, accessKeyID, secretAccessKey)

	return &Producer{
		Writer:          writer,
		cfg:             cfg,
		accessKeyID:     accessKeyID,
		secretAccessKey: secretAccessKey,
	}
}

// NewWriterWithCerd new obj
func (p *Producer) NewWriterWithCerd() *Producer {
	p.Writer = newKafkaWriterIAM(p.cfg, p.accessKeyID, p.secretAccessKey)
	return p
}

// KeepAlive keep living
func keepAlive(cfg KafkaCfg, accessKeyID, secretAccessKey string, writer *kafka.Writer) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			writer = newKafkaWriterIAM(cfg, accessKeyID, secretAccessKey)
		}
	}
}

func newKafkaWriter(cfg KafkaCfg) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(cfg.Brokers...),
		Topic:    cfg.Topic,
		Balancer: &kafka.Hash{},
	}
}

func newKafkaWriterIAM(cfg KafkaCfg, accessKeyID, secretAccessKey string) *kafka.Writer {
	creds := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")
	m := &aws_msk_iam_v2.Mechanism{
		Signer:      sigv4.NewSigner(),
		Credentials: creds,
		Region:      "us-east-1",
		SignTime:    time.Now(),
		Expiry:      time.Hour * 24,
	}

	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  cfg.Brokers,
		Topic:    cfg.Topic,
		Balancer: &kafka.Hash{},
		Dialer: &kafka.Dialer{
			Timeout:       50 * time.Second,
			KeepAlive:     60 * time.Second,
			DualStack:     true,
			SASLMechanism: m,
			TLS: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	})

}

func newCredentialMechanism(accessKeyID, secretAccessKey, region string) *aws_msk_iam_v2.Mechanism {
	creds := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")
	return &aws_msk_iam_v2.Mechanism{
		Signer:      sigv4.NewSigner(),
		Credentials: creds,
		Region:      region,
		SignTime:    time.Now(),
		Expiry:      time.Minute * 15,
	}
}

func newDialer(m *aws_msk_iam_v2.Mechanism) *kafka.Dialer {
	return &kafka.Dialer{
		Timeout:       50 * time.Second,
		DualStack:     true,
		SASLMechanism: m,
		TLS: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}
}

func newKafkaWriterSCRAM(cfg KafkaCfg) *kafka.Writer {
	mechanism, err := scram.Mechanism(scram.SHA512, "test-dao", "metadao@kafka")
	if err != nil {
		panic(err)
	}

	// Transports are responsible for managing connection pools and other resources,
	// it's generally best to create a few of these and share them across your
	// application.
	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}
	return &kafka.Writer{
		Addr:      kafka.TCP(cfg.Brokers...),
		Topic:     cfg.Topic,
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
	}
}

// PublishMsg msg
func (p *Producer) PublishMsg(key []byte, data []byte) error {
	logrus.Infof("push msg: %s from contract_address: %s, to topic: %s", string(data), hex.EncodeToString(key), p.cfg.Topic)
	msg := kafka.Message{
		Key:   key,
		Value: data,
	}
	return p.WriteMessages(context.Background(), msg)
}

func test(cfg KafkaCfg, secretAccessID, secretAccessKey string) {
	writer := newKafkaWriterIAM(cfg, secretAccessID, secretAccessKey)
	defer writer.Close()
	fmt.Println("start producing ... !!")
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		time.Sleep(1 * time.Second)
	}
}
