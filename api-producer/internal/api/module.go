package api

import (
	"go.uber.org/fx"

	app_handler "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/handler"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
)

func newService(publisher *rabbit.Rabbit) *service {
	return NewService(publisher)
}

func newHandler(service *service) app_handler.HandlerResult {
	return app_handler.HandlerResult{
		Handler: NewHandler(service),
	}
}

var Module = fx.Provide(
	newService,
	newHandler,
)
