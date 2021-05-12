package input_port

import "github.com/streadway/amqp"

type (
	TransactionSubscriber interface {
		Subscribe() (<-chan amqp.Delivery, error)
	}
)
