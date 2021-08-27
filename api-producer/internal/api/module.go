package api

import (
	"log"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/decode"
	app_handler "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/handler"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/marshal"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
	"go.uber.org/fx"
)

func newService(publisher *rabbit.Rabbit, marshal *marshal.Marshal, logger *log.Logger) *service {
	return NewService(publisher, marshal, logger)
}

func newHandler(service *service, decoder *decode.Decode) app_handler.HandlerResult {
	return app_handler.HandlerResult{
		Handler: NewHandler(service, decoder),
	}
}

var Module = fx.Provide(
	newService,
	newHandler,
)
