package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"web01/settings"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.Conf.MySqlConfig.Username,
		settings.Conf.MySqlConfig.Password,
		settings.Conf.MySqlConfig.Host,
		settings.Conf.MySqlConfig.Port,
		settings.Conf.MySqlConfig.DbName,
	)

	open, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = open

	return
}
