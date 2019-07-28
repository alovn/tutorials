package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"crypto/tls"
	"crypto/x509"

	"google.golang.org/grpc/credentials"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
    "golang.org/x/net/context"
	"google.golang.org/grpc"
	
	pb "grpc_demo/grpc_gateway/proto"
	"grpc_demo/grpc_gateway/insecure"
)

type helloService struct {}

func (service *helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
		Reply: fmt.Sprintf("hello, %s", req.Greeting),
	}
	fmt.Println(req.Greeting)
	return response, nil
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			log.Println("server grpc")
			grpcServer.ServeHTTP(w, r)
		} else {
			log.Println("server http")
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	port := 8888
	demoAddr := fmt.Sprintf("localhost:%d", port)

	pair, err := tls.X509KeyPair([]byte(insecure.Cert), []byte(insecure.Key))
	if err != nil {
		panic(err)
	}
	demoKeyPair := &pair
	demoCertPool := x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM([]byte(insecure.Cert))
	if !ok {
		panic("bad certs")
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, demoAddr))}


	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(grpcServer, &helloService{})
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs:    demoCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	err = pb.RegisterHelloServiceHandlerFromEndpoint(ctx, gwmux, demoAddr, dopts)
	if err != nil {
		log.Fatalf("serve: %v\n", err)
	}
	mux.Handle("/", gwmux)

	conn, err := net.Listen("tcp", demoAddr)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    demoAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
			NextProtos:   []string{"h2"}, //to force http2
		},
	}

	fmt.Printf("grpc on port: %d\n", port)
	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}