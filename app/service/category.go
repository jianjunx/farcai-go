package service

import (
	"database/sql"
	"farcai-go/app/dao"
	"farcai-go/app/model"
)

var Category = categoryService{}

type categoryService struct{}

func (*categoryService) AddCategory(categoryName *string) (sql.Result, error) {
	return dao.Category.AddCategory(categoryName)
}

func (*categoryService) GetCategorys() (*[]model.Category, error) {
	categorys := []model.Category{}
	result, err := dao.Category.GetCategorys()
	if err != nil {
		return nil, err
	}
	err = result.Structs(&categorys)
	if err != nil {
		return nil, err
	}
	return &categorys, nil
}
