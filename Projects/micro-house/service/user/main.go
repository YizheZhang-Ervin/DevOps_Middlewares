package main

import (
	"micro-house/service/user/handler"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"micro-house/service/user/model"
	user "micro-house/service/user/proto/user"

	"github.com/micro/go-micro/registry/consul"
)

func main() {
	//使用consul做服务发现
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9984"),
	)

	// Initialise service
	service.Init()
	model.InitDb()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
