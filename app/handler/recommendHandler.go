package handler

import (
	"app/database"
	"app/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetMyRecommend(c *gin.Context) {
	param := c.Query("sender_id")
	id, _ := strconv.Atoi(param)
	recommends := models.GetMyRecommend(database.GetDB(), id)

	c.JSON(200, recommends)
}

func GetMyRecommended(c *gin.Context) {
	param := c.Query("receiver_id")
	id, _ := strconv.Atoi(param)
	recommends := models.GetMyRecommended(database.GetDB(), id)

	c.JSON(200, recommends)
}

func SendRecommend(c *gin.Context) {
	recommend := &models.Recommend{}
	err := c.BindJSON(recommend)
	// err := c.Bind(recommend)
	if err != nil {
		println(err)
	}

	recommend.SendReccomend(database.GetDB())

	c.JSON(200, recommend)
}

// type RecommendInfo struct {
// 	Username       string
// 	BookImages     []string
// 	ReactionImages []string
// }

// func GetRecommendInfo(c *gin.Context) {
// 	param := c.Query("sender_id")
// 	id, _ := strconv.Atoi(param)
// 	recommends := models.GetMyRecommend(database.GetDB(), id)

// 	recommends_size = len(recommends)

// 	// var recommend_infos []RecommendInfo
// 	// for i, recommend := range recommends {
// 	// 	var recommend_info RecommendInfo

// 	// 	messages := GetMessageByRecommendId(database.GetDB(), recommend.Id)
// 	// 	message := messages[0]

// 	// }
// }
