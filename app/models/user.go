package models

import "gorm.io/gorm"

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
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

type TwitterDmRequest struct {
	TwitterToken string `json:"twitter_token"`
	DmMsg        string `json:"dm_msg"`
}

type TwitterDmResponse struct {
	Status  string           `json:"status"`
	Msg     string           `json:"message"`
	ReqData TwitterDmRequest `json:"req_data"`
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

	user := User{Name: req.Name, TwitterToken: req.TwitterToken}

	result := db.Create(&user)

	if err := result.Error; err != nil {
		println(err)
		res.Status = "error"
		res.Msg += "db create error"
	}

	return res
}

func (req *TwitterDmRequest) SendTwitterDm(db *gorm.DB) TwitterDmResponse {

	res := TwitterDmResponse{"error", "function isnt exist", *req}

	return res
}
