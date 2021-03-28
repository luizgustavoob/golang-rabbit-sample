package person_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/domain"
	"github.com/golang-rabbit-sample/api-producer/domain/person"
	client "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/person"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/rabbit"
	"github.com/stretchr/testify/assert"
)

func TestPersonService_AddPerson(t *testing.T) {

	t.Run("should add person", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				*fakeQueue = append(*fakeQueue, *message)
				return nil
			},
		}

		clientMock := client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				personBytes, _ := json.Marshal(&person)
				fakeRabbit.Publish("queue-name", string(personBytes))
				return person, nil
			},
		}

		serviceMock := &person.ServiceMock{
			PersonClient: clientMock,
			AddPersonFn: func(person *domain.Person, client *client.PersonClientMock) (*domain.Person, error) {
				return client.AddNewPerson(person)
			},
		}

		person := &domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111",
		}

		p, err := serviceMock.AddPerson(person)

		personJs, _ := json.Marshal(&person)
		pJs, _ := json.Marshal(&p)

		assert.NotNil(t, p)
		assert.NoError(t, err)
		assert.Equal(t, string(personJs), string(pJs))
		assert.Equal(t, 1, serviceMock.AddPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.AddNewPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 1, len(serviceMock.PersonClient.FakeRabbit.FakeQueue))
	})

	t.Run("not should add person by error in service", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				return errors.New("Oooops, error!")
			},
		}

		clientMock := client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				return nil, errors.New("Ooops, error!")
			},
		}

		serviceMock := &person.ServiceMock{
			PersonClient: clientMock,
			AddPersonFn: func(person *domain.Person, client *client.PersonClientMock) (*domain.Person, error) {
				return nil, errors.New("Oooops, error!")
			},
		}

		p, err := serviceMock.AddPerson(&domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111"})

		assert.Nil(t, p)
		assert.Error(t, err)
		assert.Equal(t, 1, serviceMock.AddPersonInvokedCount)
		assert.Equal(t, 0, serviceMock.PersonClient.AddNewPersonInvokedCount)
		assert.Equal(t, 0, serviceMock.PersonClient.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 0, len(serviceMock.PersonClient.FakeRabbit.FakeQueue))
	})

	t.Run("not should add person by error in client", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				return errors.New("Oooops, error!")
			},
		}

		clientMock := client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				return nil, errors.New("Ooops, error!")
			},
		}

		serviceMock := &person.ServiceMock{
			PersonClient: clientMock,
			AddPersonFn: func(person *domain.Person, client *client.PersonClientMock) (*domain.Person, error) {
				return client.AddNewPerson(person)
			},
		}

		p, err := serviceMock.AddPerson(&domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111"})

		assert.Nil(t, p)
		assert.Error(t, err)
		assert.Equal(t, 1, serviceMock.AddPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.AddNewPersonInvokedCount)
		assert.Equal(t, 0, serviceMock.PersonClient.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 0, len(serviceMock.PersonClient.FakeRabbit.FakeQueue))
	})

	t.Run("not should add person by error in rabbit", func(t *testing.T) {
		rabbitMock := rabbit.RabbitMQMock{
			PublishFn: func(fakeQueue *[]string, queueName string, message *string) (err error) {
				return errors.New("Oooops, error!")
			},
		}

		clientMock := client.PersonClientMock{
			FakeRabbit: rabbitMock,
			AddNewPersonFn: func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error) {
				err := fakeRabbit.Publish("queue-name", "person-serialize")
				if err != nil {
					return nil, err
				}

				return person, nil
			},
		}

		serviceMock := &person.ServiceMock{
			PersonClient: clientMock,
			AddPersonFn: func(person *domain.Person, client *client.PersonClientMock) (*domain.Person, error) {
				return client.AddNewPerson(person)
			},
		}

		p, err := serviceMock.AddPerson(&domain.Person{
			ID:    "1",
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111"})

		assert.Nil(t, p)
		assert.Error(t, err)
		assert.Equal(t, 1, serviceMock.AddPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.AddNewPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 0, len(serviceMock.PersonClient.FakeRabbit.FakeQueue))
	})
}
