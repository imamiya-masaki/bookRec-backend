package handler

import (
	"app/database"
	"app/models"

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

func calcRecommendCount(id int) {

}

func GetDashBoardInfo(c *gin.Context) {
	reaction_infos := getReactionInfo(1)

	c.JSON(200, reaction_infos)
}
