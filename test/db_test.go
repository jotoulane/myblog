package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
	"web01/model"
	"web01/util"
)

func TestName(t *testing.T) {
	dsn := "root:lihuawei@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	open.AutoMigrate(&model.User{})
}

func TestArticle(t *testing.T) {
	dsn := "root:lihuawei@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	open.AutoMigrate(&model.Article{})
}

func TestTime(t *testing.T) {
	now := time.Now() //获取当前时间
	fmt.Printf("current time:%v\n", now)
	fmt.Printf("current time:%T\n", now)
	transform := util.TimeTransform(now)
	fmt.Printf("transform: %v\n", transform)
	fmt.Printf("transform: %T\n", transform)
}
