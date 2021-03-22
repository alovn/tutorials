package main

import (
	"context"
	"log"
	"os"
	"time"
	"crypto/x509"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc_demo/grpc_gateway/proto"
	"grpc_demo/grpc_gateway/insecure"
)

const (
	address     = "localhost:8888"
	defaultName = "world"
)

func main() {
	demoCertPool := x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(insecure.Cert))
	if !ok {
		panic("bad certs")
	}
	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(demoCertPool, address)
	opts = append(opts, grpc.WithTransportCredentials(creds))
    conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s, %v", r.Reply, r.Number)
}