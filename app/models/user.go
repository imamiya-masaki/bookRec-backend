package models

import (
	"app/twitter"
	"net/url"

	"gorm.io/gorm"
)

type User struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	TwitterToken string `json:"twitter_token"`
}

type RegistUserRequest struct {
	Name         string `json:"name"`
	TwitterToken string `json:"twitter_token"`
}

type RegistUserResponse struct {
	Status  string            `json:"status"`
	Msg     string            `json:"message"`
	ReqData RegistUserRequest `json:"req_data"`
}

type TwitterPostRequest struct {
	PostMsg string `json:"post_msg"`
}

type TwitterPostResponse struct {
	Status  string             `json:"status"`
	Msg     string             `json:"message"`
	ReqData TwitterPostRequest `json:"req_data"`
}

func GetAllUsers(db *gorm.DB) ([]User, int64) {
	var users []User
	result := db.Find(&users)
	return users, result.RowsAffected
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

func (req *RegistUserRequest) RegistUser(db *gorm.DB) RegistUserResponse {
	if req.Name == "" || req.TwitterToken == "" {
		r := RegistUserResponse{"error", "field error", *req}
		return r
	}

	res := RegistUserResponse{"ok", "", *req}

	user := User{Username: req.Name, TwitterToken: req.TwitterToken}

	result := db.Create(&user)

	if err := result.Error; err != nil {
		println(err)
		res.Status = "error"
		res.Msg += "db create error"
	}

	return res
}

func (req *TwitterPostRequest) PostTwitterTweet() TwitterPostResponse {
	v := url.Values{}

	api := twitter.GetTwitterApi()

	tweet, err := api.PostTweet(req.PostMsg, v)

	res := TwitterPostResponse{"ok", "tweet send...  id:" + tweet.IdStr, *req}

	if err != nil {
		println(err.Error())
		res.Status = "error"
		res.Msg = "twitter tweet error"
	}

	return res
}
