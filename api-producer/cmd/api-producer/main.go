package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure"
	"github.com/streadway/amqp"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	fx.New(
		api.Module,
		infrastructure.Module,
		fx.Invoke(startServer),
	).Run()
}

func startServer(lc fx.Lifecycle, handler http.Handler, conn *amqp.Connection, ch *amqp.Channel) {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(
		fx.StartStopHook(
			func(ctx context.Context) error {
				slog.Info("Starting server")
				go srv.ListenAndServe()
				return nil
			},
			func(ctx context.Context) error {
				slog.Info("Stopping server...")
				ch.Close()
				conn.Close()
				return srv.Shutdown(ctx)
			},
		),
	)
}
