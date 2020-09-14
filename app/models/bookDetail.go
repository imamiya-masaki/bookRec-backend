package models

import "gorm.io/gorm"

type Book struct {
	id          int
	bookGroupId int
	title       string
	author      string
	price       int
	releaseDate int
}

type BookContent struct {
	id     int
	bookId int
	page   int
	URI    string
}

func (book *Book) createBook(db *gorm.DB) {
	result := db.Create(&book)
	if err := result.Error; err != nil {
		println(err)
	}
}

func getBook(db *gorm.DB, id int) Book {
	var book Book
	db.First(&book, id)
	return book
}

func getAllBook(db *gorm.DB) Book {
	var book Book
	db.Find(&book)
	return book
}

func (book *Book) content(db *gorm.DB) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"bookId": book.id}).Find(&bookContents)
	return bookContents
}

func getBookContent(db *gorm.DB, id int) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"bookId": id}).Find(&bookContents)
	return bookContents
}
