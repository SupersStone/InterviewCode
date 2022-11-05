package consumer

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"dao-exchange/apps/fsm/handler"
	"dao-exchange/apps/fsm/handler/event"
	"dao-exchange/config"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// VerifyKafkaCfg verify Kafka config, if invalid panic.
func VerifyKafkaCfg(cfg config.KafkaConf) {
	if len(cfg.Brokers) == 0 {
		panic("no Kafka bootstrap brokers defined, please set the -brokers flag")
	}

	if len(cfg.Topics) == 0 {
		panic("no topics given to be consumed, please set the -topics flag")
	}

	if len(cfg.Group) == 0 {
		panic("no Kafka consumer group defined, please set the -group flag")
	}
}

// Start consumer from kafka
func Start(cfg config.KafkaConf, db *gorm.DB) {
	VerifyKafkaCfg(cfg)

	keepRunning := true
	log.Println("Starting a new Sarama consumer")

	if cfg.Verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(cfg.Version)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	saramaCfg := sarama.NewConfig()
	saramaCfg.Version = version

	switch cfg.Assignor {
	case "sticky":
		saramaCfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		saramaCfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		saramaCfg.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		log.Panicf("Unrecognized consumer group partition assignor: %s", cfg.Assignor)
	}

	if cfg.Oldest {
		saramaCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	/**
	 * Setup a new Sarama consumer group
	 */
	consumer := Consumer{
		ready: make(chan bool),
		db:    db,
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.Group, saramaCfg)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, cfg.Topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		log.Println("Resuming consumption")
	} else {
		client.PauseAll()
		log.Println("Pausing consumption")
	}

	*isPaused = !*isPaused
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready   chan bool
	db      *gorm.DB
	actions map[string]event.Actions
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		resp, err := handler.ConsumeMsg(message.Value, nil)
		if err != nil {
			logrus.Warnf("EventHandler: not match event err = %s", err.Error())
			session.MarkMessage(message, "")
			continue
		}

		if err := handler.SyncToDB(consumer.db, resp); err != nil {
			logrus.Warn("ConsumeClaim Process DB Error")
			continue
		}

		/*
			处理消息成功后标记为处理, 然后会自动提交
			consumerConfig.Consumer.Offsets.AutoCommit.Enable = true

			如果设置false 需要加
			session.Commit()
		*/
		session.MarkMessage(message, "")
	}

	return nil
}
