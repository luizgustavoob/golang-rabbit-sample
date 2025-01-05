package app

import (
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/postgres"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/rabbit"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		newRepository,
		newService,
		newConsumer,
	),
)

func newRepository(db *postgres.Postgres) *repository {
	return NewRepository(db)
}

func newService(repository *repository) *service {
	return NewService(repository)
}

func newConsumer(service *service) rabbit.ConsumersOut {
	return rabbit.ConsumersOut{
		Consumer: rabbit.NewConsumer(rabbit.PersonQueue.String(), AddPersonFn(service)),
	}
}
