package main

import (
    "github.com/micro/go-log"
    "github.com/micro/go-micro"
    "ShopHome/DataInfoSrv/handler"
    example "ShopHome/DataInfoSrv/proto/example"
    "github.com/micro/go-grpc"
)

func main() {
    // New Service
    service := grpc.NewService(
        micro.Name("go.micro.srv.DataInfoSrv"),
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
