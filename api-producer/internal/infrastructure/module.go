package infrastructure

import (
	"go.uber.org/fx"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/handler"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
)

var Module = fx.Options(
	rabbit.Module,
	handler.Module,
)
