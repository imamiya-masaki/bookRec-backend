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

func GetReaction(db *gorm.DB, id int) []Recommend {
	var reactions []Recommend
	db.Where(map[string]interface{}{"id": id}).Find(&reactions)

	return recommends
}
