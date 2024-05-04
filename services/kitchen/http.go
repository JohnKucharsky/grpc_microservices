package main

import (
	grpc_microservices "JohnKucharsky/grpc_microservices/services/common/genproto/orders"
	"context"
	"html/template"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{addr: addr}
}

func (s *HttpServer) Run() error {
	router := http.NewServeMux()

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := grpc_microservices.NewOrderServiceClient(conn)
		ctx, cancel := context.WithTimeout(r.Context(), time.Second)
		defer cancel()

		_, err := c.CreateOrder(ctx, &grpc_microservices.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  234,
			Quantity:   2,
		})

		if err != nil {
			log.Fatalf("CreateOrder err: %v", err)
		}

		orders, err := c.GetOrders(ctx, &grpc_microservices.GetOrdersRequest{
			CustomerID: 24,
		})
		if err != nil {
			log.Fatalf("CreateOrder err: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := t.Execute(w, orders.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err.Error())
		}

	})

	log.Printf("Listening on %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
