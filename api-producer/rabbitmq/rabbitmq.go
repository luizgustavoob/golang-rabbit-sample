package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
}

func (self *RabbitMQ) Connect(user string, password string, hostname string, port int) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, hostname, port))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	self.connection = conn
}

func (self *RabbitMQ) Publish(queueName string, message string) {
	ch, err := self.connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	queue, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	err = ch.Publish("", queue.Name, false, false, amqp.Publishing{ContentType: "application/json", Body: []byte(message)})
	if err != nil {
		log.Fatalf("Failed to publish a message: %s", err)
	}
}
