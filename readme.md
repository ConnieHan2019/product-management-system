# product-management-system


# Project Structure
```
PRODUCT-MANAGEMENT-SYSTEM
├── .vscode
│   └── launch.json        # VS Code 配置文件
├── bin                    # 可执行文件目录
├── config                 # 配置文件目录
│   └── go.config.go       # Go 配置文件
├── docker                 # Docker 相关文件目录
│   └── pkg
│       └── config
│           └── database
│               ├── GO.connect.go    # 数据库连接相关
│               ├── GO.migrate.go   # 数据库迁移相关
│               └── GO.seed.go      # 数据库初始化种子数据
├── error                  # 错误处理相关代码
│   └── GO.error.go
├── log                    # 日志相关文件
├── model                  # 数据模型定义
│   ├── GO.catagory.go
│   ├── GO.product.go
│   └── GO.user.go
├── request                # 请求处理相关代码
│   ├── GO.common.go
│   └── GO.product.go
├── router                 # 路由定义
│   ├── GO.cors.go
│   └── GO.router.go
├── service                # 业务逻辑实现
│   ├── GO.product_service_test.go
│   └── GO.product_service.go
├── vendor                 # 第三方依赖库
│   └── docker-compose-db.yml # Docker Compose 文件
├── config.yaml            # 配置文件
├── go.mod                 # Go 模块管理文件
├── go.sum                 # Go 模块依赖文件
└── main.go                # 程序入口文件
```