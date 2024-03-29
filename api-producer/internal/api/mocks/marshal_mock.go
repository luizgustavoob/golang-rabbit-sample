package mocks

import "github.com/stretchr/testify/mock"

type MarshalMock struct {
	mock.Mock
}

func (m *MarshalMock) SerializeJSON(obj interface{}) ([]byte, error) {
	args := m.Called(obj)
	arg0 := args.Get(0)
	if arg0 != nil {
		return arg0.([]byte), nil
	}
	return nil, args.Error(1)
}
