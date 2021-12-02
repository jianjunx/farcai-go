package service

import (
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"math"
	"sync"
)

var Html = &htmlService{}

type htmlService struct{}

func (*htmlService) Home(ctgId, page int64) (*[]model.ArticleItem, *[]model.CategoryItem, int64, error) {
	var (
		ws    sync.WaitGroup
		total int64
		err   error
	)
	articles := []model.ArticleItem{}
	categorys := []model.CategoryItem{}

	ws.Add(2)
	go func() {
		defer ws.Done()
		total, err = dao.Article.GetArticlePages(&articles, ctgId, page)
	}()
	go func() {
		defer ws.Done()
		err = dao.Category.GetCategorys(&categorys)
	}()
	ws.Wait()
	return &articles, &categorys, total, err
}

func (*htmlService) GetPages(total float64) (pages []float64) {
	var size float64 = 10
	pages = []float64{}
	for i, len := 0.0, math.Ceil(total/size); i < len; i++ {
		pages = append(pages, i+1)
	}
	return
}

func (*htmlService) Detail(id int64) (*model.ArticleItem, error) {
	article := model.ArticleItem{}
	return &article, dao.Article.GetArticleItem(&article, id)
}
