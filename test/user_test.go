package test

import (
	"context"
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	daoImpl := impl.NewUserDaoImpl()

	daoImpl.CreateUser(context.TODO(), model.User{Name: "ls", LoginName: "ls", Role: 1, Pwd: "123", CreateTime: time.Time{}})

}
