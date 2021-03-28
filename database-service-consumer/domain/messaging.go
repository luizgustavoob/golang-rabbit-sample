package domain

import "github.com/streadway/amqp"

type Messaging interface {
	Consume(queueName string) <-chan amqp.Delivery
}
