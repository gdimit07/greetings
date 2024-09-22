package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string, customMessage string) (string, error) {
	if name == "" {
		return name, errors.New("received empty name")
	}
	welcome_message := randomWelcomeMessage(name)
	full_message := fmt.Sprintf("%s %s", welcome_message, customMessage)
	return full_message, nil
}

func randomWelcomeMessage(name string) string {
	messages := []string{
		"Hello %s, welcome!",
		"Great to have you here, %s!",
		"Nice to see you %s!",
	}

	return fmt.Sprintf(messages[rand.Intn(len(messages))], name)
}
