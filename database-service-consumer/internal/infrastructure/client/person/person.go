package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	pkgRabbit "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/client/rabbit"
)

type personMonitor struct {
	service domain.PersonService
}

func NewPersonMonitor(service domain.PersonService) *personMonitor {
	return &personMonitor{service}
}

func (self *personMonitor) StartMonitoring(user string, password string, hostname string, port int) {
	rabbit := pkgRabbit.NewRabbitMQ(user, password, hostname, port)

	for msgPerson := range rabbit.Consume("person-queue") {
		log.Println(fmt.Sprintf("MENSAGEM: %s", msgPerson.Body))

		person := &domain.Person{}
		err := json.Unmarshal([]byte(msgPerson.Body), &person)
		if err != nil {
			log.Fatalf("Failed to unmarshal person: %s", err)
		}

		self.service.AddPerson(person)
	}
}
