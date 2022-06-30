package dao

import (
	"context"
	"count_num/pkg/model"
)

type UserDao interface {
	CreateUser(ctx context.Context, user model.User) bool
}
