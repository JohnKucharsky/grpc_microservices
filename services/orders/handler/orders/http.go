package handler

import (
	grpc_microservices "JohnKucharsky/grpc_microservices/services/common/genproto/orders"
	"JohnKucharsky/grpc_microservices/services/common/utils"
	"JohnKucharsky/grpc_microservices/services/orders/models"
	"net/http"
)

type OrdersHttpHandler struct {
	ordersService models.OrderService
}

func NewHttpOrdersHandler(orderService models.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (handler *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", handler.CreateOrder)
}

func (handler *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req grpc_microservices.CreateOrderRequest
	err := utils.ParseJSON(r, &req)
	if err != nil {
		err := utils.WriteError(w, http.StatusBadRequest, err)
		if err != nil {
			return
		}
		return
	}

	order := &grpc_microservices.Order{
		OrderID:    234,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = handler.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		err := utils.WriteError(w, http.StatusInternalServerError, err)
		if err != nil {
			return
		}
		return
	}

	res := &grpc_microservices.CreateOrderResponse{Status: "success"}
	utils.WriteJSON(w, http.StatusCreated, res)
}
