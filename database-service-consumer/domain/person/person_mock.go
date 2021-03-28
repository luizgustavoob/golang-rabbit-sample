package person

import (
	"github.com/golang-rabbit-sample/database-service-consumer/domain"
	pkgStorage "github.com/golang-rabbit-sample/database-service-consumer/internal/infrastructure/storage/person"
)

type PersonServiceMock struct {
	AddPersonInvokedCount int
	StorageFake           pkgStorage.PersonStorageMock
	AddPersonFn           func(*pkgStorage.PersonStorageMock, *domain.Person) error
}

func (self *PersonServiceMock) AddPerson(person *domain.Person) error {
	self.AddPersonInvokedCount++
	return self.AddPersonFn(&self.StorageFake, person)
}
