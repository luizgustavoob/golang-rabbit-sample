package person

import (
	"github.com/golang-rabbit-sample/api-producer/domain"
	client "github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/person"
)

type ServiceMock struct {
	AddPersonInvokedCount int
	PersonClient          client.PersonClientMock
	AddPersonFn           func(person *domain.Person, client *client.PersonClientMock) (*domain.Person, error)
}

func (self *ServiceMock) AddPerson(person *domain.Person) (*domain.Person, error) {
	self.AddPersonInvokedCount++
	return self.AddPersonFn(person, &self.PersonClient)
}
