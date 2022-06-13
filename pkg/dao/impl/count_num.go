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
	return true
}

func (impl CountNumDAOImpl) GetNumInfoByKey(ctx context.Context, key string) entity.NumInfo {
	var info entity.NumInfo
	impl.db.First(&info, "info_key", key)
	return info
}

func (impl CountNumDAOImpl) FindAllNumInfo(ctx context.Context) []entity.NumInfo {
	var infos []entity.NumInfo
	impl.db.Find(&infos)
	return infos
}

func (impl CountNumDAOImpl) UpdateNumInfoByKey(ctx context.Context, info entity.NumInfo) bool {
	impl.db.Model(&entity.NumInfo{}).Where("info_key = ?", info.InfoKey).Update("info_num", info.InfoNum)
	return true
}

func (impl CountNumDAOImpl) DeleteNumInfoById(ctx context.Context, id int64) bool {
	impl.db.Delete(&entity.NumInfo{}, id)
	return true
}
