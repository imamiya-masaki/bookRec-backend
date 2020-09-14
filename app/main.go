package main

import (
	"app/database"
	"app/models"
	"encoding/json"
	"strconv"

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
	r.GET("/book/:id", apiGetBook)
	r.GET("/book", apiGetAllBook)
	r.GET("/book/:id/Content", apiGetContent)
	r.GET("/users/:id", func(c *gin.Context) {
		getUserData(c, c.Param("id"))
	})

	r.POST("/recommend", sendRecommend)
	r.GET("/my_recommend", getMyRecommend)
	r.GET("/my_recommended", getMyRecommended)

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
func getUserData(c *gin.Context, id string) {
	println("getting userdata")
	var i int
	i, _ = strconv.Atoi(id)
	data := models.GetUserDataById(database.GetDB(), i)
	enc, _ := json.Marshal(data)
	c.JSON(200, string(enc))
}
