package main

import (
	"log"
	pb "grpc_demo/grpc_middleware/proto" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	address = "127.0.0.1:50001"
)

// customCredential 自定义认证
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "10001",
		"appkey": "mysecret",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(address, opts...)

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloServiceClient(conn)

	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Greeting = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Greeting: %s, %v", r.Reply, r.Number)
}