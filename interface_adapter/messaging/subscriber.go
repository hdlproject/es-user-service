package messaging

import (
	"github.com/hdlproject/es-user-service/config"
	"log"
)

func Subscribe(eventBusConfig config.EventBus) {
	go func() {
		rabbitMQClient, _ := GetRabbitMQClient(eventBusConfig)
		eventSubscriber, _ := GetRabbitMQSubscriber(rabbitMQClient)

		transactionSubscriber, err := NewTransactionSubscriber(eventSubscriber)
		if err != nil {
			panic(err)
		}

		messages, err := transactionSubscriber.Subscribe()
		if err != nil {
			panic(err)
		}

		for message := range messages {
			log.Printf("Transaction Subscriber: %s", message.Body)
		}
	}()
}
