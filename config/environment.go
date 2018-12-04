
package config

import (
	"os"

	"github.com/jinzhu/gorm"
)

// Db is the global variable that must be used to access the db
var Db *gorm.DB
var err error

// Env defines environmental variables
func Env() {
	os.Setenv("HOST", "localhost")
	os.Setenv("DBPORT", "5432")
	os.Setenv("USER", "tonis")
	os.Setenv("PASSWORD", "123456")
	os.Setenv("DBNAME", "golangblog")
	os.Setenv("PORT", "7070")
}

// OpenDb open database connection
func OpenDb(){
	conf := ("host=" + os.Getenv("HOST") + " port=" + os.Getenv("DBPORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=disable password=" + os.Getenv("PASSWORD"))
	Db, err = gorm.Open(
		"postgres",
		conf)

	if err != nil {
		panic("database failed to connect")
	}
}