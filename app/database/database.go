package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// Init initializes database
func Init() {
	dsn := "docker:docker@tcp(mysql:3306)/book_db?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:sql_daishin_1131@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
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
