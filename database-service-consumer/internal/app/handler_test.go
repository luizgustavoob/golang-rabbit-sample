package app_test

import (
	"bytes"
	"errors"
	"log"
	"testing"

	"github.com/golang-rabbit-sample/database-service-consumer/internal/app"
	appmocks "github.com/golang-rabbit-sample/database-service-consumer/internal/app/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler(t *testing.T) {

	buffer := &bytes.Buffer{}
	logger := log.New(buffer, "", log.LstdFlags)

	t.Run("should return success", func(t *testing.T) {
		decoder := new(appmocks.DecodeMock)
		decoder.On("DecodeJSON", mock.Anything, mock.Anything).Return(nil)
		service := new(appmocks.ServiceMock)
		service.On("AddPerson", mock.Anything).Return(nil)

		handler := app.NewHandler(logger, decoder, service)
		err := handler.HandleMessage([]byte(`{"chave": "content"}`))

		assert.Nil(t, err)
		decoder.AssertExpectations(t)
		service.AssertExpectations(t)
	})

	t.Run("should return error because unexpected behavior on decoder", func(t *testing.T) {
		decoder := new(appmocks.DecodeMock)
		decoder.On("DecodeJSON", mock.Anything, mock.Anything).Return(errors.New("decoder error"))

		handler := app.NewHandler(logger, decoder, nil)
		err := handler.HandleMessage([]byte(`{"chave": "content"}`))

		assert.NotNil(t, err)
		assert.Equal(t, "decoder error", err.Error())
		decoder.AssertExpectations(t)
	})

	t.Run("should return error because unexpected behavior on service", func(t *testing.T) {
		decoder := new(appmocks.DecodeMock)
		decoder.On("DecodeJSON", mock.Anything, mock.Anything).Return(nil)
		service := new(appmocks.ServiceMock)
		service.On("AddPerson", mock.Anything).Return(errors.New("service error"))

		handler := app.NewHandler(logger, decoder, service)
		err := handler.HandleMessage([]byte(`{"chave": "content"}`))

		assert.NotNil(t, err)
		assert.Equal(t, "service error", err.Error())
		decoder.AssertExpectations(t)
		service.AssertExpectations(t)
	})
}
