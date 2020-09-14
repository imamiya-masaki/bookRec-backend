package main

import (
	"app/database"
	"app/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})
	r.POST("/posts", createPost)
	r.GET("/posts", getAllPost)

	r.Run()
}

func getAllPost(c *gin.Context) {
	println("get all posts")
	posts := models.GetAllPosts(database.GetDB())

	for _, v := range posts {
		println(v.Title)
	}

	c.JSON(200, posts)
}

func createPost(c *gin.Context) {
	post := &models.Post{}
	err := c.Bind(post)
	if err != nil {
		println(err)
	}

	post.CreatePost(database.GetDB())
	// err = post.CreatePost(database.GetDB())
	// if err != nil {
	// 	println(err)
	// }

	c.JSON(200, post)
}
