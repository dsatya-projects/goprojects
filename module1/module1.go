package module1

import (
	"errors"
	"fmt"
)

/*
// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
*/
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat())
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func BestMascot() string {
	return "All the best"
}
