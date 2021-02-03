package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to check connection on PostgreSQL: %s", err)
	}

	log.Println("PostgreSQL connection ok!")
	return db
}

func Disconnect(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Failed to close connection on PostgreSQL: %s", err)
	}
}
