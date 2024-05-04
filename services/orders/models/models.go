package models

import (
	grpcmicroservices "JohnKucharsky/grpc_microservices/services/common/genproto/orders"
	"context"
)

type OrderService interface {
	CreateOrder(context.Context, *grpcmicroservices.Order) error
	GetOrders(ctx context.Context) []*grpcmicroservices.Order
}
