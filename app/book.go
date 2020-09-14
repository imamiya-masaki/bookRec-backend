package main

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func apiGetBook(c *gin.Context) {
	println("get all posts")
	param := c.Param("id")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetBook(database.GetDB(), intParam)
	c.JSON(200, posts)
}
func apiGetAllBook(c *gin.Context) {
	posts := models.GetAllBook(database.GetDB())
	c.JSON(200, posts)
}
func apiGetContent(c *gin.Context) {
	param := c.Param("id")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetBookContent(database.GetDB(), intParam)
	c.JSON(200, posts)
}
