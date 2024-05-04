package main

import (
	"JohnKucharsky/grpc_microservices/services/common/utils"
	handler "JohnKucharsky/grpc_microservices/services/orders/handler/orders"
	"JohnKucharsky/grpc_microservices/services/orders/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{addr: addr}
}

func (srv *GRPCServer) Run() error {
	lis, err := net.Listen("tcp", srv.addr)
	utils.PrintError(err)

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	log.Println("Starting gRPC Server")

	return grpcServer.Serve(lis)
}
