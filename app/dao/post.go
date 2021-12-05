package dao

import (
	"database/sql"
	"farcai-go/app/model"
	"sync"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var Post = postDao{}

type postDao struct{}

// 分页列表
func (*postDao) GetPostPages(cid, page *int) (*gdb.Result, int, error) {
	var (
		result gdb.Result
		total  int
		err    error
		ws     sync.WaitGroup
	)
	ws.Add(2)
	go func() {
		defer ws.Done()
		mod := postModel()
		if *cid > 0 {
			mod = mod.Where("category_id", cid)
		}
		total, err = mod.Count()
	}()
	go func() {
		defer ws.Done()
		mod := postModel("p")
		if *cid > 0 {
			mod = mod.Where("category_id", cid)
		}
		result, err = mod.LeftJoin("tbl_blog_user u", "p.user_id=u.uid").LeftJoin("tbl_blog_category c", "p.category_id=c.cid").Fields("p.*,c.name as category_name,u.user_name as user_name").OrderDesc("pid").Offset(10 * (*page - 1)).Limit(10).All()
	}()
	ws.Wait()
	return &result, total, err
}

// 根据id获取文章
func (*postDao) GetPostItem(pid *int) (gdb.Record, error) {
	return postModel("p").LeftJoin("tbl_blog_user u", "p.user_id=u.uid").LeftJoin("tbl_blog_category c", "p.category_id=c.cid").Fields("p.*,c.name as category_name,u.user_name as user_name").WherePri(pid).One()
}

// 添加文章
func (*postDao) AddPostItem(row *g.Map) (sql.Result, error) {
	return postModel().Data(row).Insert()
}

func (*postDao) UpdatePostItem(post *model.PostReq) (sql.Result, error) {
	return postModel().Data(g.Map{
		"title":       post.Title,
		"content":     post.Content,
		"markdown":    post.Markdown,
		"category_id": post.CategoryId,
		"user_id":     post.UserId,
	}).WherePri(post.Pid).Update()
}

// 删除文章
func (*postDao) DeletePostItem(pid *int) (sql.Result, error) {
	return postModel().WherePri(pid).Delete()
}

// AddViewCount
func (*postDao) AddViewCount(pid *int) (sql.Result, error) {
	return postModel().WherePri(pid).Increment("view_count", 1)
}
