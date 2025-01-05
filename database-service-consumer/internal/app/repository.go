package app

import "log/slog"

type (
	DB interface {
		Exec(query string, args ...any) error
	}

	repository struct {
		db DB
	}
)

func NewRepository(db DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddPerson(person *Person) error {
	err := r.db.Exec(`INSERT INTO person(id, nome, idade, email, telefone) VALUES ($1, $2, $3, $4, $5)`,
		&person.ID,
		&person.Name,
		&person.Age,
		&person.Email,
		&person.Phone,
	)
	if err != nil {
		slog.Error("Error inserting person", slog.String("error", err.Error()))
		return err
	}

	return nil
}
