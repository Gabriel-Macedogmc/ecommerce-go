package dto

type OrderDTO struct {
	UserID    int     `json:"user_id" binding:"required"`
	ProductID int     `json:"product_id" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}
