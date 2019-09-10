package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Article ... struct definition
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("Starting rest apis ... ")
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {

	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit : returnAllArticles")

	articles := []Article{
		Article{Title: "ZNMD", Desc: "Zindagi Na Milegi Dobara", Content: "Best Movie Ever"},
		Article{Title: "CCOB", Desc: "Curious Case of Benjamin Button", Content: "Best Movie Ever 2"},
	}

	json.NewEncoder(w).Encode(articles)
}
