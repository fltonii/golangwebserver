package main

import (
	"net/http"

	"github.com/joaotonial/golangwebserver/controllers"
)

func main() {
	http.HandleFunc("/posts", controllers.Get)
}
