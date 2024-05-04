package handler

import (
	grpcmicroservices "JohnKucharsky/grpc_microservices/services/common/genproto/orders"
	"JohnKucharsky/grpc_microservices/services/orders/models"
	"context"
	"google.golang.org/grpc"
)

type OrdersGRPCHandler struct {
	orderService models.OrderService
	grpcmicroservices.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService models.OrderService) {
	gRPCHandler := &OrdersGRPCHandler{
		orderService: orderService,
	}

	grpcmicroservices.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrdersGRPCHandler) GetOrders(ctx context.Context, req *grpcmicroservices.GetOrdersRequest) (*grpcmicroservices.GetOrderResponse, error) {

	o := h.orderService.GetOrders(ctx)
	resp := &grpcmicroservices.GetOrderResponse{
		Orders: o,
	}

	return resp, nil
}

func (h *OrdersGRPCHandler) CreateOrder(ctx context.Context, req *grpcmicroservices.CreateOrderRequest) (*grpcmicroservices.CreateOrderResponse, error) {
	order := &grpcmicroservices.Order{
		OrderID:    23,
		CustomerID: 3,
		ProductID:  4,
		Quantity:   2,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &grpcmicroservices.CreateOrderResponse{Status: "success"}

	return res, nil
}
