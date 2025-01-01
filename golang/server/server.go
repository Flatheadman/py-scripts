/*
编写思路与运行：
1.定位proto中的服务定义，找到服务名
2.定义一个结构体，嵌入未实现的服务结构体，名称为pb.Unimplemented服务名Server
3.找到pb文件中的Unimplemented服务名Server结构体和其下的方法，将其拷贝出来并实现具体逻辑
4.切换到server.go目录下: go run server.go
*/
package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Flatheadman/py-scripts/golang/rpc/hello" // 导入生成的 protobuf 和 gRPC 代码
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server 是 GreeterServer 的实现.
/*
service Greeter {
  // 定义一个方法
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}
*/
type server struct {
	pb.UnimplementedGreeterServer // 嵌入未实现的 GreeterServer。结构体嵌入：在定义结构体时，直接将另一个结构体类型作为字段嵌入，不需要指定字段名称, 效果等价于js对象中的键值同名省略写法。
	// 为什么使用未实现的 GreeterServer？向后兼容性: 如果 .proto 文件添加了新的 RPC 方法到 Greeter 服务中，而你的服务端代码没有及时更新，那么使用 pb.UnimplementedGreeterServer 可以避免编译错误。
	// 因为 pb.UnimplementedGreeterServer 会提供了所有gRPC 方法的空实现，客户端调用新的方法时，会收到一个 "unimplemented" 错误，这比直接导致服务端崩溃要好得多。
}

/*
	  // proto定义：
		service Greeter {
		  // 参考SayHello方法的proto定义
		  rpc SayHello (HelloRequest) returns (HelloReply) {}
		}
		// 生成的代码（以本处为准）
		func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
			return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
		}
*/
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{}) // 注册服务
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil { // 启动服务
		log.Fatalf("failed to serve: %v", err)
	}
}
