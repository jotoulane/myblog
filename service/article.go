package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"web01/dao/mysql"
	"web01/dao/redis"
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
	articles, err := mysql.ListArticle()
	for i := 0; i < len(articles); i++ {
		nums, _ := redis.GetViewNums(strconv.FormatInt(articles[i].ArticleID, 10))
		numsLike, _ := redis.GetLikeNums(strconv.FormatInt(articles[i].ArticleID, 10))
		articles[i].ReadNum = nums
		articles[i].LikeNum = numsLike
	}
	return articles, err
}

// ListTagsArticle 根据标签查询文章列表
func ListTagsArticle(tags string) ([]model.Article, error) {
	return mysql.ListTagsArticle(tags)
}

// LatestArticles 查询最新文章
func LatestArticles() ([]model.Article, error) {
	return mysql.LatestArticles()
}

func DetailArticle(articleID int64) (model.Article, error) {
	return mysql.DetailArticle(articleID)
}

func DeleteArticle(id int64) int64 {
	return mysql.DeleteArticle(id)
}

func LikeArticle(id string) error {
	return redis.AddLikeNums(id)
}
