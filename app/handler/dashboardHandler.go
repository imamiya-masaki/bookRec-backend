package handler

import "app/database"

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

func getReactionInfo(id) {
	// id : 自分のID
	recommends := models.GetMyRecommend(database.GetDB(), id)

	var reaction_infos []ReactionInfo
	for _, recommend := range recommends {
		user_id := recommend.ReceiverId
		reaction_id := recommend.ReactionContentId

		username := models.GetUserDataById(database.GetDB, user_id).Username

		reaction_info := ReactionInfo{
			Username: username, 
			TwitterToken: req.TwitterToken
		}
	}
}
