# gRPC示例

```bash
protoc --go_out=plugins=grpc:. hello.proto

go run server.go
go run client.go
```
