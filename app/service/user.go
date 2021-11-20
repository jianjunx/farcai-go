package service

import (
	"errors"
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"farcai-go/library/jwt"
	"farcai-go/library/snowflake"

	"github.com/gogf/gf/os/gtime"
)

var User = userService{}

type userService struct{}

func (*userService) Register(req *model.RegisterReq) error {
	users := []model.UserItem{}
	err := dao.User.GetUserItem(req.Name, &users)
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return errors.New("用户已存在")
	}
	return dao.User.PutUserItem(&model.UserItem{
		UserID:   snowflake.GenerateId(),
		UserName: req.Name,
		Passwd:   req.Passwd,
		CreateAt: gtime.Datetime(),
	})
}

func (*userService) Login(req *model.LoginReq) (interface{}, error) {
	users := []model.UserItem{}
	err := dao.User.GetUserItem(req.Name, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("用户不存在")
	}
	if users[0].Passwd != req.Passwd {
		return nil, errors.New("用户密码不正确")
	}
	token, err := jwt.Award(users[0].UserName)
	if err != nil {
		return nil, err
	}
	return &model.LoginResp{Token: token, UserInfo: users[0]}, nil
}

func (*userService) GetUserInfo(userName string) (*model.UserItem, error) {
	users := []model.UserItem{}
	err := dao.User.GetUserItem(userName, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("用户不存在")
	}
	return &users[0], nil
}
