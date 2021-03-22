package main

import (
	"fmt"
	pb "grpc_demo/grpc_generic/proto"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.CommonResponse, error) {
	ret := new(pb.CommonResponse)
	ret.Err = 1
	ret.Msg = "test"
	ret.DataValue = &structpb.Value{Kind: &structpb.Value_StringValue{StringValue: "test"}}

	if req.JsCode == "value-struct" {
		data := &pb.LoginResult{
			OpenId:     "openid...",
			SessionKey: "sessionkey...",
		}
		ret.DataValue = pb.ToValue(data)
		// or
		// ret.Data = &structpb.Value{
		// 	Kind: &structpb.Value_StructValue{
		// 		StructValue: &structpb.Struct{
		// 			Fields: fields,
		// 		},
		// 	},
		// }
	} else if req.JsCode == "any" {
		// any
		anyResult := &pb.LoginResult{
			OpenId:     "openid--any",
			SessionKey: "sessionkey-any",
		}
		ret.DataAny, _ = ptypes.MarshalAny(anyResult)
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
