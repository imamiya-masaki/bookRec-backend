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

	r.POST("/recommends", sendRecommend)
	r.GET("/recommends", getRecommends)

	r.Run()
}

func getRecommends(c *gin.Context) {
	recommends := models.GetMyRecommend(database.GetDB())

	c.JSON(200, recommends)
}

func sendRecommend(c *gin.Context) {
	recommend := &models.Recommend{}
	err := c.Bind(recommend)
	if err != nil {
		println(err)
	}

	recommend.SendReccomend(database.GetDB())

	c.JSON(200, recommend)
}
