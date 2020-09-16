package handler

import (
	"app/database"
	"app/models"

	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ApiGetUserdata(c *gin.Context) {
	println("getting userdata")
	param := c.Param("id")
	var id int
	id, _ = strconv.Atoi(param)
	data := models.GetUserDataById(database.GetDB(), id)
	enc, _ := json.Marshal(data)
	c.JSON(200, string(enc))
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
