package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
}

var Articles []Article

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey there! welcome to go!")
	fmt.Println("endpoint hit: Homepage")
}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func returnArticle(w http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	id := val["id"]
	fmt.Println("Endpoint Hit: return Article")
	//fmt.Fprintf(w, "ID : "+id)
	for _, item := range Articles {
		if item.Id == id {
			json.NewEncoder(w).Encode(item)
		}
	}
}
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	fmt.Println("Endpoint Hit: create Article")
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Println("Endpoint Hit: delete Article")
	for i, item := range Articles {
		if item.Id == id {
			Articles = append(Articles[:i], Articles[i+1:]...)
		}
	}
}
func main() {
	Articles = []Article{
		Article{"1", "west winds", "west Indian drama", "the great gambler"},
		Article{"2", "ode to the", "australian guy", "hey mate"},
	}
	handleRequests()
}
