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

func getUserData(c *gin.Context, id string) {
	println("getting userdata")
	var i int
	i, _ = strconv.Atoi(id)
	data := models.GetUserDataById(database.GetDB(), i)
	enc, _ := json.Marshal(data)
	c.JSON(200, string(enc))
}
