package api

import (
	"errors"
	"log/slog"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/rabbit"
)

var ErrInvalidPerson = errors.New("person contains invalid fields")

type service struct {
	producer rabbit.Producer
}

func NewService(producer rabbit.Producer) *service {
	return &service{
		producer: producer,
	}
}

func (s *service) AddPerson(person *Person) (*Person, error) {
	person.GenerateID()

	if !person.IsValid() {
		return nil, ErrInvalidPerson
	}

	slog.Debug("Sending person to queue..")

	err := s.producer.Produce(rabbit.PersonQueue.String(), person)
	if err != nil {
		slog.Error("Error producing message in queue", slog.String("error", err.Error()))
		return nil, err
	}

	slog.Info("SUCCESS. Person has been produced")

	return person, nil
}
