package rabbit

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/streadway/amqp"
)

type (
	Consumer interface {
		Start(context.Context, *amqp.Channel)
	}

	Processor[T any] func(T) error

	consumer[T any] struct {
		queue     string
		processor Processor[T]
	}
)

func NewConsumer[T any](queue string, processor Processor[T]) *consumer[T] {
	return &consumer[T]{
		queue:     queue,
		processor: processor,
	}
}

func (c *consumer[T]) Start(ctx context.Context, ch *amqp.Channel) {
	go func() {
		msgs, err := ch.Consume(
			c.queue, // queue
			"",      // consumer
			true,    // autoAck
			false,   // exclusive
			false,   // noLocal
			false,   // noWait
			nil,     // args
		)
		if err != nil {
			slog.Error("Error starting consumer", slog.String("error", err.Error()))
			return
		}

		sem := newSemaphore(semaphoreSizePerConsumer)

		for {
			select {
			case <-ctx.Done():
				sem.Close()
				return

			case msg, ok := <-msgs:
				if !ok {
					sem.Close()
					return
				}

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
		}
	}()
}

const semaphoreSizePerConsumer = 20

type semaphore struct {
	ch chan bool
}

func newSemaphore(size int) *semaphore {
	return &semaphore{ch: make(chan bool, size)}
}

func (s *semaphore) Acquire() { s.ch <- true }

func (s *semaphore) Release() { <-s.ch }

func (s *semaphore) Close() { close(s.ch) }
