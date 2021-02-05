package person

import (
	"log"

	"github.com/golang-rabbit-sample/api-producer/domain"
)

type service struct {
	client domain.PersonClient
}

func NewService(client domain.PersonClient) *service {
	return &service{client}
}

func (s *service) AddPerson(person *domain.Person) (*domain.Person, error) {
	log.Println("Sending person to queue..")
	person.GenerateID()
	return s.client.AddNewPerson(person)
}
