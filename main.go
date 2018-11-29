package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("\/posts", api.tonis)
}
