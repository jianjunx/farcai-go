package service

import (
	"database/sql"
	"farcai-go/app/dao"
	"farcai-go/app/model"

	"github.com/gogf/gf/frame/g"
)

var Post = &postService{}

type postService struct{}

func (*postService) AddPost(post *model.PostReq) (sql.Result, error) {
	return dao.Post.AddPostItem(&g.Map{
		"title":       post.Title,
		"content":     post.Content,
		"markdown":    post.Markdown,
		"category_id": post.CategoryId,
		"user_id":     post.UserId,
		"type":        post.Type,
		"slug":        post.Slug,
	})
}

func (*postService) UpdatePost(post *model.PostReq) error {
	_, err := dao.Post.UpdatePostItem(post)
	return err
}

func (*postService) DeletePost(pid *int) error {
	_, err := dao.Post.DeletePostItem(pid)
	return err
}

func (*postService) GetPostItem(pid *int) (*model.Post, error) {
	post := &model.Post{}
	record, err := dao.Post.GetPostItem(pid)
	if err != nil {
		return nil, err
	}
	err = record.Struct(post)
	return post, err
}

func (*postService) SearchPost(search *string) (*[]model.SearchResp, error) {
	posts := []model.SearchResp{}
	res, err := dao.Post.Search(search)
	if err != nil {
		return nil, err
	}
	if res.IsEmpty() {
		return &posts, nil
	}
	err = res.Structs(&posts)
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
