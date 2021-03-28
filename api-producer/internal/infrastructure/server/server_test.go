package server_test

import (
	"testing"
	"time"

	"github.com/golang-rabbit-sample/api-producer/internal/infrastructure/server"
	"github.com/stretchr/testify/assert"
)

func TestServer_ListenAndServe(t *testing.T) {
	server := server.New("9000", nil)
	server.ListenAndServe()

	stopChan := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		stopChan <- true
	}()

	var result bool
	result = <-stopChan
	server.Shutdown()

	assert.True(t, result)
}
