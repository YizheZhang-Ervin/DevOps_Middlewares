# my-gomicro2

1. 技术栈
- 微服务 gomicro
- web框架 gin
- 远程服务调用 grpc + protobuf
- 服务发现 etcd
- 服务熔断 hystrix
- 数据库 gorm + mysql
- 消息队列 amqp + rabbitMQ
- 鉴权 jwt
- 加密 crypto
- 跨域 cors
- 日志
- 全局异常
- 会话状态 cookie&session

2. 功能
- 用户注册登录(jwt鉴权)
- 新增/删除/修改/查询 备忘录

3. 部署
- 运行RabbitMQ
- 运行ETCD
    - go run main.go --registry=etcd --registry_address=127.0.0.1:2379
- 运行各模块