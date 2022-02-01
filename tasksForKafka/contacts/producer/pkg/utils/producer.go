package utils

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func PostMessage(topic string, message []byte) {
	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	data := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic (%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
