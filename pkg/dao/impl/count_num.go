package impl

import (
	"context"
	"count_num/pkg/cache"
	"count_num/pkg/config"
	"count_num/pkg/entity"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var cacheTime = time.Second * 3600

type CountNumDAOImpl struct {
	db    *gorm.DB
	cache *cache.CountNumCacheDAOImpl
}

func NewCountNumDAOImpl() *CountNumDAOImpl {
	return &CountNumDAOImpl{db: config.DB, cache: cache.NewCountNumCacheDAOImpl()}
}

func (impl CountNumDAOImpl) AddNumInfo(ctx context.Context, info entity.NumInfo) bool {
	var in entity.NumInfo
	impl.db.First(&in, "info_key", info.InfoKey)
	if in.InfoKey == info.InfoKey { //去重
		return false
	}
	impl.db.Save(&info) //要使用指针,Id可以回显
	fmt.Println(info)
	impl.cache.SetNumInfo(ctx, string(info.Id), info, cacheTime)
	return true
}

func (impl CountNumDAOImpl) GetNumInfoByKey(ctx context.Context, key string) entity.NumInfo {
	var info entity.NumInfo
	impl.db.First(&info, "info_key", key)
	return info
}

func (impl CountNumDAOImpl) FindAllNumInfo(ctx context.Context, page int, limit int) []entity.NumInfo {
	var infos []entity.NumInfo
	if page == 0 && limit == 0 {
		impl.db.Find(&infos)
	} else {
		impl.db.Limit(limit).Offset((page - 1) * limit).Find(&infos)
	}
	return infos
}

func (impl CountNumDAOImpl) UpdateNumInfoByKey(ctx context.Context, info entity.NumInfo) bool {
	impl.db.Model(&entity.NumInfo{}).Where("info_key = ?", info.InfoKey).Update("info_num", info.InfoNum)
	return true
}

func (impl CountNumDAOImpl) DeleteNumInfoById(ctx context.Context, id int64) bool {
	impl.db.Delete(&entity.NumInfo{}, id)
	impl.cache.SetNumInfo(ctx, string(id), entity.NumInfo{}, cacheTime)
	return true
}

func (impl CountNumDAOImpl) GetNumInfoById(ctx context.Context, id int64) entity.NumInfo {
	var info entity.NumInfo
	numInfoById := impl.cache.GetNumInfoById(ctx, string(id))
	if numInfoById.InfoKey != "" {
		return numInfoById
	}
	impl.db.First(&info, "id", id)
	return info
}

func (impl CountNumDAOImpl) UpdateNumInfoById(ctx context.Context, info entity.NumInfo) bool {
	impl.db.Model(&entity.NumInfo{}).Where("id", info.Id).Updates(entity.NumInfo{Name: info.Name, InfoKey: info.InfoKey, InfoNum: info.InfoNum})
	impl.cache.SetNumInfo(ctx, string(info.Id), info, cacheTime)
	return true
}
