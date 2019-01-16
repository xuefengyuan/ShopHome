package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"ShopHome/UserInfoSrv/handler"
	"ShopHome/UserInfoSrv/subscriber"

	example "ShopHome/UserInfoSrv/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.UserInfoSrv"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.UserInfoSrv", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.UserInfoSrv", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
