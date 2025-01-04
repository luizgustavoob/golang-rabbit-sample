package main

import (
	"log/slog"
	"os"

	"go.uber.org/fx"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	fx.New(
		api.Module,
		infrastructure.Module,
	).Run()
}
