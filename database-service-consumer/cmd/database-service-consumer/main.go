package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	domain "github.com/golang-rabbit-sample/database-service-consumer/domain"
	service "github.com/golang-rabbit-sample/database-service-consumer/domain/person"
	messaging "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/client/rabbit"
	postgres_storage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage"
	storage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage/person"
)

func main() {
	db, err := postgres_storage.NewConnection(getDatabase())
	if err != nil {
		return
	}
	defer db.Close()

	personStorage := storage.NewPersonStorage(db)
	personService := service.NewService(personStorage)

	forever := make(chan bool)
	go func(user string, password string, hostname string, port int) {
		personQueue := messaging.NewRabbitMQ(user, password, hostname, port).Consume("person-queue")

		for msgPerson := range personQueue {
			log.Println(fmt.Sprintf("MENSAGEM: %s", msgPerson.Body))
			person := &domain.Person{}
			err := json.Unmarshal([]byte(msgPerson.Body), &person)
			if err != nil {
				log.Fatalf("Failed to unmarshal person: %s", err)
			}
			go personService.AddPerson(person)
		}

	}(getRabbitUser(), getRabbitPassword(), getRabbitHostname(), getRabbitPort())

	log.Println("Service running...")
	<-forever
}

func getRabbitUser() string {
	user := os.Getenv("RABBIT_USER")
	if user == "" {
		return "guest"
	}
	return user
}

func getRabbitPassword() string {
	pass := os.Getenv("RABBIT_PASS")
	if pass == "" {
		return "guest"
	}
	return pass
}

func getRabbitHostname() string {
	hostname := os.Getenv("RABBIT_HOSTNAME")
	if hostname == "" {
		return "localhost"
	}
	return hostname
}

func getRabbitPort() int {
	portStr := os.Getenv("RABBIT_PORT")
	if portStr == "" {
		return 5672
	}

	port, _ := strconv.Atoi(portStr)
	return port
}

func getDatabase() string {
	database := os.Getenv("DATABASE")
	if database == "" {
		return "host=localhost port=5440 user=postgres password=postgres dbname=mydb sslmode=disable"
	}
	return database
}
