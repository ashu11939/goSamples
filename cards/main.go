package main

// import "fmt"
//Not an object oriented language

func main() {
	cards := newDeck()
	cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("myCards")

	// newCards := newDeckFromFile("myCards")
	// newCards.print()
	cards.shuffle()
	cards.print()

}

