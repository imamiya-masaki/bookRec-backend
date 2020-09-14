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

func GetMyRecommend(db *gorm.DB) []Recommend {
	var recommends []Recommend
	// result := db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&recommends)
	result := db.Find(&recommends)
	if err := result.Error; err != nil {
		println(err)
	}

	return recommends
}
