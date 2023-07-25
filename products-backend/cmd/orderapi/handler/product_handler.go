package handler

import (
	"net/http"

	"strconv"

	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/dto"
	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService}
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	var productDTO dto.ProductDTO
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(productDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := ph.productService.CreateProduct(&productDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (ph *ProductHandler) GetProducts(c *gin.Context) {
	products, err := ph.productService.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ph *ProductHandler) GetProductById(c *gin.Context) {
	convertId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	ui := uint(convertId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := ph.productService.GetProductById(ui)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	convertId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	ui := uint(convertId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ph.productService.DeleteProduct(ui)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Product deleted"})
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	convertId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	ui := uint(convertId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productDTO dto.ProductDTO
	if err := c.ShouldBindJSON(&productDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := ph.productService.UpdateProduct(ui, &productDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) CreateOrder(c *gin.Context) {
	var orderDTO dto.OrderDTO
	if err := c.ShouldBindJSON(&orderDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ph.productService.CreateOrder(&orderDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}
