package person

import (
	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	storage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage/person"
)

type PersonServiceMock struct {
	AddPersonInvokedCount int
	StorageFake           storage.PersonStorageMock
	AddPersonFn           func(*storage.PersonStorageMock, *domain.Person) error
}

func (self *PersonServiceMock) AddPerson(person *domain.Person) error {
	self.AddPersonInvokedCount++
	return self.AddPersonFn(&self.StorageFake, person)
}
