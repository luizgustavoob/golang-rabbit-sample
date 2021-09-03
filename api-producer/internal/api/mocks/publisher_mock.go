package mocks

import "github.com/stretchr/testify/mock"

type PublisherMock struct {
	mock.Mock
}

func (m *PublisherMock) Publish(queueName string, message string) error {
	args := m.Called(queueName, message)
	return args.Error(0)
}
