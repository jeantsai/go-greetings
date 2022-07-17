package greetings

import "fmt"

func hello(name string) string {
	msg := fmt.Sprintf("Hello %v", name)
	return msg
}
