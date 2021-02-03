package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-rabbit-sample/api-producer/models"
	"github.com/golang-rabbit-sample/api-producer/rabbitmq"
)

var publisher *rabbitmq.RabbitMQ

func SetPublisher(publ *rabbitmq.RabbitMQ) {
	publisher = publ
}

func AddPerson(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writer.Header().Set("Allow", http.MethodPost)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	var person *models.Person
	err = json.Unmarshal(bodyBytes, &person)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	person.GenerateID()
	personBytes, err := json.Marshal(person)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	publisher.Publish("person-table", string(personBytes))

	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, string(personBytes))
}
