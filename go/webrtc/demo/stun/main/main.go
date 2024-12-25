package main

import (
	"fmt"
	"github.com/pion/stun"
	"log"
	"net"
	"time"
)

func main() {
	// STUN服务器地址
	stunServerAddr := "8.130.14.29:53478"
	//stunServerAddr := "stun:stun.l.google.com:19302"

	// 创建UDP连接
	conn, err := net.Dial("udp", stunServerAddr)
	if err != nil {
		log.Fatalf("Failed to connect to STUN server: %v", err)
	}
	defer conn.Close()

	// 创建STUN客户端
	client, err := stun.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create STUN client: %v", err)
	}
	defer client.Close()

	// 设置用户名和密码
	username := "ZZGEDA"
	password := "tal1024"
	realm := "ZZGEDA.com"

	// 生成HA1
	ha1 := stun.NewLongTermIntegrity(username, realm, password)

	stunUdp(ha1, client)

	fmt.Println("准备进入轮询")

	for i := 0; i < 100; i++ {
		time.Sleep(time.Second * 3)
		stunUdp(ha1, client)
	}

	fmt.Println("轮询结束")
}

func stunUdp(ha1 stun.MessageIntegrity, client *stun.Client) {
	// 发送绑定请求
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest, ha1)
	if err := client.Do(message, func(res stun.Event) {
		if res.Error != nil {
			log.Fatalf("STUN request failed: %v", res.Error)
		}

		// 解析绑定响应
		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			log.Fatalf("Failed to get XOR-MAPPED-ADDRESS: %v", err)
		}

		// 打印公共IP地址和端口号
		fmt.Printf("Public IP: %s, Port: %d\n", xorAddr.IP, xorAddr.Port)
	}); err != nil {
		log.Fatalf("Failed to send STUN request: %v", err)
	}
}
