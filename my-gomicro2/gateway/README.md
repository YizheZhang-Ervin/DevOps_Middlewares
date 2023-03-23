# Gateway Service

1. 结构
```
gateway/
├── pkg
│  ├── e
│  ├── logging
│  └── util
├── services
│  └── proto
├── weblib
│  ├── handlers
│  └──  middleware
└── wrappers

- pkg/e : 封装错误码
- pkg/logging : 日志文件
- pkg/util : 工具函数
- service/proto : 放置proto文件以及生成的pb文件
- weblib/handlers : 各个服务的接口
- weblib/middleware : http服务器的中间件
- wrappers : 放置服务熔断的配置
```

2. 命令
```
micro new gateway
make proto
micro run .
```