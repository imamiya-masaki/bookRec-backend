package models

import (
	"gorm.io/gorm"
)

type Coupon struct {
	Id      int
	UserId  int
	Percent int
}

func (coupon *Coupon) PostCoupon(db *gorm.DB) {
	result := db.Create(&coupon)
	println(result)
	if err := result.Error; err != nil {
		println(err)
	}
}

func GetCouponByUser(db *gorm.DB, id int) []Coupon {
	var coupons []Coupon
	db.Where(map[string]interface{}{"user_id": id}).Find(&coupons)

	return coupons
}

func GetCoupon(db *gorm.DB) []Coupon {
	var coupon []Coupon
	db.Find(&coupon)
	return coupon
}
