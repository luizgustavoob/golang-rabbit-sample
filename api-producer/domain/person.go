package domain

import (
	"math/rand"
	"time"
)

type PersonService interface {
	AddPerson(person *Person) (*Person, error)
}

type PersonClient interface {
	AddNewPerson(person *Person) (*Person, error)
}

const (
	size    = 8
	simbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"nome,omitempty"`
	Age   int    `json:"idade,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"telefone,omitempty"`
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

func (self *Person) IsValid() bool {
	return (self.ID != "" && self.Name != "" &&
		self.Age > 0 && self.Email != "" && self.Phone != "")
}
