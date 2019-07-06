package main

import (
    "google.golang.org/grpc/grpclog"
    "google.golang.org/grpc/status"
    "net"
    "log"

    pb "grpc_demo/grpc_interceptor/proto"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"       // grpc 响应状态码
    //"google.golang.org/grpc/credentials" // grpc认证包
    "google.golang.org/grpc/metadata" // grpc metadata包
)


// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService ...
var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    resp := new(pb.HelloResponse)
    resp.Reply = "Hello " + in.Greeting + "."

    return resp, nil
}

// auth 验证Token
func auth(ctx context.Context) error {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return status.Errorf(codes.Unauthenticated, "无Token认证信息")
    }

    var (
        appid  string
        appkey string
    )

    if val, ok := md["appid"]; ok {
        appid = val[0]
    }

    if val, ok := md["appkey"]; ok {
        appkey = val[0]
    }

    if appid != "10001" || appkey != "mysecret" {
        return status.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
    }

    return nil
}

func main() {
    listen, err := net.Listen("tcp", "127.0.0.1:50052")
    if err != nil {
        grpclog.Fatalf("Failed to listen: %v", err)
    }

    var opts []grpc.ServerOption

    // TLS认证
    //creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
    //if err != nil {
    //    grpclog.Fatalf("Failed to generate credentials %v", err)
    //}

    //opts = append(opts, grpc.Creds(creds))

    // 注册interceptor
    var interceptor grpc.UnaryServerInterceptor
    interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
        err = auth(ctx)
        if err != nil {
            return
        }
        // 继续处理请求
        return handler(ctx, req)
    }
    opts = append(opts, grpc.UnaryInterceptor(interceptor))

    // 实例化grpc Server
    s := grpc.NewServer(opts...)

	// 注册HelloService
    pb.RegisterHelloServiceServer(s, HelloService)

    log.Println("Listen on 50052 with Token + Interceptor")

    s.Serve(listen)
}