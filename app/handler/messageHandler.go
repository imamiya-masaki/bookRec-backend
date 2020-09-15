package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"app/database"
	"app/models"
)

func SendMessage(c *gin.Context) {
	message := &models.Message{}
	err := c.BindJSON(message)
	if err != nil {
		println(err)
	}

	message.SendMessage(database.GetDB())

	c.JSON(200, message)
}

func GetMessageById(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	messages := models.GetMessageById(database.GetDB(), id)

	c.JSON(200, messages)
}

func GetMessageByRecommendId(c *gin.Context) {
	param := c.Query("recommend_id")
	id, _ := strconv.Atoi(param)
	messages := models.GetMessageByRecommendId(database.GetDB(), id)

	c.JSON(200, messages)
}
