package dao

import (
	"context"
	"count_num/pkg/model"
)

type UserDao interface {
	// 添加一个
	CreateUser(ctx context.Context, user model.User) bool
	// 根据ID查找一个
	GetUserByUid(ctx context.Context, uId int64) model.User
	// 查找全部
	GetAll(ctx context.Context, page int, limit int) []model.User
	// 根据ID修改一个
	UpdateUserById(ctx context.Context, user model.User) bool
	// 根据登录名查找一个
	GetUserByLoginName(ctx context.Context, loginName string) model.User
}
