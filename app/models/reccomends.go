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

func GetMyRecommend(db *gorm.DB, sender_id int) []Recommend {
	var recommends []Recommend
	db.Where(map[string]interface{}{"sender_id": sender_id}).Find(&recommends)

	return recommends
}

func GetMyRecommended(db *gorm.DB, reciver_id int) []Recommend {
	var recommends []Recommend
	db.Where(map[string]interface{}{"reciever_id": reciver_id}).Find(&recommends)

	return recommends
}
