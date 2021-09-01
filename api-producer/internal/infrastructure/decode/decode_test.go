package decode_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/decode"
	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	something := struct {
		Key string `json:"chave"`
	}{Key: "content"}

	decoder := decode.New()
	assert.NotNil(t, decoder)

	js, _ := json.Marshal(&something)
	target := make(map[string]string)

	err := decoder.DecodeJSON(bytes.NewReader(js), &target)
	assert.Nil(t, err)
	assert.Equal(t, "content", target["chave"])
}
