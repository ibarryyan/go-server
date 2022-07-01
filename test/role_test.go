package test

import (
	"context"
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"fmt"
	"testing"
)

func TestRole(t *testing.T) {
	daoImpl := impl.NewRoleDaoImpl()
	daoImpl.CreateRole(context.TODO(), model.Role{Name: "管理员"})

	role := daoImpl.GetRoleById(context.Background(), 1)
	fmt.Println(role)
}
