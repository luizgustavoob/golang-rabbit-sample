package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	connection *amqp.Connection
}

func NewRabbitMQ(user string, password string, hostname string, port int) *rabbitMQ {
	rabbit := &rabbitMQ{}
	rabbit.connect(user, password, hostname, port)
	return rabbit
}

func (self *rabbitMQ) connect(user string, password string, hostname string, port int) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, hostname, port))
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	self.connection = conn
}

func (self *rabbitMQ) Consume(queueName string) <-chan amqp.Delivery {
	ch, err := self.connection.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	queue, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil)       // args

	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // autoAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil)        //args

	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	return msgs
}
