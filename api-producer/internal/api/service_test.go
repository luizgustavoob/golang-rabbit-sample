package api_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/golang-rabbit-sample/api-producer/internal/api/mocks"
)

func TestService_AddPerson(t *testing.T) {
	t.Run("should add person successfully", func(t *testing.T) {
		person := &api.Person{
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		publisher := mocks.NewPublisher(t)
		publisher.On("Publish", mock.Anything, person).Return(nil)

		service := api.NewService(publisher)
		p, err := service.AddPerson(person)

		assert.NoError(t, err)
		assert.NotEmpty(t, p.ID)
		assert.Equal(t, "name", p.Name)
		assert.Equal(t, 25, p.Age)
		assert.Equal(t, "email@email.com", p.Email)
		assert.Equal(t, "12345678", p.Phone)
	})

	t.Run("should return error due to invalid fields", func(t *testing.T) {
		service := api.NewService(nil)
		p, err := service.AddPerson(&api.Person{
			Name:  "",
			Age:   -1,
			Email: "",
			Phone: "",
		})

		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("should return error due to unexpected behavior on publisher", func(t *testing.T) {
		person := &api.Person{
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		expectedErr := errors.New("something wrong has happened")

		publisher := mocks.NewPublisher(t)
		publisher.On("Publish", mock.Anything, person).Return(expectedErr)

		service := api.NewService(publisher)
		p, err := service.AddPerson(person)

		assert.Nil(t, p)
		assert.Equal(t, expectedErr, err)
	})
}
