package api_test

import (
	"testing"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestEntities_Person(t *testing.T) {

	t.Run("should generate ID", func(t *testing.T) {
		person := &api.Person{
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		person.GenerateID()

		assert.NotEmpty(t, person.ID)
	})

	t.Run("should validate person", func(t *testing.T) {
		person := &api.Person{
			ID:    "1",
			Name:  "name",
			Age:   25,
			Email: "email@email.com",
			Phone: "12345678",
		}

		assert.True(t, person.IsValid())
	})

	t.Run("should return not valid person", func(t *testing.T) {
		person := &api.Person{}
		assert.False(t, person.IsValid())
	})
}
