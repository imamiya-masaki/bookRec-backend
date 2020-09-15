package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"app/database"
	"app/models"
)

func sendReaction(c *gin.Context) {
	reaction := &models.Reaction{}
	err := c.BindJSON(reaction)
	if err != nil {
		println(err)
	}
}

func getReaction(c *gin.Context) {
	param := c.Query("id")
	id, _ := strconv.Atoi(param)
	reactions := models.GetReaction(database.GetDB(), id)

	c.JSON(200, reactions)
}
