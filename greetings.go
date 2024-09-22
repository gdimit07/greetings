package greetings

import "fmt"

func Hello(name string, message string) string {
	return fmt.Sprintf("Hello %s, welcome! %s", name, message)
}
