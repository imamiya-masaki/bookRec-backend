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
	println(coupon)
	coupon.PostCoupon(database.GetDB())

	c.JSON(200, coupon)
}

func GetCoupon(c *gin.Context) {
	coupon := models.GetCoupon(database.GetDB())
	c.JSON(200, coupon)
}

func GetCouponByUser(c *gin.Context) {
	param := c.Query("user_id")
	id, _ := strconv.Atoi(param)
	coupons := models.GetCouponByUser(database.GetDB(), id)

	c.JSON(200, coupons)
}
