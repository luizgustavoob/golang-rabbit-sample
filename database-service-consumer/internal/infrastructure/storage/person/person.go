package storage

import (
	"database/sql"
	"log"

	"github.com/golang-rabbit-sample/database-service-consumer/domain"
)

type personStorage struct {
	db *sql.DB
}

func NewPersonStorage(db *sql.DB) *personStorage {
	return &personStorage{
		db: db,
	}
}

func (self *personStorage) AddPerson(person *domain.Person) error {
	_, err := self.db.Exec(`INSERT INTO person(id, nome, idade, email, telefone) VALUES ($1, $2, $3, $4, $5)`,
		&person.ID, &person.Nome, &person.Idade, &person.Email, &person.Telefone)

	if err != nil {
		log.Printf("Failed to insert person: %s\n", err)
		return err
	}

	return nil
}
