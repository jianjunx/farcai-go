package dao

import (
	"strings"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

func userModel() *gdb.Model {
	return g.DB().Model("tbl_blog_user")
}

// strs 别名
func postModel(strs ...string) *gdb.Model {
	names := []string{"tbl_blog_post"}
	// 是否传了别名
	if len(strs) > 0 {
		names = append(names, strs...)
	}
	return g.DB().Model(strings.Join(names, " "))
}

func categoryModel() *gdb.Model {
	return g.DB().Model("tbl_blog_category")
}
