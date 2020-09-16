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

type RecommendInfo struct {
	Username       string
	BookImages     []string
	ReactionImages []string
}

func check_exist(arr []RecommendInfo, username string) (bool, int) {
	for index, v := range arr {
		if v.Username == username {
			return true, index
		}
	}
	return false, -1
}

func GetRecommendInfo(c *gin.Context) {
	param := c.Query("sender_id")
	id, _ := strconv.Atoi(param)
	recommends := models.GetMyRecommend(database.GetDB(), id)

	var recommend_infos []RecommendInfo
	for _, recommend := range recommends {
		var reciever_id = recommend.ReceiverId
		var book_id = recommend.BookId
		var reaction_content_id = recommend.ReactionContentId

		var username = models.GetUserDataById(database.GetDB(), reciever_id).Name
		var book_image = models.GetBookContent(database.GetDB(), book_id)[0].URI
		var reaction_image = models.GetReaction(database.GetDB(), reaction_content_id)[0].Uri

		isexist, index := check_exist(recommend_infos, username)

		if isexist {
			book_images := recommend_infos[index].BookImages
			book_images = append(book_images, book_image)
			recommend_infos[index].BookImages = book_images

			reaction_images := recommend_infos[index].ReactionImages
			reaction_images = append(reaction_images, reaction_image)
			recommend_infos[index].ReactionImages = reaction_images
		} else {
			var recommend_info RecommendInfo
			recommend_info.Username = username
			recommend_info.BookImages = []string{book_image}
			recommend_info.ReactionImages = []string{reaction_image}

			recommend_infos = append(recommend_infos, recommend_info)
		}
	}

	c.JSON(200, recommend_infos)
}
