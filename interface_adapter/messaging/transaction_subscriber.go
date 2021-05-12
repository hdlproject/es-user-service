package messaging

import (
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/input_port"
	"github.com/streadway/amqp"
)

type (
	transactionSubscriber struct {
		subscriber *RabbitMQSubscriber
		queue      amqp.Queue
	}
)

func NewTransactionSubscriber(subscriber *RabbitMQSubscriber) (input_port.TransactionSubscriber, error) {
	err := subscriber.channel.ExchangeDeclare(
		"transactions_topic",
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

	return &transactionSubscriber{
		subscriber: subscriber,
		queue:      queue,
	}, nil
}

func (instance *transactionSubscriber) Subscribe() (<-chan amqp.Delivery, error) {
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
