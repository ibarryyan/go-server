package cache

import (
	"context"
	"count_num/pkg/config"
	"count_num/pkg/model"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

type CountNumCacheDAOImpl struct {
	db *redis.Client
}

type CountNumCacheDAO interface {
	// set一个
	SetNumInfo(ctx context.Context, key string, info model.NumInfo, t time.Duration) bool
	// 根据ID获取一个
	GetNumInfoById(ctx context.Context, key string) model.NumInfo
}

func NewCountNumCacheDAOImpl() *CountNumCacheDAOImpl {
	return &CountNumCacheDAOImpl{db: config.RDB}
}

func (impl CountNumCacheDAOImpl) SetNumInfo(ctx context.Context, key string, info model.NumInfo, t time.Duration) bool {
	res := impl.db.Set(ctx, key, info, t)
	result, _ := res.Result()
	if result != "OK" {
		return false
	}
	return true
}

func (impl CountNumCacheDAOImpl) GetNumInfoById(ctx context.Context, key string) model.NumInfo {
	res := impl.db.Get(ctx, key)
	var info model.NumInfo
	j := res.Val()
	json.Unmarshal([]byte(j), &info)
	return info
}
