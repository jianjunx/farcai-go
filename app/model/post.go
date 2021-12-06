package model

type Post struct {
	Pid        int    `orm:"pid" json:"pid"`
	Title      string `orm:"title" json:"title"`
	Slug       string `orm:"slug" json:"slug"`
	Content    string `orm:"content" json:"content"`
	Markdown   string `orm:"markdown" json:"markdown"`
	CategoryId int    `orm:"category_id" json:"categoryId"`
	UserId     int    `orm:"user_id" json:"userId"`
	ViewCount  int    `orm:"view_count" json:"viewCount"`
	Type       int    `orm:"type" json:"type"`
	CreateAt   string `orm:"create_at" json:"createAt"`
	UpdateAt   string `orm:"update_at" json:"updateAt"`
}

type PostMore struct {
	Pid          int    `orm:"pid" json:"pid"`
	Title        string `orm:"title" json:"title"`
	Slug         string `orm:"slug" json:"slug"`
	Content      string `orm:"content" json:"content"`
	CategoryId   int    `orm:"category_id" json:"categoryId"`
	CategoryName string `orm:"category_name" json:"categoryName"`
	UserId       int    `orm:"user_id" json:"userId"`
	UserName     string `orm:"user_name" json:"userName"`
	ViewCount    int    `orm:"view_count" json:"viewCount"`
	Type         int    `orm:"type" json:"type"`
	CreateAt     string `orm:"create_at" json:"createAt"`
	UpdateAt     string `orm:"update_at" json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}
