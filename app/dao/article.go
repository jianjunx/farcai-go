package dao

import (
	"farcai-go/app/model"
	"farcai-go/library/dynamodb"
	"sync"

	"github.com/guregu/dynamo"
)

var Article = articleDao{}

type articleDao struct{}

// 分页列表
func (*articleDao) GetArticlePages(articles *[]model.ArticleItem, categoryId, page int64) (total int64, err error) {
	table := dynamodb.ArticleTable()
	scan := table.Scan()
	if categoryId != 0 {
		scan = table.Scan().Filter("'category_id' = ?", categoryId)
	}
	var ws sync.WaitGroup
	ws.Add(2)
	// 在协程中获取列表
	go func() {
		defer ws.Done()
		var size int64 = 10
		if page == 1 {
			scan.Limit(size).All(articles)
			return
		}
		var lastKey dynamo.PagingKey
		lastKey, err = scan.Limit((page - 1) * size).AllWithLastEvaluatedKey(&[]model.ArticleItemSimpl{})
		if err != nil {
			return
		}
		scan.StartFrom(lastKey).Limit(size).All(articles)
	}()
	// 获取总数
	go func() {
		defer ws.Done()
		total, err = scan.Count()
	}()
	ws.Wait()
	return total, err
}

// 根据id获取文章
func (*articleDao) GetArticleItem(articleItem *[]model.ArticleItem, articleId int64) error {
	table := dynamodb.ArticleTable()
	return table.Scan().Filter("'article_id' = ?", articleId).All(articleItem)
}

// 修改添加文章
func (*articleDao) PutArticleItem(row *model.ArticleItem) error {
	return dynamodb.ArticleTable().Put(row).Run()
}

// 删除文章
func (*articleDao) DeleteArticleItem(articleId int64) error {
	return dynamodb.ArticleTable().Delete("article_id", articleId).Run()
}
