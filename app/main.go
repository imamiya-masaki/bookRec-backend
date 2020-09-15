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

	r.GET("/book/:id", apiGetBook)
	r.GET("/book", apiGetAllBook)
	r.GET("/book/:id/Content", apiGetContent)
	r.GET("/users/:id", apiGetUserdata)

	r.POST("/recommend", sendRecommend)
	r.GET("/my_recommend", getMyRecommend)
	r.GET("/my_recommended", getMyRecommended)

	r.Run()
}
