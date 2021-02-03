package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-rabbit-sample/api-producer/rabbitmq"
	"github.com/golang-rabbit-sample/api-producer/routes"
)

var publisher *rabbitmq.RabbitMQ

func init() {
	publisher = &rabbitmq.RabbitMQ{}
	port, _ := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	publisher.Connect(os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"), os.Getenv("RABBIT_HOSTNAME"), port)
}

func main() {
	routes.SetPublisher(publisher)
	http.HandleFunc("/people", routes.AddPerson)
	log.Printf("Servidor iniciando na porta 8889...")
	log.Fatal(http.ListenAndServe(":8889", nil))
}
