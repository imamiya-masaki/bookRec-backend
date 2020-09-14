package models

import "gorm.io/gorm"

type Post struct {
	Id      int
	Title   string
	Content string
}

func (post *Post) CreatePost(db *gorm.DB) {
	result := db.Create(&post)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetAllPosts(db *gorm.DB) []Post {
	var posts []Post
	result := db.Find(&posts)
	println(result.RowsAffected)

	return posts
}
