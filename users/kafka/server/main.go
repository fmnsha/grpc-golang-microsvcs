package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {

	// to produce messages
	// topic := "topic-A"
	// partition := 0

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9094", topic, partition)
	// if err != nil {
	// 	log.Fatal("failed to dial leader:", err)
	// }

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	// if err != nil {
	// 	log.Fatal("failed to write messages:", err)
	// }

	// if err := conn.Close(); err != nil {
	// 	log.Fatal("failed to close writer:", err)
	// }

	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:    "topic-A",
		Balancer: &kafka.LeastBytes{},
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			//Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			//Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			//Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
