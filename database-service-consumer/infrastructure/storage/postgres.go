package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewConnection(url string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %s", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to check connection on PostgreSQL: %s", err)
		return nil, err
	}

	log.Println("PostgreSQL connection ok!")
	return &PostgresDB{DB: db}, nil
}
