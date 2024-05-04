package service

import (
	grpcmicroservices "JohnKucharsky/grpc_microservices/services/common/genproto/orders"
	"context"
)

var orders = make([]*grpcmicroservices.Order, 0)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *grpcmicroservices.Order) error {
	orders = append(orders, order)

	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*grpcmicroservices.Order {
	return orders
}
