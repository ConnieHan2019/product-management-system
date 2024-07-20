package router

import (
	"fmt"
	"net/http"

	//"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"

	"product-management-system/pkg/request"
	"product-management-system/pkg/service"
)

var routerLogger logr.Logger
var productService *service.ProductService

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
			product.POST("/create", createProduct)

			product.GET("/list", listProduct)

			product.GET("/getById", getProductByUUID)
			product.GET("/getByName", getProductByName)

			product.POST("/update", updateProduct)

			product.POST("/deleteById", deleteProductById)
			product.POST("/deleteByName", deleteProductByName)
		}

	}
	return router
}

func getListProductParams(c *gin.Context) *request.ListProductOptions {
	productName := c.PostForm("productName")
	maxPrice := c.GetFloat64("maxPrice")
	minPrice := c.GetFloat64("minPrice")
	intMInPrice := c.GetInt("minPrice")
	intMaxPrice := c.GetInt("maxPrice")
	if intMInPrice != 0 {
		minPrice = float64(intMInPrice)
	}
	if intMaxPrice != 0 {
		maxPrice = float64(intMaxPrice)
	}
	onlyAvailable := c.GetBool("onlyAvailable")
	category := c.PostForm("category")
	providerName := c.PostForm("providerName")

	return &request.ListProductOptions{
		ProductName:   productName,
		Category:      category,
		MaxPrice:      maxPrice,
		MinPrice:      minPrice,
		OnlyAvailable: onlyAvailable,
		ProviderName:  providerName,
	}
}
func createProduct(c *gin.Context) {
	var product request.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := product.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := productService.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.ProductName)

}

func listProduct(c *gin.Context) {
	//尝试获取request.ProductL
	laptopistProductOptions := getListProductParams(c)

	routerLogger.Info("list product options", "options", laptopistProductOptions)
	res, err := productService.ListProducts(laptopistProductOptions)
	if err != nil {
		resErr := fmt.Errorf("error listing products: %w", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func getProductByName(c *gin.Context) {
	productName := c.Query("productName")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productName is required"})
		return
	}
	res, err := productService.GetProductByName(productName)
	if err != nil {
		resErr := fmt.Errorf("error getting product: %w, product name:%s", err, productName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func getProductByUUID(c *gin.Context) {
	uuid := c.Query("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid is required"})
		return
	}
	res, err := productService.GetProductByUUID(uuid)
	if err != nil {
		resErr := fmt.Errorf("error getting product: %w, product uuid:%s", err, uuid)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func updateProduct(c *gin.Context) {
	product := &request.Product{}
	if err := c.ShouldBindJSON(product); err != nil {
		resErr := fmt.Errorf("error binding product: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": resErr.Error()})
	}
	if err := product.Validate(); err != nil {
		resErr := fmt.Errorf("error validating product: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": resErr.Error()})
		return
	}
	res, err := productService.UpdateProduct(product)
	if err != nil {
		resErr := fmt.Errorf("error updating product: %w, product name:%s", err, product.ProductName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
	}
	c.JSON(http.StatusOK, res.ProductName)
}

func deleteProductById(c *gin.Context) {
	uuid := c.PostForm("uuid")
	if uuid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid is required"})
		return
	}
	err := productService.DeleteProductByUUID(uuid)
	if err != nil {
		resErr := fmt.Errorf("error deleting product: %w, product uuid:%s", err, uuid)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
		return
	}
	c.JSON(http.StatusOK, "success")
}

func deleteProductByName(c *gin.Context) {
	productName := c.PostForm("productName")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productName is required"})
		return
	}
	err := productService.DeleteProductByName(productName)
	if err != nil {
		resErr := fmt.Errorf("error deleting product: %w, product name:%s", err, productName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
		return
	}
	c.JSON(http.StatusOK, "success")
}
