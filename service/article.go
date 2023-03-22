package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web01/dao/mysql"
	"web01/model"
	"web01/util"
)

func AddArticle(p *model.ParamNewArticle) int64 {
	article := new(model.Article)
	article.ArticleID = util.GenID()
	article.Title = p.Title
	article.Content = p.Content
	article.Short = p.Short
	article.Tags = p.Tags
	article.Author = "zhangsan"
	return mysql.AddArticle(article)
}

func UpdateArticle(c *gin.Context, p *model.ParamNewArticle) int64 {
	article := new(model.Article)
	article.ArticleID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	article.Title = p.Title
	article.Content = p.Content
	article.Short = p.Short
	article.Tags = p.Tags
	return mysql.UpdateArticle(article)
}

func ListArticle() ([]model.Article, error) {
	return mysql.ListArticle()
}

func DetailArticle(articleID int64) (model.Article, error) {
	return mysql.DetailArticle(articleID)
}

func DeleteArticle(id int64) int64 {
	return mysql.DeleteArticle(id)
}
