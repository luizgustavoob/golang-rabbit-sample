package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/golang-rabbit-sample/api-producer/internal/api/mocks"
)

func TestHandler(t *testing.T) {
	t.Run("should create handler", func(t *testing.T) {
		service := mocks.NewService(t)
		handler := api.NewHandler(service)

		assert.Equal(t, "/people", handler.GetPattern())
		assert.Equal(t, http.MethodPost, handler.GetMethod())
	})

	t.Run("should return status created", func(t *testing.T) {
		person := &api.Person{
			ID:    "id",
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		service := mocks.NewService(t)
		service.On("AddPerson", mock.MatchedBy(func(p *api.Person) bool {
			return p.Name == person.Name && p.Age == person.Age &&
				p.Email == person.Email && p.Phone == person.Phone
		})).Return(person, nil)

		personJS, err := json.Marshal(person)
		require.NoError(t, err)

		res := httptest.NewRecorder()

		req, err := http.NewRequest("POST", "/people", bytes.NewReader(personJS))
		require.NoError(t, err)

		handler := api.NewHandler(service)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
	})

	t.Run("should return status bad request", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/people", strings.NewReader(`invalid`))

		handler := api.NewHandler(nil)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return status bad request due to invalid field", func(t *testing.T) {
		service := mocks.NewService(t)
		service.On("AddPerson", mock.MatchedBy(func(p *api.Person) bool {
			return p.Name == "" && p.Age == 29 && p.Email == "abc@email.com" && p.Phone == "1234"
		})).Return(nil, api.ErrInvalidPerson)

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/people",
			strings.NewReader(`{"nome":"","idade":29,"email":"abc@email.com","telefone":"1234"}`),
		)

		handler := api.NewHandler(service)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return internal server error", func(t *testing.T) {
		service := mocks.NewService(t)
		service.On("AddPerson", mock.MatchedBy(func(p *api.Person) bool {
			return p.Name == "fulano" && p.Age == 29 && p.Email == "abc@email.com" && p.Phone == "1234"
		})).Return(nil, errors.New("something wrong has happened"))

		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/people",
			strings.NewReader(`{"nome":"fulano","idade":29,"email":"abc@email.com","telefone":"1234"}`),
		)

		handler := api.NewHandler(service)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}
