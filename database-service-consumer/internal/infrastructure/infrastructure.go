package infrastructure

import (
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/decode"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/postgres"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/rabbit"
	"go.uber.org/fx"
)

var Module = fx.Options(
	decode.Module,
	postgres.Module,
	rabbit.Module,
)
