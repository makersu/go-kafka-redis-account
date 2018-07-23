package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

func mainConsumer(partition int32) {
	kafka := newKafkaConsumer()
	defer kafka.Close()

	// consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)
	consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	go consumeEvents(consumer)

	fmt.Println("Press [enter] to exit consumer")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")
}

func newKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return consumer
}

func consumeEvents(consumer sarama.PartitionConsumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}
	var account *Account
	var err error

	for {
		select {
		case err := <-consumer.Errors():
			fmt.Printf("Kafka error: %s\n", err)
		case msg := <-consumer.Messages():
			msgVal = msg.Value
			if err = json.Unmarshal(msgVal, &log); err != nil {
				fmt.Printf("Failed parsing: %s", err)
			} else {
				logMap = log.(map[string]interface{})
				logType := logMap["Type"]
				fmt.Printf("Processing %s:\n%s\n", logMap["Type"], string(msgVal))

				switch logType {
				case "CreateEvent":
					event := new(CreateEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						account, err = event.Process()
					}
				default:
					fmt.Println("Unknown command: ", logType)
				}

				if err != nil {
					fmt.Printf("Error processing: %s\n", err)
				} else {
					fmt.Printf("Redis: %+v\n\n", *account)
				}
			}
		}
	}
}
