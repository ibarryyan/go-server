package main

import (
	"count_num/cmd/grpc"
	"count_num/cmd/web"
)

func main() {
	go grpc.Run()
	go web.Run()
	select {}
}
