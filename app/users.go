package main

import (
	"app/database"
	"app/models"
	"strconv"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func apiGetUserdata(c *gin.Context) {
	println("getting userdata")
	param := c.Param("id")
	var id int
	id, _ = strconv.Atoi(param)
	data := models.GetUserDataById(database.GetDB(), id)
	enc, _ := json.Marshal(data)
	c.JSON(200, string(enc))
}
