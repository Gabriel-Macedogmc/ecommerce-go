package model

import "time"

type Order struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	ProductID int       `json:"product_id" gorm:"column:product_id"`
	Price     float64   `json:"price" gorm:"column:price"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
}
