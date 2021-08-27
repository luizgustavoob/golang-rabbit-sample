package api

import (
	"errors"
)

const (
	queue = "person-queue"
)

type (
	Marshaller interface {
		MarshalValue(obj interface{}) ([]byte, error)
	}

	Logger interface {
		Printf(format string, values ...interface{})
		Println(values ...interface{})
	}

	Publisher interface {
		Publish(queueName string, message string) error
	}

	service struct {
		publisher  Publisher
		marshaller Marshaller
		logger     Logger
	}
)

func (s *service) AddPerson(person *Person) (*Person, error) {
	person.GenerateID()

	if !person.IsValid() {
		return nil, errors.New("person contains invalid fields")
	}

	personBytes, err := s.marshaller.MarshalValue(&person)
	if err != nil {
		s.logger.Printf("Failed in marshaling person: %s", err.Error())
		return nil, err
	}

	s.logger.Println("Sending person to queue..")
	err = s.publisher.Publish(queue, string(personBytes))
	if err != nil {
		s.logger.Printf("Failed to publish message in queue: %s", err.Error())
		return nil, err
	}

	s.logger.Println("success")

	return person, nil
}

func NewService(publisher Publisher, marshaller Marshaller, logger Logger) *service {
	return &service{
		publisher:  publisher,
		marshaller: marshaller,
		logger:     logger,
	}
}
