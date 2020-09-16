package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"app/database"
	"app/models"
)

func PostCoupon(c *gin.Context) {
	coupon := &models.Coupon{}
	err := c.BindJSON(coupon)
	if err != nil {
		println(err)
	}

	coupon.PostCoupon(database.GetDB())

	c.JSON(200, coupon)
}

func GetCoupon(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	coupon := models.GetCoupon(database.GetDB(), id)

	c.JSON(200, coupon)
}

func GetCouponByUser(c *gin.Context) {
	param := c.Query("user_id")
	id, _ := strconv.Atoi(param)
	coupons := models.GetCouponByUser(database.GetDB(), id)

	c.JSON(200, coupons)
}
