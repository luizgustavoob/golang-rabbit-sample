package api

import (
	"go.uber.org/fx"

	app_handler "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/handler"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
)

var Module = fx.Provide(
	newService,
	newHandler,
)

func newService(producer rabbit.Producer) *service {
	return NewService(producer)
}

func newHandler(service *service) app_handler.HandlerOut {
	return app_handler.HandlerOut{
		Handler: NewHandler(service),
	}
}
