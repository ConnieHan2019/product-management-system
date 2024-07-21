package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"product-management-system/pkg/request"
)

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

func CreateProduct(c *gin.Context) {
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
	c.JSON(http.StatusOK, res.UUID)

}

func ListProduct(c *gin.Context) {
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

func GetProductByName(c *gin.Context) {
	productName := c.Query("productName")
	if productName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "productName is required"})
		return
	}
	productName = strings.TrimSpace(productName)
	res, err := productService.GetProductByName(productName)
	if err != nil {
		resErr := fmt.Errorf("error getting product: %w, product name:%s", err, productName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
		return
	}
	c.JSON(http.StatusOK, res)

}

func GetProductByUUID(c *gin.Context) {
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

func UpdateProduct(c *gin.Context) {
	product := &request.Product{}
	if err := c.ShouldBindJSON(product); err != nil {
		resErr := fmt.Errorf("error binding product: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": resErr.Error()})
	}
	if product.UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "uuid is required"})
		return
	}
	res, err := productService.UpdateProduct(product)
	if err != nil {
		resErr := fmt.Errorf("error updating product: %w, product name:%s", err, product.ProductName)
		c.JSON(http.StatusInternalServerError, gin.H{"error": resErr.Error()})
	}
	c.JSON(http.StatusOK, res.UUID)
}

func DeleteProductById(c *gin.Context) {
	uuid := c.Query("uuid")
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
