package model

var CategoryTableName string = "TBL_BLOG_CATEGORY"

type CategoryItem struct {
	CategoryID   int64  `dynamo:"category_id" json:"categoryId"`
	CategoryName string `dynamo:"category_name" json:"categoryName" p:"categoryName"  v:"required#请输入分类名"`
	CreateAt     string `dynamo:"create_at" json:"-"`
}

type Category struct {
	Cid      int    `orm:"uid" json:"cid"`
	Name     string `orm:"name" json:"name"`
	CreateAt string `orm:"create_at" json:"createAt"`
}
