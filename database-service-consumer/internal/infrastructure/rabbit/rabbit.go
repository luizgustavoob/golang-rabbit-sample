package rabbit

import (
	"log"

	"github.com/streadway/amqp"
)

type (
	Logger interface {
		Printf(format string, values ...interface{})
	}

	Rabbit struct {
		connection *amqp.Connection
		logger     Logger
	}
)

func (r *Rabbit) Consume(queueName string) <-chan amqp.Delivery {
	ch, err := r.connection.Channel()
	if err != nil {
		r.logger.Printf("Failed to open a channel: %s", err.Error())
		return nil
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil)       // args

	if err != nil {
		r.logger.Printf("Failed to declare a queue: %s", err.Error())
		return nil
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // autoAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil)        //args

	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	return msgs
}

func New(connection *amqp.Connection) *Rabbit {
	return &Rabbit{
		connection: connection,
	}
}
