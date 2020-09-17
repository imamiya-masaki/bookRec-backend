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

type SendRecommendReqest struct {
	Id                int    `json: "id"`
	SenderId          int    `json: "sender_id"`
	ReceiverId        int    `json: "reciever_id"`
	BookId            int    `json: "book_id"`
	ReactionContentId int    `json: "reaction_content_id"`
	TwitterToken      string `json: "twitter_token"`
	Message           string `json: "message"`
}

func SendRecommend(c *gin.Context) {
	req := &SendRecommendReqest{}
	err := c.BindJSON(req)
	if err != nil {
		println(err)
	}

	if req.ReceiverId == 0 {
		userRequest := &models.RegistUserRequest{
			Name:         "unknown",
			TwitterToken: req.TwitterToken,
		}
		regist_user_res := userRequest.RegistUser(database.GetDB())
		if regist_user_res.Status == "ok" {
			println("success regist user")
		} else {
			println("error, regist user")
		}

		twitter_post_req := &models.TwitterPostRequest{
			PostMsg: req.Message,
		}
		twitter_post_res := twitter_post_req.PostTwitterTweet()

		// c.JSON(200, twitter_post_res)
		if twitter_post_res.Status == "ok" {
			println("success tweet")
		} else {
			println("error, cannot tweet")
		}
	}

	recommend := &models.Recommend{}
	err = c.BindJSON(recommend)
	if err != nil {
		println(err)
	}

	recommend.SendReccomend(database.GetDB())

	c.JSON(200, recommend)
}

func UpdateRecommend(c *gin.Context) {
	recommend := &models.Recommend{}
	err := c.BindJSON(recommend)
	if err != nil {
		println(err)
	}

	recommend.UpdateRecommend(database.GetDB())

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

		var username = models.GetUserDataById(database.GetDB(), reciever_id).Username
		var book_image = models.GetBook(database.GetDB(), book_id).URI
		var reaction_image = models.GetReaction(database.GetDB(), reaction_content_id).Uri

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
