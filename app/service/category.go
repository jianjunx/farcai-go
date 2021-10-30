package service

import (
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"farcai-go/library/snowflake"

	"github.com/gogf/gf/os/gtime"
)

var Category = categoryService{}

type categoryService struct{}

func (*categoryService) AddCategory(categoryName string) error {
	return dao.Category.PutCategoryItem(&model.CategoryItem{
		CategoryID:   snowflake.GenerateId(),
		CategoryName: categoryName,
		CreateAt:     gtime.Datetime(),
	})
}

func (*categoryService) GetCategorys() (*[]model.CategoryItem, error) {
	categorys := []model.CategoryItem{}
	err := dao.Category.GetCategorys(&categorys)
	if err != nil {
		return nil, err
	}
	return &categorys, nil
}
