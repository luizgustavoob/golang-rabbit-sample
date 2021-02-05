package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	"github.com/streadway/amqp"
)

type personMonitor struct {
	service domain.PersonService
}

func NewPersonMonitor(service domain.PersonService) *personMonitor {
	return &personMonitor{service}
}

func (self *personMonitor) StartMonitoring(people <-chan amqp.Delivery) {
	for msgPerson := range people {
		log.Println(fmt.Sprintf("MENSAGEM: %s", msgPerson.Body))

		person := &domain.Person{}
		err := json.Unmarshal([]byte(msgPerson.Body), &person)
		if err != nil {
			log.Fatalf("Failed to unmarshal person: %s", err)
		}

		self.service.AddPerson(person)
	}
}
