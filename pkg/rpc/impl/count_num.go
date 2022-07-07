package impl

import (
	"context"
	"count_num/pkg/dao/impl"
	"count_num/pkg/model"
	"count_num/proto"
	"encoding/json"
)

type NumInfoRPCImpl struct {
	dao *impl.CountNumDAOImpl
}

func NewNumInfoControllerImpl() *NumInfoRPCImpl {
	return &NumInfoRPCImpl{dao: impl.NewCountNumDAOImpl()}
}

func (impl *NumInfoRPCImpl) AddNumByKey(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	key := request.GetInfoKey()
	id := request.GetId()
	name := request.GetName()
	num := request.GetInfoNum()
	impl.dao.UpdateNumInfoByKey(ctx, model.NumInfo{
		id,
		name,
		key,
		num,
	})
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

func (impl *NumInfoRPCImpl) FindNumInfoByKey(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	key := request.GetInfoKey()
	numInfo := impl.dao.GetNumInfoByKey(ctx, key)
	info, _ := json.Marshal(numInfo)
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
}

func (impl *NumInfoRPCImpl) SaveNumInfo(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	key := request.GetInfoKey()
	id := request.GetId()
	name := request.GetName()
	num := request.GetInfoNum()
	impl.dao.AddNumInfo(ctx, model.NumInfo{
		id,
		name,
		key,
		num,
	})
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

func (impl *NumInfoRPCImpl) DeleteById(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	id := request.GetId()
	impl.dao.DeleteNumInfoById(ctx, id)
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

func (impl *NumInfoRPCImpl) FindAll(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	numInfos := impl.dao.FindAllNumInfo(ctx, 0, 0)
	infos, _ := json.Marshal(numInfos)
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *NumInfoRPCImpl) GetNumInfoById(ctx context.Context, req *proto.InfoRequest) (resp *proto.InfoResponse, err error) {
	id := req.GetId()
	numInfo := impl.dao.GetNumInfoById(ctx, id)
	info, _ := json.Marshal(numInfo)
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
}
