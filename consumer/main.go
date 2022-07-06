package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

var (
	BROKERS = []string{"localhost:9092"}
)

const (
	TOPIC     = "MESSAGE_TOPIC"
	PARTITION = 0
)

func main() {
	conn := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   BROKERS,
		Topic:     TOPIC,
		Partition: PARTITION,
	})
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()
	for {
		msg, err := conn.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("GET MESSAGE: %s\n", string(msg.Value))
	}
}
