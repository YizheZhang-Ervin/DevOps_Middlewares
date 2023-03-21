# MicroServices_GoMicro

## Coding
1. my-gomicro
- web框架 gin
- 配置 viper
- 微服务 gomicro
- 服务发现 etcd/k8s
- 服务熔断/负载均衡 wrapper(hystrix&rate)
- 网关
- 中间件
    - mysql(gorm&validate)
    - redis(redigo)
    - rabbitMQ
    - elasticSearch
    - ceph
    - tidb
- 会话状态 cookie&session
- 身份认证 jwt
- 跨域 cors
- 全局异常
- 日志

2. 服务治理
- Serverless
- Istio
- 各中间件operator

3. k8s运维
- velero
- kubevirt

4. 理论 (云原生&分布式)
- 分布式协议
- 云原生安全
- k8s题目

5. 实战
- go实战
- 秒杀系统
- MIT课程

## 命令
```
# 1. golang
go mod init xxMod
go mod tidy
go get xx/xx
go build
go run

# 2. etcd 
etcdctl--endpoints=http://127.0.0.1:2379 put xxKey"xxVal"
etcdctl--endpoints=http://127.0.0.1:2379 get xxKey
```
