package dao

import (
	"context"
	"count_num/pkg/entity"
)

type CountNumDAO interface {
	//添加一个
	AddNumInfo(ctx context.Context, info entity.NumInfo) bool
	//根据Key获取一个
	GetNumInfoByKey(ctx context.Context, url string) entity.NumInfo
	//查看全部
	FindAllNumInfo(ctx context.Context) []entity.NumInfo
	//根据Key修改
	UpdateNumInfoByKey(ctx context.Context, info entity.NumInfo) bool
}
