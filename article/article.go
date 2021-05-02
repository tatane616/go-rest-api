package article

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tatane616/go-rest-api/database"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID      string `json:id`
	Title   string `json:title`
	Desc    string `json:desc`
	Content string `json:content`
}

// Get all articles
func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.DBConn
	var articles []Article
	db.Find(&articles)

	json.NewEncoder(w).Encode(articles)
}

// Get single article
func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.DBConn
	params := mux.Vars(r)
	var article Article
	db.Find(&article, params["id"])

	json.NewEncoder(w).Encode(article)
}

// Create new article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.DBConn
	var article Article
	_ = json.NewDecoder(r.Body).Decode(&article)
	db.Create(&article)

	json.NewEncoder(w).Encode(article)
}

// Update article
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	// TODO
}

// Delete article
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.DBConn
	params := mux.Vars(r)
	var article Article
	db.First(&article, params["id"])
	if article.Title == "" {
		return
	}
	db.Delete(&article)

	var articles []Article
	db.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}
