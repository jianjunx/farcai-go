package service

import (
	"errors"
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"farcai-go/library/jwt"
	"farcai-go/library/snowflake"

	"github.com/gogf/gf/frame/g"
)

var User = userService{}

type userService struct{}

func (*userService) Register(req *model.RegisterReq) error {
	user := model.User{}
	count, err := dao.User.UserCount(&g.Map{"user_name": user.UserName})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	_, err = dao.User.AddUserItem(&model.User{
		Uid:      int(snowflake.GenerateId()),
		UserName: req.Name,
		Passwd:   req.Passwd,
	})
	return err
}

func (*userService) Login(req *model.LoginReq, audience *string) (interface{}, error) {
	user := model.User{}
	record, err := dao.User.UserLogin(req)
	if err != nil {
		return nil, err
	}
	if record.IsEmpty() {
		return nil, errors.New("用户名或密码不正确")
	}
	if err = record.Struct(&user); err != nil {
		return nil, errors.New("用户名或密码不正确")
	}
	token, err := jwt.Award(&user.Uid, audience)
	if err != nil {
		return nil, err
	}
	return &model.LoginResp{Token: token, UserInfo: user}, nil
}

func (*userService) GetUserInfo(uid *int) (*model.User, error) {
	user := model.User{}
	err := dao.User.GetUserItem(uid, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
