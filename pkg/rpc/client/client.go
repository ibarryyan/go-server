package main

import (
	"context"
	"count_num/proto"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

type NumInfoServiceClientImpl struct {
	proto.NumInfoServiceClient
}

func NewNumInfoServiceClientImpl() *NumInfoServiceClientImpl {
	conn, err := grpc.Dial(":6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("正在监听服务端 : %v\n", err)
	}
	return &NumInfoServiceClientImpl{proto.NewNumInfoServiceClient(conn)}
}

func main() {
	//1 配置grpc服务端的端口作为客户端的监听
	conn, err := grpc.Dial(":6666", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("正在监听服务端 : %v\n", err)
	}
	defer conn.Close()
	//2 实例化 UserInfoService 服务的客户端
	client := proto.NewNumInfoServiceClient(conn)
	//3 调用grpc服务
	req := new(proto.InfoRequest)
	req.Id = 20
	resp, err := client.FindAll(context.Background(), req)
	if err != nil {
		log.Fatalf("请求错误 : %v\n", err)
	}
	fmt.Printf("响应内容 : %v\n", resp)
}
