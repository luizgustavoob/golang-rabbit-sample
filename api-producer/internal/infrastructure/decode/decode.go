package decode

import (
	"io"

	"github.com/go-chi/render"
)

type Decode struct{}

func (d *Decode) DecodeRequestBody(r io.Reader, target interface{}) error {
	err := render.DecodeJSON(r, target)
	return err
}

func New() *Decode {
	return &Decode{}
}
