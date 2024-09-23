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

func TestHelloEmptyName(t *testing.T) {
	name := ""
	customMessage := "fake custom"
	msg, err := Hello(name, customMessage)

	assert.Equal(t, msg, name)
	assert.Error(t, err)
}

func TestHelloEmptyCustomMessage(t *testing.T) {
	name := "Bob"
	customMessage := ""
	msg, err := Hello(name, customMessage)

	assert.Contains(t, msg, name)
	assert.Nil(t, err)
}
