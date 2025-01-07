package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/rabbit"
	"github.com/streadway/amqp"
	"go.uber.org/fx"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	fx.New(
		app.Module,
		infrastructure.Module,
		fx.Invoke(startApp),
	).Run()
}

func startApp(lc fx.Lifecycle, consumerIn rabbit.ConsumersIn, conn *amqp.Connection, ch *amqp.Channel) {
	lc.Append(
		fx.StartHook(func(ctx context.Context) error {
			slog.Info("Service running...")
			for _, c := range consumerIn.Consumers {
				c.Start(ctx, ch)
			}
			return nil
		}),
	)

	lc.Append(
		fx.StopHook(func(c context.Context) error {
			slog.Info("Service stopping...")
			ch.Close()
			conn.Close()
			return nil
		}),
	)
}
