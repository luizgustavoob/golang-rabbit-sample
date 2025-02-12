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

	HandlersIn struct {
		fx.In
		Handlers []Handler `group:"handlers"`
	}

	HandlerOut struct {
		fx.Out
		Handler Handler `group:"handlers"`
	}
)

func New(params HandlersIn) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	for _, h := range params.Handlers {
		r.Method(h.GetMethod(), h.GetPattern(), h)
	}

	return r
}
