package main

import (
	"context"
	"count_num/pkg/dao/impl"
	"count_num/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
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

func main() {
	//1 添加监听的端口
	port := ":6666"
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("端口监听错误 : %v\n", err)
	}
	fmt.Printf("正在监听： %s 端口\n", port)
	//2 启动grpc服务
	s := grpc.NewServer()
	//3 将服务注册到gRPC中 ,注意第二个参数是接口类型的变量，需要取地址传参
	proto.RegisterNumInfoServiceServer(s, NewNumInfoControllerImpl())
	s.Serve(l)
}
