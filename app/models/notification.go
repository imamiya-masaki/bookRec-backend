package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	Id         int       `json:"id"`
	SenderId   int       `json:"sender_id"`
	RecieverId int       `json:"reciever_id"`
	ActionType string    `json:"action_type"`
	IdValue    int       `json:"id_value"`
	CreatedAt  time.Time `json: "created_at"`
}

func (notification *Notification) SendNotification(db *gorm.DB) {
	result := db.Create(&notification)

	if err := result.Error; err != nil {
		println(err)
	}
}

func GetMyNotification(db *gorm.DB, id int) []Notification {
	var notifications []Notification
	db.Where(map[string]interface{}{"sender_id": id}).Find(&notifications)

	return notifications
}

func GetMyNotificated(db *gorm.DB, id int) []Notification {
	var notifications []Notification
	db.Where(map[string]interface{}{"reciever_id": id}).Find(&notifications)

	return notifications
}
