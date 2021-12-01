package dynamodb

import (
	"farcai-go/app/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gogf/gf/frame/g"
	"github.com/guregu/dynamo"
)

var db *dynamo.DB

// 用户表
var userTable dynamo.Table

// 分类表
var categoryTable dynamo.Table

// 文章表
var articleTable dynamo.Table

func init() {
	sess := session.Must(session.NewSession())
	db = dynamo.New(sess, &aws.Config{Region: aws.String(g.Cfg().GetString("dynamodb.region"))})
	// 创建表
	initTable()
}

func initTable() {
	// user 表
	userTable = db.Table(model.UserTableName)
	// 分类表
	categoryTable = db.Table(model.CategoryTableName)
	// 文章表
	articleTable = db.Table(model.ArticleTableName)
}

func UserTable() *dynamo.Table {
	return &userTable
}

func CategoryTable() *dynamo.Table {
	return &categoryTable
}

func ArticleTable() *dynamo.Table {
	return &articleTable
}
