package main

import (
	"fmt"
	"net"
	"log"
	"golang.org/x/net/context"
	"github.com/golang/protobuf/ptypes/struct"
	grpc "google.golang.org/grpc"
	pb "grpc_demo/grpc_generic/proto"
)

type server struct{}

func(s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.CommonResponse, error){
	ret := new(pb.CommonResponse)
	ret.Err = 1
	ret.Msg ="test"
	ret.Data = &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "test"}}

	if req.JsCode == "ret" {
		data := &pb.LoginResult{
			OpenId: "openid...",
			SessionKey: "sessionkey...",
		}
		ret.Data = pb.ToValue(data)
		// ret.Data = &structpb.Value{
		// 	Kind: &structpb.Value_StructValue{
		// 		StructValue: &structpb.Struct{
		// 			Fields: fields,
		// 		},
		// 	},
		// }
	}
	fmt.Println("ret: %v", ret)
	return ret, nil
}

func main() {
	port := ":50000"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	fmt.Println("server start", port)
	s.Serve(lis)
}