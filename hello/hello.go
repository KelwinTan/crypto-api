package main

import (
	"fmt"
	"crypto.com/greetings"
	"log"
)

// import "rsc.io/quote"

func main(){
	// fmt.Println("Hello World!")
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	//request a greeting message
	// message, err := greetings.Hello("Gladys")
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(quote.Opt())
	// message := greetings.Hello("Gladys")
	// for _, message := range messages{
	// 	// message := messages
	// 	fmt.Println(message)
	// }
	fmt.Println(messages)
}