package person

import (
	"encoding/json"
	"log"

	"github.com/golang-rabbit-sample/api-producer/domain"
)

type service struct {
	publisher domain.Messaging
}

func NewService(publisher domain.Messaging) *service {
	return &service{publisher}
}

func (self *service) AddPerson(person *domain.Person) (*domain.Person, error) {
	log.Println("Sending person to queue..")
	person.GenerateID()

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
