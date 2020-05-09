package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "github.com/kawabatas/go-grpc-sample2/proto/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("localhost:5001")
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		return err
	}
	return http.ListenAndServe(":5000", mux)
}

func main() {
	log.Println("gateway starting...")
	if err := run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
