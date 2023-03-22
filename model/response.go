package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"` //没有值忽略
}

func ResponseDetail(c *gin.Context, code int64, msg string, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
