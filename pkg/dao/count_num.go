package dao

import (
	"context"
	"count_num/pkg/model"
)

type CountNumDAO interface {
	//添加一个
	AddNumInfo(ctx context.Context, info model.NumInfo) bool
	//根据Key获取一个
	GetNumInfoByKey(ctx context.Context, url string) model.NumInfo
	//查看全部
	FindAllNumInfo(ctx context.Context, page int, limit int) []model.NumInfo
	//根据Key修改
	UpdateNumInfoByKey(ctx context.Context, info model.NumInfo) bool
	//删除一个
	DeleteNumInfoById(ctx context.Context, id int64) bool
	//根据ID获取一个
	GetNumInfoById(ctx context.Context, id int64) model.NumInfo
	//根据ID修改
	UpdateNumInfoById(ctx context.Context, info model.NumInfo) bool
}
