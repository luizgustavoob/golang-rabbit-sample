package main

import (
	"log"
	"os"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(newLogger),
		app.Module,
		infrastructure.Module,
	).Run()
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}
