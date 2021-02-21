package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// set logger properties
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// loop()
	msg, err := greetings.Hello("Anton")
	// exit if exception occured
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}

func loop() {
	for i := 0; i < 10; i++ {
		msg, err := greetings.Hello("Anton")
		// exit if exception occured
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(msg)
	}
}
