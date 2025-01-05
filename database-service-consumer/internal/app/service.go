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
	return s.repository.AddPerson(person)
}

func AddPersonFn(service *service) func(p *Person) {
	return func(p *Person) {
		slog.Debug("Adding person...")

		if err := service.AddPerson(p); err != nil {
			slog.Error("Error adding person", slog.String("error", err.Error()))
			return
		}

		slog.Info("SUCCESS. Person has been added.")
	}
}
