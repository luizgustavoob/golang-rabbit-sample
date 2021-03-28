package server_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/domain"
	"github.com/golang-rabbit-sample/api-producer/domain/person"
	client "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/person"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/rabbit"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/server"
	"github.com/stretchr/testify/assert"
)

const inputPerson = `
	{
		"nome": "Daiane",
		"idade": 26,
		"email": "email@gmail.com",
		"telefone": "11111111"
	}
`

func TestPersonHandler_AddPerson(t *testing.T) {

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
				personJs, _ := json.Marshal(person)
				err := fakeRabbit.Publish("queue-name", string(personJs))
				if err != nil {
					return nil, err
				}

				return person, nil
			},
		}

		serviceMock := &person.ServiceMock{
			PersonClient: clientMock,
			AddPersonFn: func(person *domain.Person, client *client.PersonClientMock) (*domain.Person, error) {
				person.ID = "abcd1234"
				return client.AddNewPerson(person)
			},
		}

		server := httptest.NewServer(server.NewHandler(serviceMock))
		defer server.Close()

		URL, _ := url.Parse(server.URL)
		req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/people", URL), strings.NewReader(inputPerson))
		res, err := http.DefaultClient.Do(req)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)

		responsePost, err := ioutil.ReadAll(res.Body)
		assert.NoError(t, err)

		var person *domain.Person
		err = json.Unmarshal(responsePost, &person)
		assert.NoError(t, err)

		assert.True(t, person.IsValid())

		assert.Equal(t, 1, serviceMock.AddPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.AddNewPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.PersonClient.FakeRabbit.PublishInvokedCount)
		assert.Equal(t, 1, len(serviceMock.PersonClient.FakeRabbit.FakeQueue))

		message := serviceMock.PersonClient.FakeRabbit.FakeQueue[0]

		assert.Equal(t, string(responsePost), message)
	})
}
