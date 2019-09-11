package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Article ... struct definition
type Article struct {
	ID      string `json:"id"`
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

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/addArticle", createNewArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit : returnAllArticles")
	articles := returnArticles()
	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "error in conversion ")
	}
	articles := returnArticles()
	//fmt.Fprintf(w, "key : "+key)
	json.NewEncoder(w).Encode(articles[key])

}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	articles := returnArticles()
	articles = append(articles, article)
	json.NewEncoder(w).Encode(articles)
}

func returnArticles() []Article {
	articles := []Article{
		Article{ID: "1", Title: "ZNMD", Desc: "Zindagi Na Milegi Dobara", Content: "Best Movie Ever"},
		Article{ID: "2", Title: "CCOB", Desc: "Curious Case of Benjamin Button", Content: "Best Movie Ever 2"},
	}
	return articles
}
