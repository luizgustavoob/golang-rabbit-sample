package app

type (
	Repository interface {
		AddPerson(person *Person) error
	}

	service struct {
		repository Repository
	}
)

func (s *service) AddPerson(person *Person) error {
	return s.repository.AddPerson(person)
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}
