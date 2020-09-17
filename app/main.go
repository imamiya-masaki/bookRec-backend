package main

import (
	"app/database"
	"app/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Origin",
			"Content-Type",
			"Content-Length",
		},
		AllowOrigins: []string{
			"*",
		},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})

	r.GET("/book/:id", handler.ApiGetBook)
	r.GET("/book", handler.ApiGetAllBook)
	r.GET("/book/:id/Content", handler.ApiGetContent)
	r.POST("book_regist", handler.ApiRegistBook)

	r.GET("/users/", handler.ApiGetAllUser)
	r.GET("/users/:id", handler.ApiGetUserdata)
	r.GET("/users/:id/library", handler.ApiGetUsersBook)
	r.POST("/regist", handler.ApiRegistUser)

	r.POST("/posttweet", handler.ApiSendTwitterTweet)

	r.POST("/recommend", handler.SendRecommend)
	r.PUT("/recommend", handler.UpdateRecommend)
	r.GET("/my_recommend", handler.GetMyRecommend)
	r.GET("/my_recommended", handler.GetMyRecommended)
	r.GET("/recommend_info", handler.GetRecommendInfo)

	r.GET("/message/:id", handler.GetMessageById)
	r.GET("/message_by_recommend_id", handler.GetMessageByRecommendId)
	r.POST("/message", handler.SendMessage)

	r.POST("/reaction", handler.PostReaction)
	r.GET("/reaction/:id", handler.GetReaction)

	r.POST("/notification", handler.SendNotification)
	r.GET("/my_notification/", handler.GetMyNotification)
	r.GET("/my_notificated/", handler.GetMyNotificated)

	r.POST("/buy", handler.ApiBuy)

	r.GET("/couponAll", handler.GetCoupon)
	r.GET("/mycoupon/:user_id", handler.GetCouponByUser)
	r.POST("/coupon", handler.PostCoupon)

	r.GET("/dashboard_info", handler.GetDashBoardInfo)

	r.Run()
}
