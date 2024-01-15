package main

import (
	"user/conf"
	"user/core"
	"user/services"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/go-micro/plugins/v3/registry/etcd"
)

func main() {
	conf.Init()
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	_ = microService.Run()
}
