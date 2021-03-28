package client_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/domain"
	client "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/person"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/rabbit"
	"github.com/stretchr/testify/assert"
)

func TestPerson_AddNewPerson(t *testing.T) {

	t.Run("should add new person successfully", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				*fakeQueue = append(*fakeQueue, *message)
				return nil
			},
		}

		clientMock := &client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				personBytes, _ := json.Marshal(&person)
				err := fakeRabbit.Publish("queue-name", string(personBytes))
				if err != nil {
					return nil, err
				}
				return person, nil
			},
		}

		person := &domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111",
		}

		newPerson, err := clientMock.AddNewPerson(person)

		personJs, _ := json.Marshal(person)
		newPersonJs, _ := json.Marshal(newPerson)

		assert.NoError(t, err)
		assert.NotNil(t, newPerson)
		assert.Equal(t, string(newPersonJs), string(personJs))
		assert.Equal(t, 1, clientMock.AddNewPersonInvokedCount)
		assert.Equal(t, 1, len(clientMock.FakeRabbit.FakeQueue))
	})

	t.Run("not should add new person by error in client", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				return errors.New("Oooops, error!")
			},
		}

		clientMock := &client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				return nil, errors.New("Ooops, error!")
			},
		}

		newPerson, err := clientMock.AddNewPerson(&domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111",
		})

		assert.Error(t, err)
		assert.Nil(t, newPerson)
		assert.Equal(t, 1, clientMock.AddNewPersonInvokedCount)
		assert.Equal(t, 0, clientMock.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 0, len(clientMock.FakeRabbit.FakeQueue))
	})

	t.Run("not should add new person by error in rabbit", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				return errors.New("Oooops, error!")
			},
		}

		clientMock := &client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				err := fakeRabbit.Publish("queue-name", "person-serialize")
				if err != nil {
					return nil, err
				}

				return person, nil
			},
		}

		newPerson, err := clientMock.AddNewPerson(&domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111",
		})

		assert.Error(t, err)
		assert.Nil(t, newPerson)
		assert.Equal(t, 1, clientMock.AddNewPersonInvokedCount)
		assert.Equal(t, 1, clientMock.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 0, len(clientMock.FakeRabbit.FakeQueue))
	})
}
