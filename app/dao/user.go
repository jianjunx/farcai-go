package dao

import (
	"farcai-go/app/model"
	"farcai-go/library/dynamodb"
)

var User = userDao{}

type userDao struct{}

func (*userDao) GetUserItem(username string, user *[]model.UserItem) error {
	userTable := dynamodb.UserTable()
	return userTable.Scan().Filter("'user_name' = ?", username).All(user)
}

func (*userDao) PutUserItem(user *model.UserItem) error {
	return dynamodb.UserTable().Put(user).Run()
}
