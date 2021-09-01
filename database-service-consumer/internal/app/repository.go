package app

type (
	DB interface {
		Exec(query string, args ...interface{}) error
	}

	repository struct {
		logger Logger
		db     DB
	}
)

func (r *repository) AddPerson(person *Person) error {
	err := r.db.Exec(`INSERT INTO person(id, nome, idade, email, telefone) VALUES ($1, $2, $3, $4, $5)`,
		&person.ID,
		&person.Nome,
		&person.Idade,
		&person.Email,
		&person.Telefone)

	if err != nil {
		r.logger.Printf("Failed to insert person: %s\n", err.Error())
		return err
	}

	return nil
}

func NewRepository(logger Logger, db DB) *repository {
	return &repository{
		logger: logger,
		db:     db,
	}
}
