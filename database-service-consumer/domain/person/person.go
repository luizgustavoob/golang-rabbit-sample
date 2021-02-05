package person

import (
	"log"

	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	"github.com/golang-rabbit-sample/database-service-consumer/infrastructure/storage"
)

type service struct {
	postgres *storage.PostgresDB
}

func NewService(db *storage.PostgresDB) *service {
	return &service{postgres: db}
}

func (s *service) AddPerson(person *domain.Person) error {
	_, err := s.postgres.DB.Exec(`INSERT INTO person(id, nome, idade, email, telefone) VALUES ($1, $2, $3, $4, $5)`,
		&person.ID, &person.Nome, &person.Idade, &person.Email, &person.Telefone)

	if err != nil {
		log.Printf("Failed to insert person: %s\n", err)
		return err
	}

	log.Println("SUCCESS! Person was added")
	return nil
}
