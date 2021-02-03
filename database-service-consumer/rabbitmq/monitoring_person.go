package rabbitmq

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-rabbit-sample/database-service-consumer/models"
	"github.com/streadway/amqp"
)

func MonitoringPerson(db *sql.DB, people <-chan amqp.Delivery) {
	for msgPerson := range people {
		log.Println(fmt.Sprintf("MENSAGEM: %s", msgPerson.Body))

		person := &models.Person{}
		err := json.Unmarshal([]byte(msgPerson.Body), &person)
		if err != nil {
			log.Fatalf("Failed to unmarshal person: %s", err)
		}

		_, err = db.Exec(`INSERT INTO person(id, nome, idade, email, telefone) VALUES ($1, $2, $3, $4, $5)`,
			person.ID, person.Nome, person.Idade, person.Email, person.Telefone)

		if err != nil {
			log.Fatalf("Failed to insert person: %s", err)
		}

		log.Println("SUCCESS! Person was added")
	}
}
