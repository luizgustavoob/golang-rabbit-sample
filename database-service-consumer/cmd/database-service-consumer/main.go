package main

import (
	"log"
	"os"
	"strconv"

	pkgPersonService "github.com/golang-rabbit-sample/database-service-consumer/domain/person"
	pkgPersonClient "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/client/person"
	pkgStorage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage"
	pkgPersonStorage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage/person"
)

func main() {
	db, err := pkgStorage.NewConnection(getDatabase())
	if err != nil {
		return
	}
	defer db.Close()

	storage := pkgPersonStorage.NewPersonStorage(db)
	service := pkgPersonService.NewService(storage)
	client := pkgPersonClient.NewPersonMonitor(service)

	forever := make(chan bool)
	go client.StartMonitoring(getRabbitUser(), getRabbitPassword(), getRabbitHostName(), getRabbitPort())
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
