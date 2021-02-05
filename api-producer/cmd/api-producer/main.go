package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/golang-rabbit-sample/api-producer/domain/person"
	"github.com/golang-rabbit-sample/api-producer/infrastructure/client"
	"github.com/golang-rabbit-sample/api-producer/infrastructure/server"
)

func main() {
	publisher := client.NewRabbitMQ(getRabbitUser(), getRabbitPassword(), getRabbitHostName(), getRabbitPort())
	personClient := client.NewPersonClient(publisher)
	personService := person.NewService(personClient)
	handler := server.NewHandler(personService)
	server := server.New("8889", handler)
	server.ListenAndServe()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
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
