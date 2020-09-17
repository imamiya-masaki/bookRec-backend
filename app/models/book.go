package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id          int       `json:"id"`
	BookGroupId int       `json:"bookGroupId"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Price       int       `json:"price"`
	ReleaseDate time.Time `json:"releaseDate"`
	URI         string    `json:"uri"`
}

type BookContent struct {
	Id     int    `json:"id"`
	BookId int    `json:"bookId"`
	Page   int    `json:"page"`
	URI    string `json:"uri"`
}

type MyBook struct {
	Id     int
	UserId int
	BookId int
}

type BookResponse struct {
	Quantity int64  `json:"quantity"`
	Books    []Book `json:"books"`
}

type BookRequest struct {
	BookGroupId int    `json:"bookGroupId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	URI         string `json:"uri"`
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

func GetUsersBook(db *gorm.DB, userId int) BookResponse {
	var mybooks []MyBook

	result := db.Where("user_id = ?", userId).Find(&mybooks)
	
	res := BookResponse{result.RowsAffected, []Book{}}

	for _, mybook := range mybooks {
		res.Books = append(res.Books, GetBook(db, mybook.BookId))
	}
	
	return res
}

func (book *Book) Content(db *gorm.DB) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"book_id": book.Id}).Find(&bookContents)
	return bookContents
}

func GetBookContent(db *gorm.DB, id int) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"book_id": id}).Find(&bookContents)
	return bookContents
}

func (req *BookRequest) RegistBook(db *gorm.DB) {
	date := time.Now()
	book := Book{BookGroupId: req.BookGroupId, Title: req.Title, Author: req.Author, Price: req.Price, URI: req.URI, ReleaseDate: date}
	db.Create(&book)
}

func GetPagesByPageCount(db *gorm.DB, count int) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"book_id": 1}).Find(&bookContents)
	getCount := count
	if getCount > len(bookContents) {
		getCount = len(bookContents)
	}
	return bookContents[0:count]
}
func GetPages(db *gorm.DB) []BookContent {
	var bookContents []BookContent
	db.Where(map[string]interface{}{"book_id": 1}).Find(&bookContents)
	return bookContents
}

/*
type Book struct {
	Id          int    `json:"id"`
	BookGroupId int    `json:"bookGroupId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	ReleaseDate int    `json:"releaseDate"`
}

*/
