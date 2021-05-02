package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	ID      string `json:id`
	Title   string `json:title`
	Desc    string `json:desc`
	Content string `json:content`
}

var articles []Article

// Get all articles
func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Get single article
func getArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Article{})
}

// Create new article
func createArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	// Mock ID
	article.ID = strconv.Itoa(rand.Intn(1000000000))
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

// Update article
func updateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			var article Article
			_ = json.NewDecoder(r.Body).Decode(&article)
			article.ID = params["id"]
			articles = append(articles, article)
			return
		}
	}
}

// Delete article
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(articles)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home!")
}

func handleRequests() {
	// Init router
	r := mux.NewRouter().StrictSlash(true)

	// TODO: add DB
	articles = append(articles, Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"})
	articles = append(articles, Article{ID: "2", Title: "Hello", Desc: "Article Description", Content: "Article Content"})

	// Endpoints
	r.HandleFunc("/", home)
	r.HandleFunc("/articles", getArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	r.HandleFunc("/articles", createArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

	// Start server
	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	fmt.Println("Hello from main")
	handleRequests()
}
