package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"

	"github.com/nandhasuhendra/mock-debezium-kafka/internal/consumer"
)

func main() {
	godotenv.Load()

	var (
		brokers = []string{"localhost:9092"}
		groupID = os.Getenv("KAFKA_CONSUMER_GROUP")
		inTopic = os.Getenv("KAFKA_IN_TOPIC")
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		consumer.StartConsumer(ctx, brokers, groupID, inTopic)
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	<-sigchan

	log.Println("Shutting down gracefully...")
	cancel()
}
