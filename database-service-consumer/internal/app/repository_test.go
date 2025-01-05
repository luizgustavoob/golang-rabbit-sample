package app_test

import (
	"errors"
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRepository(t *testing.T) {
	t.Run("should exec with success", func(t *testing.T) {
		person := &app.Person{
			ID:    "id",
			Name:  "nome",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "12345678",
		}

		dbMock := mocks.NewDB(t)
		dbMock.On("Exec", mock.Anything, &person.ID, &person.Name, &person.Age, &person.Email, &person.Phone).Return(nil)

		repo := app.NewRepository(dbMock)
		err := repo.AddPerson(person)

		assert.NoError(t, err)
	})

	t.Run("should return error on exec", func(t *testing.T) {
		person := &app.Person{
			ID:    "id",
			Name:  "nome",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "12345678",
		}

		dbMock := mocks.NewDB(t)
		dbMock.On("Exec", mock.Anything, &person.ID, &person.Name, &person.Age, &person.Email, &person.Phone).Return(errors.New("db error"))

		repo := app.NewRepository(dbMock)
		err := repo.AddPerson(person)

		assert.Error(t, err)
	})
}
