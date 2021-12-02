package model

var ArticleTableName = "TBL_BLOG_ARTICLE"

type ArticleItem struct {
	ArticleID  int64        `dynamo:"article_id" json:"articleId"`
	Title      string       `dynamo:"title" json:"title" v:"required"`
	Markdown   string       `dynamo:"markdown" json:"markdown" v:"required"`
	CategoryID int64        `dynamo:"category_id" json:"categoryId" v:"required"`
	Category   CategoryItem `dynamo:"category" json:"category"`
	UserName   string       `dynamo:"user_name" json:"userName"`
	ViewCount  int          `dynamo:"view_count" json:"viewCount"`
	CreateAt   string       `dynamo:"create_at" json:"createAt"`
	UpdateAt   string       `dynamo:"update_at" json:"updateAt"`
}

type ArticleItemSimpl struct {
	ArticleID  int64 `dynamo:"article_id" json:"articleId"`
	CategoryID int64 `dynamo:"category_id" json:"categoryId"`
}
