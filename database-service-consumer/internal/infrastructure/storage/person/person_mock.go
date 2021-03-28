package storage

import "github.com/golang-rabbit-sample/database-service-consumer/domain"

type PersonStorageMock struct {
	AddPersonInvokedCount int
	AddPersonFn           func(*[]domain.Person, *domain.Person) error
	FakeDB                []domain.Person
}

func (self *PersonStorageMock) AddPerson(person *domain.Person) error {
	self.AddPersonInvokedCount++
	return self.AddPersonFn(&self.FakeDB, person)
}
