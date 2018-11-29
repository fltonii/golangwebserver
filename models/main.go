package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Post struct {
	gorm.Model
	Title   string
	Content string
	ID      int
}

func models() {
	db, err := gorm.Open("sqlite3", "/tmp/goblog.db")
	if err {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&Post{})
}
