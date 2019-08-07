package main

import (
	"log"

	pb "grpc_demo/grpc_generic/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main(){
	address := ":50000"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserClient(conn)
	req := new(pb.LoginRequest)
	req.JsCode = "code"
	r, err := c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("str generic failed: %v", err)
	}
	log.Printf("ret: %v\n", r)

	req.JsCode = "ret"
	r, err = c.Login(context.Background(), req)
	if err != nil {
		log.Fatalf("str generic failed: %v", err)
	}
	

	log.Printf("ret: %v\n", r.Msg)
	log.Printf("ret: %v\n", r.Err)
	log.Printf("ret: %v\n", r.Data.GetStructValue().Fields["open_id"].GetStringValue())
	log.Printf("ret: %v\n", r.Data.GetStructValue().Fields["session_key"].GetStringValue())
}