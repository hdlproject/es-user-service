package input_port

import "github.com/streadway/amqp"

type (
	TopUpSubscriber interface {
		Subscribe() (<-chan amqp.Delivery, error)
		HandleMessage(message string) error
	}
)
