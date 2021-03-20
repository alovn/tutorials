package main

import (
	"context"
	"log"
	"net"

	_ "grpc_json/codec"
	pb "grpc_json/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Greeting)

	return &pb.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	hServer := &server{}
	pb.RegisterHelloServiceServer(s, hServer)

	// Serve gRPC Server
	log.Println("Serving gRPC Server")
	log.Fatal(s.Serve(lis))
}
