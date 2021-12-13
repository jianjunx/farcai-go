package utils

import (
	"farcai-go/app/model"
	"fmt"
	"os"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

// 备份文章
func BackupPost(list *[]model.Post) (*string, *string, error) {
	var (
		path, name = getBackupPath("tbl_blog_post")
		listLen    = len(*list)
	)
	file, err := os.Create(*path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	// 字符串切片
	strs := make([]string, listLen+2)
	strs[0] = "BEGIN;\n"
	strs[listLen+1] = "COMMIT;\n"

	for i, item := range *list {
		sql := fmt.Sprintf("INSERT INTO `tbl_blog_post` (pid,title,content,markdown,category_id,user_id,slug,type,view_count,create_at,update_at) VALUES (%d,'%s','%s','%s',%d,%d,'%s',%d,%d,'%s','%s');\n", item.Pid, gstr.AddSlashes(item.Title), gstr.AddSlashes(item.Content), gstr.AddSlashes(item.Markdown), item.CategoryId, item.UserId, item.Slug, item.Type, item.ViewCount, item.CreateAt, item.UpdateAt)
		strs[i+1] = sql
	}
	text := strings.Join(strs, "")
	// 写入
	_, err = file.WriteString(text)
	if err != nil {
		return nil, nil, err
	}
	return name, &text, nil
}

// 备份分类
func BackupCategory(list *[]model.Category) (*string, *string, error) {
	var (
		path, name = getBackupPath("tbl_blog_category")
		listLen    = len(*list)
	)
	file, err := os.Create(*path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	// 创建字符串切片
	strs := make([]string, listLen+2)
	strs[0] = "BEGIN;\n"
	strs[listLen+1] = "COMMIT;\n"
	for i, item := range *list {
		sql := fmt.Sprintf("INSERT INTO `tbl_blog_category` (cid,name,create_at,update_at) VALUES (%d,'%s','%s','%s');\n", item.Cid, item.Name, item.CreateAt, item.UpdateAt)
		strs[i+1] = sql
	}
	text := strings.Join(strs, "")
	_, err = file.WriteString(text)
	if err != nil {
		return nil, nil, err
	}
	return name, &text, nil
}

// 备份用户
func BackupUser(list *[]model.User) (*string, *string, error) {
	var (
		path, name = getBackupPath("tbl_blog_user")
		listLen    = len(*list)
	)
	file, err := os.Create(*path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	// 创建字符串切片
	strs := make([]string, listLen+2)
	strs[0] = "BEGIN;\n"
	strs[listLen+1] = "COMMIT;\n"
	for i, item := range *list {
		sql := fmt.Sprintf("INSERT INTO `tbl_blog_user` (uid,user_name,passwd,avatar,create_at,update_at) VALUES (%d,'%s','%s','%s','%s','%s');\n", item.Uid, item.UserName, item.Passwd, item.Avatar, item.CreateAt, item.UpdateAt)
		strs[i+1] = sql
	}
	text := strings.Join(strs, "")
	_, err = file.WriteString(text)
	if err != nil {
		return nil, nil, err
	}
	return name, &text, nil
}

// 备份路径
func getBackupPath(tblName string) (*string, *string) {
	name := fmt.Sprintf("%s_%s.sql", tblName, gtime.Now().Format("Y-m-d_His"))
	path := fmt.Sprintf("%s/%s", g.Cfg().GetString("database.backupPath"), name)
	return &path, &name
}
