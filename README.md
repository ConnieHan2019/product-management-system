# Product-management-system
## Get Started

Just one command

```bash
docker-compose -f docker-compose.yml up --build
```

> Mysql address: localhost:3306 </br >
> Product-management-system address: lcoalhost:8080</br>
## Unity Test and Test Integration

The location of the unit test file is pkg/service/product_service.go. </br>
If you use goland, you can run them directly. </br>

```bash
cd  pkg/service
go test -v 
```


# Project Structure
```
PRODUCT-MANAGEMENT-SYSTEM
├── .vscode
│   └── launch.json        # VS Code 配置文件
├── bin                    # executable files
├── cmd
│   └── main.go            # The entry of the application
├── pkg
│   ├── config
│   │   └── GO.config.go   # Define the configuration structure and read the configuration file
│   ├── database
│   │   ├── GO.database.go # Database connection
│   │   ├── GO.migrate.go  # Database migration
│   │   └── GO.seed.go     # Seed data
│   ├── log
│   │   └── GO.log.go      # Log related
│   ├── model
│   │   ├── GO.catagory.go # Category model
│   │   ├── GO.product.go  # Product model
│   │   └── GO.user.go     # User model
│   ├── request 
│   │   ├── GO.common.go   # Define common request structure
│   │   └── GO.product.go  # Define product request structure and related validation
│   ├── router
│   │   ├── GO.cors.go     # Cross-domain settings of the router
│   │   ├── GO.router.go   # Main router
│   │   └── GO.product.go  # Product related router
│   └── service
│       ├── GO.product_service_test.go # Product service unit test
│       └── GO.product_service.go       # Product service
├── config.yaml            # configuration file
├── docker-compose.yml     # Docker Compose file for product-system-management app and mysql
├── docker-compose-backend.yml # Docker Compose file for product-system-management app
├── docker-compose-db.yml  # Docker Compose file for mysql
├── Dockerfile             # Dockerfile for product-system-management app
├── go.mod                 # Go module file
├── go.sum                 # Go sum file
├── README.md              # README file
```