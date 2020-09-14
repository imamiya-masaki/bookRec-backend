package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes database
func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DATABASE := os.Getenv("DATABASE")
	USERNAME := os.Getenv("USERNAME")
	USERPASS := os.Getenv("USERPASS")

	// dsn := "docker:docker@tcp(mysql:3306)/book_db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := USERNAME + ":" + USERPASS + "@tcp(mysql:3306)/" + DATABASE + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println(err)
	} else {
		println(("db connection success !!"))
	}
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}
