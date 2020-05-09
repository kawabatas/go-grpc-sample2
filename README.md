# go-grpc-sample2

## generate

```
export GRPC_GATEWAY_PROTO_PATH="$(go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)"

protoc \
-I proto/ \
-I${GRPC_GATEWAY_PROTO_PATH}/third_party/googleapis \
--go_out=plugins=grpc:./proto \
--grpc-gateway_out=logtostderr=true:./proto \
--swagger_out=logtostderr=true:./proto \
proto/hello.proto
```

## run

Server

```
go run server/main.go
go run gateway/main.go
```

Client

```
go run client/main.go

curl -X POST http://localhost:5000/v1/hello/name
```

## ref

https://github.com/grpc/grpc-go

https://github.com/golang/protobuf

https://github.com/grpc-ecosystem/grpc-gateway
