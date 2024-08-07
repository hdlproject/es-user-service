package messaging

import (
	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/interface_adapter/database"
	"github.com/hdlproject/es-user-service/use_case/interactor"
	"github.com/streadway/amqp"

	"log"
)

type (
	TopUpSubscriber interface {
		Subscribe() (<-chan amqp.Delivery, error)
		HandleMessage(message string) error
	}
)

func Subscribe() {
	go func() {
		configInstance, _ := config.GetInstance()

		rabbitMQClient, _ := GetRabbitMQClient(configInstance.EventBus)
		eventSubscriber, _ := GetRabbitMQSubscriber(rabbitMQClient)

		postgresClient, _ := database.GetPostgresClient(configInstance.Database)

		redisClient := database.GetRedisClient(configInstance.Redis)

		mongoClient, _ := database.GetMongoDB(configInstance.EventStorage)

		transactionSubscriber, err := NewTopUpSubscriber(
			eventSubscriber,
			interactor.NewTopUpUseCase(
				database.NewUserRepo(postgresClient, redisClient),
				database.NewTransactionEventRepo(mongoClient)),
		)
		if err != nil {
			panic(err)
		}

		messages, err := transactionSubscriber.Subscribe()
		if err != nil {
			panic(err)
		}

		for message := range messages {
			log.Printf("Transaction Subscriber: %s", message.Body)
			err = transactionSubscriber.HandleMessage(string(message.Body))
			if err != nil {
				log.Println(err)
			}
		}
	}()
}
