package infra

import (
	"fmt"

	"github.com/Gabriel-Macedogmc/orders-backend/internal/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConnection() (*gorm.DB, error) {
	dsn := "host=postgres-db-golang user=postgres password=docker-postgres dbname=orders port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	db.AutoMigrate(&model.Order{})

	return db, nil
}
