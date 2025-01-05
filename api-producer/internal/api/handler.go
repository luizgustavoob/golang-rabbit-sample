package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type (
	Service interface {
		AddPerson(person *Person) (*Person, error)
	}

	handler struct {
		service Service
	}
)

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetMethod() string {
	return http.MethodPost
}

func (h *handler) GetPattern() string {
	return "/people"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		writeErr(w, http.StatusBadRequest, err)
		return
	}

	newPerson, err := h.service.AddPerson(&person)
	if err != nil {
		sc := http.StatusInternalServerError
		if errors.Is(err, ErrInvalidPerson) {
			sc = http.StatusBadRequest
		}

		writeErr(w, sc, err)
		return
	}

	js, _ := json.Marshal(newPerson)

	w.WriteHeader(http.StatusCreated)
	w.Write(js)
}

func writeErr(w http.ResponseWriter, sc int, err error) {
	error := make(map[string]string)
	error["error"] = err.Error()
	js, _ := json.Marshal(error)

	w.WriteHeader(sc)
	w.Write(js)
}
