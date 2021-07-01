package person

import (
	"github.com/golang-rabbit-sample/api-producer/domain"
	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/client/rabbit"
)

type PersonServiceMock struct {
	AddPersonInvokedCount int
	FakeRabbit            rabbit.RabbitMQMock
	AddPersonFn           func(person *domain.Person, fakeRabbit *rabbit.RabbitMQMock) (*domain.Person, error)
}

func (self *PersonServiceMock) AddPerson(person *domain.Person) (*domain.Person, error) {
	self.AddPersonInvokedCount++
	return self.AddPersonFn(person, &self.FakeRabbit)
}
