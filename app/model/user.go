package model

var UserTableName string = "TBL_BLOG_USER"

type UserAction struct {
	UserName string `dynamo:"user_name,hash" index:"Seq-ID-index,range"`
}

type UserItem struct {
	UserID   int64  `dynamo:"user_id" json:"userId"`
	UserName string `dynamo:"user_name" json:"userName"`
	Avatar   string `dynamp:"avatar" json:"avatar"`
	Passwd   string `dynamo:"passwd" json:"-"`
	CreateAt string `dynamo:"create_at" json:"-"`
}

type RegisterReq struct {
	Name     string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Passwd   string `p:"passwd" v:"required|length:6,30#请输入密码|密码长度不够"`
	Repasswd string `p:"repasswd" v:"required|length:6,30|same:passwd#请确认密码|密码长度不够|两次密码不一致"`
}
type LoginReq struct {
	Name   string `p:"username"  v:"required#请输入账号"`
	Passwd string `p:"passwd" v:"required#请输入密码"`
}

type LoginResp struct {
	Token    string   `json:"token"`
	UserInfo UserItem `json:"userInfo"`
}
