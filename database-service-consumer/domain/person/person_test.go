package person_test

import (
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	"github.com/golang-rabbit-sample/database-service-consumer/domain/person"
	pkgPersonStorage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage/person"
	"github.com/stretchr/testify/assert"
)

func TestPersonService_AddPerson(t *testing.T) {

	t.Run("should add person", func(t *testing.T) {
		storageMock := pkgPersonStorage.PersonStorageMock{
			AddPersonFn: func(fakeDB *[]domain.Person, person *domain.Person) error {
				*fakeDB = append(*fakeDB, *person)
				return nil
			},
		}

		serviceMock := person.PersonServiceMock{
			StorageFake: storageMock,
			AddPersonFn: func(fakeStorage *pkgPersonStorage.PersonStorageMock, person *domain.Person) error {
				return fakeStorage.AddPerson(person)
			},
		}

		person := &domain.Person{
			ID:       "1",
			Nome:     "Luiz",
			Idade:    25,
			Email:    "email@gmail.com",
			Telefone: "11111111",
		}

		err := serviceMock.AddPerson(person)

		assert.NoError(t, err)
		assert.Equal(t, 1, serviceMock.AddPersonInvokedCount)
		assert.Equal(t, 1, serviceMock.StorageFake.AddPersonInvokedCount)
		assert.Equal(t, 1, len(serviceMock.StorageFake.FakeDB))
	})
}
