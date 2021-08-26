package app

import (
	"bytes"
	"io"
)

type (
	LoggerHandler interface {
		Printf(format string, values ...interface{})
	}

	Decoder interface {
		DecodeJSON(r io.Reader, target interface{}) error
	}

	Service interface {
		AddPerson(person *Person) error
	}

	handler struct {
		logger  LoggerHandler
		decoder Decoder
		service Service
	}
)

func (h *handler) HandleMessage(message []byte) error {
	var person Person
	err := h.decoder.DecodeJSON(bytes.NewReader(message), &person)
	if err != nil {
		h.logger.Printf("Error reading message: %s", err.Error())
		return err
	}
	return h.service.AddPerson(&person)
}

func NewHandler(logger LoggerHandler, decoder Decoder, service Service) *handler {
	return &handler{
		logger:  logger,
		decoder: decoder,
		service: service,
	}
}
