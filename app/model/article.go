package model

var ArticleTableName = "TBL_BLOG_ARTICLE"

type ArticleItem struct {
	ArticleID  int64        `dynamo:"article_id" json:"articleId"`
	Title      string       `dynamo:"title" json:"title"`
	Content    string       `dynamo:"Content" json:"content"`
	CategoryID int64        `dynamo:"category_id" json:"categoryId"`
	Category   CategoryItem `dynamo:"category" json:"category"`
	UserID     int64        `dynamo:"user_id" json:"userId"`
	User       UserItem     `dynamo:"user" json:"user"`
	Publish    int          `dynamo:"publish" json:"publish"` // 0 草稿 1 发布
	Deleted    int          `dynamo:"deleted" json:"deleted"` // 0 未删除 1 已删除
	CreateAt   string       `dynamo:"create_at" json:"createAt"`
	UpdateAt   string       `dynamo:"update_at" json:"updateAt"`
}
