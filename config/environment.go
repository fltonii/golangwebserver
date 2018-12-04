package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db is the global variable that must be used to access the db
var Db *gorm.DB
var err error

// Env defines environmental variables
func Env() {
	os.Setenv("HOST", "localhost")
	os.Setenv("DBPORT", "5432")
	os.Setenv("USER", "postgres")
	os.Setenv("PASSWORD", "sbsistemas5846")
	os.Setenv("DBNAME", "trainingblog")
	os.Setenv("PORT", "7070")
}

// OpenDb open database connection
func OpenDb() {
	conf := ("host=" + os.Getenv("HOST") + " port=" + os.Getenv("DBPORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=enabled password=" + os.Getenv("PASSWORD"))
	fmt.Println(conf)
	Db, err = gorm.Open(
		"postgres",
		conf)

	if err != nil {
		panic("database failed to connect")
	}
}
