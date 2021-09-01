package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/internal/api"

	apimocks "github.com/golang-rabbit-sample/api-producer/internal/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_AddPerson(t *testing.T) {

	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "", log.LstdFlags)

	t.Run("should add person successfully", func(t *testing.T) {
		person := &api.Person{
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		personBytes, _ := json.Marshal(&person)

		publisher := new(apimocks.PublisherMock)
		publisher.On("Publish", mock.Anything, mock.Anything).Return(nil)
		marshaller := new(apimocks.MarshalMock)
		marshaller.On("SerializeJSON", mock.Anything).Return(personBytes, nil)
		service := api.NewService(publisher, marshaller, logger)

		p, err := service.AddPerson(person)

		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.NotEmpty(t, p.ID)
		assert.Equal(t, "name", p.Name)
		assert.Equal(t, 25, p.Age)
		assert.Equal(t, "email@email.com", p.Email)
		assert.Equal(t, "12345678", p.Phone)
		assert.Contains(t, buffer.String(), "success")
	})

	t.Run("should return error because invalid fields", func(t *testing.T) {
		service := api.NewService(nil, nil, nil)

		p, err := service.AddPerson(&api.Person{
			Name:  "",
			Age:   -1,
			Email: "",
			Phone: "",
		})

		assert.NotNil(t, err)
		assert.Nil(t, p)
	})

	t.Run("should return error because unexpected behavior on marshaller", func(t *testing.T) {
		marshaller := new(apimocks.MarshalMock)
		marshaller.On("SerializeJSON", mock.Anything).Return(nil, errors.New("marshal error"))
		service := api.NewService(nil, marshaller, logger)

		p, err := service.AddPerson(&api.Person{
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		})

		assert.Nil(t, p)
		assert.NotNil(t, err)
		assert.Equal(t, "marshal error", err.Error())
	})

	t.Run("should return error because unexpected behavior on publisher", func(t *testing.T) {
		person := &api.Person{
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		personBytes, _ := json.Marshal(&person)

		publisher := new(apimocks.PublisherMock)
		publisher.On("Publish", mock.Anything, mock.Anything).Return(errors.New("publisher error"))
		marshaller := new(apimocks.MarshalMock)
		marshaller.On("SerializeJSON", mock.Anything).Return(personBytes, nil)
		service := api.NewService(publisher, marshaller, logger)

		p, err := service.AddPerson(person)
		assert.Nil(t, p)
		assert.NotNil(t, err)
		assert.Equal(t, "publisher error", err.Error())
	})
}
