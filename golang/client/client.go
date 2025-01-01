/*
切换到client.go目录并运行：go run client.go
应该看到输出：2024/12/28 03:11:51 Greeting: Hello world
*/
package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Flatheadman/py-scripts/golang/rpc/hello" // 导入生成的 protobuf 和 gRPC 代码
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:50051" // 服务端地址
	defaultName = "world"
)

func main() {
	// 建立连接，创建参数：服务端地址，客户端证书。客户端证书：insecure.NewCredentials()，不安全的证书，相当于不使用证书
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn)

	name := defaultName
	// 创建一个上下文（带有超时时间）
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 调用服务端的方法
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
