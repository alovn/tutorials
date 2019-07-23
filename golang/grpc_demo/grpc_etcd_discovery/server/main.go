package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/etcd/clientv3"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/reflection"

	pb "grpc_demo/grpc_etcd_discovery/proto"

	"github.com/coreos/etcd/clientv3/concurrency"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Greeting)
	return &pb.HelloResponse{Reply: "Hello " + in.Greeting, Number: []int32{1, 2}}, nil
}

func etcdRegister(c *clientv3.Client, service, addr string, ttl int) error {
	ss, err := concurrency.NewSession(c, concurrency.WithTTL(ttl))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	r := &etcdnaming.GRPCResolver{Client: c}
	if err = r.Update(c.Ctx(), service, naming.Update{Op: naming.Add, Addr: addr}, clientv3.WithLease(ss.Lease())); err != nil {
		log.Fatalf("error: %s", err)
	}
	return err
}

func main() {
	port := flag.Int("port", 7788, "the grpc server port")
	flag.Parse()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	config := clientv3.Config{
		Endpoints:   []string{"http://s1004.lab.org:2379"},
		DialTimeout: 3 * time.Second,
		//Username:    username,
		//Password:    password,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	etcdRegister(cli, "/_grpc/service/hello-service", fmt.Sprintf(":%d", *port), 2)

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterHelloServiceServer(s, &server{})

	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		} else {
			log.Printf("grpc service start at port: %v", *port)
		}
	}()

	// wait for stop signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sign := <-signalChan
	fmt.Printf("receive signal (%v) ,grpc server will stop", sign)
	s.GracefulStop()
}
