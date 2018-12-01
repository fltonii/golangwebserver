package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/joaotonial/golangwebserver/config"
)

var db *gorm.DB
var err error

type Post struct {
	gorm.Model
	Title   string
	Content string
}

func InitialMigration() {
	fmt.Println("Opening connection to database...")
	config.OpenDb()

	config.Db.AutoMigrate(&Post{})
}
