package models

import "gorm.io/gorm"

type BuyData struct {
	UserId  int   `json:"user_id"`
	BookIds []int `json:"book_id"`
}

type BuyResponse struct {
	Status string  `json:"status"`
	Msg    string  `json:"message"`
	Data   BuyData `json:"data"`
}

func (d *BuyData) BuyBooks(db *gorm.DB) BuyResponse {
	if d.UserId <= 0 || len(d.BookIds) == 0 {
		r := BuyResponse{"error", "invalid request", *d}
		return r
	}

	r := BuyResponse{"ok", "", *d}
	for _, v := range d.BookIds {
		mybook := MyBook{d.UserId, v}
		
		var count int64
		db.Where(mybook).Count(&count)
		if count > 0 {
			msg := "cant buy book " + string(v) + "(Already have)"
			println(msg)
			r.Msg += msg
			continue;
		}
		
		result := db.Create(&mybook)
		if err := result.Error; err != nil {
			println(err)
			r.Status = "error"
			r.Msg += " db create error in " + string(v)
		}
	}

	return r
}
