package main

import (
	"net/http"
	"fmt"
	"time"
)

func main() {

	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	//Used to communicate between routines
	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel) 
		//New go child routine engine executes the code
		//Create a new thread go routine
		//Use go keyword only infront of a function calls
	}
	
	//<- receiver
	// fmt.Println(<- channel) // Blocking line of code for main routine
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)
	// fmt.Println(<- channel)

	for {

		//Never access the variables from child routines of main routines instead pass the values 
		//inside functions
		go func() { //function literal
			time.Sleep(time.Second)
			checkLink(<- channel, channel)
		}()
	}

}

/**
Concurrency is not equal to parallelism
We have multiple threads executing code. If one thread blocks
another one is picked up and worked on
in case of one core CPU or single CPU

Parallelism
Multiple threads executed at the same time. Requires multiple CPU's
*/


func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link
		fmt.Println(link + "might be down!!")
		return
	}
	//Two way binding types messaging feature
	c <- link// Send msg 
	fmt.Println(link + " is up !!")
}