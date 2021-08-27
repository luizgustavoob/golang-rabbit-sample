package marshal

import "encoding/json"

type Marshal struct{}

func (m *Marshal) MarshalValue(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func New() *Marshal {
	return &Marshal{}
}
