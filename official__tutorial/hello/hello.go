package main 

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){

	// set logger properties
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	msg, err := greetings.Hello("Anton")
	// exit if exception occured
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}