package mocks

import (
	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) AddPerson(person *api.Person) (*api.Person, error) {
	args := m.Called(person)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.(*api.Person), nil
	}
	return nil, args.Error(1)
}
