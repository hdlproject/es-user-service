package messaging

import (
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/input_port"
	"github.com/hdlproject/es-user-service/use_case/interactor"
	"github.com/streadway/amqp"
)

type (
	topUpSubscriber struct {
		subscriber   *RabbitMQSubscriber
		queue        amqp.Queue
		exchangeName string
		routingKey   string
		topUpService *topUpService
	}
)

func NewTopUpSubscriber(subscriber *RabbitMQSubscriber,
	topUpUseCase *interactor.TopUp) (input_port.TopUpSubscriber, error) {

	exchangeName := "transactions_direct"
	routingKey := "top_up"

	err := subscriber.channel.ExchangeDeclare(
		exchangeName,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	queue, err := subscriber.channel.QueueDeclare(
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	err = subscriber.channel.QueueBind(
		queue.Name,
		routingKey,
		exchangeName,
		false,
		nil)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return &topUpSubscriber{
		subscriber:   subscriber,
		queue:        queue,
		exchangeName: exchangeName,
		routingKey:   routingKey,
		topUpService: newTopUpService(topUpUseCase),
	}, nil
}

func (instance *topUpSubscriber) Subscribe() (<-chan amqp.Delivery, error) {
	messages, err := instance.subscriber.channel.Consume(
		instance.queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return messages, nil
}

func (instance *topUpSubscriber) HandleMessage(message string) error {
	request, err := topUpRequest{}.parse(message)
	if err != nil {
		return helper.WrapError(err)
	}

	_, err = instance.topUpService.topup(request)
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}
