package postgres

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var Module = fx.Provide(newPostgres)

func newPostgres() *Postgres {
	url := os.Getenv("DATABASE")

	db, err := sql.Open("postgres", url)
	if err != nil {
		slog.Error("Error connecting to PostgreSQL", slog.String("error", err.Error()))
		fx.Error(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		slog.Error("Error checking the Postgres connection", slog.String("error", err.Error()))
		fx.Error(err)
		return nil
	}

	return New(db)
}
