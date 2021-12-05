package service

import (
	"errors"
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"math"
	"sync"

	"github.com/gogf/gf/database/gdb"
)

var Html = &htmlService{}

type htmlService struct{}

func (*htmlService) Home(ctgId, page *int) (*[]model.PostMore, *[]model.Category, int, error) {
	var (
		ws    sync.WaitGroup
		total int
		err   error
	)
	posts := []model.PostMore{}
	categorys := []model.Category{}

	ws.Add(2)
	go func() {
		defer ws.Done()
		var result *gdb.Result
		result, total, err = dao.Post.GetPostPages(ctgId, page)
		result.Structs(&posts)
	}()
	go func() {
		defer ws.Done()
		var result gdb.Result
		result, err = dao.Category.GetCategorys()
		result.Structs(&categorys)
	}()
	ws.Wait()
	return &posts, &categorys, total, err
}

func (*htmlService) GetPages(total float64) (pages []float64) {
	var size float64 = 10
	pages = []float64{}
	for i, len := 0.0, math.Ceil(total/size); i < len; i++ {
		pages = append(pages, i+1)
	}
	return
}

func (*htmlService) Detail(id *int) (*model.PostMore, error) {
	post := model.PostMore{}
	record, err := dao.Post.GetPostItem(id)
	if err != nil {
		return nil, err
	}
	if record.IsEmpty() {
		return nil, errors.New("没有找到该文章")
	}
	err = record.Struct(&post)
	return &post, err
}

func (*htmlService) AddViewCount(pid *int) {
	dao.Post.AddViewCount(pid)
}
