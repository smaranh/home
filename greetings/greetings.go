package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {

	// If no name is provided then return an error
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randGreeting(), name)
	return message, nil
}

// Input a slice of names and the Hellos func will return the mapped
// messages to each name
func Hellos(names []string) (map[string]string, error) {
	// instantiate a map object that will map the
	// names and messages
	messages := make(map[string]string)

	//loop through the names and create update messages map
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	// rand needs to be seeded like this to get random values everytime
	rand.Seed(time.Now().UnixNano())
}

func randGreeting() string {
	//A slice of messages
	messages := []string{
		"Hi, %v. How is life?",
		"Good to see you %v",
		"Welcome %v. Time to shine.",
	}

	// return randomly selected message from the slice
	return messages[rand.Intn(len(messages))]
}
