package main

import (
	"fmt"
	"log"
	"os"
	"psideris/message"
	"psideris/personService"
)

func init() {
	log.Println("Runing init")
	log.SetOutput(os.Stdout)
}

func main() {
	var msg = message.GetMessage("from main")
	fmt.Println(msg)

	p1 := personService.Create("John", 34)
	personService.Update(&p1, "Doe", 34)

	p2 := personService.Create("Marry", 30)
	p2 = personService.GetUpdated(p2, "Poppins", 40)

	personService.Print(p1)
	personService.Print(p2)
}
