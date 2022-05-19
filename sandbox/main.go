package main

import (
	"fmt"
	"log"
	"os"
	"psideris/sandbox/message"
	"psideris/sandbox/personService"
)

func init() {
	log.Println("Running init")
	log.SetOutput(os.Stdout)
}

func main() {
	var msg = message.GetMessage("from main")
	fmt.Println(msg)

	p1 := personService.NewPerson("John", "Doe", 34)
	p1.Update("Johnny", "Doe", 30)

	p2 := personService.NewPerson("Marry", "Poppins", 20)
	p2.UpdateAge(29)

	p1.Print()
	p2.Print()

	// personService.PrintBooking()
}
