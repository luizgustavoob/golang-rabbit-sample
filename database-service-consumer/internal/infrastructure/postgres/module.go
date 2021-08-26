package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func newPostgres(logger *log.Logger) *Postgres {
	url := os.Getenv("DATABASE")

	db, err := sql.Open("postgres", url)
	if err != nil {
		logger.Printf("Failed to connect on PostgreSQL: %s", err.Error())
		fx.Error(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		logger.Printf("Failed to check connection on PostgreSQL: %s", err.Error())
		fx.Error(err)
		return nil
	}

	return New(db)
}

var Module = fx.Provide(newPostgres)
