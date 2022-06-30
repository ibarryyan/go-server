package impl

import (
	"context"
	"count_num/pkg/cache"
	"count_num/pkg/config"
	"count_num/pkg/model"
	"gorm.io/gorm"
)

type UserDaoImpl struct {
	db    *gorm.DB
	cache *cache.CountNumCacheDAOImpl
}

func NewUserDaoImpl() *UserDaoImpl {
	return &UserDaoImpl{db: config.DB, cache: cache.NewCountNumCacheDAOImpl()}
}

func (impl *UserDaoImpl) CreateUser(ctx context.Context, user model.User) bool {
	return true
}

func (impl *UserDaoImpl) GetUserByUid(ctx context.Context, uId int64) model.User {
	var user model.User

	return user
}

func (impl *UserDaoImpl) GetAll(ctx context.Context, page int, limit int) []model.User {
	users := make([]model.User, 0)

	return users
}

func (impl *UserDaoImpl) UpdateUserById(ctx context.Context, user model.User) bool {

	return true
}

func (impl *UserDaoImpl) GetUserByLoginName(ctx context.Context, loginName string) model.User {
	var user model.User
	return user
}
