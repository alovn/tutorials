package main

import (
	"log"
	"net/http"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
    "golang.org/x/net/context"
	"google.golang.org/grpc"
	
	pb "grpc_demo/grpc_gateway/proto"
)

func main(){
	ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, ":7777", opts)
	if err != nil {
        log.Fatal(err)
    }

    log.Println("grpc gateway 服务开启: http:8080 -> grpc:7777")
	if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
	}
}