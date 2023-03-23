package utils

import (
	"github.com/asim/go-micro/v3"
	"github.com/go-micro/plugins/v3/registry/consul"
)

// 初始化micro
func InitMicro() micro.Service {
	// 初始化客户端
	consulReg := consul.NewRegistry()

	return micro.NewService(
		micro.Registry(consulReg),
	)
}
