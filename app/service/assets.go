package service

import (
	"farcai-go/app/dao"
	"farcai-go/app/model"
	"farcai-go/library/cos"
	"farcai-go/library/utils"

	"github.com/gogf/gf/frame/g"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

var Assets = &assetsService{}

type assetsService struct{}

func (*assetsService) COSCredentials() (*sts.CredentialResult, error) {
	return cos.GetClient().GetCredential(cos.Option)
}

// 备份文章数据
func (*assetsService) BackupPost() {
	posts := []model.Post{}
	res, err := dao.Post.GetAll()
	if err != nil {
		g.Log().Error(err) // 输出log
		return
	}
	err = res.Structs(&posts)
	if err != nil {
		g.Log().Error(err) // 输出log
		return
	}
	_, err = utils.BackupPost(&posts)
	if err != nil {
		g.Log().Error(err) // 输出log
	}
}

// 备份分类数据
func (*assetsService) BackupCategory() {
	ctgs := []model.Category{}
	res, err := dao.Category.GetCategorys()
	if err != nil {
		g.Log().Error(err) // 输出log
		return
	}
	err = res.Structs(&ctgs)
	if err != nil {
		g.Log().Error(err) // 输出log
		return
	}
	_, err = utils.BackupCategory(&ctgs)
	if err != nil {
		g.Log().Error(err) // 输出log
	}
}

// 备份用户数据
func (*assetsService) BackupUser() {
	users := []model.User{}
	res, err := dao.User.GetUserAll()
	if err != nil {
		g.Log().Error(err) // 输出log
		return
	}
	err = res.Structs(&users)
	if err != nil {
		g.Log().Error(err) // 输出log
		return
	}
	_, err = utils.BackupUser(&users)
	if err != nil {
		g.Log().Error(err) // 输出log
	}
}
