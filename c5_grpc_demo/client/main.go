package main

import (
	pb "GoRun/c5_grpc_demo/client/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/**
 * Author: tyza66
 * Date: 2023/04/20 13:49
 * Github: https://github.com/tyza66
 **/
func main() {
	//grpc连接到server端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//建立连接
	client := pb.NewSayHelloClient(conn)
	//执行rpc调用 （这个方法在服务端实现并且返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "tyza66"})
	fmt.Println(resp.GetResponseMsg())
}
