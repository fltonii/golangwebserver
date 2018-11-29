package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaotonial/golangwebserver/controllers"
)

func Routes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", controllers.GetPost).Methods("GET")
	router.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/api/posts", controllers.EditPost).Methods("PUT")
	router.HandleFunc("/api/posts/{id}", controllers.DeletePost).Methods("DELETE")
	http.ListenAndServe(":7070", router)
}
