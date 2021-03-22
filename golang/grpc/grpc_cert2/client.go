package main

import (
	"context"
	"log"
	"os"
	"io/ioutil"
	"crypto/tls"
	"crypto/x509"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc_demo/grpc_cert2/hello"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	certificate, err := tls.LoadX509KeyPair("cert/client.crt", "cert/client.key")
    if err != nil {
        log.Fatal(err)
    }

    certPool := x509.NewCertPool()
    ca, err := ioutil.ReadFile("cert/ca.crt")
    if err != nil {
        log.Fatal(err)
    }
    if ok := certPool.AppendCertsFromPEM(ca); !ok {
        log.Fatal("failed to append ca certs")
    }

    creds := credentials.NewTLS(&tls.Config{
        Certificates:       []tls.Certificate{certificate},
        ServerName:         "server.io", // NOTE: this is required!
        RootCAs:            certPool,
    })
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
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