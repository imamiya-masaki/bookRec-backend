package main

import (
	"app/database"
	"app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello, World")
	})

	r.GET("/book/:id", handler.ApiGetBook)
	r.GET("/book", handler.ApiGetAllBook)
	r.GET("/book/:id/Content", handler.ApiGetContent)

	r.GET("/users/:id", handler.ApiGetUserdata)
	r.POST("/regist", handler.ApiRegistUser)

	r.POST("/recommend", handler.SendRecommend)
	r.GET("/my_recommend", handler.GetMyRecommend)
	r.GET("/my_recommended", handler.GetMyRecommended)

	r.GET("/message/:id", handler.GetMessageById)
	r.GET("/message_by_recommend_id", handler.GetMessageByRecommendId)
	r.POST("/message", handler.SendMessage)

	r.POST("/reaction", handler.SendReaction)
	r.GET("/reaction/:id", handler.GetReaction)

	r.POST("/notification", handler.SendNotification)
	r.GET("/my_notification/", handler.GetMyNotification)
	r.GET("/my_notificated/", handler.GetMyNotificated)

	r.POST("/buy", handler.ApiBuy)

	r.Run()
}
