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

func addWhere(mod *gdb.Model, cid *int) *gdb.Model {
	if *cid > 0 {
		return mod.Where("type", 0).Where("category_id", *cid)
	}
	return mod.Where("type", 0)
}

// 分页列表
func (*postDao) GetPostPages(cid, page *int) (*gdb.Result, int, error) {
	var (
		result gdb.Result
		total  int
		err    error
		ws     sync.WaitGroup
	)
	ws.Add(2)
	// 查总数
	go func() {
		defer ws.Done()
		total, err = addWhere(postModel(), cid).Count()
	}()
	// 查列表数据
	go func() {
		defer ws.Done()
		mod := addWhere(postModel("p"), cid)
		field := "p.pid,p.title,p.content,p.view_count,p.create_at,p.user_id,p.category_id,c.name as category_name,u.user_name as user_name"
		result, err = mod.LeftJoin("tbl_blog_user u", "p.user_id=u.uid").LeftJoin("tbl_blog_category c", "p.category_id=c.cid").Fields(field).OrderDesc("pid").Offset(10 * (*page - 1)).Limit(10).All()
	}()
	ws.Wait() // 等待所有协程执行完
	return &result, total, err
}

func (*postDao) GetPostAll() (gdb.Result, error) {
	return postModel().Fields("pid", "title", "create_at").Where("type=0").OrderDesc("pid").All()
}

// 根据id获取文章
func (*postDao) GetPostItem(pid *int) (gdb.Record, error) {
	field := "p.*,c.name as category_name,u.user_name as user_name"
	return postModel("p").LeftJoin("tbl_blog_user u", "p.user_id=u.uid").LeftJoin("tbl_blog_category c", "p.category_id=c.cid").Fields(field).WherePri(pid).One()
}

// 获取自定义文章
func (*postDao) GetCustomPost(slug *string) (gdb.Record, error) {
	return postModel().Fields("pid,title,content,view_count,create_at").Where("slug", slug).One()
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
		"type":        post.Type,
		"slug":        post.Slug,
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

// 搜索结果
func (*postDao) Search(search *string) (gdb.Result, error) {
	return postModel().Where("type=? AND title like ?", g.Slice{0, "%" + *search + "%"}).All()
}

// 备份文章
func (*postDao) GetAll() (gdb.Result, error) {
	return postModel().All()
}
