package main // define this package

import (
	"fmt"
	"log"

	"goLearn/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Nick", "Foo/that-foo Adrian", "Captain Anders", "ChrispyNips"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
