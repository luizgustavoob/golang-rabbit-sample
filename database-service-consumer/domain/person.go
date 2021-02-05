package domain

type Person struct {
	ID       string `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Idade    int    `json:"idade,omitempty"`
	Email    string `json:"email,omitempty"`
	Telefone string `json:"telefone,omitempty"`
}

type PersonService interface {
	AddPerson(person *Person) error
}
