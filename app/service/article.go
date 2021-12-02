package service

import (
	"errors"
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
	ctgs := []model.CategoryItem{}
	err := dao.Category.GetCategoryItem(&ctgs, article.CategoryID)
	if err != nil {
		return err
	}
	if len(ctgs) == 0 {
		return errors.New("分类不正确")
	}
	article.Category = ctgs[0]
	return dao.Article.PutArticleItem(article)
}

func (*articleService) UpdateArticle(article *model.ArticleItem) error {
	article.UpdateAt = gtime.Datetime()
	ctgs := []model.CategoryItem{}
	err := dao.Category.GetCategoryItem(&ctgs, article.CategoryID)
	if err != nil {
		return err
	}
	if len(ctgs) == 0 {
		return errors.New("分类不正确")
	}
	article.Category = ctgs[0]
	return dao.Article.PutArticleItem(article)
}

func (*articleService) DeleteArticle(articleId int64) error {
	return dao.Article.DeleteArticleItem(articleId)
}

func (*articleService) GetArticleItem(articleId int64) (*model.ArticleItem, error) {
	articles := []model.ArticleItem{}
	err := dao.Article.GetArticleItem(&articles, articleId)
	if err != nil {
		return nil, err
	}
	if len(articles) == 0 {
		return nil, errors.New("没有找到该文章")
	}
	return &articles[0], nil
}
