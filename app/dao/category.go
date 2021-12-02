package dao

import (
	"farcai-go/app/model"
	"farcai-go/library/dynamodb"

	"github.com/guregu/dynamo"
)

var Category = categoryDao{}

type categoryDao struct{}

func (*categoryDao) GetCategorys(categorys *[]model.CategoryItem) error {
	return dynamodb.CategoryTable().Scan().All(categorys)
}

func (*categoryDao) GetCategoryItem(category *model.CategoryItem, categoryId int64) error {
	return dynamodb.CategoryTable().Get("category_id", categoryId).One(dynamo.AWSEncoding(category))
}

func (*categoryDao) PutCategoryItem(category *model.CategoryItem) error {
	return dynamodb.CategoryTable().Put(category).Run()
}
// TBL_BLOG_USER
// 087ec3bb94de26b7d2498f19b04fafee