package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	Id          int
	ReccomendId int
	Sendedtime  time.Time
	Content     string
}

func (message *Message) SendMessage(db *gorm.DB) {
	message.Sendedtime = time.Now()
	result := db.Create(&message)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetMessageById(db *gorm.DB, id int) []Message {
	var message []Message
	db.Where(map[string]interface{}{"id": id}).Find(&message)

	return message
}

func GetMessageByRecommendId(db *gorm.DB, id int) []Message {
	var messages []Message
	db.Where(map[string]interface{}{"reccomend_id": id}).Find(&messages)

	return messages
}
