package model

type User struct {
	Uid      int    `orm:"uid" json:"uid"`            // 用户id
	UserName string `orm:"user_name" json:"userName"` // 用户名
	Avatar   string `orm:"avatar" json:"avatar"`      // 头像
	Passwd   string `orm:"passwd" json:"-"`           // 密码
	CreateAt string `orm:"create_at" json:"createAt"` // 创建时间
}

type RegisterReq struct {
	Name     string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Passwd   string `p:"passwd" v:"required|length:6,255#请输入密码|密码长度不够"`
	Repasswd string `p:"repasswd" v:"required|length:6,255|same:passwd#请确认密码|密码长度不够|两次密码不一致"`
}
type LoginReq struct {
	Name   string `p:"username"  v:"required#请输入账号"`
	Passwd string `p:"passwd" v:"required#请输入密码"`
}

type LoginResp struct {
	Token    string `json:"token"`    // token
	UserInfo User   `json:"userInfo"` // 用户信息
}
