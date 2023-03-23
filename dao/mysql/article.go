package mysql

import (
	"log"
	"web01/model"
)

func AddArticle(article *model.Article) (nums int64) {
	res := db.Create(article)
	if res.Error != nil {
		log.Printf("db.Create(article) failed, err: %v\n", res.Error)
		return
	}
	return res.RowsAffected
}

func UpdateArticle(article *model.Article) (nums int64) {
	res := db.Model(&article).Where("article_id=?", article.ArticleID).Updates(model.Article{
		Title:   article.Title,
		Content: article.Content,
		Short:   article.Short,
		Tags:    article.Tags,
	})
	if res.Error != nil {
		log.Printf("db.Model(&article).Updates failed, err: %v\n", res.Error)
		return
	}
	return res.RowsAffected
}

func ListArticle() (articles []model.Article, err error) {
	articles = make([]model.Article, 0)
	tx := db.Find(&articles)
	return articles, tx.Error
}

func ListTagsArticle(tags string) (articles []model.Article, err error) {
	articles = make([]model.Article, 0)
	tx := db.Where("tags=?", tags).Find(&articles)
	return articles, tx.Error
}

func LatestArticles() (articles []model.Article, err error) {
	articles = make([]model.Article, 0)
	//tx := db.Select("article_id", "title").Order("created_at").Limit(3).Find(&articles)
	tx := db.Order("created_at desc").Limit(4).Find(&articles)
	return articles, tx.Error
}

func DetailArticle(articleID int64) (article model.Article, err error) {
	tx := db.Where("article_id=?", articleID).Find(&article)
	return article, tx.Error
}

func DeleteArticle(articleID int64) int64 {
	tx := db.Where("article_id=?", articleID).Delete(&model.Article{})
	return tx.RowsAffected
}
