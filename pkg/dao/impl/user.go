package impl

import (
	"context"
	"count_num/pkg/cache"
	"count_num/pkg/config"
	"count_num/pkg/model"
	"count_num/pkg/utils"
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
	var u model.User
	impl.db.First(&u, "login_name", user.LoginName)
	if u.LoginName == user.LoginName {
		return false
	}
	user.Pwd = utils.GetMd5Str(user.Pwd)
	impl.db.Save(&user)
	return true
}

func (impl *UserDaoImpl) GetUserByUid(ctx context.Context, uId int64) model.User {
	var user model.User
	impl.db.First(&user, "id", uId)
	return user
}

func (impl *UserDaoImpl) GetAll(ctx context.Context, page int, limit int) []model.User {
	users := make([]model.User, 0)
	if page <= 0 || limit <= 0 {
		impl.db.Find(&users)
	} else {
		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&users)
	}
	return users
}

func (impl *UserDaoImpl) UpdateUserById(ctx context.Context, user model.User) bool {
	impl.db.Model(&model.User{}).Where("id = ?", user.Id).Updates(user)
	return true
}

func (impl *UserDaoImpl) GetUserByLoginName(ctx context.Context, loginName string) model.User {
	var user model.User
	impl.db.First(&user, "login_name", loginName)
	return user
}
