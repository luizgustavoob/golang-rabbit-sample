package api

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"nome,omitempty"`
	Age   int    `json:"idade,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"telefone,omitempty"`
}

func (p *Person) GenerateID() {
	const (
		size    = 8
		simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
	)

	newID := func() string {
		id := make([]byte, size)
		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}
		return string(id)
	}

	p.ID = newID()
}

func (p *Person) IsValid() bool {
	return (p.ID != "" && p.Name != "" &&
		p.Age > 0 && p.Email != "" && p.Phone != "")
}
