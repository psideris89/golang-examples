package main

import (
	"fmt"
	"log"

	// "rsc.io/quote"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("From greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("John")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
	// fmt.Println("Hello from the other side")
	// fmt.Println(quote.Glass())
}
