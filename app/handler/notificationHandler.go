package handler

import (
	"app/database"
	"app/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMyNotification(c *gin.Context) {
	param := c.Query("sender_id")
	id, _ := strconv.Atoi(param)
	notification := models.GetMyNotification(database.GetDB(), id)

	c.JSON(200, notification)
}

func GetMyNotificated(c *gin.Context) {
	param := c.Query("reciever_id")
	id, _ := strconv.Atoi(param)
	notification := models.GetMyNotificated(database.GetDB(), id)

	c.JSON(200, notification)
}

func SendNotification(c *gin.Context) {
	notification := &models.Notification{}
	err := c.BindJSON(notification)
	if err != nil {
		println(err)
	}

	notification.SendNotification(database.GetDB())

	c.JSON(200, notification)
}
