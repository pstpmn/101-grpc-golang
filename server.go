package main

import (
	"101-grpc-golang/hello"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	"google.golang.org/grpc"
)

type helloServer struct{}

func (hs *helloServer) Say(ctx context.Context, in *hello.SayRequest) (*hello.SayResponse, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	log.Println(headers, ok)
	log.Println(ctx)
	return &hello.SayResponse{Message: fmt.Sprintf("Hello %s eiei", in.Name)}, nil
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hello.RegisterHelloServer(s, &helloServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
