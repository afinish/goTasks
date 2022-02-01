package subscribe

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/afinish/goTasks/tasksForKafka/contacts/consumer/pkg/models"
)

func Put() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := []string{"localhost:9092"}

	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	consumer, err := master.ConsumePartition("PUT", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	msgCount := 0

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				updateContact := &models.Contact{}
				message := string(msg.Value)
				log.Println([]byte(message))
				log.Println(json.Unmarshal([]byte(message), updateContact))
				models.UpdateContact(updateContact)

			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
}
