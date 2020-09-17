package handler

import (
	"app/database"
	"app/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiGetBook(c *gin.Context) {
	println("get all posts")
	param := c.Param("id")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetBook(database.GetDB(), intParam)
	c.JSON(200, posts)
}

func ApiGetUsersBook(c *gin.Context) {
	param := c.Param("id")
	intParam, _ := strconv.Atoi(param)
	res := models.GetUsersBook(database.GetDB(), intParam)
	c.JSON(200, res)
}

func ApiGetAllBook(c *gin.Context) {
	posts := models.GetAllBook(database.GetDB())
	c.JSON(200, posts)
}

func ApiGetContent(c *gin.Context) {
	param := c.Param("id")
	intParam, _ := strconv.Atoi(param)
	posts := models.GetBookContent(database.GetDB(), intParam)
	c.JSON(200, posts)
}
func ApiRegistBook(c *gin.Context) {
	req := &models.BookRequest{}
	err := c.BindJSON(req)
	if err != nil {
		println(err)
	}
	println("req", req)
	req.RegistBook(database.GetDB())

	c.JSON(200, req)
}

func GetBooksByTwitter(c *gin.Context) {
	param := c.Param("twitter_id")
	user, _ := models.GetUserDataByTwitterToken(database.GetDB(), param)
	res := models.GetUsersBook(database.GetDB(), user.Id)

	c.JSON(200, res)
}
