package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {

	// Set the prefix and the flag to not display time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Joe", "Sam", "Neville"}

	messages, err := greetings.Hellos(names)

	// if an error occured, print it to console and exit
	if err != nil {
		log.Fatal(err)
	}

	// If not error was reported then print the message
	fmt.Println(messages)
}
