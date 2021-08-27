package appmocks

import "github.com/stretchr/testify/mock"

type DatabaseMock struct {
	mock.Mock
}

func (m *DatabaseMock) Exec(query string, args ...interface{}) error {
	arg := m.Called(query, args)
	return arg.Error(0)
}
