package rabbit

import (
	"encoding/json"
	"log/slog"

	"github.com/streadway/amqp"
)

type (
	Consumer interface {
		Start(*amqp.Channel)
	}

	Processor[T any] func(T)

	consumer[T any] struct {
		queueName string
		processor Processor[T]
	}
)

func NewConsumer[T any](queueName string, processor Processor[T]) *consumer[T] {
	return &consumer[T]{
		queueName: queueName,
		processor: processor,
	}
}

func (c *consumer[T]) Start(ch *amqp.Channel) {
	go func() {
		msgs, err := ch.Consume(
			c.queueName, // queue
			"",          // consumer
			true,        // autoAck
			false,       // exclusive
			false,       // noLocal
			false,       // noWait
			nil,         // args
		)
		if err != nil {
			slog.Error("Error starting consumer", slog.String("error", err.Error()))
			return
		}

		for msg := range msgs {
			var msgIn T
			if err := json.Unmarshal(msg.Body, &msgIn); err != nil {
				slog.Error("Error parsing queue msg", slog.String("error", err.Error()))
				return
			}

			go c.processor(msgIn)
		}
	}()
}
