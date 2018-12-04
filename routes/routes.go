package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaotonial/golangwebserver/controllers"
)

// Routes defines the routes, helper functions, and methods for the server
func Routes() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controllers.EditPost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("DELETE")
	return
}
