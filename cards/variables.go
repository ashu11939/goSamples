package main

import (
	"fmt"
)

// Variables can be initialized but cannot be assigned outside a function
var sample string

func variable() {
	//var card string = "Ace of Spades" // same as
	card := newCard() // only when defining a new variable
	fmt.Println(card)
}

//Return type to be declared beside the functionName
func newCard() string {
	return "Five of Diamonds"
}
