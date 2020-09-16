package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"app/database"
	"app/models"
)

func SendReaction(c *gin.Context) {
	reaction := &models.ReactionContent{}
	err := c.BindJSON(reaction)
	if err != nil {
		println(err)
	}

	reaction.SendReaction(database.GetDB())

	c.JSON(200, reaction)
}

func GetReaction(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	reactions := models.GetReaction(database.GetDB(), id)

	c.JSON(200, reactions)
}
