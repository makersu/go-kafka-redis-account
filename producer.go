package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

func mainProducer() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	kafka := newKafkaSyncProducer()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, "###")
		cmd := args[0]

		switch cmd {
		case "create":
			if len(args) == 2 {
				accName := args[1]
				event := NewCreateAccountEvent(accName)
				sendMsg(kafka, event)
			} else {
				fmt.Println("Only specify create###Account Name")
			}
		default:
			fmt.Printf("Unknown command %s, only: create, deposit, withdraw, transfer\n", cmd)
		}

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			err = nil
		}
	}
}

func sendMsg(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := kafka.SendMessage(msgLog)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
	}

	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n",
		partition, offset)

	return nil
}
