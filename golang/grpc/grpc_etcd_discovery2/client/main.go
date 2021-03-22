package main

import (
	"fmt"
	"time"

	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "grpc_demo/grpc_etcd_discovery2/proto"
)

//GrpcConnection interface
type GrpcConnection interface {
	getGrpcConnection() *grpc.ClientConn
	connect()
}

type EtcdResolverConnection struct {
	grpcConn    *grpc.ClientConn
	serviceName string
}

func (e *EtcdResolverConnection) connect() {
	etcdClient, err:= clientv3.New(clientv3.Config{
		Endpoints: []string{"http://s1004.lab.org:2379"},
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		panic(err)
	}

	r := &etcdnaming.GRPCResolver{Client: etcdClient}
	b := grpc.RoundRobin(r)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()
	conn, gerr := grpc.DialContext(
		ctx,
		e.serviceName,
		grpc.WithInsecure(),
		grpc.WithBalancer(b),
		grpc.WithTimeout(time.Second*5),
		grpc.WithBlock(),
	)
	if gerr != nil {
		fmt.Printf("dial service(%s) by etcd resolver server error (%v)", e.serviceName, gerr.Error())
		panic(gerr)
	}
	e.grpcConn = conn

}

func (e *EtcdResolverConnection) getGrpcConnection() *grpc.ClientConn {
	return e.grpcConn
}

func main() {
	conn := &EtcdResolverConnection{
		serviceName: "/_grpc/service/hello-service",
	}
	conn.connect()

	for i := 0; i < 100; i++ {
		request := &pb.HelloRequest{Greeting: fmt.Sprintf("%d", i)}

		client := pb.NewHelloServiceClient(conn.getGrpcConnection())

		resp, err := client.SayHello(context.Background(), request)
		fmt.Printf("resp: %+v, err: %+v\n", resp, err)
		time.Sleep(time.Second)
	}
}
