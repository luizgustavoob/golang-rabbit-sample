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

	Processor[T any] func(T) error

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

		sem := newSemaphore(semaphoreSize)

		for msg := range msgs {
			sem.Acquire()

			go func(msg amqp.Delivery, sem *semaphore) {
				defer sem.Release()

				var msgIn T
				if err := json.Unmarshal(msg.Body, &msgIn); err != nil {
					slog.Error("Error parsing queue msg", slog.String("error", err.Error()))
					return
				}

				if err := c.processor(msgIn); err != nil {
					slog.Error("Error processing msg", slog.String("error", err.Error()))
				}
			}(msg, sem)
		}
	}()
}

const semaphoreSize = 30

type semaphore struct {
	ch chan bool
}

func newSemaphore(size int) *semaphore {
	return &semaphore{ch: make(chan bool, size)}
}

func (s *semaphore) Acquire() { s.ch <- true }

func (s *semaphore) Release() { <-s.ch }
