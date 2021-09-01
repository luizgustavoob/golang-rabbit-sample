package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type (
	Decoder interface {
		DecodeJSON(r io.Reader, target interface{}) error
	}

	Service interface {
		AddPerson(person *Person) (*Person, error)
	}

	handler struct {
		service Service
		decoder Decoder
	}
)

func (h *handler) GetMethod() string {
	return http.MethodPost
}

func (h *handler) GetPattern() string {
	return "/people"
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var person Person

	if err := h.decoder.DecodeJSON(r.Body, &person); err != nil {
		error := make(map[string]string)
		error["error"] = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		js, _ := json.Marshal(error)
		w.Write(js)
		return
	}

	newPerson, err := h.service.AddPerson(&person)
	if err != nil {
		error := make(map[string]string)
		error["error"] = err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		js, _ := json.Marshal(error)
		w.Write(js)
		return
	}

	w.WriteHeader(http.StatusCreated)
	js, _ := json.Marshal(newPerson)
	w.Write(js)
}

func NewHandler(service Service, decoder Decoder) *handler {
	return &handler{
		service: service,
		decoder: decoder,
	}
}
