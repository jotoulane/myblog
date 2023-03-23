package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"log"
	"net/http"
	"strconv"
	"web01/model"
	"web01/service"
)

func HomePage(c *gin.Context) {
	//articles := make([]model.ParamLatestArticles, 3)
	articles, err := service.LatestArticles()
	if err != nil {
		log.Printf("service.LatestArticles() failed, err: %v\n", err)
		return
	}
	c.HTML(http.StatusOK, "home.tmpl", model.ResponseData{
		Code: 200,
		Msg:  nil,
		Data: articles,
	})
}

// NewArticle 新建文章
func NewArticle(c *gin.Context) {
	//获取参数
	p := new(model.ParamNewArticle)
	if err := c.ShouldBind(p); err != nil {
		log.Printf("c.ShouldBind(p) failed, err: %v\n", err)
	}
	//业务处理
	nums := service.AddArticle(p)
	if nums == 0 {
		log.Printf("添加失败，请重试！")
		c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
			Code: 200,
			Msg:  "新建失败,请重新创建",
			Data: "/new",
		})
	}
	//返回相应
	c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
		Code: 200,
		Msg:  "新建成功",
		Data: "/lists",
	})
}

// ListArticle 文章列表
func ListArticle(c *gin.Context) {
	articles, err := service.ListArticle()
	if err != nil {
		log.Printf("service.ListArticle() failed, err: %v\n", err)
	}
	c.HTML(http.StatusOK, "articleList.tmpl", model.ResponseData{
		Code: 200,
		Msg:  "查询成功",
		Data: articles,
	})
}

// ListTagsArticle 文章列表，根据tags
func ListTagsArticle(c *gin.Context) {
	param := c.Param("tags")
	articles, err := service.ListTagsArticle(param)
	if err != nil {
		log.Printf("service.ListTagsArticle(params) failed, err: %v\n", err)
	}
	c.HTML(http.StatusOK, "articleList.tmpl", model.ResponseData{
		Code: 200,
		Msg:  "查询成功",
		Data: articles,
	})
}

// ManageListArticle 管理员文章列表
func ManageListArticle(c *gin.Context) {
	articles, err := service.ListArticle()
	if err != nil {
		log.Printf("service.ListArticle() failed, err: %v\n", err)
	}
	c.HTML(http.StatusOK, "articleListManage.tmpl", model.ResponseData{
		Code: 200,
		Msg:  "查询成功",
		Data: articles,
	})
}

// DetailArticle 文章详情列表
func DetailArticle(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt failed,err:%v\n", err)
	}
	detailArticle, err := service.DetailArticle(id)
	if err != nil {
		log.Printf("service.DetailArticle(article.ArticleID) failed,err:%v\n", err)
	}
	//将输入的markdown格式的内容转化为html元素进行渲染
	html := blackfriday.MarkdownCommon([]byte(detailArticle.Content))
	detailArticle.Content = string(html)
	c.HTML(http.StatusOK, "detail.tmpl", model.ResponseData{
		Code: 200,
		Msg:  nil,
		Data: detailArticle,
	})
}

// ManageDetailArticle 管理猿文章详情列表
func ManageDetailArticle(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt failed,err:%v\n", err)
	}
	detailArticle, err := service.DetailArticle(id)
	if err != nil {
		log.Printf("service.DetailArticle(article.ArticleID) failed,err:%v\n", err)
	}
	//将输入的makedown格式的内容转化为html元素进行渲染
	html := blackfriday.MarkdownCommon([]byte(detailArticle.Content))
	detailArticle.Content = string(html)
	c.HTML(http.StatusOK, "detailManage.tmpl", model.ResponseData{
		Code: 200,
		Msg:  nil,
		Data: detailArticle,
	})
}

// ManageUpdate 文章详情页面
func ManageUpdate(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt failed,err:%v\n", err)
	}
	detailArticle, err := service.DetailArticle(id)
	if err != nil {
		log.Printf("service.DetailArticle(article.ArticleID) failed,err:%v\n", err)
	}
	c.HTML(http.StatusOK, "detailUpdate.tmpl", model.ResponseData{
		Code: 200,
		Msg:  nil,
		Data: detailArticle,
	})
}

// ManageUpdateSubmit 更新文章
func ManageUpdateSubmit(c *gin.Context) {
	id := c.Param("id")
	//获取参数
	p := new(model.ParamNewArticle)
	if err := c.ShouldBind(p); err != nil {
		log.Printf("c.ShouldBind(p) failed, err: %v\n", err)
	}
	//业务处理
	nums := service.UpdateArticle(c, p)
	if nums == 0 {
		log.Printf("更新失败，请重试！")
		c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
			Code: 200,
			Msg:  "新建失败,请重新创建",
			//Data: "/update",
			Data: fmt.Sprintf("/update/%v", id),
		})
	}
	c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
		Code: 200,
		Msg:  "更新成功，点击返回查看详情",
		Data: fmt.Sprintf("/manage/detail/%v", id),
	})
}

func ManageDeleteSubmit(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Printf("strconv.ParseInt failed,err:%v\n", err)
	}
	deleteArticle := service.DeleteArticle(id)
	if deleteArticle == 0 {
		log.Printf("service.DeleteArticle(article) failed,err:%v\n", err)
		c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
			Code: 200,
			Msg:  "删除失败",
			Data: "/manage/lists",
		})
		return
	}
	c.HTML(http.StatusOK, "alarm.tmpl", model.ResponseData{
		Code: 200,
		Msg:  "删除成功",
		Data: "/manage/lists",
	})
}
