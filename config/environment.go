package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db is the global variable that must be used to access the db
var Db *gorm.DB
var err error

// Env defines environmental variables
func Env() {
	os.Setenv("HOST", "127.0.0.1")
	os.Setenv("DBPORT", "5432")
	os.Setenv("USER", "postgres")
	os.Setenv("PASSWORD", "sbsistemas5846")
	os.Setenv("DBNAME", "trainingblog")
	os.Setenv("PORT", "7070")
	os.Setenv("SSLMODE", "disable")
}

// OpenDb open database connection
func OpenDb() {
	conf := ("host=" + os.Getenv("HOST") + " port=" + os.Getenv("DBPORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=" + os.Getenv("SSLMODE") + " password=" + os.Getenv("PASSWORD"))
	Db, err = gorm.Open(
		"postgres",
		conf)

	if err != nil {
		panic(err)
	}
}
