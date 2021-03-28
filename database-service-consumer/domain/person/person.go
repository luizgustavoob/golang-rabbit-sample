package person

import (
	"github.com/golang-rabbit-sample/database-service-consumer/domain"
)

type service struct {
	storage domain.PersonStorage
}

func NewService(storage domain.PersonStorage) *service {
	return &service{
		storage: storage,
	}
}

func (self *service) AddPerson(person *domain.Person) error {
	return self.storage.AddPerson(person)
}
