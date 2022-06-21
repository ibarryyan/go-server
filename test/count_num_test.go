package test

import (
	"context"
	"count_num/pkg/cache"
	"count_num/pkg/config"
	"count_num/pkg/dao/impl"
	"count_num/pkg/entity"
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	info := entity.NumInfo{InfoNum: 22, InfoKey: "2", Name: "4"}
	daoImpl.AddNumInfo(context.Background(), info)
	fmt.Println(config.DB)
}

func TestFindByKey(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	key := daoImpl.GetNumInfoByKey(context.Background(), "123")
	fmt.Println(key)
}

func TestUpdate(t *testing.T) {
	daoImpl := impl.NewCountNumDAOImpl()
	daoImpl.UpdateNumInfoByKey(context.Background(), entity.NumInfo{Id: 1, InfoKey: "ymx", InfoNum: 22121223})
}

func TestDelete(t *testing.T) {
	impl.NewCountNumDAOImpl().DeleteNumInfoById(context.Background(), 1)
}

func TestCache(t *testing.T) {
	cacheDAOImpl := cache.NewCountNumCacheDAOImpl()
	cacheDAOImpl.SetNumInfo(context.Background(), "1", entity.NumInfo{1, "zs", "12", 2}, time.Second*1100)
	cacheDAOImpl.GetNumInfoById(context.Background(), "1")
}

func TestFindByPage(t *testing.T) {
	allNumInfo := impl.NewCountNumDAOImpl().FindAllNumInfo(context.Background(), 2, 10)
	fmt.Println(allNumInfo)
}
