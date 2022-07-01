package impl

import (
	"context"
	"count_num/pkg/cache"
	"count_num/pkg/config"
	"count_num/pkg/model"
	"gorm.io/gorm"
)

type RoleDaoImpl struct {
	db    *gorm.DB
	cache *cache.CountNumCacheDAOImpl
}

func NewRoleDaoImpl() *RoleDaoImpl {
	return &RoleDaoImpl{db: config.DB, cache: cache.NewCountNumCacheDAOImpl()}
}

func (impl *RoleDaoImpl) CreateRole(ctx context.Context, role model.Role) bool {
	var r model.Role
	impl.db.First(&r, "name", role.Name)
	if r.Name == role.Name {
		return false
	}
	impl.db.Save(&role)
	return true
}

func (impl *RoleDaoImpl) GetRoleById(ctx context.Context, id int64) model.Role {
	var role model.Role
	impl.db.First(&role, "id", id)
	return role
}
