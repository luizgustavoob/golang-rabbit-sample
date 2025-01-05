package api_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang-rabbit-sample/api-producer/internal/api"
)

func TestPerson(t *testing.T) {
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

	t.Run("should return false when validate empty person", func(t *testing.T) {
		person := &api.Person{}
		assert.False(t, person.IsValid())
	})

	t.Run("should serialize person", func(t *testing.T) {
		p := &api.Person{ID: "1", Name: "Name", Age: 10, Email: "email@email.com", Phone: "12345"}
		pJS, err := p.Serialize()

		assert.NoError(t, err)
		assert.Equal(t, "{\"id\":\"1\",\"nome\":\"Name\",\"idade\":10,\"email\":\"email@email.com\",\"telefone\":\"12345\"}", string(pJS))
	})
}
