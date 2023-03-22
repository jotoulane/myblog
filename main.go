package main

import (
	"fmt"
	"web01/dao/mysql"
	"web01/routes"
	"web01/settings"
	"web01/util"
)

func main() {
	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err: %v\n", err)
		return
	}

	//初始化雪花算法
	if err := util.Init(settings.Conf.AppConfig); err != nil {
		fmt.Printf("init sonwflake failed, err: %v\n", err)
		return
	}

	//初始化MySQL连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}

	// 注册路由
	r := routes.SetUp()

	r.Run(fmt.Sprintf(":%v", settings.Conf.AppConfig.Port))
}
