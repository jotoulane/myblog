package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web01/model"
	"web01/service"
)

func RegisterHandler(c *gin.Context) {
	//获取参数
	p := new(model.ParamRegister)
	err := c.ShouldBind(p)
	if err != nil {
		log.Printf("c.ShouldBind(p) failed, err: %v\n", err)
	}
	//创建用户
	res := service.CreateUser(p)
	log.Printf("res: %v\n", res)
	//返回响应
	if res == 0 {
		//已经存在 或者创建失败
		c.HTML(http.StatusOK, "alarm.tmpl", &model.ResponseData{
			Code: 200,
			Msg:  "注册失败,用户名已经存在!",
			Data: "/register",
		})
	} else {
		//创建成功
		c.HTML(http.StatusOK, "alarm.tmpl", &model.ResponseData{
			Code: 200,
			Msg:  "注册成功,请返回登录",
			Data: "/login",
		})
	}
}

func LoginHandler(c *gin.Context) {
	//获取参数
	p := new(model.ParamLogin)
	err := c.ShouldBind(p)
	if err != nil {
		log.Printf("c.ShouldBind(p) failed, err: %v\n", err)
	}
	//判断用户,返回响应
	err, user := service.QueryUserByUsername(p)
	if err != nil {
		c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
			Code: 200,
			Msg:  "密码错误",
			Data: "/login",
		})
	} else {
		c.Set("user", user)
		c.HTML(http.StatusOK, "home.tmpl", model.ResponseData{
			Code: 200,
			Msg:  "登录成功!",
			Data: user.Token,
		})
	}
}
