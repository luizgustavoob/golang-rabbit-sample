package api

import (
	"encoding/json"
	"math/rand"
	"time"
)

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
		seed := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := make([]byte, size)
		for i := range id {
			id[i] = simbols[seed.Intn(len(simbols))]
		}
		return string(id)
	}

	p.ID = newID()
}

func (p *Person) IsValid() bool {
	return (p.ID != "" && p.Name != "" &&
		p.Age > 0 && p.Email != "" && p.Phone != "")
}

func (p *Person) Serialize() ([]byte, error) {
	return json.Marshal(p)
}
