package controllers

import (
	"encoding/json"
	"net/http"
)

var Tonis int

func Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("DIGDIN")
}
