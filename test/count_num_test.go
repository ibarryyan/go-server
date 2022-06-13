package test

import (
	"context"
	"count_num/pkg/config"
	"count_num/pkg/dao/impl"
	"count_num/pkg/entity"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	info := entity.NumInfo{InfoNum: 22, InfoKey: "123", Name: "123", Id: 3}
	daoImpl.AddNumInfo(context.Background(), info)
	fmt.Println(config.DB)
}

func TestFindByKey(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	key := daoImpl.GetNumInfoByKey(context.Background(), "ymx")
	fmt.Println(key)
}

func TestUpdate(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	daoImpl.UpdateNumInfoByKey(context.Background(), entity.NumInfo{Id: 1, InfoKey: "ymx", InfoNum: 22121223})
}

func TestDelete(t *testing.T) {
	impl.NewCountNumDAOImpl().DeleteNumInfoById(context.Background(), 1)
}
