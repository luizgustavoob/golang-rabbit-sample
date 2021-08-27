package rabbit

import (
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

func (r *Rabbit) Publish(queueName string, message string) (err error) {
	ch, err := r.connection.Channel()
	if err != nil {
		r.logger.Printf("Failed to open a channel: %s", err.Error())
		return
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil)       // args

	if err != nil {
		r.logger.Printf("Failed to open a channel: %s", err.Error())
		return
	}

	err = ch.Publish(
		"",         // exchange
		queue.Name, // key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{ // message
			ContentType: "application/json",
			Body:        []byte(message),
		})

	if err != nil {
		r.logger.Printf("Failed to publish a message: %s", err.Error())
		return
	}

	return nil
}

func New(connection *amqp.Connection) *Rabbit {
	return &Rabbit{
		connection: connection,
	}
}
