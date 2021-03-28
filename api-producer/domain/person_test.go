package domain_test

import (
	"testing"

	"github.com/golang-rabbit-sample/api-producer/domain"
	"github.com/stretchr/testify/assert"
)

func TestDomainPerson_GenerateID(t *testing.T) {

	t.Run("should generate id and validate struct", func(t *testing.T) {
		p := &domain.Person{
			Name:  "Luiz Gustavo",
			Age:   25,
			Email: "email@gmail.com",
			Phone: "11111111",
		}
		p.GenerateID()
		assert.NotEqual(t, "", p.ID)
		assert.True(t, p.IsValid())
	})

	t.Run("not should validate struct", func(t *testing.T) {
		p := &domain.Person{
			Name:  "Luiz Gustavo",
			Phone: "11111111",
		}

		assert.Empty(t, p.ID)
		assert.False(t, p.IsValid())
	})

	t.Run("should generate ID", func(t *testing.T) {
		p := &domain.Person{
			Name:  "Luiz Gustavo",
			Phone: "11111111",
		}
		p.GenerateID()
		assert.NotEmpty(t, p.ID)
	})
}
