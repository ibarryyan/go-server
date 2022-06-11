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
	info := entity.NumInfo{Num: 22, Key: "123", Name: "zs",Id: 2}
	daoImpl.AddNumInfo(context.Background(),info)
	fmt.Println(config.DB)
}

func TestFindByKey(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	daoImpl.GetNumInfoByKey(context.Background(),"ymx")
}

func TestUpdate(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	daoImpl.UpdateNumInfoByKey(context.Background(),entity.NumInfo{Key: "ymx",Num: 2223})
}


