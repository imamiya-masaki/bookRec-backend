package handler

import (
	"app/database"
	"app/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReactionInfo struct {
	Username      string
	ReactionImage string
}

type DashBoardInfo struct {
	Reactions        []ReactionInfo
	RecommendCount   int
	RecommendedCount int
	ReactionCount    int
	CouponCount      int
	BookCount        int
}

func getReactionInfo(id int) []ReactionInfo {
	// id : 自分のID
	recommends := models.GetMyRecommend(database.GetDB(), id)

	var reaction_infos []ReactionInfo
	for _, recommend := range recommends {
		user_id := recommend.ReceiverId
		reaction_id := recommend.ReactionContentId

		username := models.GetUserDataById(database.GetDB(), user_id).Username
		reaction_image := models.GetReaction(database.GetDB(), reaction_id).ReactionName

		reaction_info := ReactionInfo{
			Username:      username,
			ReactionImage: reaction_image,
		}

		reaction_infos = append(reaction_infos, reaction_info)
	}

	return reaction_infos
}

func calcRecommendCount(id int) int {
	recommends := models.GetMyRecommend(database.GetDB(), id)

	return len(recommends)
}

func calcRecommendedCount(id int) int {
	recommendeds := models.GetMyRecommended(database.GetDB(), id)

	return len(recommendeds)
}

func calcReactionCount(id int) int {
	recommends := models.GetMyRecommend(database.GetDB(), id)

	var count int
	for _, recommend := range recommends {
		if recommend.ReactionContentId != -1 {
			count += 1
		}
	}

	return count
}

func calcCouponCount(id int) int {
	coupons := models.GetCouponByUser(database.GetDB(), id)

	return len(coupons)
}

func calcBookCount(id int) int {
	books := models.GetUsersBook(database.GetDB(), id)

	return len(books.Books)
}

func GetDashBoardInfo(c *gin.Context) {
	param := c.Query("id")
	id, _ := strconv.Atoi(param)
	reaction_infos := getReactionInfo(id)
	recommend_count := calcRecommendCount(id)
	recommended_count := calcRecommendedCount(id)
	reaction_count := calcReactionCount((id))
	coupon_count := calcCouponCount((id))
	book_count := calcBookCount(id)

	dashboard_info := DashBoardInfo{
		Reactions:        reaction_infos,
		RecommendCount:   recommend_count,
		RecommendedCount: recommended_count,
		ReactionCount:    reaction_count,
		CouponCount:      coupon_count,
		BookCount:        book_count,
	}

	c.JSON(200, dashboard_info)
}
