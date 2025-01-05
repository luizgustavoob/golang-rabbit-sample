package app

import "log/slog"

type (
	Repository interface {
		AddPerson(person *Person) error
	}

	service struct {
		repository Repository
	}
)

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) AddPerson(person *Person) error {
	slog.Debug("Adding person...")

	if err := s.repository.AddPerson(person); err != nil {
		return err
	}

	slog.Info("SUCCESS. Person has been added.")

	return nil
}
