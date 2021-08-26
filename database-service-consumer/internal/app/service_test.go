package app_test

import (
	"errors"
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	appmocks "github.com/golang-rabbit-sample/database-service-consumer/internal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService(t *testing.T) {

	t.Run("should return success", func(t *testing.T) {
		repo := new(appmocks.RepositoryMock)
		repo.On("AddPerson", mock.Anything).Return(nil)
		srv := app.NewService(repo)

		err := srv.AddPerson(&app.Person{
			ID:       "id",
			Nome:     "nome",
			Idade:    25,
			Email:    "email@gmail.com",
			Telefone: "12345678",
		})

		assert.Nil(t, err)
	})

	t.Run("should return error", func(t *testing.T) {
		repo := new(appmocks.RepositoryMock)
		repo.On("AddPerson", mock.Anything).Return(errors.New("repo error"))
		srv := app.NewService(repo)

		err := srv.AddPerson(&app.Person{
			ID:       "id",
			Nome:     "nome",
			Idade:    25,
			Email:    "email@gmail.com",
			Telefone: "12345678",
		})

		assert.NotNil(t, err)
		assert.Equal(t, "repo error", err.Error())
	})
}
