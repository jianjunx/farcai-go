package service

import (
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"time"

	"github.com/gogf/gf/os/gtime"
)

var Article = &articleService{}

type articleService struct{}

func (*articleService) AddArticle(article *model.ArticleItem) error {
	article.ArticleID = time.Now().Unix()
	article.CreateAt = gtime.Datetime()
	article.UpdateAt = gtime.Datetime()
	err := dao.Category.GetCategoryItem(&article.Category, article.CategoryID)
	if err != nil {
		return err
	}
	return dao.Article.PutArticleItem(article)
}

func (*articleService) UpdateArticle(article *model.ArticleItem) error {
	article.UpdateAt = gtime.Datetime()
	return dao.Article.PutArticleItem(article)
}

func (*articleService) DeleteArticle(articleId int64) error {
	return dao.Article.DeleteArticleItem(articleId)
}
