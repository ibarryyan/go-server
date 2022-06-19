package cache

import (
	"context"
	"count_num/pkg/config"
	"count_num/pkg/entity"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type CountNumCacheDAOImpl struct {
	db *redis.Client
}

type CountNumCacheDAO interface {
	// set一个
	SetNumInfo(ctx context.Context, key string, info entity.NumInfo, t time.Duration) bool
	// 查看全部
	GetAllNumInfo(ctx context.Context) []entity.NumInfo
	// 根据ID获取一个
	GetNumInfoById(ctx context.Context, id int64) entity.NumInfo
}

func NewCountNumCacheDAOImpl() *CountNumCacheDAOImpl {
	return &CountNumCacheDAOImpl{db: config.RDB}
}

func (impl CountNumCacheDAOImpl) SetNumInfo(ctx context.Context, key string, info entity.NumInfo, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	fmt.Println(res)
	return true
}

func (impl CountNumCacheDAOImpl) GetAllNumInfo(ctx context.Context) []entity.NumInfo {
	return []entity.NumInfo{}
}

func (impl CountNumCacheDAOImpl) GetNumInfoById(ctx context.Context, id int64) entity.NumInfo {
	return entity.NumInfo{}
}
