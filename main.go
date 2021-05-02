package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tatane616/go-rest-api/article"
	"github.com/tatane616/go-rest-api/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error

	// in-memory DB
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.DBConn = db

	db.AutoMigrate(&article.Article{})
	db.Create(&article.Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"})

	fmt.Println("Connection Opened to Database")
}

func handleRequests() {
	// Init router
	r := mux.NewRouter().StrictSlash(true)

	// Endpoints
	r.HandleFunc("/articles", article.GetArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", article.GetArticle).Methods("GET")
	r.HandleFunc("/articles", article.CreateArticle).Methods("POST")
	r.HandleFunc("/articles/{id}", article.UpdateArticle).Methods("PUT")
	r.HandleFunc("/articles/{id}", article.DeleteArticle).Methods("DELETE")

	// Start server
	fmt.Println("Start server")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	initDatabase()
	handleRequests()
}
