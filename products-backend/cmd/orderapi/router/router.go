package router

import (
	"github.com/Gabriel-Macedogmc/products-backend/cmd/orderapi/handler"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/services"
	"github.com/gin-gonic/gin"
)

func NewRouter(productService services.ProductService) *gin.Engine {
	r := gin.Default()

	productHandler := handler.NewProductHandler(productService)

	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.GetProducts)
	r.GET("/products/id/:id", productHandler.GetProductById)
	r.DELETE("/products/:id", productHandler.DeleteProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.POST("/orders", productHandler.CreateOrder)
	return r
}
