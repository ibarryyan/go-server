package impl

import (
	"context"
	"count_num/pkg/dao/impl"
	"count_num/proto"
	"fmt"
)

//定义服务端 实现 约定的接口
type NumInfoRPCImpl struct {
	dao *impl.CountNumDAOImpl
}

func NewNumInfoControllerImpl() *NumInfoRPCImpl {
	return &NumInfoRPCImpl{dao: impl.NewCountNumDAOImpl()}
}

func (impl *NumInfoRPCImpl) AddNumByKey(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	panic("implement me")
}

func (impl *NumInfoRPCImpl) FindNumInfoByKey(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	panic("implement me")
}

func (impl *NumInfoRPCImpl) SaveNumInfo(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	panic("implement me")
}

func (impl *NumInfoRPCImpl) DeleteById(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	panic("implement me")
}

func (impl *NumInfoRPCImpl) FindAll(ctx context.Context, request *proto.InfoRequest) (*proto.InfoResponse, error) {
	panic("implement me")
}

func (impl *NumInfoRPCImpl) GetNumInfoById(ctx context.Context, req *proto.InfoRequest) (resp *proto.InfoResponse, err error) {
	id := req.GetId()
	numInfo := impl.dao.GetNumInfoById(ctx, id)
	fmt.Println(numInfo)
	return &proto.InfoResponse{Code: 0, Msg: "", Count: 1, Data: numInfo.Name}, err
}
