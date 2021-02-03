package models

import (
	"math/rand"
	"time"
)

const (
	size    = 5
	simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

type Person struct {
	ID       string `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Idade    int    `json:"idade,omitempty"`
	Email    string `json:"email,omitempty"`
	Telefone string `json:"telefone,omitempty"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (self *Person) GenerateID() {
	newID := func() string {
		id := make([]byte, size, size)
		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}
		return string(id)
	}

	self.ID = newID()
}
