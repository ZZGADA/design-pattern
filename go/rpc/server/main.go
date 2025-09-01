package main

import (
	pb "ZZGADA.com/grpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	// UnimplementedSayHelloServer提供了所有接口方法的默认实现, 通过嵌入, Server结构体获取其所有方法并实现需要重写的方法
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("收到请求:", in.RequestName, ctx)
	return &pb.HelloResponse{ResponseMsg: fmt.Sprintf("你好, %s", in.RequestName)}, nil
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("无法监听端口: %v", err)
	}

	// 创建gRPC服务器
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterSayHelloServer(s, &server{})

	log.Println("服务器启动在 :50051")

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
