package dao

import (
	"context"
	"count_num/pkg/model"
)

type RoleDao interface {
	//添加一个
	CreateRole(ctx context.Context, role model.Role) bool
	//根据ID查看
	GetRoleById(ctx context.Context, id int64) model.Role
}
