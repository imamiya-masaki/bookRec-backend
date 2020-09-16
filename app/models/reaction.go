package models

import "gorm.io/gorm"

type ReactionContent struct {
	Id           int    `json: "id"`
	ReactionName string `json: "reaction_name"`
	Uri          string `json: "uri"`
}

func (reaction *ReactionContent) SendReaction(db *gorm.DB) {
	result := db.Create(&reaction)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetReaction(db *gorm.DB, id int) []ReactionContent {
	var reactions []ReactionContent
	db.Where(map[string]interface{}{"id": id}).Find(&reactions)

	return reactions
}
