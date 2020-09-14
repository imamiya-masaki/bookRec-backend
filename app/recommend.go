package main

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})

	r.POST("/recommend", sendRecommend)
	r.GET("/my_recommend", getMyRecommend)
	r.GET("/my_recommended", getMyRecommended)

	r.Run()
}

func getMyRecommend(c *gin.Context) {
	param := c.Query("sender_id")
	id, _ := strconv.Atoi(param)
	recommends := models.GetMyRecommend(database.GetDB(), id)

	c.JSON(200, recommends)
}

func getMyRecommended(c *gin.Context) {
	param := c.Query("reciever_id")
	id, _ := strconv.Atoi(param)
	recommends := models.GetMyRecommended(database.GetDB(), id)

	c.JSON(200, recommends)
}

func sendRecommend(c *gin.Context) {
	recommend := &models.Recommend{}
	err := c.BindJSON(recommend)
	// err := c.Bind(recommend)
	if err != nil {
		println(err)
	}

	recommend.SendReccomend(database.GetDB())

	c.JSON(200, recommend)
}
