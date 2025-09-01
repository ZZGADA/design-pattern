package main

import (
	pb "ZZGADA.com/grpc/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// 这里创建不安全的传输凭证, 即告诉gRPC客户端跳过加密, 明文通信
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	sayHelloClient := pb.NewSayHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := sayHelloClient.SayHello(ctx, &pb.HelloRequest{RequestName: "ZZGEDA"})
	if err != nil {
		log.Fatalf("SayHello error: %v", err)
	}

	log.Println(resp)
}
