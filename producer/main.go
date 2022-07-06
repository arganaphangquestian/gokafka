package main

import (
	"context"
	"fmt"
	"log"
	"time"

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
	conn := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      BROKERS,
		Topic:        TOPIC,
		BatchTimeout: time.Millisecond, // Set this to millisecond, because the default is 1 second
	})
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	i := 1
	ctx := context.Background()
	for {
		send(conn, ctx, i)
		i++
		time.Sleep(time.Second)
	}
}

func send(conn *kafka.Writer, ctx context.Context, counter int) {
	err := conn.WriteMessages(ctx,
		kafka.Message{Value: []byte(fmt.Sprintf("MESSAGE %d PUSHED!", counter))},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	log.Printf("Message Sent %d 🎉\n", counter)
}
