package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"

	"product-management-system/pkg/service"
)

var routerLogger logr.Logger
var productService *service.ProductService

// InitRouter initializes the router
func InitRouter(log logr.Logger, productSvc *service.ProductService) *gin.Engine {
	routerLogger = log
	productService = productSvc
	routerLogger.Info("start init router")

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[FORMATTER TEST] %v | %3d | %13v | %15s | %-7s %#v\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)
	}))

	// support cors
	router.Use(Cors())
	//pprof.Register(router)
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello from product management system")
	})
	v1 := router.Group("/api")
	{
		v1.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "Hello from product management system api")
		})

		product := v1.Group("/product")
		{
			product.POST("/create", CreateProduct)

			product.GET("/list", ListProduct)

			product.GET("/getById", GetProductByUUID)
			product.GET("/getByName", GetProductByName)

			product.POST("/update", UpdateProduct)

			product.DELETE("/deleteById", DeleteProductById)

		}

	}
	return router
}
