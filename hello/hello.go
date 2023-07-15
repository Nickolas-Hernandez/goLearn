package main // define this package

import (
	"fmt"
	"log"

	"goLearn/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Nick")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
