package models

import "gorm.io/gorm"

type Reaction struct {
	Id           int
	ReactionName string
	Uri          string
}

func (reaction *Reaction) SendReaction(db *gorm.DB) {
	result := db.Create(&reaction)

	if err := result.Error; err != nil {
		println(err)
	}
}
