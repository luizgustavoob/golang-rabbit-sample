package marshal

import "encoding/json"

type Marshal struct{}

func (m *Marshal) SerializeJSON(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

func New() *Marshal {
	return &Marshal{}
}
