package infra

import (
	"fmt"

	"github.com/Gabriel-Macedogmc/products-backend/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	dsn := "host=postgres-db-golang user=postgres password=docker-postgres dbname=products port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	db.AutoMigrate(&model.Product{})

	return db, nil
}
