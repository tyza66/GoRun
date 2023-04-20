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

// token认证需要实现接口 我们先把这个接口提出来
type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}

type ClientTokenAuth struct {
}
//这个地方发送了验证的token内容 只有令牌正确才有值给
func (c *ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"appId":"giao","appKey":"giao"},nil
}
//这个地方决定是否使用ssl证书加密验证
func (c *ClientTokenAuth) RequireTransportSecurity() bool {
	//返回false 就是不开启安全验证
	return false
}

func main() {
	var opts []grpc.DialOption
	//拼接一个基础的
	opts =append(opts,grpc.WithTransportCredentials(insecure.NewCredentials()))
	//加入我们自定义的实现接口的那个结构体
	opts = append(opts,grpc.WithPerRPCCredentials(new(ClientTokenAuth)))
	//grpc连接到server端，此处禁用安全传输，没有加密和验证
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
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
