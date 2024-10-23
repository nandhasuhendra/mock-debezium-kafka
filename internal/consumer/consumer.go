package consumer

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"

	"github.com/nandhasuhendra/mock-debezium-kafka/internal/processor"
	"github.com/nandhasuhendra/mock-debezium-kafka/internal/producer"
)

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	consumer.ready = make(chan bool)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)

		processedMessage, err := processor.ProcessMessage(message.Value)
		if err != nil {
			log.Printf("Error processing message: %v", err)
			continue
		}

		var (
			brokers  = []string{"localhost:9092"}
			outTopic = os.Getenv("KAFKA_OUT_TOPIC")
		)

		err = producer.PublishMessage(brokers, outTopic, processedMessage)
		if err != nil {
			log.Printf("Error publishing message: %v", err)
		} else {
			log.Printf("Message successfully published to debezium-topic")
		}

		session.MarkMessage(message, "")
	}
	return nil
}

func StartConsumer(ctx context.Context, brokers []string, groupID, topic string) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer := &Consumer{}
	client, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	log.Println("Kafka consumer is listening: localhost:9092")
	for {
		err := client.Consume(ctx, []string{topic}, consumer)
		if err != nil {
			log.Panicf("Error during consumption: %v", err)
		}
	}
}
