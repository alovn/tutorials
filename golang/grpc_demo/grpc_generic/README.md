# gRPC generic

```shell
go mod init
protoc -I . -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis  --go_out=plugins=grpc:. user.proto
```
