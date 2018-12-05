package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joaotonial/golangwebserver/config"
	"github.com/joaotonial/golangwebserver/models"
	"github.com/joaotonial/golangwebserver/routes"
)

// Post struct is the posts model

func main() {
	fmt.Println("Starting execution...")
	config.Env()
	models.InitialMigration()
	fmt.Println("Starting Router...")
	fmt.Printf("Running on port %s\n", os.Getenv("PORT"))
	server := &http.Server{
		Addr:         ":7070",
		Handler:      routes.Routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
	defer config.Db.Close()
}
