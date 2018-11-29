package api

import (
	"net/http"
)

var tonis int

func get(w http.ResponseWriter, r *http.Request) {
	tonis = 4

}
