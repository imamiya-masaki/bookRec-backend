package models

import "gorm.io/gorm"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	TwitterToken string `json:"twitter_token"`
}

func GetUserDataById(db *gorm.DB, id int) User {
	var user User
	result := db.Where("id = ?", id).Find(&user)
	
	println(result.RowsAffected)
	
    //println("GetUserDataById")
    //println(id)
    //user.Id = id
    //user.Name = "test"
    //user.TwitterToken = "token"
    
	return user
}