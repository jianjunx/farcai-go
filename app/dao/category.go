package dao

import (
	"database/sql"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var Category = categoryDao{}

type categoryDao struct{}

func (*categoryDao) GetCategorys() (gdb.Result, error) {
	return categoryModel().All()
}

func (*categoryDao) GetCategoryItem(cid *int) (gdb.Record, error) {
	return categoryModel().WherePri(cid).One()
}

func (*categoryDao) AddCategory(categoryName *string) (sql.Result, error) {
	return categoryModel().Data(g.Map{"name": *categoryName}).Insert()
}
