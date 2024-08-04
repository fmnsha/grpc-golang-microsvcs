package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	// if err != nil {
	//     log.Fatal("failed to dial leader:", err)
	// }

	// to consume messages
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID: "consumer-group-id",
		Topic:   "topic-A",
		//Partition: 0,
		MaxBytes: 10e6, // 10MB
	})
	//r.SetOffset(42)
	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		break
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
