package app

type Person struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"nome,omitempty"`
	Age   int    `json:"idade,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"telefone,omitempty"`
}
