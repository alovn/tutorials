package main

import (
	"context"
	"time"
	"fmt"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"

	"google.golang.org/grpc"
	pb "grpc_demo/grpc_discovery/proto"
)

type GrpcConnection interface {
	getGrpcConnection() *grpc.ClientConn
	connect()
}


type EtcdResolverConnection struct {
	grpcConn    *grpc.ClientConn
	serviceName string
}


func (e *EtcdResolverConnection) connect() {
	config := clientv3.Config{
		Endpoints:   []string{"http://s1004.lab.org:2379"},
		DialTimeout: 5 * time.Second,
		//Username:    username,
		//Password:    password,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}

	r := &etcdnaming.GRPCResolver{Client: cli}
	b := grpc.RoundRobin(r)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		e.serviceName,
		grpc.WithInsecure(), 
		grpc.WithTimeout(time.Second * 5),
		grpc.WithBalancer(b),
		grpc.WithBlock(),
	)
	if err != nil {
		fmt.Printf("dial service(%s) by etcd resolver server error (%v)", e.serviceName, err.Error())
		panic(err)
	}
	e.grpcConn = conn
}

func main() {
	conn := &EtcdResolverConnection{
		serviceName: "/_grpc/service/hello-service",
	}
	conn.connect()

	defer conn.grpcConn.Close()
	for i := 0; i < 100; i++ {
		request := &pb.HelloRequest{Greeting: fmt.Sprintf("send: %d", i)}
		client := pb.NewHelloServiceClient(conn.grpcConn)
		resp, err := client.SayHello(context.Background(), request)
		fmt.Printf("resp: %+v, err: %+v\n", resp, err)
		time.Sleep(time.Second)
	}
}
