package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	i := 1
	for {
		send(conn, i)
		i++
		time.Sleep(time.Second * 1)
	}
}

func send(conn *kafka.Conn, counter int) {
	_, err := conn.WriteMessages(
		kafka.Message{Value: []byte(fmt.Sprintf("MESSAGE %d PUSHED!", counter))},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	log.Printf("Message Sent %d 🎉\n", counter)
}
