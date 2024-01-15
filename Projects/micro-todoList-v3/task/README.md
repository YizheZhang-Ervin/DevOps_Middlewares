# Task Service
1. 结构
```
task/
├── conf
├── core
├── model
└── service

- conf：配置信息
- core：业务逻辑
- model：数据库模型
- service：proto文件以及各服务

2. 命令
```
micro new task
make proto
micro run .
```