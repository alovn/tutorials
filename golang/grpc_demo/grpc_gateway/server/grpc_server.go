package main

import (
	"net"
	"fmt"
	"log"
    "golang.org/x/net/context"
	"google.golang.org/grpc"
	
	pb "grpc_demo/grpc_gateway/proto"
)

type helloService struct {}

func (service *helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
		Reply: fmt.Sprintf("hello, %s", req.Greeting),
	}
	fmt.Println(req.Greeting)
	return response, nil
}

func main(){
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &helloService{})

	port := 7777
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("grpc server start at port: %d", port)
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	} else {
		fmt.Printf("grpc service start at prot: %d", port)
	}
}