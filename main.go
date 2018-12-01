package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), (routes.Routes())))
	defer config.Db.Close()
}
