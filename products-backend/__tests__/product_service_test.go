package tests

import (
	"testing"

	"github.com/Gabriel-Macedogmc/ecommerce-go01~/internal/domain/dto"
	"github.com/Gabriel-Macedogmc/ecommerce-go01~/internal/domain/model"
	"github.com/Gabriel-Macedogmc/ecommerce-go01~/internal/domain/repository"
	"github.com/Gabriel-Macedogmc/ecommerce-go01~/internal/domain/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {

	db, errDB := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if errDB != nil {
		t.Error(errDB)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.AutoMigrate(&model.Product{})

	repo := repository.NewProductRepository(db)
	service := services.NewProductService(*repo)

	productData := &dto.ProductDTO{Name: "Test Product", Price: 100, Quantity: 10}

	product, err := service.CreateProduct(productData)
	if err != nil {
		t.Error(err)
	}

	if product.ID == 0 {
		t.Errorf("Product ID not created")
	}
}
