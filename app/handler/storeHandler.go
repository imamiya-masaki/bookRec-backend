package handler

import (
	"app/database"
	"app/models"

	"github.com/gin-gonic/gin"
)

func ApiBuy(c *gin.Context) {
	b := &models.BuyData{}
	err := c.BindJSON(b)

	if err != nil {
		println(err)
	}

	res := b.BuyBooks(database.GetDB())

	c.JSON(200, res)
}
