package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

const (
	TOPIC     = "MESSAGE_TOPIC"
	PARTITION = 0
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", TOPIC, PARTITION)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()
	for {
		msg, err := conn.ReadMessage(10e3)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("GET MESSAGE: %s\n", string(msg.Value))
	}
}
