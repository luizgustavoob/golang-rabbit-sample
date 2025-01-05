package api

import (
	"errors"
	"log/slog"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
)

var ErrInvalidPerson = errors.New("person contains invalid fields")

type service struct {
	publisher rabbit.Publisher
}

func NewService(publisher rabbit.Publisher) *service {
	return &service{
		publisher: publisher,
	}
}

func (s *service) AddPerson(person *Person) (*Person, error) {
	person.GenerateID()

	if !person.IsValid() {
		return nil, ErrInvalidPerson
	}

	slog.Debug("Sending person to queue..")

	err := s.publisher.Publish(rabbit.PersonQueue.String(), person)
	if err != nil {
		slog.Error("Error publishing message in queue", slog.String("error", err.Error()))
		return nil, err
	}

	slog.Info("SUCCESS. Person has been published")

	return person, nil
}
