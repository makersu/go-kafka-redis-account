package main

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

var (
	brokers = []string{"127.0.0.1:9092"}
	topic   = "go-kafka-redis-account"
	topics  = []string{topic}
)

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func newKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return kafka
}
