package marshal_test

import (
	"testing"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/marshal"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	something := struct {
		Key string `json:"chave"`
	}{Key: "content"}

	marshaller := marshal.New()
	assert.NotNil(t, marshaller)

	response, err := marshaller.MarshalValue(&something)
	assert.Nil(t, err)
	assert.NotNil(t, response)
}
