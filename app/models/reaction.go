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

func GetReaction(db *gorm.DB, id int) ReactionContent {
	var reaction_content ReactionContent
	db.First(&reaction_content, id)
	return reaction_content
}
