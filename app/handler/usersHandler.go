package handler

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiGetAllUser(c *gin.Context) {
	users, rows := models.GetAllUsers(database.GetDB())
	c.JSON(200, gin.H{
		"count": rows,
		"users": users,
	})
}

func ApiGetUserdata(c *gin.Context) {
	println("getting userdata")
	param := c.Param("id")
	var id int
	id, _ = strconv.Atoi(param)
	data := models.GetUserDataById(database.GetDB(), id)
	c.JSON(200, data)
}

func ApiRegistUser(c *gin.Context) {
	req := &models.RegistUserRequest{}
	err := c.BindJSON(req)

	if err != nil {
		println(err)
	}

	res := req.RegistUser(database.GetDB())

	c.JSON(200, res)
}

func ApiSendTwitterTweet(c *gin.Context) {
	req := &models.TwitterPostRequest{}
	err := c.BindJSON(req)

	if err != nil {
		println(err)
	}

	res := req.PostTwitterTweet()

	c.JSON(200, res)
}

func GetUserDataByTwitter(c *gin.Context) {
	param := c.Param("twitter_id")
	user, _ := models.GetUserDataByTwitterToken(database.GetDB(), param)

	c.JSON(200, user)
}
