package main

import (
    "github.com/micro/go-log"
    "github.com/micro/go-micro"
    "github.com/micro/go-grpc"
    "ShopHome/OrderSrv/handler"
    example "ShopHome/OrderSrv/proto/example"
)

func main() {
    // New Service
    service := grpc.NewService(
        micro.Name("go.micro.srv.OrderSrv"),
        micro.Version("latest"),
    )

    // Initialise service
    service.Init()

    // Register Handler
    example.RegisterExampleHandler(service.Server(), new(handler.Example))
    // Run service
    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}
