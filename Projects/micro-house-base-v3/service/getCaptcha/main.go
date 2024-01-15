package main

import (
	"micro-house-base-v3/service/get-captcha/handler"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"

	getCaptcha "micro-house-base-v3/service/get-captcha/proto/getCaptcha"

	"github.com/go-micro/plugins/v3/registry/consul"
)

func main() {
	// 初始化consul
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Address("192.168.6.108:12341"), // 防止随机生成 port
		micro.Name("getCaptcha"),
		micro.Registry(consulReg), // 添加注册
		micro.Version("latest"),
	)

	// Register Handler
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
