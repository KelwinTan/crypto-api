package main

import (
	"fmt"
	"crypto.com/greetings"
)

// import "rsc.io/quote"

func main(){
	// fmt.Println("Hello World!")
	// fmt.Println(quote.Opt())
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}