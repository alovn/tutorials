# gRPC gateway

    protoc -I . -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis  --go_out=plugins=grpc:. hello.proto

    protoc -I . -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/hello.proto

## 一、启动测试

    //启动 grpc server
    go run server/grpc_server.go
    
    //启动 grpc gateway
    go run server/grpc_gateway.go

然后可以通过Postman 或 curl 进行请求测试：

    curl -X POST http://localhost:8080/hello/say \
        -H 'Content-Type: application/json' \
        -d '{"greeting": "a"}'


openssl genrsa -out ca-key.pem 1024

