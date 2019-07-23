package main

import (
	"flag"
	"log"
	"net"
	"fmt"
	"os"
	"time"
	"os/signal"
	"syscall"

	"github.com/coreos/etcd/proxy/grpcproxy"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc_demo/grpc_etcd_discovery2/proto"
)

//HelloServer type
type HelloServer struct{}

//SayHello func
func (h *HelloServer) SayHello(ctx context.Context, req *pb.HelloRequest)(*pb.HelloResponse, error){
	response := &pb.HelloResponse{
		Reply: fmt.Sprintf("hello, %s", req.Greeting),
	}
	fmt.Println(req.Greeting)
	return response, nil
}

func main(){
	port := flag.Int("port", 7788, "the grpc server port")
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatal(err)
	}
	
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &HelloServer{})
	go func(){
		if err := grpcServer.Serve(listen); err != nil {
			panic(err)
		} else {
			fmt.Printf("grpc service start at prot: %d\n", *port)
		}
	}()
	
	etcdClient, err:= clientv3.New(clientv3.Config{
		Endpoints: []string{"http://s1004.lab.org:2379"},
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("%s:%d", "localhost", *port)//TODO: get ip
	registTTL := 2
	serviceName  := "hello-service"
	grpcproxy.Register(etcdClient, fmt.Sprintf("/_grpc/service/%s", serviceName), addr, registTTL)

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	sign := <- signalChannel
	fmt.Printf("receive signal %v, grpc serve will stop\n", sign)
	etcdClient.Close()
	grpcServer.GracefulStop()
	listen.Close()
}