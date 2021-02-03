package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/golang-rabbit-sample/database-service-consumer/database"
	"github.com/golang-rabbit-sample/database-service-consumer/rabbitmq"
)

var (
	consumer *rabbitmq.RabbitMQ
	db       *sql.DB
)

func init() {
	consumer = &rabbitmq.RabbitMQ{}
	port, _ := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	consumer.Connect(os.Getenv("RABBIT_USER"), os.Getenv("RABBIT_PASS"), os.Getenv("RABBIT_HOSTNAME"), port)
	db = database.Connect(os.Getenv("DATABASE"))
}

func main() {
	defer database.Disconnect(db)
	forever := make(chan bool)
	people := consumer.Consume("person-table")
	go rabbitmq.MonitoringPerson(db, people)
	log.Println("Service running...")
	<-forever
}
