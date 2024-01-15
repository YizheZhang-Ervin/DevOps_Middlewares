# Projects

- gin-mall
    - web框架 gin
    - 数据库 mysql + gorm + dbresolver读写分离
    - 缓存 redis (存储商品的浏览次数)
    - 鉴权 jwt
    - 加密 crypto
    - 日志 logrus
    - ELK

- grpc-todoList
    - 服务发现 etcd
    - web框架 gin 
    - 远程服务调用 grpc + protobuf
    - 数据库 gorm + mysql
    - 鉴权 jwt
    - 日志 logrus
    - 配置库 viper

- micro-house
    - 功能：用户、房屋、订单
    - 启动：redis、fastdfs(trackerd/storaged)、nginx、consul、微服务

- micro-house-base-v2
    - web 客户端
        - mysql + gorm
        - redis + redigo
        - cookie & session
        - fastDFS + fdfd_client + fastdfs-nginx-module + nginx
        - consul
    - service 微服务
        - getcaptcha 图片验证码
        - user 用户+短信验证码

- micro-house-base-v3 (micro-house-base-v2升级v3)
    - gomicro/v3 + gin + consul
    - mysql(gorm) + redis(redigo) + fdfs
    - captcha + dysmsapi
    - grpc + protobuf
    - crypto

- micro-todoList-v2
    - 微服务 gomicro/v2
    - 服务发现 etcd
    - 鉴权 jwt
    - web框架 gin 
    - 数据库 gorm + mysql
    - 远程服务调用 grpc + protobuf
    - 消息队列 amqp
    - 服务熔断 hystrix
    - 加密 crypto

- micro-todoList-v3 (micro-todoList-v2升级v3)
    - gomicro/v3 + gin + etcd
    - mysql(gorm) + rabbitMQ(amqp)
    - grpc + protobuf
    - hystrix + jwt + crypto + cors + cookie&session

