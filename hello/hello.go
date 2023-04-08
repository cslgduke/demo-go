package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("Dave")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"James", "Michal", "Lucas"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
