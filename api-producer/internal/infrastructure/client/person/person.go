package client

import (
	"encoding/json"
	"log"

	"github.com/golang-rabbit-sample/api-producer/domain"
)

type personClient struct {
	publisher domain.Messaging
}

func NewPersonClient(publisher domain.Messaging) *personClient {
	return &personClient{publisher}
}

func (self *personClient) AddNewPerson(person *domain.Person) (*domain.Person, error) {
	personBytes, err := json.Marshal(&person)
	if err != nil {
		log.Fatalf("Failed in marshaling person: %s", err)
		return nil, err
	}

	err = self.publisher.Publish("person-queue", string(personBytes))
	if err != nil {
		log.Fatalf("Failed to publish message in queue: %s", err)
		return nil, err
	}

	return person, nil
}
