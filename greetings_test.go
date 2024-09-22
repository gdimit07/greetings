package greetings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloName(t *testing.T) {
	name := "Bob"
	customMessage := "fake custom"
	msg, err := Hello(name, customMessage)

	assert.Contains(t, msg, name)
	assert.Contains(t, msg, customMessage)
	assert.Nil(t, err)
}
