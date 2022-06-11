package impl

import (
	"context"
	"count_num/pkg/config"
	"count_num/pkg/entity"
	"gorm.io/gorm"
)

type CountNumDAOImpl struct {
	db *gorm.DB
}

func NewCountNumDAOImpl() *CountNumDAOImpl {
	return &CountNumDAOImpl{db: config.DB}
}

func (impl CountNumDAOImpl) AddNumInfo(ctx context.Context, info entity.NumInfo) bool {
	impl.db.Save(info)
	return false
}

func (impl CountNumDAOImpl) GetNumInfoByKey(ctx context.Context, key string) entity.NumInfo {
	var info entity.NumInfo
	impl.db.First(&info, "key", key)
	return info
}

func (impl CountNumDAOImpl) FindAllNumInfo(ctx context.Context) []entity.NumInfo {
	var infos []entity.NumInfo
	impl.db.Find(&infos)
	return infos
}

func (impl CountNumDAOImpl) UpdateNumInfoByKey(ctx context.Context, info entity.NumInfo) bool {
	impl.db.Model(&entity.NumInfo{}).Where("key = ?", info.Key).Update("num", info.Num)
	return false
}
