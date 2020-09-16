package models

import "gorm.io/gorm"

type Reaction struct {
	Id         int
	ReactionId string
	SenderId   int
	RecieverId int
}

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

func GetReaction(db *gorm.DB, id int) ReactionContent {
	var reaction_content ReactionContent
	db.First(&reaction_content, id)
	return reaction_content
}
