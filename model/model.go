package model

import (
	"gorm.io/gorm"
)

type User struct {
	UserID   int64  `json:"user_id" gorm:"type:varchar(100);unique_index"`
	Username string `json:"username" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	Grade    string `json:"grade" gorm:"type:varchar(100)"`
	gorm.Model
	Token string
}

type Article struct {
	ArticleID  int64  `json:"article_id" gorm:"type:varchar(200);unique_index"`
	Title      string `json:"title" gorm:"type:varchar(20)"`
	Content    string `json:"content" gorm:"type:longtext"`
	Short      string `json:"short" gorm:"type:varchar(20)"`
	Tags       string `json:"tags" gorm:"type:varchar(10)"`
	Author     string `json:"author" gorm:"type:varchar(10)"`
	ReadNum    int64  `json:"read_num" gorm:"type:int"`
	LikeNum    int64  `json:"like_num" gorm:"type:int"`
	CollectNum int64  `json:"collect_num" gorm:"type:int"`
	gorm.Model
}
