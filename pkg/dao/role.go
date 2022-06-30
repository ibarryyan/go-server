package dao

import (
	"context"
	"count_num/pkg/model"
)

type RoleDao interface {
	CreateRole(ctx context.Context, role model.Role) bool
	GetRoleById(ctx context.Context, id int64) model.Role
}
