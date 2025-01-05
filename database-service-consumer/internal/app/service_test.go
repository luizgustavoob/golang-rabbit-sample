package app_test

import (
	"errors"
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/app/mocks"
	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	t.Run("should return success", func(t *testing.T) {
		person := &app.Person{
			ID:    "id",
			Name:  "nome",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "12345678",
		}

		repo := mocks.NewRepository(t)
		repo.On("AddPerson", person).Return(nil)

		srv := app.NewService(repo)
		err := srv.AddPerson(person)

		assert.NoError(t, err)
	})

	t.Run("should return error", func(t *testing.T) {
		person := &app.Person{
			ID:    "id",
			Name:  "nome",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "12345678",
		}

		repo := mocks.NewRepository(t)
		repo.On("AddPerson", person).Return(errors.New("repo error"))

		srv := app.NewService(repo)
		err := srv.AddPerson(person)

		assert.Error(t, err)
	})
}
