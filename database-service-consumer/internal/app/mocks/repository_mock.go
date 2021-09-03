package mocks

import (
	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) AddPerson(person *app.Person) error {
	args := m.Called(person)
	return args.Error(0)
}
