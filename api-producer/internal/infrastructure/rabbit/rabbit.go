package rabbit

import (
	"log/slog"

	"github.com/streadway/amqp"
)

type (
	Serializable interface {
		Serialize() ([]byte, error)
	}

	Publisher interface {
		Publish(queue string, msg Serializable) error
	}

	Rabbit struct {
		connection *amqp.Connection
	}
)

func New(connection *amqp.Connection) *Rabbit {
	return &Rabbit{
		connection: connection,
	}
}

func (r *Rabbit) Publish(queueName string, message Serializable) error {
	ch, err := r.connection.Channel()
	if err != nil {
		slog.Error("Error opening the channel", slog.String("error", err.Error()))
		return err
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)
	if err != nil {
		slog.Error("Error declaring queue", slog.String("error", err.Error()))
		return err
	}

	msg, err := message.Serialize()
	if err != nil {
		slog.Error("Error serializing msg", slog.String("error", err.Error()))
		return err
	}

	err = ch.Publish(
		"",         // exchange
		queue.Name, // key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{ // message
			ContentType: "application/json",
			Body:        msg,
		},
	)
	if err != nil {
		slog.Error("Error publishing the message", slog.String("error", err.Error()))
		return err
	}

	return nil
}
