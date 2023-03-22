package mysql

import (
	"log"
	"web01/model"
)

func InsertUser(user *model.User) (nums int64) {
	res := db.Create(user)
	if res.Error != nil {
		log.Printf("db.Create(user) failed, err: %v\n", res.Error)
		return
	}
	return res.RowsAffected
}

func FindUserOrNot(username string) int64 {
	user := new(model.User)
	res := db.Find(user, "username=?", username)
	if res.Error != nil {
		log.Printf("dao FindUserOrNot(username string) failed, err: %v\n", res.Error)
	}
	return res.RowsAffected
}

func QueryUserByUsername(username string) (err error, user *model.User) {
	res := db.Find(&user, "username=?", username)
	return res.Error, user
}
