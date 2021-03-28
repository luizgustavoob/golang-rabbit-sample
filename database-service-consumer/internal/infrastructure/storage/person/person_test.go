package storage_test

import (
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	storage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage/person"
	"github.com/stretchr/testify/assert"
)

func TestStorage_AddPerson(t *testing.T) {

	t.Run("should add person on DB", func(t *testing.T) {
		storageMock := storage.PersonStorageMock{
			AddPersonFn: func(db *[]domain.Person, person *domain.Person) error {
				*db = append(*db, *person)
				return nil
			},
		}

		err := storageMock.AddPerson(&domain.Person{
			ID:       "1",
			Nome:     "Luiz",
			Idade:    25,
			Email:    "email@gmail.com",
			Telefone: "11111111",
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, storageMock.AddPersonInvokedCount)
		assert.Equal(t, 1, len(storageMock.FakeDB))

		p := storageMock.FakeDB[0]

		assert.Equal(t, "1", p.ID)
		assert.Equal(t, "Luiz", p.Nome)
		assert.Equal(t, 25, p.Idade)
		assert.Equal(t, "email@gmail.com", p.Email)
		assert.Equal(t, "11111111", p.Telefone)
	})
}
