package infrastructure

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/handler"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
)

func startServer(lc fx.Lifecycle, router http.Handler) {
	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			slog.Info("Starting server", slog.String("port", port))
			go srv.ListenAndServe()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			slog.Info("Stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}

var Module = fx.Options(
	rabbit.Module,
	handler.Module,
	fx.Invoke(startServer),
)
