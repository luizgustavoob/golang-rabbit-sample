package decode

import (
	"encoding/json"
	"io"
)

type Decode struct{}

func (d *Decode) DecodeJSON(r io.Reader, target interface{}) error {
	return json.NewDecoder(r).Decode(target)
}

func New() *Decode {
	return &Decode{}
}
