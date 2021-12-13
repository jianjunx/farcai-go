package model

type Category struct {
	Cid      int    `orm:"uid" json:"cid"`
	Name     string `orm:"name" json:"name"`
	CreateAt string `orm:"create_at" json:"createAt"`
	UpdateAt string `orm:"update_at" json:"-"`
}
