# gRPC generic

```shell
go mod tidy

cd proto && protoc -I . -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis  --go_out=. --go-grpc_out=. user.proto

go run server.go

go run client.go
```
