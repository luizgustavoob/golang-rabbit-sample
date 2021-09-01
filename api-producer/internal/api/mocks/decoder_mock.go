package apimocks

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type DecodeMock struct {
	mock.Mock
}

func (m *DecodeMock) DecodeJSON(r io.Reader, target interface{}) error {
	args := m.Called(r, target)
	return args.Error(0)
}
