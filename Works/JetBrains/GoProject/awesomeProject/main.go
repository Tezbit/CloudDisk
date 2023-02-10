package main

import (
	"awesomeProject/example"
	"awesomeProject/token"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithKeepaliveParams(kacp))
	if err != nil {
		log.Fatal("连接 gPRC 服务失败,", err)
	}

	defer conn.Close()

	// 创建 gRPC 客户端
	grpcClient := example.NewLoginServiceClient(conn)

	// 创建请求参数
	request := example.LoginRequest{
		Account:  "name",
		Password: "pwd",
	}

	// 发送请求，调用 MyTest 接口
	response, err := grpcClient.Login(context.Background(), &request)
	if response.Code != 100 {
		log.Fatal("发送请求失败，原因是:", err)
	}
	log.Println(response)

	request = example.LoginRequest{
		Account:  "xiaoming",
		Password: "123456",
	}

	// 发送请求，调用 MyTest 接口
	response, err = grpcClient.Login(context.Background(), &request)
	if response.Code != 100 {
		log.Fatal("发送请求失败，原因是:", err)
	}
	requestToken := new(token.AuthToken)
	requestToken.Token = response.Token

}
