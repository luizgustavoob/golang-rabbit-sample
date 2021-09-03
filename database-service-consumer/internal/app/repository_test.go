package app_test

import (
	"bytes"
	"errors"
	"log"
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/golang-rabbit-sample/database-service-consumer/internal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRepository(t *testing.T) {

	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "", log.LstdFlags)

	t.Run("should exec with success", func(t *testing.T) {
		dbMock := new(mocks.DatabaseMock)
		dbMock.On("Exec", mock.Anything, mock.Anything).Return(nil)
		repo := app.NewRepository(logger, dbMock)

		err := repo.AddPerson(&app.Person{
			ID:       "id",
			Nome:     "nome",
			Idade:    25,
			Email:    "email@gmail.com",
			Telefone: "12345678",
		})

		assert.Nil(t, err)
	})

	t.Run("should return error on exec", func(t *testing.T) {
		dbMock := new(mocks.DatabaseMock)
		dbMock.On("Exec", mock.Anything, mock.Anything).Return(errors.New("db error"))
		repo := app.NewRepository(logger, dbMock)

		err := repo.AddPerson(&app.Person{
			ID:       "id",
			Nome:     "nome",
			Idade:    25,
			Email:    "email@gmail.com",
			Telefone: "12345678",
		})

		assert.NotNil(t, err)
		assert.Equal(t, "db error", err.Error())
		assert.Contains(t, buffer.String(), "Failed to insert")
	})
}
