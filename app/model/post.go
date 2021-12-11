package model

type Post struct {
	Pid        int    `orm:"pid" json:"pid"`                // 文章ID
	Title      string `orm:"title" json:"title"`            // 文章ID
	Slug       string `orm:"slug" json:"slug"`              // 自定也页面 path
	Content    string `orm:"content" json:"content"`        // 文章的html
	Markdown   string `orm:"markdown" json:"markdown"`      // 文章的Markdown
	CategoryId int    `orm:"category_id" json:"categoryId"` //分类id
	UserId     int    `orm:"user_id" json:"userId"`         //用户id
	ViewCount  int    `orm:"view_count" json:"viewCount"`   //查看次数
	Type       int    `orm:"type" json:"type"`              //文章类型 0 普通，1 自定义文章
	CreateAt   string `orm:"create_at" json:"createAt"`     // 创建时间
	UpdateAt   string `orm:"update_at" json:"updateAt"`     // 更新时间
}

type PostMore struct {
	Pid          int    `orm:"pid" json:"pid"`                    // 文章ID
	Title        string `orm:"title" json:"title"`                // 文章ID
	Slug         string `orm:"slug" json:"slug"`                  // 自定也页面 path
	Content      string `orm:"content" json:"content"`            // 文章的html
	CategoryId   int    `orm:"category_id" json:"categoryId"`     // 文章的Markdown
	CategoryName string `orm:"category_name" json:"categoryName"` // 分类名
	UserId       int    `orm:"user_id" json:"userId"`             // 用户id
	UserName     string `orm:"user_name" json:"userName"`         // 用户名
	ViewCount    int    `orm:"view_count" json:"viewCount"`       // 查看次数
	Type         int    `orm:"type" json:"type"`                  // 文章类型 0 普通，1 自定义文章
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

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` // 文章ID
	Title string `orm:"title" json:"title"`
}
