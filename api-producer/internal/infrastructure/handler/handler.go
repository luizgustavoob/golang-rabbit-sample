package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/fx"
)

type (
	Handler interface {
		GetMethod() string
		GetPattern() string
		http.Handler
	}

	HandlerParams struct {
		fx.In
		Handlers []Handler `group:"handlers"`
	}

	HandlerResult struct {
		fx.Out
		Handler Handler `group:"handlers"`
	}
)

func New(params HandlerParams) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	for _, h := range params.Handlers {
		router.Method(h.GetMethod(), h.GetPattern(), h)
	}

	return router
}
