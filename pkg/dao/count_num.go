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
	//删除一个
	DeleteNumInfoById(ctx context.Context, id int64) bool
	//根据ID获取一个
	GetNumInfoById(ctx context.Context, id int64) entity.NumInfo
	//根据ID修改
	UpdateNumInfoById(ctx context.Context, info entity.NumInfo) bool
}
