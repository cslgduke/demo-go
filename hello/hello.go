package main

import (
	"example.com/greetings"
	"fmt"
	"golang.org/x/example/stringutil"
	"log"
)
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())

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

	fmt.Println(stringutil.Reverse("Hello"))
}
