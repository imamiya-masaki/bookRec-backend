package models

import "gorm.io/gorm"

type Book struct {
	Id          int    `json:"id"`
	BookGroupId int    `json:"bookGroupId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	ReleaseDate int    `json:"releaseDate"`
}

type BookContent struct {
	Id     int    `json:"id"`
	BookId int    `json:"bookId"`
	Page   int    `json:"page"`
	URI    string `json:"uri"`
}

func (book *Book) CreateBook(db *gorm.DB) {
	result := db.Create(&book)
	if err := result.Error; err != nil {
		println(err)
	}
}

func GetBook(db *gorm.DB, id int) Book {
	var book Book
	db.First(&book, id)
	return book
}

func GetAllBook(db *gorm.DB) []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (book *Book) Content(db *gorm.DB) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"bookId": book.Id}).Find(&bookContents)
	return bookContents
}

func GetBookContent(db *gorm.DB, id int) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"bookId": id}).Find(&bookContents)
	return bookContents
}
