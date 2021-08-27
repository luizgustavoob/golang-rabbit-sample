package infrastructure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/decode"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/handler"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/marshal"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
	"go.uber.org/fx"
)

func startServer(lc fx.Lifecycle, handler http.Handler, logger *log.Logger) {
	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Printf("Starting server at port %s\n", port)
			go srv.ListenAndServe()
			return nil
		},

		OnStop: func(ctx context.Context) error {
			logger.Println("Stopping server...")
			return srv.Shutdown(ctx)
		},
	})
}

var Module = fx.Options(
	marshal.Module,
	decode.Module,
	rabbit.Module,
	handler.Module,
	fx.Invoke(startServer),
)
