package models

import "gorm.io/gorm"

type ReactionContent struct {
	Id           int    `json: "id"`
	ReactionName string `json: "reaction_name"`
	Uri          string `json: "uri"`
}

func (reaction *ReactionContent) PostReaction(db *gorm.DB) {
	result := db.Create(&reaction)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetAllReactionContents(db *gorm.DB) []ReactionContent {
	var reaction_contents []ReactionContent
	result := db.Find(&reaction_contents)
	println(result.RowsAffected)

	return reaction_contents
}

func GetReaction(db *gorm.DB, id int) ReactionContent {
	var reaction_content ReactionContent
	db.First(&reaction_content, id)
	return reaction_content
}
