package main

import (
	"fmt"
	"log"
	"os"
	"psideris/message"
)

func init() {
	log.Println("Runing init")
	log.SetOutput(os.Stdout)
}

func main() {
	var msg = message.GetMessage("Pablo")
	fmt.Println(msg)
}
