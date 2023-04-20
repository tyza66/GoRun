package main

import (
	pb "GoRun/c5_grpc_demo/server/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

// 重写这个方法 实现服务端被调用的方法 业务逻辑
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	//获取元数据的信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	//验证token
	fmt.Println(appId, appKey)
	if appId != "giao" || appKey != "giao" {
		return nil,errors.New("token不一致")
	}

	return &pb.HelloResponse{ResponseMsg: "hello," + req.RequestName}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()
	//在grpc客户端注册我们自己写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})
	//启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return

	}
}
