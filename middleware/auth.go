package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"web01/model"
	"web01/util"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	log.Printf("hahahahha")
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.HTML(http.StatusOK, "alarm.tmpl", &model.ResponseData{
				Code: 200,
				Msg:  "访问失败,请登录!",
				Data: "/login",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			log.Printf("访问失败,无效的token,请登录!")
			c.HTML(http.StatusOK, "alarm.tmpl", &model.ResponseData{
				Code: 200,
				Msg:  "访问失败,无效的token,请登录!",
				Data: "/login",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.ParseToken(parts[1])
		if err != nil {
			c.HTML(http.StatusOK, "alarm.tmpl", &model.ResponseData{
				Code: 200,
				Msg:  "访问失败,无效的token,请登录!",
				Data: "/login",
			})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		log.Printf("mc.UserID: %v\n", mc.UserID)
		c.Set("userID", mc.UserID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
