package postgres

import (
	"database/sql"
)

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) *Postgres {
	return &Postgres{
		db: db,
	}
}

func (p *Postgres) Exec(query string, args ...any) error {
	_, err := p.db.Exec(query, args...)
	return err
}
