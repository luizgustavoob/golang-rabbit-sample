package rabbit

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/streadway/amqp"
	"go.uber.org/fx"
)

func newRabbit(logger *log.Logger) *Rabbit {
	user := os.Getenv("RABBIT_USER")
	pass := os.Getenv("RABBIT_PASS")
	hostname := os.Getenv("RABBIT_HOSTNAME")
	port, _ := strconv.Atoi(os.Getenv("RABBIT_PORT"))

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, pass, hostname, port))
	if err != nil {
		logger.Printf("Failed to connect on RabbitMQ: %s", err.Error())
		fx.Error(err)
		return nil
	}

	return New(conn)
}

var Module = fx.Provide(newRabbit)
