package main

import (
	"log"
	"os"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(newLogger),
		api.Module,
		infrastructure.Module,
	).Run()
}

func newLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags)
}
