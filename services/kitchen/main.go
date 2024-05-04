package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func NewGRPCClient(addr string) *grpc.ClientConn {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err.Error())
	}

	return conn
}

func main() {
	httpServer := NewHttpServer(":1000")
	err := httpServer.Run()
	if err != nil {
		log.Fatalf("did not connect: %v", err.Error())
	}
}
