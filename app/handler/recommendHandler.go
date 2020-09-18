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
	SenderId          int    `json: "sender_id"`
	ReceiverId        int    `json: "reciever_id"`
	BookIds           []int  `json: "book_id"`
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

	var new_reciever_id int
	if req.ReceiverId == 0 {
		userRequest := &models.RegistUserRequest{
			Name:         "unknown",
			TwitterToken: req.TwitterToken,
		}
		regist_user_res := userRequest.RegistUser(database.GetDB())
		new_reciever_id = regist_user_res.User.Id
		if regist_user_res.Status == "ok" {
			println("success regist user")
		} else {
			println("error, regist user")
		}
		url := "http://bookrec-litemode.s3-website-ap-northeast-1.amazonaws.com/dist/" + "?=%22" + req.TwitterToken + "%22"
		msg := "@" + req.TwitterToken + "さん\nbookRecにて本をおすすめされました！\nメッセージ：" + req.Message + "\n" + url

		twitter_post_req := &models.TwitterPostRequest{
			PostMsg: msg,
		}

		twitter_post_res := twitter_post_req.PostTwitterTweet()

		if twitter_post_res.Status == "ok" {
			println("success tweet")
		} else {
			println("error, cannot tweet")
		}
	}

	for _, bookid := range req.BookIds {
		var reciever_id int
		if req.ReceiverId == 0 {
			reciever_id = new_reciever_id
		} else {
			reciever_id = req.ReceiverId
		}

		recommend := &models.Recommend{
			SenderId:          req.SenderId,
			ReceiverId:        reciever_id,
			BookId:            bookid,
			ReactionContentId: req.ReactionContentId,
		}
		// err = c.BindJSON(recommend)
		// println(recommend.SenderId)
		// if err != nil {
		// 	println("error")
		// 	println(err.Error())
		// } else {
		// 	println("success")
		// }

		recommend.SendReccomend(database.GetDB())
	}

	c.JSON(200, gin.H{"status": "response..."})
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
	User      models.User
	Books     []models.Book
	Reactions []models.ReactionContent
}

func check_exist(arr []RecommendInfo, username string) (bool, int) {
	for index, v := range arr {
		if v.User.Username == username {
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

		var user = models.GetUserDataById(database.GetDB(), reciever_id)
		var book = models.GetBook(database.GetDB(), book_id)
		var reaction = models.GetReaction(database.GetDB(), reaction_content_id)

		isexist, index := check_exist(recommend_infos, user.Username)

		if isexist {
			books := recommend_infos[index].Books
			books = append(books, book)
			recommend_infos[index].Books = books

			reactions := recommend_infos[index].Reactions
			reactions = append(reactions, reaction)
			recommend_infos[index].Reactions = reactions
		} else {
			var recommend_info RecommendInfo
			recommend_info.User = user
			recommend_info.Books = []models.Book{book}
			recommend_info.Reactions = []models.ReactionContent{reaction}

			recommend_infos = append(recommend_infos, recommend_info)
		}
	}

	c.JSON(200, recommend_infos)
}

func GetRecommendedInfo(c *gin.Context) {
	param := c.Query("reciever_id")
	id, _ := strconv.Atoi(param)
	recommends := models.GetMyRecommend(database.GetDB(), id)

	var recommend_infos []RecommendInfo
	for _, recommend := range recommends {
		var sender_id = recommend.ReceiverId
		var book_id = recommend.BookId
		var reaction_content_id = recommend.ReactionContentId

		var user = models.GetUserDataById(database.GetDB(), sender_id)
		var book = models.GetBook(database.GetDB(), book_id)
		var reaction = models.GetReaction(database.GetDB(), reaction_content_id)

		isexist, index := check_exist(recommend_infos, user.Username)

		if isexist {
			books := recommend_infos[index].Books
			books = append(books, book)
			recommend_infos[index].Books = books

			reactions := recommend_infos[index].Reactions
			reactions = append(reactions, reaction)
			recommend_infos[index].Reactions = reactions
		} else {
			var recommend_info RecommendInfo
			recommend_info.User = user
			recommend_info.Books = []models.Book{book}
			recommend_info.Reactions = []models.ReactionContent{reaction}

			recommend_infos = append(recommend_infos, recommend_info)
		}
	}

	c.JSON(200, recommend_infos)
}

func ApiUnanonymizeByReccomendId(c *gin.Context) {
	p := &models.Recommend{}
	err := c.BindJSON(p)
	if err != nil {
		println(err)
	}

	if res := p.Unanonymize(database.GetDB()); res {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Unanonymized ReccomendId:" + string(p.Id),
			"req_id":  p.Id,
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "error",
			"message": "Unanonymizing Failed. ReccomendId:" + string(p.Id),
			"req_id":  p.Id,
		})
	}

}
