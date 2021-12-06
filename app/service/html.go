package service

import (
	"errors"
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"math"
	"sync"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
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

func (*htmlService) Pigeonhole() (*map[string][]model.Post, *[]model.Category, error) {
	var (
		ws        sync.WaitGroup
		lines     = make(map[string][]model.Post)
		posts     []model.Post
		categorys []model.Category
		err       error
	)
	ws.Add(2)
	go func() {
		defer ws.Done()
		var record gdb.Result
		record, err = dao.Post.GetPostAll()
		if err != nil {
			return
		}
		if err = record.Structs(&posts); err != nil {
			return
		}
		for _, value := range posts {
			// 创建日期格式化成月
			key := gtime.NewFromStr(value.CreateAt).Format("Y-m")
			// 追加切片
			lines[key] = append(lines[key], value)
		}
	}()
	go func() {
		defer ws.Done()
		var record gdb.Result
		record, err = dao.Category.GetCategorys()
		err = record.Structs(&categorys)
	}()
	ws.Wait()
	return &lines, &categorys, err
}

func (*htmlService) Custom(slug *string) (bool, *model.Post, error) {
	var post model.Post
	record, err := dao.Post.GetCustomPost(slug)
	if err != nil {
		return true, nil, err
	}
	if record.IsEmpty() {
		return true, nil, nil
	}
	err = record.Struct(&post)
	if err != nil {
		return true, nil, err
	}
	return false, &post, nil
}
