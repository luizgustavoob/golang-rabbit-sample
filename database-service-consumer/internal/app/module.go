package app

import (
	"context"
	"log"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/decode"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/postgres"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/rabbit"
	"go.uber.org/fx"
)

const (
	queueName = "person-queue"
)

func newRepository(logger *log.Logger, db *postgres.Postgres) *repository {
	return NewRepository(logger, db)
}

func newService(repository *repository) *service {
	return NewService(repository)
}

func newHandler(logger *log.Logger, decoder *decode.Decode, service *service) *handler {
	return NewHandler(logger, decoder, service)
}

func setHooks(lc fx.Lifecycle, rabbit *rabbit.Rabbit, handler *handler, logger *log.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			logger.Println("Service running...")
			go func() {
				msgs := rabbit.Consume(queueName)
				for msg := range msgs {
					go handler.HandleMessage(msg.Body)
				}
			}()
			return nil
		},
		OnStop: func(c context.Context) error {
			logger.Println("Service stopping...")
			return nil
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		newRepository,
		newService,
		newHandler,
	),
	fx.Invoke(setHooks),
)
