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
		Publish(queue string, message Serializable) error
	}

	publisher struct {
		queueName string
		ch        *amqp.Channel
	}
)

func NewPublisher(queueName string, ch *amqp.Channel) *publisher {
	return &publisher{
		queueName: queueName,
		ch:        ch,
	}
}

func (p *publisher) Publish(queueName string, message Serializable) error {
	msg, err := message.Serialize()
	if err != nil {
		slog.Error("Error serializing msg", slog.String("error", err.Error()))
		return err
	}

	err = p.ch.Publish(
		"",        // exchange
		queueName, // key
		false,     // mandatory
		false,     // immediate
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
