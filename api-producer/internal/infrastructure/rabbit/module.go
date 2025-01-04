package rabbit

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"go.uber.org/fx"

	"github.com/streadway/amqp"
)

func newRabbit() *Rabbit {
	user := os.Getenv("RABBIT_USER")
	pass := os.Getenv("RABBIT_PASS")
	hostname := os.Getenv("RABBIT_HOSTNAME")
	port, _ := strconv.Atoi(os.Getenv("RABBIT_PORT"))

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, pass, hostname, port))
	if err != nil {
		slog.Error("Error connecting to RabbitMQ", slog.String("error", err.Error()))
		fx.Error(err)
	}

	return New(conn)
}

var Module = fx.Provide(newRabbit)
