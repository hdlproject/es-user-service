package main

import (
	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/interface_adapter/api"
	"github.com/hdlproject/es-user-service/interface_adapter/database"
	"github.com/hdlproject/es-user-service/interface_adapter/messaging"
	"github.com/hdlproject/es-user-service/interface_adapter/security"
)

func init() {
	configInstance, err := config.GetInstance()
	if err != nil {
		panic(err)
	}

	_, err = database.GetPostgresClient(configInstance.Database)
	if err != nil {
		panic(err)
	}

	rabbitMQClient, err := messaging.GetRabbitMQClient(configInstance.EventBus)
	if err != nil {
		panic(err)
	}

	_, err = messaging.GetRabbitMQPublisher(rabbitMQClient)
	if err != nil {
		panic(err)
	}

	_, err = messaging.GetRabbitMQSubscriber(rabbitMQClient)
	if err != nil {
		panic(err)
	}

	_, err = database.GetMongoDB(configInstance.EventStorage)
	if err != nil {
		panic(err)
	}

	kmsClient, err := security.NewKMSClient(configInstance.AWS.ID, configInstance.AWS.Secret)
	if err != nil {
		panic(err)
	}

	_ = security.NewJWT(kmsClient)

	_, err = messaging.GetCentrifugeClient(configInstance.Centrifuge.ServerUrl, configInstance.Centrifuge.Token, "")
	if err != nil {
		panic(err)
	}
}

func main() {
	configInstance, _ := config.GetInstance()

	messaging.Subscribe()

	httpServer := api.GetHTTPServer(configInstance.Port)
	api.RegisterUserAPI()
	httpServer.Serve()
}
