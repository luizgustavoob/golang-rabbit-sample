package rabbit

import (
	"log/slog"

	"github.com/streadway/amqp"
)

type (
	Serializable interface {
		Serialize() ([]byte, error)
	}

	Producer interface {
		Produce(queue string, message Serializable) error
	}

	producer struct {
		ch *amqp.Channel
	}
)

func NewProducer(ch *amqp.Channel) *producer {
	return &producer{ch: ch}
}

func (p *producer) Produce(queue string, message Serializable) error {
	msg, err := message.Serialize()
	if err != nil {
		slog.Error("Error serializing msg", slog.String("error", err.Error()))
		return err
	}

	err = p.ch.Publish(
		"",    // exchange
		queue, // key
		false, // mandatory
		false, // immediate
		amqp.Publishing{ // message
			ContentType: "application/json",
			Body:        msg,
		},
	)
	if err != nil {
		slog.Error("Error producing the message", slog.String("error", err.Error()))
		return err
	}

	return nil
}
