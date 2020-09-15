package main

import (
	"github.com/gin-gonic/gin"

	"app/models"
)

func sendReaction(c *gin.Context) {
	reaction := &models.Reaction{}
	err := c.BindJSON(reaction)
	if err != nil {
		println(err)
	}
}
