package rabbit

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/streadway/amqp"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	newConnection,
	newChannel,
)

type (
	ConsumersIn struct {
		fx.In
		Consumers []Consumer `group:"consumers"`
	}

	ConsumersOut struct {
		fx.Out
		Consumer Consumer `group:"consumers"`
	}
)

func newConnection() *amqp.Connection {
	user := os.Getenv("RABBIT_USER")
	pass := os.Getenv("RABBIT_PASS")
	hostname := os.Getenv("RABBIT_HOSTNAME")
	port, err := strconv.Atoi(os.Getenv("RABBIT_PORT"))
	if err != nil {
		fx.Error(err)
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, pass, hostname, port))
	if err != nil {
		slog.Error("Error connecting to RabbitMQ", slog.String("error", err.Error()))
		fx.Error(err)
		return nil
	}

	return conn
}

func newChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		slog.Error("Error opening the AMQP channel", slog.String("error", err.Error()))
		fx.Error(err)
		return nil
	}

	return ch
}
