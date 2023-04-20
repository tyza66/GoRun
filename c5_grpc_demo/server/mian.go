package main

import (
	pb "GoRun/c5_grpc_demo/server/proto/"
	"context"
	"google.golang.org/grpc"
	"net"
)

/**
 * Author: tyza66
 * Date: 2023/04/20 13:49
 * Github: https://github.com/tyza66
 **/
//helloServer
type server struct {
	pb.UnimplementedSayHelloServer
}
//重写这个方法 实现服务端被调用的方法
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "hello," + req.RequestName}, nil
}

func main() {
	//开启端口
	listen,_ := net.Listen("tcp","9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
}
