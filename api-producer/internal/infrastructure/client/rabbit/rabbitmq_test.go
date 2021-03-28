package rabbit_test

import (
	"errors"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/rabbit"
	"github.com/stretchr/testify/assert"
)

func TestRabbit_Publish(t *testing.T) {

	t.Run("should publish a message successfully", func(t *testing.T) {
		rabbitMock := &rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				*fakeQueue = append(*fakeQueue, *message)
				return nil
			},
		}

		err := rabbitMock.Publish("queue-name", "new-message")
		assert.NoError(t, err)
		assert.Equal(t, 1, rabbitMock.PublishInvokedCount)
		assert.Equal(t, 1, len(rabbitMock.FakeQueue))

		message := rabbitMock.FakeQueue[0]
		assert.Equal(t, "new-message", message)
	})

	t.Run("not should publish a message", func(t *testing.T) {
		mock := &rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				return errors.New("Ooops, error!")
			},
		}

		err := mock.Publish("queue-name", "message")
		assert.Error(t, err)
		assert.Equal(t, 1, mock.PublishInvokedCount)
		assert.Equal(t, 0, len(mock.FakeQueue))
	})
}
