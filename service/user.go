package service

import (
	"errors"
	"log"
	"web01/dao/mysql"
	"web01/model"
	"web01/util"
)

// CreateUser 注册
func CreateUser(p *model.ParamRegister) int64 {
	//检测用户名是否存在
	res := mysql.FindUserOrNot(p.Username)
	if res != 0 {
		log.Printf("mysql.QueryUserByUsername(p.Username) failed")
		return 0
	}
	u := new(model.User)
	u.UserID = util.GenID()
	u.Username = p.Username
	u.Password = util.EncryptPassword(p.Password)
	u.Grade = "normal"
	return mysql.InsertUser(u)
}

// QueryUserByUsername 登录查询用户信息是否征正确
func QueryUserByUsername(p *model.ParamLogin) (err error, user *model.User) {
	log.Printf("p: %v\n", p)
	err, user = mysql.QueryUserByUsername(p.Username)
	if err != nil {
		log.Printf("mysql.QueryUserByUsername(p.Username) failed, err: %v\n", err)
		return errors.New("service is busy"), nil
	}
	if util.EncryptPassword(p.Password) != user.Password {
		log.Printf("the user password failed, err: %v\n", errors.New("password is incorrect"))
		return errors.New("password is incorrect"), nil
	}

	token, err := util.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
