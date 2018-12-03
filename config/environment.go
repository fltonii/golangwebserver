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
	os.Setenv("DBNAME", "postgres://utelpmgiocqift:131cbb93e197e1d5f673bd703c321dba6b4aa3e95dcf997894b7c400bdadb810@ec2-184-72-239-186.compute-1.amazonaws.com:5432/d2njf86qv85sjm")
	os.Setenv("PORT", "7070")
}

// OpenDb open database connection
func OpenDb() {
	// conf := ("host=" + os.Getenv("HOST") + " port=" + os.Getenv("DBPORT") + " user=" + os.Getenv("USER") + " dbname=" + os.Getenv("DBNAME") + " sslmode=disable password=" + os.Getenv("PASSWORD"))
	Db, err = gorm.Open(
		"postgres",
		"postgres://utelpmgiocqift:131cbb93e197e1d5f673bd703c321dba6b4aa3e95dcf997894b7c400bdadb810@ec2-184-72-239-186.compute-1.amazonaws.com:5432/d2njf86qv85sjm")

	if err != nil {
		panic("database failed to connect")
	}
}
