package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	apimocks "github.com/golang-rabbit-sample/api-producer/internal/api/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler(t *testing.T) {

	t.Run("should return success", func(t *testing.T) {
		person := &api.Person{
			ID:    "id",
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}
		personBytes, _ := json.Marshal(person)
		service := new(apimocks.ServiceMock)
		service.On("AddPerson", mock.Anything).Return(person, nil)
		decoder := new(apimocks.DecodeMock)
		decoder.On("DecodeRequestBody", mock.Anything, mock.Anything).Return(nil)
		handler := api.NewHandler(service, decoder)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/people", bytes.NewReader(personBytes))

		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("should return bad request error", func(t *testing.T) {
		decoder := new(apimocks.DecodeMock)
		decoder.On("DecodeRequestBody", mock.Anything, mock.Anything).Return(errors.New("decode error"))
		handler := api.NewHandler(nil, decoder)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/people", nil)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)

		body, _ := ioutil.ReadAll(res.Body)
		assert.Contains(t, string(body), "decode error")
	})

	t.Run("should return internal server error", func(t *testing.T) {
		service := new(apimocks.ServiceMock)
		service.On("AddPerson", mock.Anything).Return(nil, errors.New("service error"))
		decoder := new(apimocks.DecodeMock)
		decoder.On("DecodeRequestBody", mock.Anything, mock.Anything).Return(nil)
		handler := api.NewHandler(service, decoder)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/people", nil)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)

		body, _ := ioutil.ReadAll(res.Body)
		assert.Contains(t, string(body), "service error")
	})
}
