package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/joaotonial/golangwebserver/config"

	"github.com/joaotonial/golangwebserver/models"
)

var Posts []*models.Post

// GetPosts gets a list of posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	config.Db.Find(&Posts)
	json.NewEncoder(w).Encode(Posts)
}

// GetPost gets a single post
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post []*models.Post
	config.Db.Where("id = ?", params["id"]).Find(&post)

	json.NewEncoder(w).Encode(post)
}

// CreatePost creates a new post with given parameters
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var newPost *models.Post
	_ = json.NewDecoder(r.Body).Decode(&newPost)
	config.Db.Create(&newPost)

	json.NewEncoder(w).Encode(newPost)
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post []*models.Post
	var updated *models.Post
	_ = json.NewDecoder(r.Body).Decode(&updated)

	config.Db.Where("id = ?", params["id"]).First(&post)

	post[0].Content = updated.Content
	config.Db.Save(&post[0])
	json.NewEncoder(w).Encode(post[0])
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var toBeDeleted []*models.Post
	config.Db.Where("id = ?", params["id"]).First(&toBeDeleted)

	config.Db.Delete(&toBeDeleted)

	config.Db.Find(&Posts)
	json.NewEncoder(w).Encode(Posts)
}
