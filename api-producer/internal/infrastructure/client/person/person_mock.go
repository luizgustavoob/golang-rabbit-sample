package client

import (
	"github.com/golang-rabbit-sample/api-producer/domain"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/rabbit"
)

type PersonClientMock struct {
	AddNewPersonInvokedCount int
	FakeRabbit               rabbit.RabbitMQMock
	AddNewPersonFn           func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error)
}

func (self *PersonClientMock) AddNewPerson(person *domain.Person) (*domain.Person, error) {
	self.AddNewPersonInvokedCount++
	return self.AddNewPersonFn(person, &self.FakeRabbit)
}
