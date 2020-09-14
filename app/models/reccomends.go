package models

import "gorm.io/gorm"

type Recommend struct {
	Id                int
	SenderId          int
	ReceiverId        int
	BookId            int
	ReactionContentId int
}

func (recommend *Recommend) SendReccomend(db *gorm.DB) {
	result := db.Create(&recommend)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetMyRecommend(db *gorm.DB, id int) []Recommend {
	var recommends []Recommend
	db.Where(map[string]interface{}{"sender_id": id}).Find(&recommends)

	return recommends
}
