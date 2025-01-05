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

type ServiceParams struct {
	fx.In
	Publisher rabbit.Publisher `name:"person-publisher"`
}

func newService(params ServiceParams) *service {
	return NewService(params.Publisher)
}

func newHandler(service *service) app_handler.HandlerResult {
	return app_handler.HandlerResult{
		Handler: NewHandler(service),
	}
}
