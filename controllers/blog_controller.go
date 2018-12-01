package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/joaotonial/golangwebserver/config"

	"github.com/jinzhu/gorm"
	"github.com/joaotonial/golangwebserver/models"
)

var Posts []*models.Post
var db *gorm.DB
var err error

func GetPosts(w http.ResponseWriter, r *http.Request) {
	config.Db.Find(&Posts)
	json.NewEncoder(w).Encode(Posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost *models.Post
	_ = json.NewDecoder(r.Body).Decode(&newPost)
	config.Db.Create(&newPost)

	json.NewEncoder(w).Encode(newPost)
}
func EditPost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
