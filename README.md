# go-grpc-sample2

## generate

```
protoc -I proto/ proto/hello.proto --go_out=plugins=grpc:./proto
```

## run

Server

```
go run server/main.go
```

Client

```
go run client/main.go
```

## ref

https://github.com/grpc/grpc-go

https://github.com/golang/protobuf
