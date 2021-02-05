package main

import (
	"log"
	"os"
	"strconv"

	"github.com/golang-rabbit-sample/database-service-consumer/domain/person"
	"github.com/golang-rabbit-sample/database-service-consumer/infrastructure/client"
	"github.com/golang-rabbit-sample/database-service-consumer/infrastructure/storage"
)

func main() {
	consumer := client.NewRabbitMQ(getRabbitUser(), getRabbitPassword(), getRabbitHostName(), getRabbitPort())
	postgresDB, err := storage.NewConnection(getDatabase())
	if err != nil {
		return
	}
	defer postgresDB.DB.Close()
	service := person.NewService(postgresDB)
	personMonitor := client.NewPersonMonitor(service)

	forever := make(chan bool)
	people := consumer.Consume("person-queue")
	go personMonitor.StartMonitoring(people)
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

func getRabbitHostName() string {
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
