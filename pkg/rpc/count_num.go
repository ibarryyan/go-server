package rpc

import (
	"count_num/pkg/rpc/impl"
	"count_num/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGRPC() {
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
	proto.RegisterNumInfoServiceServer(s, impl.NewNumInfoControllerImpl())
	s.Serve(l)
}
