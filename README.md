# MicroServices_GoMicro

## 一、项目
1. my-gomicro
    - gin + etcd
    - viper

2. my-gomicro2
    - gomicro/v3 + gin + etcd
    - mysql(gorm) + rabbitMQ(amqp)
    - grpc + protobuf
    - hystrix + jwt + crypto + cors + cookie&session

3. my-gomicro3
    - gomicro/v3 + gin + consul
    - mysql(gorm) + redis(redigo) + fdfs
    - captcha + dysmsapi
    - grpc + protobuf
    - crypto

4. my-srv-cli
    - gomicro/v4

## 二、技术栈
1. 相关技术
    - 配置
        - viper
        - gopkg.in/ini.v1
    - 微服务 
        - gomicro
        - gokit【暂无】
    - 服务发现 etcd/k8s
    - 服务熔断/负载均衡
        - hystrix
        - rate【暂无】
    - 网关
        - api gateway
        - web框架 gin
    - 中间件
        - mysql(gorm&validate)
        - redis(redigo)
        - rabbitMQ
        - elasticSearch【暂无】
        - fastdfs / ceph【暂无】
        - tidb【暂无】
    - 会话状态 cookie&session
    - 身份认证 jwt
    - 跨域 cors
    - 全局异常
    - 日志

2. 服务治理
    - Serverless (k8s的kubeless)
    - Istio
    - 各中间件operator

3. k8s运维
    - velero
    - kubevirt

4. 理论 (云原生&分布式)
    - 分布式协议
    - 云原生安全
    - k8s题目
    - 架构师

5. 实战
    - go实战
    - 秒杀系统
    - MIT课程

## 三、命令
```
# 1. golang
go clean --modcache
go mod init xxMod
go mod tidy
go mod vendor
go get xx/xx
go build
go run

# 2. etcd 
etcdctl--endpoints=http://127.0.0.1:2379 put xxKey"xxVal"
etcdctl--endpoints=http://127.0.0.1:2379 get xxKey
```
