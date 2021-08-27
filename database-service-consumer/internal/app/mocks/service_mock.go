package appmocks

import (
	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) AddPerson(person *app.Person) error {
	args := m.Called(person)
	return args.Error(0)
}
