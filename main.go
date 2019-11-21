package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/alactic/shippyproject/handler"
	"github.com/alactic/shippyproject/subscriber"

	shippyproject "github.com/alactic/shippyproject/proto/shippyproject"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.shippyproject"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	shippyproject.RegisterShippyprojectHandler(service.Server(), new(handler.Shippyproject))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.shippyproject", service.Server(), new(subscriber.Shippyproject))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.shippyproject", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
