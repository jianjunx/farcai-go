package dao

import (
	"database/sql"
	"farcai-go/app/model"
	"farcai-go/library/snowflake"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var User = userDao{}

type userDao struct{}

func (*userDao) GetUserItem(uid *int, user *model.User) error {
	record, err := userModel().Fields("uid", "user_name", "passwd", "avatar", "create_at").WherePri(uid).One()
	if err != nil {
		return err
	}
	return record.Struct(user)
}
func (*userDao) UserLogin(param *model.LoginReq) (gdb.Record, error) {
	return userModel().Fields("uid", "user_name", "avatar", "create_at").Where(g.Map{
		"user_name": param.Name,
		"passwd":    param.Passwd,
	}).One()
}

func (*userDao) UserCount(where *g.Map) (int, error) {
	return userModel().Where(where).Count()
}

func (*userDao) PutUserItem(user *model.User) (sql.Result, error) {
	return userModel().Data(g.Map{
		"user_name": user.UserName,
		"passwd":    user.Passwd,
	}).WherePri(user.Uid).Update()
}

func (*userDao) AddUserItem(user *model.User) (sql.Result, error) {
	return userModel().Data(g.Map{
		"uid":       snowflake.GenerateId(),
		"user_name": user.UserName,
		"passwd":    user.Passwd,
	}).Insert()
}

func (*userDao) GetUserAll() (gdb.Result, error) {
	return userModel().All()
}
