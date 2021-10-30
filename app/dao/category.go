package dao

import (
	"farcai-go/app/model"
	"farcai-go/library/dynamodb"
)

var Category = categoryDao{}

type categoryDao struct{}

func (*categoryDao) GetCategorys(categorys *[]model.CategoryItem) error {
	return dynamodb.CategoryTable().Scan().All(categorys)
}

func (*categoryDao) PutCategoryItem(category *model.CategoryItem) error {
	return dynamodb.CategoryTable().Put(category).Run()
}
