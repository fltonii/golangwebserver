package controllers

import (
	"encoding/json"
	"net/http"
)

var Tonis int

func GetPosts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
func GetPost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
func CreatePost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
func EditPost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
