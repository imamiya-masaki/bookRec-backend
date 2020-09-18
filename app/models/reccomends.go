package models

import "gorm.io/gorm"

type Recommend struct {
	Id                int `json: "id"`
	SenderId          int `json: "sender_id"`
	ReceiverId        int `json: "reciever_id"`
	BookId            int `json: "book_id"`
	ReactionContentId int `json: "reaction_content_id"`
	IsUnanonymized    int `json: "is_unanonymized"`
}

func (recommend *Recommend) SendReccomend(db *gorm.DB) {
	recommend.ReactionContentId = -1
	result := db.Create(&recommend)

	if err := result.Error; err != nil {
		println(err)
	}
}

func (recommend *Recommend) UpdateRecommend(db *gorm.DB) {
	// db.Model(&recommend).Update("reaction_content_id", id)
	result := db.Save(&recommend)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetMyRecommend(db *gorm.DB, sender_id int) []Recommend {
	var recommends []Recommend
	db.Where(map[string]interface{}{"sender_id": sender_id}).Find(&recommends)

	return recommends
}

func GetMyRecommended(db *gorm.DB, receiver_id int) []Recommend {
	var recommends []Recommend
	db.Where(map[string]interface{}{"receiver_id": receiver_id}).Find(&recommends)

	return recommends
}

func (recommend *Recommend) Unanonymize(db *gorm.DB) bool {
	result := db.Model(&Recommend{}).Where("id = ?", recommend.Id).Update("is_unanonymized", "1")

	return (result.RowsAffected == 1)
}
