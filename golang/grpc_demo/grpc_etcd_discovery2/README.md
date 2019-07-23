# gRPC etcd discovery

    protoc --go_out=plugins=grpc:. hello.proto 

    go run server/main.go
    go run client/main.go
