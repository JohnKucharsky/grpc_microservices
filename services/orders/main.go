package main

import (
	"JohnKucharsky/grpc_microservices/services/common/utils"
	"fmt"
)

func main() {
	httpServer := NewHttpServer(":8000")
	go func() {
		err := httpServer.Run()
		if err != nil {
			fmt.Println(err.Error())
		}
	}()

	grpcServer := NewGRPCServer(":9000")
	err := grpcServer.Run()
	utils.PrintError(err)

}
