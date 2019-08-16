package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

//Create a new type of 'deck'
//Assume it as a class which has print function and all objects can use
//its print function
type deck []string

//Receiver : any variable of type deck has access to the print method
//convention to reference to the object with single letter

//func receiver funcName returnType
func newDeck() deck {

	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, value := range cardValues {
		for _, suite := range cardSuites {

			cards = append(cards, value+" of "+suite)
		}
	}
	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	//For loop syntax
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error{
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		//log the error 
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	// Generate a random number
	for i := range d {
		position := r.Intn(len(d) - 1)
		d[i], d[position] = d[position], d[i]
	}

	return d
}
